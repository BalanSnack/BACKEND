package gorm

import "gorm.io/gorm"

type GameTagRepo struct {
	db *gorm.DB
}

func NewGameTagRepo(db *gorm.DB) *GameTagRepo {
	return &GameTagRepo{db: db}
}

// Create 게임 태그를 생성한다.
func (r *GameTagRepo) Create(gameID, tagID uint) (GameTag, error) {
	gameTag := GameTag{
		GameID: gameID,
		TagID:  tagID,
	}

	err := r.db.Create(&gameTag).Error

	return gameTag, err
}

// GetAllByGameID 게임의 태그들을 모두 조회한다.
func (r *GameTagRepo) GetAllByGameID(gameID uint) ([]GameTag, error) {
	var gameTags []GameTag

	err := r.db.Where("game_id = ?", gameID).Find(&gameTags).Error

	return gameTags, err
}

// GetAllByTagID 태그를 달고 있는 게임들을 모두 조회한다.
func (r *GameTagRepo) GetAllByTagID(tagID uint) ([]GameTag, error) {
	var gameTags []GameTag

	err := r.db.Where("tag_id = ?", tagID).Find(&gameTags).Error

	return gameTags, err
}

// DeleteByGameID 게임의 태그를 모두 제거한다.
func (r *GameTagRepo) DeleteByGameID(gameID uint) (affected int64, err error) {
	tx := r.db.Where("game_id = ?", gameID).Delete(&GameTag{})
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// DeleteByTagID 태그를 제거한다.
func (r *GameTagRepo) DeleteByTagID(tagID uint) (affected int64, err error) {
	tx := r.db.Where("tag_id = ?", tagID).Delete(&GameTag{})
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}
