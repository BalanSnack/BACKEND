package agg

import "time"

// Panel is a struct that contains the information of a panel
// Should be independent of the Game struct, so that many games can use the same panel.
type Panel struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Image     *string `json:"image"`
	VoteCount int     `json:"voteCount"`
}

// Game is a struct that contains the information of a game
type Game struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	LeftPanel  Panel     `json:"leftPanel"`
	RightPanel Panel     `json:"rightPanel"`
	Comment    []Comment `json:"comment"`
	Host       Avatar    `json:"host"`
	CreatedAt  time.Time `json:"createdAt"`
}
