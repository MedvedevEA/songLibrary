package model

import (
	"songLibrary/internal/pkg/types"

	"github.com/google/uuid"
)

type Group struct {
	GroupId *uuid.UUID `json:"group_id"`
	Name    string     `json:"name"`
}

type Song struct {
	SongId      *uuid.UUID  `json:"song_id"`
	Group       *Group      `json:"group"`
	Name        string      `json:"name"`
	ReleaseDate *types.Date `json:"release_date"`
	Text        *string     `json:"text"`
	Link        *string     `json:"link"`
}

type Verse struct {
	Index int    `json:"index"`
	Text  string `json:"text"`
}
type Pagination[T any] struct {
	AllRecordCount int  `json:"all_record_count"`
	Data           []*T `json:"data"`
}
