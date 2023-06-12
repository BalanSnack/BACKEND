package entity

type CreateGameRequest struct {
	Title       string `json:"title"`
	LeftOption  string `json:"left_option"`
	RightOption string `json:"right_option"`
	LeftDesc    string `json:"left_desc"`
	RightDesc   string `json:"right_desc"`
}

type UpdateGameRequest struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	LeftOption  string `json:"left_option"`
	RightOption string `json:"right_option"`
	LeftDesc    string `json:"left_desc"`
	RightDesc   string `json:"right_desc"`
}
