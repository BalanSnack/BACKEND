package agg

// Action is a parent structure of all interactions between users and games
// TODO: Action should be an interface
type Action struct {
	ID       *int `json:"id"`
	AvatarID int  `json:"avatar"`
}

type GameLike struct {
	Action
	GameID int `json:"gameID"`
}

type CommentLike struct {
	Action
	CommentID int `json:"commentID"`
}

type GameReport struct {
	Action
	GameID int    `json:"gameID"`
	Reason string `json:"reason"`
}

type CommentReport struct {
	Action
	CommentID int    `json:"commentID"`
	Reason    string `json:"reason"`
}

type Vote struct {
	Action
	PanelID int `json:"panelID"`
}
