package service

import (
	mocks "code-kata/client/mock"
	"code-kata/config"
	"code-kata/todo/model"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type todoServiceTestSuite struct {
	suite.Suite
	mockCntrl       *gomock.Controller
	mockHttpHandler *mocks.MockHttpRequestHandler
	mockConfig      config.Config
	todoService     TodoService
}

func TestTodoServiceTestSuite(t *testing.T) {
	suite.Run(t, new(todoServiceTestSuite))
}
func (suite *todoServiceTestSuite) SetupTest() {
	suite.mockCntrl = gomock.NewController(suite.T())
	suite.mockHttpHandler = mocks.NewMockHttpRequestHandler(suite.mockCntrl)
	suite.mockConfig = config.Config{TodoApiDetails: config.TodoApiDetails{Url: "/testurl"}}
	suite.todoService = NewTodoService(suite.mockHttpHandler, suite.mockConfig)
}

func (suite *todoServiceTestSuite) TearDownTest() {
	suite.mockCntrl.Finish()
}

func (suite *todoServiceTestSuite) TestGet_Should_Fetch_Todolist_For_The_limit() {
	base := 1
	suite.mockHttpHandler.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(func(url string, response *model.Todo) error {
		*response = getMockResponse(base)
		base = base + 1
		return nil
	}).Times(8)

	gotResponse, gotErr := suite.todoService.Get(4, true)
	suite.Nil(gotErr)
	suite.NotNil(gotResponse)
}

func (suite *todoServiceTestSuite) TestGet_Should_Fetch_Todolist_For_The_limit_when_its_not_odd() {

	base := 1
	suite.mockHttpHandler.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(func(url string, response *model.Todo) error {
		*response = getMockResponse(base)
		base = base + 1
		return nil
	}).Times(4)

	gotResponse, gotErr := suite.todoService.Get(4, false)

	suite.Nil(gotErr)
	suite.NotNil(gotResponse)

}

func (suite *todoServiceTestSuite) TestGet_Should_return_error_when_api_throws_error() {

	suite.mockHttpHandler.EXPECT().Get(gomock.Any(), gomock.Any()).DoAndReturn(func(url string, response *model.Todo) error {
		return errors.New("error from api")
	}).Times(1)

	gotResponse, gotErr := suite.todoService.Get(4, false)

	suite.Nil(gotResponse)
	suite.NotNil(gotErr)
	suite.Equal("error from api", gotErr.Error())

}

func (suite *todoServiceTestSuite) TestPrint_Should_not_throw_any_errors() {

	todos := []model.Todo{
		{UserId: 1, Id: 1, Title: "test", IsCompleted: false},
		{UserId: 2, Id: 3, Title: "test one", IsCompleted: true},
		{UserId: 3, Id: 4, Title: "test two", IsCompleted: true},
		{UserId: 4, Id: 2, Title: "test three", IsCompleted: false},
		{UserId: 5, Id: 6, Title: "test four", IsCompleted: true},
	}

	gotErr := suite.todoService.Print(todos)

	suite.Nil(gotErr)

}

func getMockResponse(id int) model.Todo {
	return model.Todo{Id: id}
}
