package app

import       (
	"fmt"
	"log"
)

func (s *Server) InitRoutes() {

	var routes []string
	s.router.HandleFunc(fmt.Sprintf("/%s%s", s.url, "/health"), s.handleHealth).Methods("GET")
	routes = append(routes, "/health")

	s.router.HandleFunc(fmt.Sprintf("/%s%s", s.url, "/reader"), s.readNATS).Methods("GET")

	log.Printf("routes: %v", routes)
}
