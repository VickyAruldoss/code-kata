package client

//go:generate mockgen -source=http_request_handler.go -destination=mock/mock_http_request_handler.go
import (
	"bytes"
	configuration "code-kata/config"
	constant "code-kata/constants"
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type HttpRequestHandler interface {
	Get(url string, repsonseModel interface{}) (err error)
}

type httpRequestHandler struct {
	config     configuration.Config
	httpClient HttpClient
}

func NewHttpRequestHandler(config configuration.Config, httpClient HttpClient) HttpRequestHandler {
	return httpRequestHandler{config, httpClient}
}
func (h httpRequestHandler) Get(url string, repsonseModel interface{}) (err error) {
	err = h.makeRequest(constant.METHOD_GET, url, repsonseModel, nil)

	if err != nil {
		log.Errorf("Error while making http request : error %v", err)
	}
	return
}

func (h httpRequestHandler) makeRequest(method string, url string, reponseModel interface{}, requestBodyBytes []byte) (err error) {

	httpRequest, err := http.NewRequest(method, url, bytes.NewBuffer(requestBodyBytes))

	if err != nil {
		log.Error("Error while building request ", err)
		return
	}
	httpRequest.Header.Set("Content-Type", "application/json")

	res, err := h.httpClient.Do(httpRequest)

	if err != nil {
		log.Errorf("client: error making http request: %s\n", err)
		return
	}

	err = h.populateResponse(res, reponseModel)
	if err != nil {
		log.Error("Error while populating response model")
	}
	return
}

func (h httpRequestHandler) populateResponse(response *http.Response, reponseModel interface{}) (err error) {
	responseBytes, err := io.ReadAll(response.Body)

	if err != nil {
		log.Errorf("while populating the response: %s\n", err)
	}
	err = json.Unmarshal(responseBytes, &reponseModel)

	if err != nil {
		log.Errorf("error while unmarshalling: %s\n", err)
	}
	return
}
