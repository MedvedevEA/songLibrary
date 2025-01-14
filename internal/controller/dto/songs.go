package dto

import (
	"songLibrary/internal/pkg/types"

	"github.com/google/uuid"
)

type AddSong struct {
	Group string `json:"group" binding:"required"`
	Song  string `json:"song" binding:"required"`
}
type GetSong struct {
	SongId string `uri:"song_id" binding:"required,uuid4_rfc4122"`
}
type GetSongText struct {
	SongId string `uri:"song_id" binding:"required,uuid4_rfc4122"`
	Limit  *int   `form:"limit"`
	Offset *int   `form:"offset"`
}
type GetSongs struct {
	Limit       *int        `form:"limit"`
	Offset      *int        `form:"offset"`
	Group       *string     `form:"group"`
	Name        *string     `form:"name"`
	ReleaseDate *types.Date `form:"release_date"`
	Text        *string     `form:"text"`
	Link        *string     `form:"link"`
}

type UpdateSong struct {
	SongId         string      `uri:"song_id" binding:"required,uuid4_rfc4122"`
	GroupId        *uuid.UUID  `json:"group_id"`
	Name           *string     `json:"name"`
	ReleaseDate    *types.Date `json:"release_date"`
	Text           *string     `json:"text"`
	Link           *string     `json:"link"`
	SetReleaseDate *bool       `json:"set_release_date"`
	SetText        *bool       `json:"set_text"`
	SetLink        *bool       `json:"set_link"`
}
type RemoveSong struct {
	SongId string `uri:"song_id" binding:"required,uuid4_rfc4122"`
}
