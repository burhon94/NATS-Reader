package app

import "net/http"

func (s *Server) handleHealth(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	writer.Write([]byte("It reader, work!"))
}
