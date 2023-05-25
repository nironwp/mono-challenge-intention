package entities_test

import (
	"intent-service/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestCreateIntentProductWithAllValidParameters(t *testing.T) {
	input := entities.NewIntentProductDto{
		Title:       "Title",
		Price:       1.0,
		Category:    "Category",
		Image:       "Image",
		ID:          primitive.NewObjectID(),
		Description: "Description",
	}

	_, err := entities.NewIntentProduct(input)

	assert.Nil(t, err)
}

func TestShouldThrowAErrorWhenProductHaveInvalidName(t *testing.T) {
	input := entities.NewIntentProductDto{
		Title:       "",
		Price:       1.0,
		Category:    "Category",
		Image:       "Image",
		ID:          primitive.NewObjectID(),
		Description: "Description",
	}

	_, err := entities.NewIntentProduct(input)

	assert.Equal(t, "Title is required", err.Error())
}

func TestShouldThrowAErrorWhenProductHaveInvalidPrice(t *testing.T) {
	input := entities.NewIntentProductDto{
		Title:       "Title",
		Price:       0,
		Category:    "Category",
		Image:       "Image",
		ID:          primitive.NewObjectID(),
		Description: "Description",
	}

	_, err := entities.NewIntentProduct(input)

	assert.Equal(t, "Price is required", err.Error())
}

func TestShouldThrowAErrorWhenProductHaveInvalidCategory(t *testing.T) {
	input := entities.NewIntentProductDto{
		Title:       "Title",
		Price:       1.0,
		Category:    "",
		Image:       "Image",
		ID:          primitive.NewObjectID(),
		Description: "Description",
	}

	_, err := entities.NewIntentProduct(input)

	assert.Equal(t, "Category is required", err.Error())
}

func TestShouldThrowAErrorWhenProductHaveInvalidImage(t *testing.T) {
	input := entities.NewIntentProductDto{
		Title:       "Title",
		Price:       1.0,
		Category:    "Category",
		Image:       "",
		ID:          primitive.NewObjectID(),
		Description: "Description",
	}

	_, err := entities.NewIntentProduct(input)

	assert.Equal(t, "Image is required", err.Error())
}

func TestShouldThrowAErrorWhenProductHaveInvalidId(t *testing.T) {
	input := entities.NewIntentProductDto{
		Title:       "Title",
		Price:       1.0,
		Category:    "Category",
		Image:       "Image",
		ID:          primitive.NilObjectID,
		Description: "Description",
	}

	_, err := entities.NewIntentProduct(input)

	assert.Equal(t, "Id is required", err.Error())
}
