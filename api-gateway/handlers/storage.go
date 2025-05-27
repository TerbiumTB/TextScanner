package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var fileStorageProxy *httputil.ReverseProxy

func init() {
	storageURL, _ := url.Parse(os.Getenv("FILE_STORAGE_URL"))
	fileStorageProxy = httputil.NewSingleHostReverseProxy(storageURL)
}

func StorageHandler(wr http.ResponseWriter, r *http.Request) {
	fileStorageProxy.ServeHTTP(wr, r)
}

//func UploadHandler(wr http.ResponseWriter, r *http.Request) {
//	fileStorageProxy.ServeHTTP(wr, r)
//}
//
//func DownloadHandler(wr http.ResponseWriter, r *http.Request) {
//	fileStorageProxy.ServeHTTP(wr, r)
//}
//
//func GetRecordHandler(wr http.ResponseWriter, r *http.Request) {
//	fileStorageProxy.ServeHTTP(wr, r)
//}
//
//func GetAllRecordsHandler(wr http.ResponseWriter, r *http.Request) {
//	fileStorageProxy.ServeHTTP(wr, r)
//}
