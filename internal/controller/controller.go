package controller

import (
	"errors"
	controllerDto "songLibrary/internal/controller/dto"
	"songLibrary/internal/logger"
	"songLibrary/internal/model"
	"songLibrary/internal/pkg/servererrors"
	storeDto "songLibrary/internal/repository/store/dto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Service interface {
	AddGroup(name *storeDto.AddGroup) (*model.Group, error)
	GetGroup(groupId *storeDto.GetGroup) (*model.Group, error)
	GetGroups(dto *storeDto.GetGroups) (*model.Pagination[model.Group], error)
	UpdateGroup(dto *storeDto.UpdateGroup) error
	RemoveGroup(groupId *storeDto.RemoveGroup) error

	AddSong(dto *storeDto.AddSong) (*model.Song, error)
	GetSong(songId *storeDto.GetSong) (*model.Song, error)
	GetSongText(dto *storeDto.GetSongText) (*model.Pagination[model.Verse], error)
	GetSongs(dto *storeDto.GetSongs) (*model.Pagination[model.Song], error)
	UpdateSong(dto *storeDto.UpdateSong) error
	RemoveSong(songId *storeDto.RemoveSong) error
}
type Controller struct {
	service Service
	logger  logger.Logger
}

func Init(router *gin.Engine, service Service, logger logger.Logger) {
	controller := &Controller{
		service: service,
		logger:  logger,
	}
	router.POST("groups", controller.addGroup)
	router.GET("groups/:group_id", controller.getGroup)
	router.GET("groups", controller.getGroups)
	router.PUT("groups/:group_id", controller.updateGroup)
	router.DELETE("groups/:group_id", controller.removeGroup)

	router.POST("songs", controller.addSong)
	router.GET("songs/:song_id", controller.getSong)
	router.GET("songs", controller.getSongs)
	router.GET("songs/:song_id/text", controller.getSongText)
	router.PUT("songs/:song_id", controller.updateSong)
	router.DELETE("songs/:song_id", controller.removeSong)

}
func (c *Controller) addGroup(ctx *gin.Context) {
	req := new(controllerDto.AddGroup)
	if err := ctx.ShouldBindJSON(req); err != nil {
		c.logger.Errorf("controller: %s: %s", "AddGroup", err)
		ctx.Status(400)
		return
	}
	res, err := c.service.AddGroup(&storeDto.AddGroup{
		Name: req.Name,
	})
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.JSON(200, res)
}
func (c *Controller) getGroup(ctx *gin.Context) {
	req := new(controllerDto.GetGroup)
	if err := ctx.ShouldBindUri(req); err != nil {
		c.logger.Errorf("controller: %s: %s", "GetGroup", err)
		ctx.Status(400)
		return
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err != nil {
		c.logger.Errorf("controller: %s: %s", "GetGroup", err)
		ctx.Status(400)
		return
	}
	res, err := c.service.GetGroup(&storeDto.GetGroup{
		GroupId: &groupId,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.Status(404)
		return
	}
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.JSON(200, res)
}

func (c *Controller) getGroups(ctx *gin.Context) {
	req := new(controllerDto.GetGroups)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		c.logger.Errorf("controller: %s: %s", "GetGroups", err)
		ctx.Status(400)
		return
	}
	res, err := c.service.GetGroups(&storeDto.GetGroups{
		Offset: req.Offset,
		Limit:  req.Limit,
		Name:   req.Name,
	})
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.JSON(200, res)
}

func (c *Controller) updateGroup(ctx *gin.Context) {
	req := new(controllerDto.UpdateGroup)
	if err := ctx.ShouldBindUri(&req); err != nil {
		c.logger.Errorf("controller: %s: %s", "UpdateGroup", err)
		ctx.Status(400)
		return
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err != nil {
		c.logger.Errorf("controller: %s: %s", "UpdateGroup", err)
		ctx.Status(400)
		return
	}
	if err := ctx.ShouldBindJSON(req); err != nil {
		c.logger.Errorf("controller: %s: %s", "UpdateGroup", err)
		ctx.Status(400)
		return
	}
	err = c.service.UpdateGroup(&storeDto.UpdateGroup{
		GroupId: &groupId,
		Name:    req.Name,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.Status(404)
		return
	}
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.Status(204)
}

func (c *Controller) removeGroup(ctx *gin.Context) {
	req := new(controllerDto.RemoveGroup)
	if err := ctx.ShouldBindUri(&req); err != nil {
		c.logger.Errorf("controller: %s: %s", "RemoveGroup", err)
		ctx.Status(400)
		return
	}
	groupId, err := uuid.Parse(req.GroupId)
	if err != nil {
		c.logger.Errorf("controller: %s: %s", "RemoveGroup", err)
		ctx.Status(400)
		return
	}
	err = c.service.RemoveGroup(&storeDto.RemoveGroup{
		GroupId: &groupId,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.Status(404)
		return
	}
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.Status(204)
}

func (c *Controller) addSong(ctx *gin.Context) {
	req := new(controllerDto.AddSong)
	if err := ctx.ShouldBindJSON(req); err != nil {
		c.logger.Errorf("controller: %s: %s", "AddSong", err)
		ctx.Status(400)
		return
	}
	res, err := c.service.AddSong(&storeDto.AddSong{
		Group: req.Group,
		Name:  req.Song,
	})
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.JSON(200, res)
}
func (c *Controller) getSong(ctx *gin.Context) {
	req := new(controllerDto.GetSong)
	if err := ctx.ShouldBindUri(req); err != nil {
		c.logger.Errorf("controller: %s: %s", "GetSong", err)
		ctx.Status(400)
		return
	}
	songId, err := uuid.Parse(req.SongId)
	if err != nil {
		c.logger.Errorf("controller: %s: %s", "GetSong", err)
		ctx.Status(400)
		return
	}
	res, err := c.service.GetSong(&storeDto.GetSong{
		SongId: &songId,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.Status(404)
		return
	}
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.JSON(200, res)
}
func (c *Controller) getSongText(ctx *gin.Context) {
	req := new(controllerDto.GetSongText)
	if err := ctx.ShouldBindUri(req); err != nil {
		c.logger.Errorf("controller: %s: %s", "GetSongText", err)
		ctx.Status(400)
		return
	}
	if err := ctx.BindQuery(req); err != nil {
		c.logger.Errorf("controller: %s: %s", "GetSongText", err)
		ctx.Status(400)
		return
	}
	songId, err := uuid.Parse(req.SongId)
	if err != nil {
		c.logger.Errorf("controller: %s: %s", "GetSongText", err)
		ctx.Status(400)
		return
	}
	res, err := c.service.GetSongText(&storeDto.GetSongText{
		SongId: &songId,
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.Status(404)
		return
	}
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.JSON(200, res)
}
func (c *Controller) getSongs(ctx *gin.Context) {
	req := new(controllerDto.GetSongs)
	if err := ctx.ShouldBindQuery(&req); err != nil {
		c.logger.Errorf("controller: %s: %s", "GetSongs", err)
		ctx.Status(400)
		return
	}
	res, err := c.service.GetSongs(&storeDto.GetSongs{
		Offset:      req.Offset,
		Limit:       req.Limit,
		Group:       req.Group,
		Name:        req.Name,
		ReleaseDate: req.ReleaseDate,
		Text:        req.Text,
		Link:        req.Link,
	})
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.JSON(200, res)
}

func (c *Controller) updateSong(ctx *gin.Context) {
	req := new(controllerDto.UpdateSong)
	if err := ctx.ShouldBindUri(&req); err != nil {
		c.logger.Errorf("controller: %s: %s", "UpdateSong", err)
		ctx.Status(400)
		return
	}
	songId, err := uuid.Parse(req.SongId)
	if err != nil {
		c.logger.Errorf("controller: %s: %s", "UpdateSong", err)
		ctx.Status(400)
		return
	}
	if err := ctx.ShouldBindJSON(req); err != nil {
		c.logger.Errorf("controller: %s: %s", "UpdateSong", err)
		ctx.Status(400)
		return
	}
	err = c.service.UpdateSong(&storeDto.UpdateSong{
		SongId:         &songId,
		GroupId:        req.GroupId,
		Name:           req.Name,
		ReleaseDate:    req.ReleaseDate,
		Text:           req.Text,
		Link:           req.Link,
		SetReleaseDate: req.SetReleaseDate,
		SetText:        req.SetText,
		SetLink:        req.SetLink,
	})
	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.Status(404)
		return
	}
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.Status(204)
}
func (c *Controller) removeSong(ctx *gin.Context) {
	req := new(controllerDto.RemoveSong)
	if err := ctx.ShouldBindUri(&req); err != nil {
		c.logger.Errorf("controller: %s: %s", "RemoveSong", err)
		ctx.Status(400)
		return
	}
	songId, err := uuid.Parse(req.SongId)
	if err != nil {
		c.logger.Errorf("controller: %s: %s", "RemoveSong", err)
		ctx.Status(400)
		return
	}
	err = c.service.RemoveSong(&storeDto.RemoveSong{
		SongId: &songId,
	})

	if errors.Is(err, servererrors.ErrorRecordNotFound) {
		ctx.Status(404)
		return
	}
	if err != nil {
		ctx.Status(500)
		return
	}
	ctx.Status(204)
}
