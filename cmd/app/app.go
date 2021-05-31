package app

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"

	nats "github.com/burhon94/NATS-Reader/pkg/events"

	"github.com/gorilla/mux"
)

const MAX_UPLOAD_FILE_SIZE = 10 * 1024 * 1024 // 10mb;
type Server struct {
	router *mux.Router
	url    string
	iEvent nats.IEvent
}

func NewServer(router *mux.Router, url string, stan nats.IEvent) *Server {
	return &Server{router: router, url: url, iEvent: stan}
}

func (s *Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	s.router.ServeHTTP(writer, request)
}


func (s *Server) readNATS(w http.ResponseWriter, r *http.Request) {
	log.Print("start to read NATS")
	s.iEvent.Subscribe("test1", HandleStan)
}

type SendSMSData struct {
	User    userStruct
	Message string
}

type userStruct struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func HandleStan(msg *stan.Msg) {
	var request SendSMSData

	err := json.Unmarshal(msg.Data, &request)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(request)
}
