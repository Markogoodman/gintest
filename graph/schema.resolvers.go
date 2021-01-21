package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"math/rand"

	"github.com/Markogoodman/gqltest/graph/generated"
	"github.com/Markogoodman/gqltest/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Text:   input.Text,
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) AddRelated(ctx context.Context, input model.Relation) (*model.Todo, error) {
	var a, b *model.Todo

	for _, todo := range r.todos {
		if todo.ID == input.A {
			a = todo
		} else if todo.ID == input.B {
			b = todo
		}
	}
	if a == nil || b == nil {
		return nil, errors.New("QQ cant find")
	}
	a.Related = append(a.Related, b.ID)
	return a, nil
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	// fmt.Printf("CollectAllFields: %+v \n\n", graphql.CollectAllFields(ctx))
	// fmt.Printf("CollectFieldsCtx: \n")
	// for _, field := range graphql.CollectFieldsCtx(ctx, nil) {
	// 	fmt.Printf("  %+v\n", field.Name)
	// 	for _, sel := range field.Selections {
	// 		fmt.Printf("    %+v\n", sel)
	// 	}
	// }

	return r.todos, nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	return &model.User{
		ID:   obj.UserID,
		Name: obj.UserID + "_name",
	}, nil
}

func (r *todoResolver) Related(ctx context.Context, obj *model.Todo, count int) ([]*model.Todo, error) {
	todos := []*model.Todo{}
	i := 0
	for _, rid := range obj.Related {
		for _, todo := range r.todos {
			if rid == todo.ID {
				todos = append(todos, todo)
				i++
				if i >= count {
					break
				}
			}
		}
	}
	return todos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
