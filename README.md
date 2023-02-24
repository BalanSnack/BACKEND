# go-mvc
## folder structure
```
.
├── config              # configuration
├── docs                # swagger documentation
├── internals           
│   ├── app             # run, dependency injection
│   ├── controller      # handler
│   ├── entity          # request, response model
│   ├── repository      # abstract storage
│   ├── service         # business logic
│   └── util            # tool
└── ...
```

## dependency injection
### layers
```golang
// repository
type UserRepository interface {
	Create(avatarId uint64, email, provider string) User
	Delete(userId uint64) error
	GetByEmailAndProvider(email, provider string) (User, error)
}

// service
type LoginService struct {
	userRepository   repository.UserRepository
	avatarRepository repository.AvatarRepository
}

// controller
type LoginController struct {
	loginService *service.LoginService
}
```
### no constructor should call another constructor
```golang
func NewLoginController(loginService *LoginService) *LoginController {
	return &LoginController{
		loginService: loginService,
	}
}

loginController := NewLoginController(loginService)
```
