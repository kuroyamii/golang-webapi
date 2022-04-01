package pingServicePkg

type pingService struct {
}

func ProvidePingService() pingService {
	return pingService{}
}

func (ps *pingService) GetPingData() string {
	return "dataaaaa"
}
