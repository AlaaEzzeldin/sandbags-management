package repo_order

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	//s.DB, err = gorm.Open("postgres", db)
	s.DB, err = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	require.NoError(s.T(), err)

}
