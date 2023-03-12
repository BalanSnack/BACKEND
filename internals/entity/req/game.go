package req

type CreateGameRequest struct {
	Title     string `json:"title"`
	TagId     uint64 `json:"tag_id"`
	Left      string `json:"left"`
	Right     string `json:"right"`
	LeftDesc  string `json:"left_desc"`
	RightDesc string `json:"right_desc"`
}

type UpdateGameRequest struct {
	GameId    uint64 `json:"game_id"`
	Title     string `json:"title"`
	TagId     uint64 `json:"tag_id"`
	Left      string `json:"left"`
	Right     string `json:"right"`
	LeftDesc  string `json:"left_desc"`
	RightDesc string `json:"right_desc"`
}
