package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"testing"

	"code-kata/config"

	mocks "code-kata/client/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type httpRequestHandlerTestSuite struct {
	suite.Suite
	url                string
	mockCntrl          *gomock.Controller
	mockHttpClient     *mocks.MockHttpClient
	mockConfig         config.Config
	httpRequestHandler HttpRequestHandler
}

type dummyResponse struct {
	ResponseFieldA string `json:"responseFieldA" validate:"required"`
}

func TestHttpClientTestSuite(t *testing.T) {
	suite.Run(t, new(httpRequestHandlerTestSuite))
}

func (suite *httpRequestHandlerTestSuite) SetupTest() {
	suite.mockCntrl = gomock.NewController(suite.T())
	suite.mockHttpClient = mocks.NewMockHttpClient(suite.mockCntrl)
	suite.url = "https://jsonplaceholder.typicode.com/todos/1"
	suite.mockConfig = config.Config{HttpClientSettings: config.HttpClientSettings{TimeOutInSecs: 10}}
	suite.httpRequestHandler = NewHttpRequestHandler(suite.mockConfig, suite.mockHttpClient)
}

func (suite *httpRequestHandlerTestSuite) TearDownTest() {
	suite.mockCntrl.Finish()
}

func (suite *httpRequestHandlerTestSuite) TestHttpClient_should_make_api_call_without_any_error() {
	expectedResponse := dummyResponse{ResponseFieldA: "sample response value"}
	responseJsonString, _ := json.Marshal(expectedResponse)
	mockHttpResponse := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBuffer(responseJsonString))}
	suite.mockHttpClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(actualRequest *http.Request) (*http.Response, error) {
		return mockHttpResponse, nil
	}).Times(1)

	var response dummyResponse
	gotErr := suite.httpRequestHandler.Get(suite.url, &response)

	suite.Equal(expectedResponse, response)
	suite.Nil(gotErr)
}

func (suite *httpRequestHandlerTestSuite) TestHttpClient_should_Throw_Error_When_Http_Client_Throws_Error() {

	suite.mockHttpClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(actualRequest *http.Request) (*http.Response, error) {
		return &http.Response{}, errors.New("error from the client")
	}).Times(1)

	var response dummyResponse
	gotErr := suite.httpRequestHandler.Get(suite.url, &response)

	suite.NotNil(gotErr)
	suite.Equal("error from the client", gotErr.Error())
}

func (suite *httpRequestHandlerTestSuite) TestHttpClient_should_throw_error_when_response_model_conversion_fails() {
	expectedResponse := dummyResponse{ResponseFieldA: "sample response value"}
	responseJsonString, _ := json.Marshal(expectedResponse)
	mockHttpResponse := &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(bytes.NewBuffer(responseJsonString))}
	suite.mockHttpClient.EXPECT().Do(gomock.Any()).DoAndReturn(func(actualRequest *http.Request) (*http.Response, error) {
		return mockHttpResponse, nil
	}).Times(1)

	var response string
	gotErr := suite.httpRequestHandler.Get(suite.url, &response)

	suite.NotNil(gotErr)
}
