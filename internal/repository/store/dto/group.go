package dto

import "github.com/google/uuid"

type AddGroup struct {
	Name string `json:"name"`
}
type GetGroup struct {
	GroupId *uuid.UUID `json:"group_id"`
}
type GetGroups struct {
	Limit  *int    `json:"limit"`
	Offset *int    `json:"offset"`
	Name   *string `json:"name"`
}
type UpdateGroup struct {
	GroupId *uuid.UUID `json:"group_id"`
	Name    *string    `json:"name"`
}
type RemoveGroup struct {
	GroupId *uuid.UUID `json:"group_id"`
}
