package mysql

import "gorm.io/gorm"

type AvatarRepo struct {
	db *gorm.DB
}

func NewAvatarRepo(db *gorm.DB) *AvatarRepo {
	return &AvatarRepo{db: db}
}

func (r *AvatarRepo) Create(nick string, profile string) (Avatar, error) {
	avatar := Avatar{Nick: nick, Profile: profile}

	err := r.db.Create(&avatar).Error

	return avatar, err
}

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

func (r *AvatarRepo) Delete(id uint) (affected int64, err error) {
	tx := r.db.Delete(&Avatar{}, id)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

func (r *AvatarRepo) GetByID(id uint) (Avatar, error) {
	var avatar Avatar

	err := r.db.First(&avatar, id).Error

	return avatar, err
}
