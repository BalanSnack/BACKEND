package mysql

import "gorm.io/gorm"

type ActivityRepo struct {
	db *gorm.DB
}

func NewActivityRepo(db *gorm.DB) *ActivityRepo {
	return &ActivityRepo{db: db}
}

// CreateJoinGame Setting the choice value to true means left, and setting false means right.
func (r *ActivityRepo) CreateJoinGame(avatarID, gameID uint, choice bool) error {
	activity := Activity{
		AvatarID: avatarID,
		GameID:   gameID,
		Type:     JoinGame,
		Choice:   choice,
	}

	err := r.db.Create(&activity).Error

	return err
}

// CreateVoteGame Setting the choice value to true means up, and setting false means down.
func (r *ActivityRepo) CreateVoteGame(avatarID, gameID uint, choice bool) error {
	activity := Activity{
		AvatarID: avatarID,
		GameID:   gameID,
		Type:     VoteGame,
		Choice:   choice,
	}

	err := r.db.Create(&activity).Error

	return err
}

// CreateVoteComment Setting the choice value to true means up, and setting false means down.
func (r *ActivityRepo) CreateVoteComment(avatarID, gameID, commentID uint, choice bool) error {
	activity := Activity{
		AvatarID:  avatarID,
		GameID:    gameID,
		CommentID: commentID,
		Type:      VoteComment,
		Choice:    choice,
	}

	err := r.db.Create(&activity).Error

	return err
}

func (r *ActivityRepo) Delete(id uint) (affected int64, err error) {
	tx := r.db.Delete(&Activity{}, id)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

func (r *ActivityRepo) GetAllByAvatarID(avatarID uint) ([]Activity, error) {
	var activities []Activity

	err := r.db.Where("avatar_id", avatarID).Find(&activities).Error

	return activities, err
}

func (r *ActivityRepo) GetAllByAvatarIDAndGameID(avatarID, gameID uint) ([]Activity, error) {
	var activities []Activity

	err := r.db.Where(map[string]interface{}{"avatar_id": avatarID, "game_id": gameID}).Find(&activities).Error

	return activities, err
}
