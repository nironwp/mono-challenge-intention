package graph

import "intent-service/domain/entities"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Service entities.IntentService
}
