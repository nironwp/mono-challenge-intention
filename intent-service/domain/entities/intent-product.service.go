package entities

type IntentProductService struct {
	Persistence IntentProductPersistenceInterface
}

func (itp *IntentProductService) Get(id string) (IntentProductInterface, error) {
	intentProduct, err := itp.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return intentProduct, nil
}

func (itp *IntentProductService) Create(product IntentProductInterface) (IntentProductInterface, error) {
	input := NewIntentProductDto{
		Title:       product.GetTitle(),
		Price:       product.GetPrice(),
		Category:    product.GetCategory(),
		Image:       product.GetImage(),
		Description: product.GetDescription(),
		ID:          product.GetId(),
	}

	intentProduct, err := NewIntentProduct(input)

	if err != nil {
		return nil, err
	}

	saved_product, err := itp.Persistence.Save(intentProduct)

	if err != nil {
		return nil, err
	}

	return saved_product, nil
}

func NewProductService(persistence IntentProductPersistenceInterface) *IntentProductService {
	return &IntentProductService{Persistence: persistence}
}
