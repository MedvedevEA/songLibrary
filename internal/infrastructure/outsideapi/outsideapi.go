package outsideapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"songLibrary/internal/logger"
	"songLibrary/internal/model"
	"songLibrary/internal/pkg/servererrors"
	"songLibrary/internal/repository/outsideapi/dto"
)

type OutsideApi struct {
	bindAddress string
	logger      logger.Logger
}

func New(bindAddress string, logger logger.Logger) *OutsideApi {
	return &OutsideApi{
		bindAddress: bindAddress,
		logger:      logger,
	}
}
func (o *OutsideApi) GetInfo(req *dto.GetInfo) (*model.Song, error) {
	reqUrl, err := url.Parse(o.bindAddress + "/info")
	if err != nil {
		o.logger.Errorf("outsideApi: GetInfo: %s", err)
		return nil, servererrors.ErrorInternal
	}
	params := url.Values{}
	params.Add("group", req.Group)
	params.Add("song", req.Song)
	reqUrl.RawQuery = params.Encode()
	resHttp, err := http.Get(reqUrl.String())
	if err != nil {
		o.logger.Errorf("outsideApi: GetInfo: %s", err)
		return nil, servererrors.ErrorInternal
	}
	defer resHttp.Body.Close()

	if resHttp.StatusCode != 200 {
		o.logger.Errorf("outsideApi: GetInfo: outside API server response is '%d', it should be '200' ", resHttp.StatusCode)
		return nil, servererrors.ErrorInternal
	}
	body, err := io.ReadAll(resHttp.Body)
	if err != nil {
		o.logger.Errorf("outsideApi: GetInfo: %s", err)
		return nil, servererrors.ErrorInternal
	}
	res := model.Song{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		o.logger.Errorf("outsideApi: GetInfo: %s", err)
		return nil, servererrors.ErrorInternal
	}
	return &res, nil
}
