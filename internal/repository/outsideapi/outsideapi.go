package outsideapi

import (
	"songLibrary/internal/repository/outsideapi/dto"
)

type OutsizeApi interface {
	GetInfo(dto *dto.GetInfoReq) (*dto.GetInfoRes, error)
}
