package migration

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

type avatar struct {
	ID        uint   `gorm:"column:id;primaryKey;not null;"`
	Profile   string `gorm:"column:profile;type:varchar(255);not null;"`
	NickName  string `gorm:"column:nick_name;type:varchar(32);not null;"`
	Anonymity bool   `gorm:"column:anonymity;type:tinyint(1);not null;default:0"`
}

func (avatar) TableName() string {
	return "avatar"
}

type avatarAuth struct {
	ID       uint   `gorm:"column:id;primaryKey;not null;"`
	AvatarID uint   `gorm:"column:avatar_id;not null;"`
	Email    string `gorm:"column:email;type:varchar(255);not null;"`
	Token    string `gorm:"column:token;type:varchar(255);not null;"`
	Provider string `gorm:"column:provider;type:varchar(255);not null;"`
}

func (avatarAuth) TableName() string {
	return "avatar_auth"
}

type activity struct {
	ID           uint   `gorm:"column:id;primaryKey;not null;"`
	AvatarID     uint   `gorm:"column:avatar_id;not null;"`
	GameID       uint   `gorm:"column:game_id;not null;"`
	CommentID    uint   `gorm:"column:comment_id;null;"`
	OpinionID    uint   `gorm:"column:opinion_id;null;"`
	ActivityType string `gorm:"column:activity_type;type:varchar(32);not null;"`
	Choice       string `gorm:"column:choice;type:varchar(32);null;"`
}

func (activity) TableName() string {
	return "activity"
}

type game struct {
	ID           uint      `gorm:"column:id;primaryKey;not null;"`
	Title        string    `gorm:"column:title;type:varchar(255);not null;"`
	TagID        uint      `gorm:"column:tag_id;not null;"`
	LeftPanelID  uint      `gorm:"column:left_panel_id;not null;"`
	RightPanelID uint      `gorm:"column:right_panel_id;not null;"`
	ViewCount    uint      `gorm:"column:view_count;not null;default:0"`
	AuthorID     uint      `gorm:"column:author_id;not null;comment: author avatar id;"`
	CreateTime   time.Time `gorm:"column:create_time;not null;autoCreateTime;"`
	UpdateTime   time.Time `gorm:"column:update_time;not null;autoUpdateTime;"`
}

func (game) TableName() string {
	return "game"
}

type Panel struct {
	ID          uint   `gorm:"column:id;primaryKey;not null;"`
	GameID      uint   `gorm:"column:game_id;not null;"`
	Description string `gorm:"column:description;type:varchar(255);not null;"`
	VoteCount   uint   `gorm:"column:vote_count;not null;default:0"`
}

func (Panel) TableName() string {
	return "panel"
}

type tag struct {
	ID   uint   `gorm:"column:id;primaryKey;not null;"`
	Name string `gorm:"column:name;type:varchar(255);not null;"`
}

func (tag) TableName() string {
	return "tag"
}

type comment struct {
	ID          uint      `gorm:"column:id;primaryKey;not null;"`
	GameID      uint      `gorm:"column:game_id;not null;"`
	AuthorID    uint      `gorm:"column:author_id;not null;"`
	Description string    `gorm:"column:description;type:varchar(255);not null;"`
	Password    string    `gorm:"column:password;type:varchar(255);null;"`
	VoteCount   uint      `gorm:"column:vote_count;not null;default:0"`
	CreateTime  time.Time `gorm:"column:create_time;not null;autoCreateTime;"`
	UpdateTime  time.Time `gorm:"column:update_time;not null;autoUpdateTime;"`
}

func (comment) TableName() string {
	return "comment"
}

type opinion struct {
	ID          uint      `gorm:"column:id;primaryKey;not null;"`
	GameID      uint      `gorm:"column:game_id;not null;"`
	AuthorID    uint      `gorm:"column:author_id;not null;"`
	CommentID   uint      `gorm:"column:comment_id;not null;"`
	Description string    `gorm:"column:description;type:varchar(255);not null;"`
	VoteCount   uint      `gorm:"column:vote_count;not null;default:0"`
	CreateTime  time.Time `gorm:"column:create_time;not null;autoCreateTime;"`
	UpdateTime  time.Time `gorm:"column:update_time;not null;autoUpdateTime;"`
}

func (opinion) TableName() string {
	return "opinion"
}

var InitTables = &gormigrate.Migration{
	ID: "01_init_create_tables",
	Migrate: func(tx *gorm.DB) error {
		return tx.AutoMigrate(&avatar{}, &avatarAuth{}, &activity{}, &game{}, &Panel{}, &tag{}, &comment{}, &opinion{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable(&avatar{}, &avatarAuth{}, &activity{}, &game{}, &Panel{}, &tag{}, &comment{}, &opinion{})
	},
}
