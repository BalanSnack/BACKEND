package mysql

import "gorm.io/gorm"

type MemberRepo struct {
	db *gorm.DB
}

func NewMemberRepo(db *gorm.DB) *MemberRepo {
	return &MemberRepo{db: db}
}

func (r *MemberRepo) Create(avatarID uint, email string, provider string) (Member, error) {
	member := Member{
		AvatarID: avatarID,
		Email:    email,
		Provider: provider,
	}

	err := r.db.Create(&member).Error

	return member, err
}

func (r *MemberRepo) Delete(id uint) (err error) {
	tx := r.db.Delete(&Member{}, id)
	if err = tx.Error; err != nil {
		return
	}
	if tx.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}

func (r *MemberRepo) GetAvatarIDByEmailAndProvider(email string, provider string) (uint, error) {
	var member Member

	err := r.db.Where(Member{Email: email, Provider: provider}).First(&member).Error

	return member.AvatarID, err
}
