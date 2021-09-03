package clients

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
)

type RequestClient struct {
	Url string
	Headers map[string]interface{}
	response *http.Response
}

func MakeClient(url string, headers map[string]interface{}) *RequestClient {
	return &RequestClient{
		Url:     url,
		Headers: headers,
	}
}

func (r *RequestClient) prepareRequest() (*http.Request, error) {
	req, err := http.NewRequest("GET", r.Url, nil)
	if err != nil {
		return nil, err
	}

	if r.Headers != nil {
		for key, value := range r.Headers {
			req.Header.Add(key, fmt.Sprintf("%v", value))
		}
	}

	return req, nil
}

func GetHttpClient(insecureSkipVerify bool) *http.Client {
	client := &http.Client{}
	if insecureSkipVerify {
		client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	return client
}

func (r *RequestClient) initClient() *http.Client {
	return GetHttpClient(strings.HasPrefix(strings.ToLower(r.Url), "https"))
}

func (r *RequestClient) CheckConnection() (bool, error) {
	client := r.initClient()
	req, err := r.prepareRequest()
	if err != nil {
		return false, err
	}

	var res *http.Response
	res, err = client.Do(req)

	r.response = res

	if err != nil && strings.HasSuffix(strings.ToLower(err.Error()), "no such host") {
		return false, nil
	}

	return true, err
}
