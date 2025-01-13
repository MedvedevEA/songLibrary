package service

type Store interface {
}
type OutsizeApi interface {
}
type Logger interface {
	Debug(string)
	Debugf(string, []interface{})
	Info(string)
	Infof(string, []interface{})
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
