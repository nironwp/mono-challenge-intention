package entities_test

import (
	"intent-service/domain/entities"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestShouldCreateAIntentWhenGiveAllValidParameters(t *testing.T) {
	input := entities.NewIntentProductDto{
		Title:       "Title",
		Price:       1.0,
		Category:    "Category",
		Image:       "Image",
		ID:          primitive.NewObjectID(),
		Description: "Description",
	}

	intentProductOne, err := entities.NewIntentProduct(input)
	assert.Nil(t, err)

	inputTwo := entities.NewIntentProductDto{
		Title:       "Title 1",
		Price:       1.0,
		Category:    "Category 2",
		Image:       "Image 2",
		ID:          primitive.NewObjectID(),
		Description: "Description 2",
	}

	intentProductTwo, err := entities.NewIntentProduct(inputTwo)
	assert.Nil(t, err)

	intentsProducts := []entities.IntentProduct{*intentProductOne, *intentProductTwo}

	createIntentDto := entities.IntentDto{
		Itens:   intentsProducts,
		User_id: "1",
	}
	intent, err := entities.NewIntent(createIntentDto)

	assert.Nil(t, err)

	assert.Equal(t, intent.Itens, intentsProducts)
	assert.Equal(t, intent.GetUserId(), "1")
	// Further test logic...
}
