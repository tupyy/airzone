package http

import (
	"net/http"
	"time"
)

type transportWapper func(http.RoundTripper) http.RoundTripper

func GetClient() *http.Client {
	rp := createTransport(getWrappers())
	return &http.Client{
		Transport: rp,
	}
}

func getWrappers() []transportWapper {
	transportWapper := make([]transportWapper, 0, 1)
	logger := &logTransportWrapper{}
	transportWapper = append(transportWapper, logger.Wrap)
	return transportWapper
}

func createTransport(wrappers []transportWapper) http.RoundTripper {
	var rp http.RoundTripper

	rp = &http.Transport{
		Proxy:                 http.ProxyFromEnvironment,
		ResponseHeaderTimeout: 5 * time.Second,
	}

	for i := len(wrappers) - 1; i >= 0; i-- {
		rp = wrappers[i](rp)
	}

	return rp
}
