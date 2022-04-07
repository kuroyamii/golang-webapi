package pingServicePkg

type pingService struct {
}

func ProvidePingService() pingService {
	return pingService{}
}

func (ps *pingService) GetPingDataSuccess() string {
	return "dataaaaa"
}

func (ps *pingService) GetPingDataError() string {
	return "null"
}
