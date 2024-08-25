package repositories

/* import (
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hunderaweke/spher/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type TaskRepoTestSuite struct {
	suite.Suite
	db   *gorm.DB
	mock sqlmock.Sqlmock
	data []domain.Task
}

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

func (suite *TaskRepoTestSuite) SetupSuite() {
	db, mock := NewMockDB()
	suite.db = db
	suite.mock = mock
	suite.db.AutoMigrate(&domain.Task{})
	suite.data = []domain.Task{
		{
			Title:       "Complete project report",
			Tags:        []domain.Tag{{Name: "Work"}, {Name: "Urgent"}},
			Description: "Finish the quarterly project report and send it to the team",
			Status:      "In Progress",
			StartTime:   time.Now().Add(-2 * time.Hour),
			Deadline:    time.Now().Add(24 * time.Hour),
			Priority:    domain.PriorityMedium,
		},
		{
			Title:       "Prepare presentation",
			Tags:        []domain.Tag{{Name: "Work"}},
			Description: "Prepare slides for the upcoming client meeting",
			Status:      "Not Started",
			StartTime:   time.Now().Add(1 * time.Hour),
			Deadline:    time.Now().Add(48 * time.Hour),
			Priority:    domain.PriorityHigh,
		},
		{
			Title:       "Exercise",
			Tags:        []domain.Tag{{Name: "Health"}},
			Description: "Go for a 30-minute run in the morning",
			Status:      "Completed",
			StartTime:   time.Now().Add(-3 * time.Hour),
			Deadline:    time.Now().Add(-2 * time.Hour),
			Priority:    domain.PriorityMedium,
		},
		{
			Title:       "Read a book",
			Tags:        []domain.Tag{{Name: "Personal"}, {Name: "Leisure"}},
			Description: "Read at least 50 pages of the current novel",
			Status:      "Pending",
			StartTime:   time.Now().Add(2 * time.Hour),
			Deadline:    time.Now().Add(10 * time.Hour),
			Priority:    domain.PriorityLow,
		},
	}
}

func (suite *TaskRepoTestSuite) TestCreate() {
	assert := assert.New(suite.T())
	taskDB := suite.db.Model(&domain.Task{})
	repo := NewTaskRepository(taskDB)
	suite.mock.ExpectBegin()
	suite.mock.ExpectExec("^INSERT INTO `tasks` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectExec("^INSERT INTO `tags` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectExec("^INSERT INTO `task_tags` (.+)$").WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mock.ExpectCommit()
	createdTask, err := repo.Create(suite.data[0])
	assert.NoError(suite.mock.ExpectationsWereMet())
	suite.T().Log(createdTask.ID)
	assert.NoError(err)
}

func (suite *TaskRepoTestSuite) TestFetch() {
	assert := assert.New(suite.T())
	taskDB := suite.db.Model(&domain.Task{})
	repo := NewTaskRepository(taskDB)
	taskRows := sqlmock.NewRows([]string{"id", "title", "description", "status", "start_time", "deadline", "priority"})
	tagRows := sqlmock.NewRows([]string{"id", "name"})
	for i, task := range suite.data {
		taskRows.AddRow(i+1, task.Title, task.Description, task.Status, task.StartTime, task.Deadline, task.Priority)
		for j, tag := range task.Tags {
			tagRows.AddRow(j, tag.Name)
		}
	}
	suite.mock.ExpectQuery("^SELECT \\* FROM `tasks`").WillReturnRows(taskRows)
	tasks, err := repo.Fetch()
	assert.NoError(err)
	suite.T().Logf("%+v", tasks)
	assert.NoError(suite.mock.ExpectationsWereMet())
}

func TestTaskRepo(t *testing.T) {
	suite.Run(t, new(TaskRepoTestSuite))
} */
