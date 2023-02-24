package repository

type User struct {
	UserId   uint64 `json:"id"`
	AvatarId uint64 `json:"avatar_id"`
	Email    string `json:"email"`
	Provider string `json:"provider"`
}

type UserRepository interface {
	Create(avatarId uint64, email, provider string) User
	Delete(userId uint64) error
	GetByEmailAndProvider(email, provider string) (User, error)
}
