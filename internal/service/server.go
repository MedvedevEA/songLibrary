package service

type Store interface {
}
type OutsizeApi interface {
}
type Logger interface {
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})
}

type Sevice struct {
	Store      Store
	OutsideApi OutsizeApi
	Logger     Logger
}

func New(store Store, outsizeApi OutsizeApi, logger Logger) *Sevice {
	return &Sevice{
		Store:      store,
		OutsideApi: outsizeApi,
		Logger:     logger,
	}
}
