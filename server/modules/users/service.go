package users

type UserService struct {
    repo *UserRepository
}

func NewUserService(r *UserRepository) *UserService {
    return &UserService{repo: r}
}

func (s *UserService) GetUsers() []string {
    return s.repo.FindAll()
}
