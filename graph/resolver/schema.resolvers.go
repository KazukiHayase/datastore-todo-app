package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/KazukiHayase/datastore-todo-app/graph/generated"
	"github.com/KazukiHayase/datastore-todo-app/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	dsClient, err := datastore.NewClient(ctx, r.Config.GCP.ProjectID)
	if err != nil {
		return nil, err
	}
	defer dsClient.Close()

	// 本当はgqlgenのmodelと共通で使用するようにしたい
	type Todo struct {
		id        int64
		Text      string    `datastore:"text"`
		Done      bool      `datastore:"done"`
		CreatedAt time.Time `datastore:"createdAt"`
	}

	key := datastore.IncompleteKey("Todo", nil)
	newKey, err := dsClient.Put(ctx, key, &Todo{
		Text:      input.Text,
		Done:      false,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:   strconv.Itoa(int(newKey.ID)),
		Text: input.Text,
		Done: false,
		// TODO: 仮
		User: &model.User{
			ID:   "UserID",
			Name: "Name",
		},
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
