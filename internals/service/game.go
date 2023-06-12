package service

import (
	"fmt"
	"github.com/BalanSnack/BACKEND/internals/entity"
	"github.com/BalanSnack/BACKEND/internals/repository/mysql"
	"log"
)

type GameService struct {
	gameRepository    *mysql.GameRepository
	likeRepository    *mysql.LikeRepository
	voteRepository    *mysql.VoteRepository
	commentRepository *mysql.CommentRepository
}

func NewGameService(gameRepository *mysql.GameRepository, likeRepository *mysql.LikeRepository, voteRepository *mysql.VoteRepository, commentRepository *mysql.CommentRepository) *GameService {
	return &GameService{
		gameRepository:    gameRepository,
		likeRepository:    likeRepository,
		voteRepository:    voteRepository,
		commentRepository: commentRepository,
	}
}

func (s *GameService) Create(avatarID int, req entity.CreateGameRequest) (*pkg.Game, error) {
	game := pkg.Game{
		LeftOption:  req.LeftOption,
		RightOption: req.RightOption,
		LeftDesc:    req.LeftDesc,
		RightDesc:   req.RightDesc,
		AvatarID:    avatarID,
	}

	err := s.gameRepository.Create(&game)
	if err != nil {
		return nil, err
	}

	return &game, err
}

func (s *GameService) Update(req entity.UpdateGameRequest) (*pkg.Game, error) {
	game := pkg.Game{
		ID:          req.ID,
		LeftOption:  req.LeftOption,
		RightOption: req.RightOption,
		LeftDesc:    req.LeftDesc,
		RightDesc:   req.RightDesc,
	}

	err := s.gameRepository.Update(&game)
	if err != nil {
		return nil, err
	}

	return &game, err
}

func (s *GameService) Delete(id int) error {
	err := s.gameRepository.Delete(id)

	return err
}

// Get 게임 정보 조회, 파라미터로 넘어온 클래스에 따라 다음 게임의 ID도 함께 반환
func (s *GameService) Get(id, avatarID int, class string) (*entity.GetGameResponse, error) {
	res := entity.GetGameResponse{}
	var err error

	// 첫 화면에서 API 호출 시
	if id == 0 {
		id, err = s.getNext(avatarID, id, class)
		if err != nil {
			return nil, err
		}
	}

	game, err := s.gameRepository.Get(id)
	if err != nil {
		return nil, err
	}
	res.Game = *game

	votes, err := s.voteRepository.GetByGameID(id)
	if err != nil {
		return nil, err
	}
	for _, v := range votes {
		if v.Pick {
			res.RightVoteCount++
		} else {
			res.LeftVoteCount++
		}
	}

	if v, ok := votes[avatarID]; ok {
		res.Voted = true
		res.Pick = v.Pick // true: 오른쪽, false: 왼쪽
	}

	likes, err := s.likeRepository.GetLikeGameByGameID(id)
	if err != nil {
		return nil, err
	}

	res.LikeCount = len(likes)
	if _, ok := likes[avatarID]; ok {
		res.Liked = true
	}

	// 댓글 리스트 조회
	res.Comments, err = s.getComments(id, avatarID)
	if err != nil {
		return nil, err
	}

	// 다음 게임 ID
	res.Next, err = s.getNext(avatarID, game.ID, class)
	if err != nil {
		return nil, err
	}

	return &res, err
}

// getNext 해당되는 다음 게임이 없으면 id == 0 반환
func (s *GameService) getNext(avatarID, gameID int, class string) (next int, err error) {
	switch class {
	case "recent":
		next, err = s.gameRepository.GetNextRecentGame(avatarID, gameID)
	case "random":
		next, err = s.gameRepository.GetNextRandomGame(avatarID, gameID)
	default:
		err = fmt.Errorf("invalid class; %v", class)
	}
	return next, err
}

// getComments 해당 게임의 댓글들과 각 댓글의 메타 데이터를 조회하여 반환
func (s *GameService) getComments(gameID, avatarID int) ([]*entity.Comment, error) {
	comments, err := s.commentRepository.GetByGameID(gameID)
	if err != nil {
		return nil, err
	}

	m := make(map[int]*entity.Comment)
	rst := make([]*entity.Comment, 0, len(comments))
	for _, comment := range comments {
		m[comment.ID] = &entity.Comment{
			ID:       comment.ID,
			Content:  comment.Content,
			Deleted:  comment.Deleted,
			Children: make([]*entity.Comment, 0),
		}

		if comment.ParentID == 0 { // 대댓글인 경우
			rst = append(rst, m[comment.ID])
		} else { // 아닌 경우
			m[comment.ParentID].Children = append(m[comment.ParentID].Children, m[comment.ID])
		}
	}

	likes, err := s.likeRepository.GetLikeCommentByGameID(gameID)
	if err != nil {
		return nil, err
	}
	// 좋아요 개수 계산
	for _, v := range likes {
		for _, c := range v {
			log.Println(*c)
			m[c.CommentID].Likes++
		}
	}
	// 유저의 좋아요 마킹
	for _, v := range likes[avatarID] {
		m[v.CommentID].Liked = true
	}

	return rst, nil
}
