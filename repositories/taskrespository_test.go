package repositories

import (
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hunderaweke/spher/domain"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return gormDB, mock
}

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMockDB()
	taskDB := db.Model(&domain.Task{})
	repo := NewTaskRepository(taskDB)
	expectedTask := domain.Task{
		Title:       "Test Title",
		Description: "Test description",
		Deadline:    time.Now().Add(24 * time.Hour),
		Priority:    1,
		StartTime:   time.Now(),
		Tags:        []domain.Tag{{Name: "education"}},
	}
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO `tasks` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("^INSERT INTO `tags` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("^INSERT INTO `task_tags` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	_, err := repo.Create(expectedTask)
	assert.NoError(err)
}
