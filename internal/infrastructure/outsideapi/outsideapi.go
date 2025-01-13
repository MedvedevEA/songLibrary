package outsideapi

type OutsideApi struct {
	OutsideApiBindAddress string
	Logger                Logger
}
type Logger interface {
	Debug(string)
	Debugf(string, []interface{})
	Info(string)
	Infof(string, []interface{})
}

func New(outsideApiBindAddress string, logger Logger) *OutsideApi {
	return &OutsideApi{
		OutsideApiBindAddress: outsideApiBindAddress,
		Logger:                logger,
	}
}
