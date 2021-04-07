package proxy

import (
	"bytes"
	"encoding/json"
	"errors"
	"gospa/types"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type RequestPayloadStruct struct {
	ProxyCondition string `json:"proxy_condition"`
}

// Get a json decoder for a given requests body
func requestBodyDecoder(request *http.Request) *json.Decoder {
	// Read body to buffer
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		panic(err)
	}

	request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))
}

// ParseRequestBody the requests body
func ParseRequestBody(request *http.Request) RequestPayloadStruct {
	decoder := requestBodyDecoder(request)

	var requestPayload RequestPayloadStruct
	err := decoder.Decode(&requestPayload)

	if err != nil && !errors.Is(err, io.EOF) {
		panic(err)
	}

	return requestPayload
}

// GetProxyUrl the url for a given proxy condition. DEPRECATED
func GetProxyUrl(url string, conditions []types.Condition) string {

	for _, v := range conditions {
		if strings.HasPrefix(url, v.Source) {
			return v.Dest
		}
	}

	return url
}

// Serve a reverse proxy for a given url
func ServeReverseProxy(target string, res http.ResponseWriter, req *http.Request) {
	// parse the url
	url, _ := url.Parse(target)

	// create the reverse proxy
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	// Note that ServeHttp is non blocking and uses a go routine under the hood
	proxy.ServeHTTP(res, req)
}
