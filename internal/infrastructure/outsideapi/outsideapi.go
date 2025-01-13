package outsideapi

type OutsideApi struct {
	OutsideApiBindAddress string
	Logger                Logger
}
type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
}

func New(outsideApiBindAddress string, logger Logger) *OutsideApi {
	return &OutsideApi{
		OutsideApiBindAddress: outsideApiBindAddress,
		Logger:                logger,
	}
}
