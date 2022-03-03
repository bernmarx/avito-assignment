//go:generate mockgen -source $GOFILE -destination ./httpclient_mock.go -package $GOPACKAGE
package exchangerateapi

import "net/http"

type httpClient interface {
	Get(url string) (*http.Response, error)
}
