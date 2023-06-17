package utils

import (
	"crypto/tls"
	"github.com/go-resty/resty/v2"
	"net/http"
)

// http工具
type httpUtils struct {
}

func (u *utils) NewHttpUtils() *httpUtils {
	return &httpUtils{}
}

func (h *httpUtils) Client() (resp *resty.Request) {
	tr := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	resp = resty.NewWithClient(tr).R()
	return
}
