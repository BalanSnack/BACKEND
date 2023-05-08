package config

type JwtConfig struct {
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
	accessTokenSecret      string
	refreshTokenSecret     string
	refreshTokenMap        map[uint64]map[string]bool
}
