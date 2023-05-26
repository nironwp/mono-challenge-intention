package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"context"
	"intent-service/domain/entities"
	"intent-service/graph/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateIntent is the resolver for the createIntent field.
func (r *mutationResolver) CreateIntent(ctx context.Context, input model.NewIntent) (*model.Intent, error) {
	intentDto := entities.IntentDto{
		User_id: *input.UserID,
	}

	for _, item := range input.Items {
		intentProductDto := entities.NewIntentProductDto{
			Title:       *item.Title,
			Price:       *item.Price,
			Category:    *item.Category,
			Image:       *item.Image,
			Description: *item.Description,
		}
		intentProductDto.ID = primitive.NewObjectID()

		intentProduct, err := entities.NewIntentProduct(intentProductDto)

		if err != nil {
			return nil, err
		}

		intentDto.Itens = append(intentDto.Itens, *intentProduct)
	}

	intent, err := entities.NewIntent(intentDto)

	if err != nil {
		return nil, err
	}

	created_it, err := r.Service.Create(intent)

	if err != nil {
		return nil, err
	}

	intent_model := model.Intent{
		UserID: input.UserID,
	}

	for _, item := range created_it.GetItens() {
		var title = item.GetTitle()
		var price = item.GetPrice()
		var category = item.GetCategory()
		var image = item.GetImage()
		var description = item.GetDescription()

		intent_model.Items = append(intent_model.Items, &model.IntentProduct{
			ID:          item.GetId().Hex(),
			Title:       &title,
			Price:       &price,
			Category:    &category,
			Image:       &image,
			Description: &description,
		})
	}

	return &intent_model, nil
}

// ListIntents is the resolver for the listIntents field.
func (r *queryResolver) ListIntents(ctx context.Context) ([]*model.Intent, error) {
	listed_intents, err := r.Service.GetAll()

	if err != nil {
		return nil, err
	}

	var intents []*model.Intent

	for _, intent := range listed_intents {
		var user_id = intent.GetUserId()
		intent_model := model.Intent{
			UserID: &user_id,
		}

		for _, item := range intent.GetItens() {
			var title = item.GetTitle()
			var price = item.GetPrice()
			var category = item.GetCategory()
			var image = item.GetImage()
			var description = item.GetDescription()

			intent_model.Items = append(intent_model.Items, &model.IntentProduct{
				ID:          item.GetId().Hex(),
				Title:       &title,
				Price:       &price,
				Category:    &category,
				Image:       &image,
				Description: &description,
			})
		}

		intents = append(intents, &intent_model)
	}

	return intents, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
