package usecase

type UserUseCase struct {
	CreateUserRepo UserRepository
}

func NewUserCase(createUserRepo UserRepository) *UserUseCase {
	return &UserUseCase{CreateUserRepo: createUserRepo}
}
