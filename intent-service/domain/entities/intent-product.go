package entities

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type IntentProductInterface interface {
	GetId() primitive.ObjectID
	GetTitle() string
	GetPrice() float64
	GetCategory() string
	GetDescription() string
	GetImage() string
}

type IntentProductServiceInterface interface {
	Get(id string) (IntentProductInterface, error)
	GetAll() ([]IntentProductInterface, error)
	Create(product IntentProductInterface) (IntentProductInterface, error)
}

type IntentProductReader interface {
	Get(id string) (IntentProductInterface, error)
	GetAll() ([]IntentProductInterface, error)
}

type IntentProductWriter interface {
	Save(product IntentProductInterface) (IntentProductInterface, error)
}

type IntentProductPersistenceInterface interface {
	IntentProductReader
	IntentProductWriter
}

type IntentProduct struct {
	ID          primitive.ObjectID `bson:"_id" valid:"required"`
	Title       string             `bson:"title" valid:"required"`
	Price       float64            `bson:"price" valid:"required,float"`
	Category    string             `bson:"category" valid:"required"`
	Description string             `bson:"description" valid:"required"`
	Image       string             `bson:"image" valid:"required"`
}

func (ip *IntentProduct) IsValid() error {
	if ip.Title == "" {
		return errors.New("Title is required")
	}
	if ip.Price == 0 {
		return errors.New("Price is required")
	}
	if ip.Category == "" {
		return errors.New("Category is required")
	}
	if ip.Description == "" {
		return errors.New("Description is required")
	}
	if ip.Image == "" {
		return errors.New("Image is required")
	}
	if ip.ID == primitive.NilObjectID {
		return errors.New("Id is required")
	}
	return nil
}

func (ip *IntentProduct) GetId() primitive.ObjectID {
	return ip.ID
}

func (ip *IntentProduct) GetTitle() string {
	return ip.Title
}

func (ip *IntentProduct) GetPrice() float64 {
	return ip.Price
}

func (ip *IntentProduct) GetCategory() string {
	return ip.Category
}

func (ip *IntentProduct) GetDescription() string {
	return ip.Description
}

func (ip *IntentProduct) GetImage() string {
	return ip.Image
}

type NewIntentProductDto struct {
	Title       string             `json:"title" valid:"required"`
	Price       float64            `json:"price" valid:"required,float"`
	Category    string             `json:"category" valid:"required"`
	Description string             `json:"description" valid:"required"`
	Image       string             `json:"image" valid:"required"`
	ID          primitive.ObjectID `bson:"id" valid:"required"`
}

func NewIntentProduct(input NewIntentProductDto) (*IntentProduct, error) {
	intentProduct := IntentProduct{
		Title:       input.Title,
		Price:       input.Price,
		Category:    input.Category,
		Description: input.Description,
		Image:       input.Image,
		ID:          input.ID,
	}

	if err := intentProduct.IsValid(); err != nil {
		return nil, err
	}

	return &intentProduct, nil
}
