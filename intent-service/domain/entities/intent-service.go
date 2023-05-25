package entities

type IntentService struct {
	Persistence IntentPersistenceInterface
}

func (it *IntentService) Get(id string) (IntentInterface, error) {
	intent, err := it.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return intent, nil
}

func (it *IntentService) Create(intent IntentInterface) (IntentInterface, error) {
	saved_intent, err := it.Persistence.Save(intent)

	if err != nil {
		return nil, err
	}

	return saved_intent, nil
}

func (it *IntentService) GetAll() ([]IntentInterface, error) {
	intents, err := it.Persistence.GetAll()

	if err != nil {
		return nil, err
	}

	return intents, nil
}
