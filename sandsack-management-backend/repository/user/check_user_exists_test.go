package repo_user

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

type GormMockTest struct {
	suite.Suite
	DB *gorm.DB
	mock sqlmock.Sqlmock
}

func (s *GormMockTest) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T){
	suite.Run(t, new(GormMockTest))
}

func (s *GormMockTest) SetupSuite(){
	db, mock, err := sqlmock.New()
	s.mock = mock
	require.NoError(s.T(), err)
	s.DB, err = gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	require.NoError(s.T(), err)

}

func (s *GormMockTest) TestCheckUserExists() {
	var (
		count = 1
		mockExists bool
	)

	s.mock.ExpectQuery("select count(1) from public.user where email = $1;").
		WithArgs("admin@admin.com").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))

	if count == 0 {
		mockExists = false
	} else {
		mockExists = true
	}

	exists := CheckUserExists(s.DB, "admin@admin.com")

	s.T().Log("Count", count)

	if !exists || !mockExists{
		s.T().Error("dismatch in result")
	}

	require.NoError(s.T(), errors.New("user does not exist"))
}