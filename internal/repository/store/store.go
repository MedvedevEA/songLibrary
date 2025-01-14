package store

import (
	"songLibrary/internal/model"
	"songLibrary/internal/repository/store/dto"
)

type Store interface {
	AddGroup(name *dto.AddGroup) (*model.Group, error)
	GetGroup(groupId *dto.GetGroup) (*model.Group, error)
	GetGroups(dto *dto.GetGroups) (*model.Pagination[model.Group], error)
	UpdateGroup(dto *dto.UpdateGroup) error
	RemoveGroup(groupId *dto.RemoveGroup) error

	AddSong(dto *dto.AddSong) (*model.Song, error)
	GetSong(songId *dto.GetSong) (*model.Song, error)
	GetSongText(dto *dto.GetSongText) (*model.Pagination[model.Verse], error)
	GetSongs(dto *dto.GetSongs) (*model.Pagination[model.Song], error)
	UpdateSong(dto *dto.UpdateSong) error
	RemoveSong(songId *dto.RemoveSong) error
}
