package repo

type Avatar struct {
	Id        uint64 `json:"id"`
	Nickname  string `json:"nickname"`
	Profile   string `json:"profile"`
	Anonymity bool   `json:"anonymity"`
}
