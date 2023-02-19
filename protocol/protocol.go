package protocol

type LoginResponse struct {
	Id           uint64 `json:"id"`
	Nickname     string `json:"nickname"`
	Profile      string `json:"profile"`
	Anonymity    bool   `json:"anonymity"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
