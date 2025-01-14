package dto

import (
	"songLibrary/internal/pkg/types"

	"github.com/google/uuid"
)

type AddSong struct {
	Group       string      `json:"group"`
	Name        string      `json:"name"`
	ReleaseDate *types.Date `json:"release_date"`
	Text        *string     `json:"text"`
	Link        *string     `json:"link"`
}
type GetSong struct {
	SongId *uuid.UUID `json:"song_id"`
}
type GetSongText struct {
	SongId *uuid.UUID `json:"song_id"`
	Limit  *int       `json:"limit"`
	Offset *int       `json:"offset"`
}
type GetSongs struct {
	Limit       *int        `json:"limit"`
	Offset      *int        `json:"offset"`
	Group       *string     `json:"group"`
	Name        *string     `json:"name"`
	ReleaseDate *types.Date `json:"release_date"`
	Text        *string     `json:"text"`
	Link        *string     `json:"link"`
}
type UpdateSong struct {
	SongId         *uuid.UUID  `json:"song_id"`
	GroupId        *uuid.UUID  `json:"group_id"`
	Name           *string     `json:"name"`
	SetReleaseDate *bool       `json:"set_release_date"`
	ReleaseDate    *types.Date `json:"release_date"`
	SetText        *bool       `json:"set_text"`
	Text           *string     `json:"text"`
	SetLink        *bool       `json:"set_link"`
	Link           *string     `json:"link"`
}
type RemoveSong struct {
	SongId *uuid.UUID `json:"song_id"`
}
