package service

import (
	"songLibrary/internal/logger"
	"songLibrary/internal/model"
	outsideApiRepo "songLibrary/internal/repository/outsideapi"
	outsideApiDto "songLibrary/internal/repository/outsideapi/dto"
	storeRepo "songLibrary/internal/repository/store"
	storeDto "songLibrary/internal/repository/store/dto"
)

type Service struct {
	store      storeRepo.Store
	outsideApi outsideApiRepo.OutsizeApi
	logger     logger.Logger
}

func New(store storeRepo.Store, outsizeApi outsideApiRepo.OutsizeApi, logger logger.Logger) *Service {
	return &Service{
		store:      store,
		outsideApi: outsizeApi,
		logger:     logger,
	}
}

func (s *Service) AddGroup(req *storeDto.AddGroup) (*model.Group, error) {
	return s.store.AddGroup(req)
}
func (s *Service) GetGroup(req *storeDto.GetGroup) (*model.Group, error) {
	return s.store.GetGroup(req)
}
func (s *Service) GetGroups(req *storeDto.GetGroups) (*model.Pagination[model.Group], error) {
	if req.Offset == nil {
		offset := 0
		req.Offset = &offset
	}
	if req.Limit == nil {
		limit := 10
		req.Limit = &limit
	}
	return s.store.GetGroups(req)
}
func (s *Service) UpdateGroup(req *storeDto.UpdateGroup) error {
	return s.store.UpdateGroup(req)
}
func (s *Service) RemoveGroup(req *storeDto.RemoveGroup) error {
	return s.store.RemoveGroup(req)
}

func (s *Service) AddSong(req *storeDto.AddSong) (*model.Song, error) {

	//Запрос на внешний api сервер
	outsideApiRes, err := s.outsideApi.GetInfo(&outsideApiDto.GetInfo{
		Group: req.Group,
		Song:  req.Name,
	})
	if err == nil {
		s.logger.Debugf("service: AddSong: response received, creating a new record with outside API server data")
		return s.store.AddSong(&storeDto.AddSong{
			Group:       req.Group,
			Name:        req.Name,
			ReleaseDate: outsideApiRes.ReleaseDate,
			Text:        outsideApiRes.Text,
			Link:        outsideApiRes.Link,
		})

	}
	s.logger.Debugf("service: AddSong: no response received, creating a new record without outside API server data")
	return s.store.AddSong(req)
}
func (s *Service) GetSong(req *storeDto.GetSong) (*model.Song, error) {
	return s.store.GetSong(req)
}
func (s *Service) GetSongText(req *storeDto.GetSongText) (*model.Pagination[model.Verse], error) {
	if req.Offset == nil {
		offset := 0
		req.Offset = &offset
	}
	if req.Limit == nil {
		limit := 5
		req.Limit = &limit
	}
	return s.store.GetSongText(req)
}
func (s *Service) GetSongs(req *storeDto.GetSongs) (*model.Pagination[model.Song], error) {
	if req.Offset == nil {
		offset := 0
		req.Offset = &offset
	}
	if req.Limit == nil {
		limit := 10
		req.Limit = &limit
	}
	return s.store.GetSongs(req)
}
func (s *Service) UpdateSong(req *storeDto.UpdateSong) error {
	return s.store.UpdateSong(req)
}
func (s *Service) RemoveSong(req *storeDto.RemoveSong) error {
	return s.store.RemoveSong(req)
}
