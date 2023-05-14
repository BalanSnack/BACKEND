package mysql

import "gorm.io/gorm"

type GameTagRepo struct {
	db *gorm.DB
}

func NewGameTagRepo(db *gorm.DB) *GameTagRepo {
	return &GameTagRepo{db: db}
}

func (r *GameTagRepo) Create(gameID, tagID uint) (GameTag, error) {
	gameTag := GameTag{
		GameID: gameID,
		TagID:  tagID,
	}

	err := r.db.Create(&gameTag).Error

	return gameTag, err
}

func (r *GameTagRepo) GetAllByGameID(gameID uint) ([]GameTag, error) {
	var gameTags []GameTag

	err := r.db.Where("game_id = ?", gameID).Find(&gameTags).Error

	return gameTags, err
}

func (r *GameTagRepo) GetAllByTagID(tagID uint) ([]GameTag, error) {
	var gameTags []GameTag

	err := r.db.Where("tag_id = ?", tagID).Find(&gameTags).Error

	return gameTags, err
}

func (r *GameTagRepo) DeleteByGameID(gameID uint) (affected int64, err error) {
	tx := r.db.Where("game_id = ?", gameID).Delete(&GameTag{})
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

func (r *GameTagRepo) DeleteByTagID(tagID uint) (affected int64, err error) {
	tx := r.db.Where("tag_id = ?", tagID).Delete(&GameTag{})
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}
