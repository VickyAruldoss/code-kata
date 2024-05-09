package client

//go:generate mockgen -source=http_client.go -destination=mock/mock_http_client.go
import "net/http"

// HttpClient interface
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
