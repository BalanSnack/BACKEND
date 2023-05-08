package kakaoauth

import "time"

type Profile struct {
	Nickname          string `json:"nickname"`
	ThumbnailImageURL string `json:"thumbnail_image_url"`
	ProfileImageURL   string `json:"profile_image_url"`
	IsDefaultImage    bool   `json:"is_default_image"`
}

type KakaoAccount struct {
	ProfileNicknameNeedsAgreement bool    `json:"profile_nickname_needs_agreement"`
	ProfileImageNeedsAgreement    bool    `json:"profile_image_needs_agreement"`
	Profile                       Profile `json:"profile"`
	HasEmail                      bool    `json:"has_email"`
	EmailNeedsAgreement           bool    `json:"email_needs_agreement"`
	IsEmailValid                  bool    `json:"is_email_valid"`
	IsEmailVerified               bool    `json:"is_email_verified"`
	Email                         string  `json:"email"`
}

type Properties struct {
	Nickname       string `json:"nickname"`
	ProfileImage   string `json:"profile_image"`
	ThumbnailImage string `json:"thumbnail_image"`
}

type KakaoUser struct {
	ID           int64        `json:"id"`
	ConnectedAt  time.Time    `json:"connected_at"`
	Properties   Properties   `json:"properties"`
	KakaoAccount KakaoAccount `json:"kakao_account"`
}
