package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"strconv"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/KazukiHayase/datastore-todo-app/graph/generated"
	"github.com/KazukiHayase/datastore-todo-app/graph/model"
)

type Todo struct {
	Text      string    `datastore:"text"`
	Done      bool      `datastore:"done"`
	CreatedAt time.Time `datastore:"createdAt"`
}

type User struct {
	Name string `datastore:"name"`
}

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	dsClient, err := datastore.NewClient(ctx, r.Config.GCP.ProjectID)
	if err != nil {
		return nil, err
	}
	defer dsClient.Close()

	var user User
	userKey := datastore.NameKey("User", input.UserID, nil)
	err = dsClient.Get(ctx, userKey, &user)
	if err != nil {
		if err == datastore.ErrNoSuchEntity {
			return nil, errors.New("指定されたユーザーは登録されていません")
		} else {
			return nil, err
		}
	}

	todoKey := datastore.IncompleteKey("Todo", userKey)
	newTodoKey, err := dsClient.Put(ctx, todoKey, &Todo{
		Text:      input.Text,
		Done:      false,
		CreatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	return &model.Todo{
		ID:   strconv.Itoa(int(newTodoKey.ID)),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   userKey.Name,
			Name: user.Name,
		},
	}, nil
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]model.Todo, error) {
	dsClient, err := datastore.NewClient(ctx, r.Config.GCP.ProjectID)
	if err != nil {
		return nil, err
	}
	defer dsClient.Close()

	var todos []Todo
	keys, err := dsClient.GetAll(ctx, datastore.NewQuery("Todo"), &todos)
	if err != nil {
		return []model.Todo{}, err
	}

	var gqlTodos []model.Todo
	for i, key := range keys {
		todo := todos[i]
		gqlTodos = append(gqlTodos, model.Todo{
			ID:   strconv.Itoa(int(key.ID)),
			Text: todo.Text,
			Done: todo.Done,
		})
	}

	return gqlTodos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
