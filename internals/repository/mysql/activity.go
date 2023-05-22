package mysql

import "gorm.io/gorm"

type ActivityRepo struct {
	db *gorm.DB
}

func NewActivityRepo(db *gorm.DB) *ActivityRepo {
	return &ActivityRepo{db: db}
}

// CreateJoinGame 게임 참여를 기록한다. choice의 경우 true는 왼쪽, false는 오른쪽 선택을 의미한다.
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

// CreateVoteGame 게임의 좋아요 및 싫어요를 기록한다. choice의 경우 true는 up, false는 down 투표를 의미한다.
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

// CreateVoteComment 댓글의 좋아요 및 싫어요를 기록한다. choice의 경우 true는 up, false는 down 투표를 의미한다.
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

// Delete 활동 기록을 삭제한다.
func (r *ActivityRepo) Delete(id uint) (affected int64, err error) {
	tx := r.db.Delete(&Activity{}, id)
	if err = tx.Error; err != nil {
		return
	}
	affected = tx.RowsAffected

	return
}

// GetAllByAvatarID 아바타의 활동 기록을 모두 삭제한다.
func (r *ActivityRepo) GetAllByAvatarID(avatarID uint) ([]Activity, error) {
	var activities []Activity

	err := r.db.Where("avatar_id", avatarID).Find(&activities).Error

	return activities, err
}

// GetAllByAvatarIDAndGameID 아바타의 게임 참여 내역(게임 참여, 좋아요 및 싫어요)들을 조회한다.
func (r *ActivityRepo) GetAllByAvatarIDAndGameID(avatarID, gameID uint) ([]Activity, error) {
	var activities []Activity

	err := r.db.Where(map[string]interface{}{"avatar_id": avatarID, "game_id": gameID}).Find(&activities).Error

	return activities, err
}
