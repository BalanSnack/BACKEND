package mysql

import "gorm.io/gorm"

type TagRepo struct {
	db *gorm.DB
}

func NewTagRepo(db *gorm.DB) *TagRepo {
	return &TagRepo{db: db}
}

func (r *TagRepo) Create(name string) (Tag, error) {
	tag := Tag{
		Name:  name,
		Count: 0,
	}

	err := r.db.Create(&tag).Error

	return tag, err
}

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

func (r *TagRepo) Delete(id uint) (affected int64, err error) {
	tx := r.db.Delete(&Tag{}, id)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

func (r *TagRepo) GetByID(id uint) (Tag, error) {
	var tag Tag

	err := r.db.First(&tag, id).Error

	return tag, err
}
