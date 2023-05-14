package mysql

import "gorm.io/gorm"

type GameRepo struct {
	db *gorm.DB
}

func NewGameRepo(db *gorm.DB) *GameRepo {
	return &GameRepo{db: db}
}

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

func (r *GameRepo) Delete(id uint) (affected int64, err error) {
	tx := r.db.Delete(&Game{}, id)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

func (r *GameRepo) GetByID(id uint) (Game, error) {
	var game Game

	err := r.db.First(&game, id).Error

	return game, err
}

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
