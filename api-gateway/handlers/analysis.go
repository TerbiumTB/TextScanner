package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var fileAnalysisProxy *httputil.ReverseProxy

func init() {
	analysisURL, _ := url.Parse(os.Getenv("FILE_ANALYSIS_URL"))
	fileAnalysisProxy = httputil.NewSingleHostReverseProxy(analysisURL)
}

func AnalyseHandler(wr http.ResponseWriter, r *http.Request) {
	fileAnalysisProxy.ServeHTTP(wr, r)
}
