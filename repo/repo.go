package repo

type UserRepo interface {
	Create(avatarId uint64, email, provider string) User
	Delete(userId uint64) error
	GetUserByEmailAndProvider(email, provider string) (User, error)
}

type AvatarRepo interface {
	Create(nickname, picture string, anonymity bool) Avatar
	Update(avatarId uint64, nickname, picture string) (Avatar, error)
	Delete(avatarId uint64) error
	GetAvatarById(avatarId uint64) (Avatar, error)
}
