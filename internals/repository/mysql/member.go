package mysql

import "gorm.io/gorm"

type MemberRepo struct {
	db *gorm.DB
}

func NewMemberRepo(db *gorm.DB) *MemberRepo {
	return &MemberRepo{db: db}
}

// Create 회원을 생성한다.
func (r *MemberRepo) Create(avatarID uint, email string, provider string) (Member, error) {
	member := Member{
		AvatarID: avatarID,
		Email:    email,
		Provider: provider,
	}

	err := r.db.Create(&member).Error

	return member, err
}

// Delete 회원을 삭제한다.
func (r *MemberRepo) Delete(id uint) (affected int64, err error) {
	tx := r.db.Delete(&Member{}, id)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// GetAvatarIDByEmailAndProvider email과 provider가 일치하는 회원을 조회한다.
func (r *MemberRepo) GetAvatarIDByEmailAndProvider(email string, provider string) (uint, error) {
	var member Member

	err := r.db.Where(Member{Email: email, Provider: provider}).First(&member).Error

	return member.AvatarID, err
}
