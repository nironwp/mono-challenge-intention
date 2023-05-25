package entities

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type IntentInterface interface {
	GetId() primitive.ObjectID
	GetItens() []IntentProduct
	GetTotal() float64
	GetUserId() string
}

type IntentServiceInterface interface {
	Get(id string) (IntentProductInterface, error)
	GetAll() ([]IntentInterface, error)
	Create(intent IntentInterface) (IntentInterface, error)
}

type IntentReader interface {
	Get(id string) (IntentInterface, error)
	GetAll() ([]IntentInterface, error)
}

type IntentWriter interface {
	Save(intent IntentInterface) (IntentInterface, error)
}

type IntentPersistenceInterface interface {
	IntentReader
	IntentWriter
}

type Intent struct {
	ID     primitive.ObjectID `bson:"_id" valid:"required"`
	Itens  []IntentProduct    `bson:"itens" valid:"required"`
	UserID string             `bson:"user" valid:"required"`
}

func (ip *Intent) GetId() primitive.ObjectID {
	return ip.ID
}

func (ip *Intent) GetItens() []IntentProduct {
	return ip.Itens
}

func (ip *Intent) GetTotal() float64 {
	total := 0.0
	for _, item := range ip.Itens {
		total += item.Price
	}
	return total
}
func (ip *Intent) GetUserId() string {
	return ip.UserID
}

func (ip *Intent) IsValid() error {
	if ip.UserID == "" {
		return errors.New("User is required")
	}
	if ip.Itens == nil || len(ip.Itens) == 0 {
		return errors.New("Itens is required")
	}
	if ip.ID == primitive.NilObjectID {
		return errors.New("ID is required")
	}
	return nil
}

type IntentDto struct {
	Itens   []IntentProduct `bson:"itens" valid:"required" json:"itens"`
	User_id string          `bson:"user_id" valid:"required" json:"user_id"`
}

func NewIntent(input IntentDto) (*Intent, error) {
	intent := &Intent{
		ID:     primitive.NewObjectID(),
		Itens:  input.Itens,
		UserID: input.User_id,
	}

	if err := intent.IsValid(); err != nil {
		return nil, err
	}

	return intent, nil
}
