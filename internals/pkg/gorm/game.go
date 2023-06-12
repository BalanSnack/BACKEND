package gorm

import "gorm.io/gorm"

type GameRepo struct {
	db *gorm.DB
}

func NewGameRepo(db *gorm.DB) *GameRepo {
	return &GameRepo{db: db}
}

// Create 게임을 생성한다.
func (r *GameRepo) Create(avatarID uint, title, leftOption, rightOption, leftDesc, rightDesc string) (Game, error) {
	game := Game{
		AvatarID:    avatarID,
		Title:       title,
		LeftOption:  leftOption,
		RightOption: rightOption,
		LeftDesc:    leftDesc,
		RightDesc:   rightDesc,
	}

	err := r.db.Create(&game).Error

	return game, err
}

// Update 게임 정보를 수정한다.
func (r *GameRepo) Update(id uint, title, leftOption, rightOption, leftDesc, rightDesc string) (affected int64, err error) {
	var game Game
	game.ID = id

	list := make(map[string]interface{})

	if title != "" {
		list["title"] = title
	}
	if leftOption != "" {
		list["left_option"] = leftOption
	}
	if rightOption != "" {
		list["right_option"] = rightOption
	}
	if leftDesc != "" {
		list["left_desc"] = leftDesc
	}
	if rightDesc != "" {
		list["right_desc"] = rightDesc
	}

	tx := r.db.Model(&game).Updates(list)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// Delete 게임을 삭제한다.
func (r *GameRepo) Delete(id uint) (affected int64, err error) {
	tx := r.db.Delete(&Game{}, id)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// GetByID 게임 정보를 조회한다.
func (r *GameRepo) GetByID(id uint) (Game, error) {
	var game Game

	err := r.db.First(&game, id).Error

	return game, err
}

// UpdateView 조회를 기록한다.
func (r *GameRepo) UpdateView(id uint) (affected int64, err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("view", gorm.Expr("view  + ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// UpdateVoteUp 좋아요를 기록한다.
func (r *GameRepo) UpdateVoteUp(id uint) (affected int64, err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("vote", gorm.Expr("vote  + ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// UpdateVoteDown 싫어요를 기록한다.
func (r *GameRepo) UpdateVoteDown(id uint) (affected int64, err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("vote", gorm.Expr("vote  - ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// UpdateLeftCountUp 게임 참여(왼쪽)를 기록한다.
func (r *GameRepo) UpdateLeftCountUp(id uint) (affected int64, err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("left_count", gorm.Expr("left_count  + ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// UpdateLeftCountUp 게임 참여(왼쪽)를 취소한다.
func (r *GameRepo) UpdateLeftCountDown(id uint) (affected int64, err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("left_count", gorm.Expr("left_count  - ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// UpdateRightCountUp 게임 참여(오른쪽)을 기록한다.
func (r *GameRepo) UpdateRightCountUp(id uint) (affected int64, err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("right_count", gorm.Expr("right_count  + ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// UpdateRightCountDown 게임 참여(오른쪽)을 취소한다.
func (r *GameRepo) UpdateRightCountDown(id uint) (affected int64, err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("right_count", gorm.Expr("right_count  - ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}
