package entity

import "BACKEND/internals/pkg"

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GetGameResponse struct {
	Game           pkg.Game   `json:"game"`
	LeftVoteCount  int        `json:"left_vote_count"`
	RightVoteCount int        `json:"right_vote_count"`
	Voted          bool       `json:"voted"`
	Pick           bool       `json:"pick"`
	LikeCount      int        `json:"like_count"`
	Liked          bool       `json:"liked"`
	Comments       []*Comment `json:"comments"`
	Next           int        `json:"next"`
}

type Comment struct {
	ID       int
	Content  string
	AuthorID int
	Likes    int
	Liked    bool
	Deleted  bool
	Children []*Comment
}
