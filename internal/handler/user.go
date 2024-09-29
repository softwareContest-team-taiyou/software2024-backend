type UserInterface interface {
	GetUser(ctx context.Context, id int) (*domain.User, error)
	CreateUser(ctx context.Context, user *domain.User) error
}