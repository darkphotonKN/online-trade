package rating

import (
	"fmt"

	"github.com/darkphotonKN/ecommerce-server-go/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RatingRepository struct {
	DB *sqlx.DB
}

func NewRatingRepository(db *sqlx.DB) *RatingRepository {
	return &RatingRepository{
		DB: db,
	}
}

func (r *RatingRepository) CreateRating(userId uuid.UUID, ratingReq RatingRequest) error {

	// add a new rating under this product's id
	query := `
	INSERT INTO ratings (user_id, rating)
	VALUES (:user_id, :rating)
	`

	// temporary struct to hold values
	params := map[string]interface{}{
		"user_id": userId,
		"rating":  ratingReq.Rating,
	}

	_, err := r.DB.NamedQuery(query, params)

	if err != nil {
		return err
	}

	return nil
}

func (r *RatingRepository) GetAllRatingsByUserId(userId uuid.UUID) (*[]models.Rating, error) {

	var ratings []models.Rating

	query := `
	SELECT * FROM ratings
	WHERE ratings.user_id = $1
	`

	err := r.DB.Select(&ratings, query, userId)

	fmt.Println("ratings:", ratings)

	if err != nil {
		return nil, err
	}

	return &ratings, nil
}
