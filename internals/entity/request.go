package req

type CreateGameRequest struct {
	Title       string `json:"title"`
	LeftOption  string `json:"left_option"`
	RightOption string `json:"right_option"`
	LeftDesc    string `json:"left_desc"`
	RightDesc   string `json:"right_desc"`
}

//{
//"title": "더 좋아하는 음식",
//"left_option": "짜장면",
//"right_option": "짬뽕",
//"left_desc": "간짜장 X",
//"right_desc": "차돌짬뽕"
//}

type UpdateGameRequest struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	LeftOption  string `json:"left_option"`
	RightOption string `json:"right_option"`
	LeftDesc    string `json:"left_desc"`
	RightDesc   string `json:"right_desc"`
}
