package outsideapi

import (
	"songLibrary/internal/model"
	"songLibrary/internal/repository/outsideapi/dto"
)

type OutsizeApi interface {
	GetInfo(dto *dto.GetInfo) (*model.Song, error)
}
