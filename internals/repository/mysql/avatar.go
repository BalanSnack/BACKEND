package mysql

import "gorm.io/gorm"

type AvatarRepo struct {
	db *gorm.DB
}

func NewAvatarRepo(db *gorm.DB) *AvatarRepo {
	return &AvatarRepo{db: db}
}

// Create 아바타를 생성한다.
func (r *AvatarRepo) Create(nick string, profile string) (Avatar, error) {
	avatar := Avatar{Nick: nick, Profile: profile}

	err := r.db.Create(&avatar).Error

	return avatar, err
}

// UpdateNick 아바타의 닉네임을 수정한다.
func (r *AvatarRepo) UpdateNick(id uint, nick string) (affected int64, err error) {
	var avatar Avatar
	avatar.ID = id

	tx := r.db.Model(&avatar).Update("nick", nick)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// UpdateProfile 아바타의 프로필 사진을 수정한다.
func (r *AvatarRepo) UpdateProfile(id uint, profile string) (affected int64, err error) {
	var avatar Avatar
	avatar.ID = id

	tx := r.db.Model(&avatar).Update("profile", profile)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// Delete 아바타를 삭제한다.
func (r *AvatarRepo) Delete(id uint) (affected int64, err error) {
	tx := r.db.Delete(&Avatar{}, id)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// GetByID 아바타 정보를 조회한다.
func (r *AvatarRepo) GetByID(id uint) (Avatar, error) {
	var avatar Avatar

	err := r.db.First(&avatar, id).Error

	return avatar, err
}
