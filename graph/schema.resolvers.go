package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/warintorn1990/golang-graphql-mongo/database"
	"github.com/warintorn1990/golang-graphql-mongo/graph/generated"
	"github.com/warintorn1990/golang-graphql-mongo/graph/model"
)

var db = database.Connect()

func (r *mutationResolver) CreateDog(ctx context.Context, input *model.NewDog) (*model.Dog, error) {
	return db.Save(input), nil
}

func (r *queryResolver) Dog(ctx context.Context, id string) (*model.Dog, error) {
	return db.FindByID(id), nil
}

func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	return db.All(), nil
}

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
