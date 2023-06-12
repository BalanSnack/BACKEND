package gorm

import "gorm.io/gorm"

type TagRepo struct {
	db *gorm.DB
}

func NewTagRepo(db *gorm.DB) *TagRepo {
	return &TagRepo{db: db}
}

// Create 태그를 생성한다.
func (r *TagRepo) Create(name string) (Tag, error) {
	tag := Tag{
		Name:  name,
		Count: 0,
	}

	err := r.db.Create(&tag).Error

	return tag, err
}

// Update 태그 정보를 수정한다.
func (r *TagRepo) Update(id uint, name string) (affected int64, err error) {
	var tag Tag
	tag.ID = id

	tx := r.db.Model(&tag).Update("name", name)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// Delete 태그를 삭제한다.
func (r *TagRepo) Delete(id uint) (affected int64, err error) {
	tx := r.db.Delete(&Tag{}, id)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// GetByID 태그 정보를 조회한다.
func (r *TagRepo) GetByID(id uint) (Tag, error) {
	var tag Tag

	err := r.db.First(&tag, id).Error

	return tag, err
}
