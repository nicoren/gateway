package api

import "net/http"

type httpReponseWriterInterface interface {
	write(sources []http.Response) http.Response
}

type restResponseWriter struct {
}

func (w restResponseWriter) write(sources []http.Response) http.Response {
	response := new http.Response{}
	return response
}
