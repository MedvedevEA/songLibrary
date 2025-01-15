package dto

import "songLibrary/internal/pkg/types"

type GetInfoReq struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}
type GetInfoRes struct {
	ReleaseDate *types.Date `json:"releaseDate"`
	Text        string      `json:"text"`
	Link        string      `json:"link"`
}
