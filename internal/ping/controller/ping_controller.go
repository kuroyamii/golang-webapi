package pingController

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kuroyamii/golang-webapi/internal/global"
	pingServicePkg "github.com/kuroyamii/golang-webapi/internal/ping/service"
	"github.com/kuroyamii/golang-webapi/pkg/entity/response"
)

type PingController struct {
	router *mux.Router
	ps     pingServicePkg.PingService
}

func (pc *PingController) HandlePing(w http.ResponseWriter, r *http.Request) {
	pingData := pc.ps.GetPingData()

	w.WriteHeader(http.StatusOK)
	response.NewBaseResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		response.NewErrorResponseData(
			response.NewErrorResponseValue("msg1", "error value1"),
			response.NewErrorResponseValue("msg2", "error value2"),
			response.NewErrorResponseValue("msg3", "error value3"),
		),
		pingData,
	).ToJSON(w)
}

func (pc *PingController) InitializeEndPoint() {
	pc.router.HandleFunc(global.API_PATH_PING, pc.HandlePing)
}

func ProvidePingController(router *mux.Router, ps pingServicePkg.PingService) PingController {
	return PingController{
		router: router,
		ps:     ps,
	}
}
