package agg

// Avatar is a struct that do actions on game
type Avatar struct {
	ID           *int    `json:"id"`
	ProfileImage *string `json:"profileImage"`
	NickName     string  `json:"nickName"`
}

// Member is a struct that contains auth information of avatar
type Member struct {
	ID int `json:"id"`

	AvatarID *int   `json:"avatarID"`
	Token    string `json:"token"`
	Email    string `json:"email"`
	Provider string `json:"provider"`
}
