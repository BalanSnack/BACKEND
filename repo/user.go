package repo

type User struct {
	Id       uint64 `json:"id"`
	AvatarId uint64 `json:"avatar_id"`
	Email    string `json:"email"`
	Provider string `json:"provider"`
}
