package agg

import "time"

// Comment should be implemented as a tree structure with soft delete
type Comment struct {
	ID              int        `json:"id"`
	Content         string     `json:"content"`
	ChildCommentIDs []int      `json:"childCommentIDs"`
	GameID          int        `json:"gameID"`
	Author          Avatar     `json:"author"`
	CreatedAt       time.Time  `json:"createdAt"`
	DeletedAt       *time.Time `json:"deletedAt"`
}
