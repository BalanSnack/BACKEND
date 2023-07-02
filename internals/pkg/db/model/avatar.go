package model

type Avatar struct {
	ID       int    `gorm:"primaryKey"`
	Profile  string `gorm:"column:profile;type:varchar(255);not null;"`
	NickName string `gorm:"column:nick_name;type:varchar(32);not null;"`

	AvatarAuth *AvatarAuth `gorm:"foreignKey:AvatarID"`
}

func (Avatar) TableName() string {
	return "avatar"
}

type AvatarAuth struct {
	ID       int    `gorm:"primaryKey"`
	Email    string `gorm:"column:email;type:varchar(255);not null;"`
	Token    string `gorm:"column:token;type:varchar(255);not null;"`
	Provider string `gorm:"column:provider;type:varchar(255);not null;"`
}

func (AvatarAuth) TableName() string {
	return "avatar_auth"
}
