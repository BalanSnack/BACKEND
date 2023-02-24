package repository

type Avatar struct {
	AvatarId  uint64 `json:"id"`
	Nickname  string `json:"nickname"`
	Profile   string `json:"profile"`
	Anonymity bool   `json:"anonymity"`
}

type AvatarRepository interface {
	Create(nickname, picture string, anonymity bool) Avatar
	Update(avatarId uint64, nickname, picture string) (Avatar, error)
	Delete(avatarId uint64) error
	GetByAvatarId(avatarId uint64) (Avatar, error)
}
