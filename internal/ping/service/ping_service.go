package pingServicePkg

type PingService interface {
	GetPingDataSuccess() string
	GetPingDataError() string
}
