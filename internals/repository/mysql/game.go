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

func (r *GameRepo) Update(id uint, title, leftOption, rightOption, leftDesc, rightDesc string) (game Game, err error) {
	game.ID = id

	tx := r.db.Model(&game).Updates(map[string]interface{}{
		"title":        title,
		"left_option":  leftOption,
		"right_option": rightOption,
		"left_desc":    leftDesc,
		"right_desc":   rightDesc,
	})
	if err = tx.Error; err != nil {
		return
	}
	if tx.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}

func (r *GameRepo) Delete(id uint) (err error) {
	tx := r.db.Delete(&Game{}, id)
	if err = tx.Error; err != nil {
		return
	}
	if tx.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}

func (r *GameRepo) GetByID(id uint) (Game, error) {
	var game Game

	err := r.db.First(&game, id).Error

	return game, err
}

func (r *GameRepo) UpdateView(id uint) (err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("view", gorm.Expr("view  + ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	if tx.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}

func (r *GameRepo) UpdateVoteUp(id uint) (err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("vote", gorm.Expr("vote  + ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	if tx.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}

func (r *GameRepo) UpdateVoteDown(id uint) (err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("vote", gorm.Expr("vote  - ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	if tx.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}

func (r *GameRepo) UpdateLeftCountUp(id uint) (err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("left_count", gorm.Expr("left_count  + ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	if tx.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}

func (r *GameRepo) UpdateLeftCountDown(id uint) (err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("left_count", gorm.Expr("left_count  - ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	if tx.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}

func (r *GameRepo) UpdateRightCountUp(id uint) (err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("right_count", gorm.Expr("right_count  + ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	if tx.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}

func (r *GameRepo) UpdateRightCountDown(id uint) (err error) {
	var game Game
	game.ID = id

	tx := r.db.Model(&game).UpdateColumn("right_count", gorm.Expr("right_count  - ?", 1))
	if err = tx.Error; err != nil {
		return
	}
	if tx.RowsAffected == 0 {
		err = gorm.ErrRecordNotFound
	}

	return
}
