package gorm

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	// setup docker image
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not construct pool: %s", err)
	}

	// uses pool to try to connect to Docker
	err = pool.Client.Ping()
	if err != nil {
		log.Fatalf("Could not connect to Docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	resource, err := pool.Run("gorm", "5.7", []string{"MYSQL_ROOT_PASSWORD=secret", "MYSQL_DATABASE=test"})
	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		//"root:yksyjs@tcp(localhost:3306)/test?parseTime=true"
		db, err = gorm.Open(mysql.Open(fmt.Sprintf("root:secret@(localhost:%s)/test?parseTime=true", resource.GetPort("3306/tcp"))), &gorm.Config{})

		return err
	}); err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	exit := m.Run()

	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}

	os.Exit(exit)
}

func TestCommentRepo(t *testing.T) {
	db.AutoMigrate(&Comment{})

	repo := NewCommentRepo(db)

	sample := []Comment{
		{
			AvatarID: 1,
			GameID:   1,
			Content:  "first comment",
		},
		{
			AvatarID: 2,
			GameID:   1,
			ParentID: 1,
			Content:  "child of first comment",
		},
		{
			AvatarID: 1,
			GameID:   1,
			Content:  "second comment",
		},
		{
			AvatarID: 1,
			GameID:   2,
			Content:  "another game's comment",
		},
	}

	for _, v := range sample {
		_, err := repo.Create(v.AvatarID, v.ParentID, v.GameID, v.Content)
		if err != nil {
			t.Fatal(err)
		}
	}

	comments, err := repo.GetAllByGameID(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(comments))

	affected, err := repo.UpdateVoteUp(1)
	if err != nil || affected != 1 {
		t.Fatalf("repo.UpdateVoteUp(1) failed; error: %s, affected: %d", err, affected)
	}
	comment, err := repo.GetByID(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 1, comment.Vote)
}

func TestActivityRepo(t *testing.T) {
	db.AutoMigrate(&Activity{})

	repo := NewActivityRepo(db)

	sample := []Activity{
		{
			AvatarID: 1,
			GameID:   1,
			Type:     JoinGame,
			Choice:   true,
		},
		{
			AvatarID: 1,
			GameID:   1,
			Type:     VoteGame,
			Choice:   false,
		},
		{
			AvatarID:  1,
			GameID:    1,
			Type:      VoteComment,
			CommentID: 1,
			Choice:    true,
		},
		{
			AvatarID: 2,
			GameID:   1,
			Type:     JoinGame,
			Choice:   false,
		},
		{
			AvatarID: 1,
			GameID:   2,
			Type:     JoinGame,
			Choice:   false,
		},
	}

	for _, v := range sample {
		switch v.Type {
		case JoinGame:
			err := repo.CreateJoinGame(v.AvatarID, v.GameID, v.Choice)
			if err != nil {
				t.Fatal(err)
			}
		case VoteGame:
			err := repo.CreateVoteGame(v.AvatarID, v.GameID, v.Choice)
			if err != nil {
				t.Fatal(err)
			}
		case VoteComment:
			err := repo.CreateVoteComment(v.AvatarID, v.GameID, v.CommentID, v.Choice)
			if err != nil {
				t.Fatal(err)
			}
		}
	}

	activities, err := repo.GetAllByAvatarIDAndGameID(1, 1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, 3, len(activities))
}

func TestGameRepo(t *testing.T) {
	db.AutoMigrate(&Game{})

	repo := NewGameRepo(db)

	sample := Game{
		Title:       "title",
		LeftOption:  "leftOption",
		RightOption: "rightOption",
		LeftDesc:    "leftDesc",
		RightDesc:   "rightDesc",
	}

	game, err := repo.Create(1, sample.Title, sample.LeftOption, sample.RightOption, sample.LeftDesc, sample.RightDesc)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "leftDesc", game.LeftDesc)

	affected, err := repo.Update(1, "", "", "", "updated", "")
	if err != nil || affected != 1 {
		t.Fatalf("repo.Update(1, ...) failed; error: %s, affected: %d", err, affected)
	}
	game, err = repo.GetByID(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "updated", game.LeftDesc)

	affected, err = repo.UpdateRightCountUp(1)
	if err != nil || affected != 1 {
		t.Fatalf("repo.UpdateRightCountUp(1) failed; error: %s, affected: %d", err, affected)
	}
	game, err = repo.GetByID(1)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, uint(1), game.RightCount)
}
