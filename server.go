package main

import (
	"fmt"
	"net/http"

	"github.com/highlight/fit/data"
	"github.com/highlight/fit/entities"
	"github.com/highlight/fit/router"
	"github.com/highlight/fit/service"
	"github.com/highlight/fit/utils"
)

var (
	repo    data.SessionRepository = data.NewSessionRepository()
	svc     *service.Service       = service.NewSessionService(repo)
	hRouter router.Router          = router.NewChiRouter()
)

func main() {
	const port = ":3000"
	hRouter.GET("/", getSessionsHandler)
	hRouter.POST("/save", saveSessionHandler)
	hRouter.SERVE(port)

}
func getSessionsHandler(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("getting sessions...")

	s, err := svc.GetSession("001")
	if err != nil || s == nil {
		fmt.Fprintf(resp, "error: %v", err)
		return
	}
	fmt.Fprintf(resp, "We got a session id %s", s.Id)
}

func saveSessionHandler(resp http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	s := &entities.Session{}
	if err := utils.DecodeHttpBody(*req, s); err != nil {
		fmt.Println("there was an error decoding the body: %v", err)
		return
	}
	fmt.Println("this is the session id: ", s.Id)

	svc.CreateSession(s)

}
