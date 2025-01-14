package postgresql

import (
	"database/sql"
	"embed"
	"errors"

	"songLibrary/internal/logger"
	"songLibrary/internal/model"
	"songLibrary/internal/pkg/helpers"
	"songLibrary/internal/pkg/servererrors"
	"songLibrary/internal/repository/store/dto"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type Postgresql struct {
	db     *sql.DB
	logger logger.Logger
}

func New(databaseConnectString string, logger logger.Logger, embedMigrations embed.FS) (*Postgresql, error) {
	//db
	db, err := sql.Open("postgres", databaseConnectString)
	if err != nil {
		return nil, err
	}
	//migration
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}
	if err := goose.Up(db, "migrations"); err != nil {
		return nil, err
	}

	return &Postgresql{
		db:     db,
		logger: logger,
	}, nil
}
func (p *Postgresql) Close() error {
	return p.db.Close()
}

func (p *Postgresql) AddGroup(req *dto.AddGroup) (*model.Group, error) {
	res, err := helpers.StoreJsonRequest[dto.AddGroup, model.Group](p.db, p.logger, "add_group", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "AddGroup", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil
}
func (p *Postgresql) GetGroup(req *dto.GetGroup) (*model.Group, error) {
	res, err := helpers.StoreJsonRequest[dto.GetGroup, model.Group](p.db, p.logger, "get_group", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "GetGroup", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil

}
func (p *Postgresql) GetGroups(req *dto.GetGroups) (*model.Pagination[model.Group], error) {
	res, err := helpers.StoreJsonRequest[dto.GetGroups, model.Pagination[model.Group]](p.db, p.logger, "get_groups", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "GetGroups", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil
}
func (p *Postgresql) UpdateGroup(req *dto.UpdateGroup) error {
	_, err := helpers.StoreJsonRequest[dto.UpdateGroup, model.Group](p.db, p.logger, "update_group", req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return servererrors.ErrorRecordNotFound
		}
		p.logger.Errorf("store: %s: %s", "UpdateGroups", err)
		return servererrors.ErrorInternal
	}
	return nil
}
func (p *Postgresql) RemoveGroup(req *dto.RemoveGroup) error {
	_, err := helpers.StoreJsonRequest[dto.RemoveGroup, model.Group](p.db, p.logger, "remove_group", req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return servererrors.ErrorRecordNotFound
		}
		p.logger.Errorf("store: %s: %s", "RemoveGroups", err)
		return servererrors.ErrorInternal
	}
	return nil
}
func (p *Postgresql) AddSong(req *dto.AddSong) (*model.Song, error) {
	res, err := helpers.StoreJsonRequest[dto.AddSong, model.Song](p.db, p.logger, "add_song", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "AddSong", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil
}
func (p *Postgresql) GetSong(req *dto.GetSong) (*model.Song, error) {
	res, err := helpers.StoreJsonRequest[dto.GetSong, model.Song](p.db, p.logger, "get_song", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "GetSong", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil
}
func (p *Postgresql) GetSongText(req *dto.GetSongText) (*model.Pagination[model.Verse], error) {
	res, err := helpers.StoreJsonRequest[dto.GetSongText, model.Pagination[model.Verse]](p.db, p.logger, "get_song_text", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "GetSongText", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil
}
func (p *Postgresql) GetSongs(req *dto.GetSongs) (*model.Pagination[model.Song], error) {
	res, err := helpers.StoreJsonRequest[dto.GetSongs, model.Pagination[model.Song]](p.db, p.logger, "get_songs", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "GetSongs", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil
}
func (p *Postgresql) UpdateSong(req *dto.UpdateSong) error {
	_, err := helpers.StoreJsonRequest[dto.UpdateSong, model.Song](p.db, p.logger, "update_song", req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return servererrors.ErrorRecordNotFound
		}
		p.logger.Errorf("store: %s: %s", "UpdateSong", err)
		return servererrors.ErrorInternal
	}
	return nil
}
func (p *Postgresql) RemoveSong(req *dto.RemoveSong) error {
	_, err := helpers.StoreJsonRequest[dto.RemoveSong, model.Song](p.db, p.logger, "remove_song", req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return servererrors.ErrorRecordNotFound
		}
		p.logger.Errorf("store: %s: %s", "RemoveSong", err)
		return servererrors.ErrorInternal
	}
	return nil
}
