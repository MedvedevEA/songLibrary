package dto

type AddGroup struct {
	Name string `json:"name" binding:"required"`
}
type GetGroup struct {
	GroupId string `uri:"group_id" binding:"required,uuid4_rfc4122"`
}
type GetGroups struct {
	Limit  *int    `form:"limit"`
	Offset *int    `form:"offset"`
	Name   *string `form:"name"`
}
type UpdateGroup struct {
	GroupId string  `uri:"group_id" binding:"required,uuid4_rfc4122"`
	Name    *string `json:"name"`
}
type RemoveGroup struct {
	GroupId string `uri:"group_id" binding:"required,uuid4_rfc4122"`
}
