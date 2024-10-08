package usecase

import (
	"context"
	"fmt"

	"github.com/softwareContest-team-taiyou/software2024-backend/internal/domain"
)
type TodoRepository interface{
	GetTodo(ctx context.Context,id int) (*domain.Todo, error)
	CreateTodo(ctx context.Context,todo *domain.Todo)  error
	GetAllTodos(ctx context.Context) ([]*domain.Todo, error)
}

type TodoUsecase struct {
	tr TodoRepository
}

func NewTodoUsecase(todoRepository TodoRepository) *TodoUsecase {
	return &TodoUsecase{tr: todoRepository}
}
func (tu *TodoUsecase) GetTodo(ctx context.Context,id int) (*domain.Todo, error) {
	todo,err:= tu.tr.GetTodo(ctx,id)
	if err != nil {
		return nil,err
	}
	return todo,nil
}
func (tu *TodoUsecase) CreateTodo(ctx context.Context,todo *domain.Todo) error {
	fmt.Print(ctx.Value("uid"));
	if err := tu.tr.CreateTodo(ctx,todo); err != nil {
		return err
	}
	return nil	
}
func (tu *TodoUsecase) GetAllTodos(ctx context.Context) ([]*domain.Todo, error) {
	todos,err:= tu.tr.GetAllTodos(ctx)
	if err != nil {
		return nil,err
	}
	return todos,nil
}