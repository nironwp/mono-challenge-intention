package entities_test

import (
	"testing"

	"intent-service/domain/entities"
	mock_entities "intent-service/domain/entities/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetProductService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	intentProductDto := entities.NewIntentProductDto{
		Title:       "Title",
		Price:       1.0,
		Category:    "Category",
		Image:       "Image",
		ID:          primitive.NewObjectID(),
		Description: "Description",
	}

	intentProduct, _ := entities.NewIntentProduct(intentProductDto)

	persistence := mock_entities.NewMockIntentProductPersistenceInterface(ctrl)

	persistence.EXPECT().Get(gomock.Any()).Return(intentProduct, nil).AnyTimes()

	service := entities.IntentProductService{
		Persistence: persistence,
	}
	result, err := service.Get("1")

	assert.Nil(t, err)
	assert.Equal(t, intentProduct, result)
}

func TestCreateProductService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	intentProductDto := entities.NewIntentProductDto{
		Title:       "Title",
		Price:       1.0,
		Category:    "Category",
		Image:       "Image",
		ID:          primitive.NewObjectID(),
		Description: "Description",
	}

	intentProduct, _ := entities.NewIntentProduct(intentProductDto)

	persistence := mock_entities.NewMockIntentProductPersistenceInterface(ctrl)

	persistence.EXPECT().Save(gomock.Any()).Return(intentProduct, nil).AnyTimes()

	service := entities.IntentProductService{
		Persistence: persistence,
	}
	result, err := service.Create(intentProduct)

	assert.Nil(t, err)
	assert.Equal(t, intentProduct, result)
}
