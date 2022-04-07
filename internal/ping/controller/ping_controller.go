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

func (pc *PingController) HandlePingSuccess(w http.ResponseWriter, r *http.Request) {
	pingData := pc.ps.GetPingDataSuccess()

	w.WriteHeader(http.StatusOK)
	response.NewBaseResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		response.NewErrorResponseData(),
		pingData,
	).ToJSON(w)
}
func (pc *PingController) HandlePingError(w http.ResponseWriter, r *http.Request) {
	pingData := pc.ps.GetPingDataError()

	w.WriteHeader(http.StatusOK)
	response.NewBaseResponse(
		http.StatusNoContent,
		http.StatusText(http.StatusNoContent),
		response.NewErrorResponseData(
			response.NewErrorResponseValue("msg1", "error value1"),
			response.NewErrorResponseValue("msg2", "error value2"),
			response.NewErrorResponseValue("msg3", "error value3"),
		),
		pingData,
	).ToJSON(w)
}

func (pc *PingController) InitializeEndPoint() {
	pc.router.HandleFunc(global.API_PATH_PING_SUCCESS, pc.HandlePingSuccess)
	pc.router.HandleFunc(global.API_PATH_PING_ERROR, pc.HandlePingError)
}

func ProvidePingController(router *mux.Router, ps pingServicePkg.PingService) PingController {
	return PingController{
		router: router,
		ps:     ps,
	}
}
