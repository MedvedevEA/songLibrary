package postgresql

import (
	"database/sql"
	"embed"
	"encoding/json"
	"errors"
	"fmt"

	"songLibrary/internal/logger"
	"songLibrary/internal/model"
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
func JsonRequest[Req any, Res any](p *Postgresql, funcName string, req *Req) (*Res, error) {
	j, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	query := fmt.Sprintf("Select * FROM public.%s($1)", funcName)
	p.logger.Debugf("| SQL query | %s | %s |", query, string(j))
	row := p.db.QueryRow(query, j)
	if err := row.Scan(&j); err != nil {
		return nil, err
	}
	res := new(Res)
	err = json.Unmarshal(j, res)
	return res, err
}
func (p *Postgresql) Close() error {
	return p.db.Close()
}

func (p *Postgresql) AddGroup(req *dto.AddGroup) (*model.Group, error) {
	res, err := JsonRequest[dto.AddGroup, model.Group](p, "add_group", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "AddGroup", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil
}
func (p *Postgresql) GetGroup(req *dto.GetGroup) (*model.Group, error) {
	res, err := JsonRequest[dto.GetGroup, model.Group](p, "get_group", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "GetGroup", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil

}
func (p *Postgresql) GetGroups(req *dto.GetGroups) (*model.Pagination[model.Group], error) {
	res, err := JsonRequest[dto.GetGroups, model.Pagination[model.Group]](p, "get_groups", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "GetGroups", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil
}
func (p *Postgresql) UpdateGroup(req *dto.UpdateGroup) error {
	_, err := JsonRequest[dto.UpdateGroup, model.Group](p, "update_group", req)
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
	_, err := JsonRequest[dto.RemoveGroup, model.Group](p, "remove_group", req)
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
	res, err := JsonRequest[dto.AddSong, model.Song](p, "add_song", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "AddSong", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil
}
func (p *Postgresql) GetSong(req *dto.GetSong) (*model.Song, error) {
	res, err := JsonRequest[dto.GetSong, model.Song](p, "get_song", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "GetSong", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil
}
func (p *Postgresql) GetSongText(req *dto.GetSongText) (*model.Pagination[model.Verse], error) {
	res, err := JsonRequest[dto.GetSongText, model.Pagination[model.Verse]](p, "get_song_text", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "GetSongText", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil
}
func (p *Postgresql) GetSongs(req *dto.GetSongs) (*model.Pagination[model.Song], error) {
	res, err := JsonRequest[dto.GetSongs, model.Pagination[model.Song]](p, "get_songs", req)
	if err != nil {
		p.logger.Errorf("store: %s: %s", "GetSongs", err)
		return nil, servererrors.ErrorInternal
	}
	return res, nil
}
func (p *Postgresql) UpdateSong(req *dto.UpdateSong) error {
	_, err := JsonRequest[dto.UpdateSong, model.Song](p, "update_song", req)
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
	_, err := JsonRequest[dto.RemoveSong, model.Song](p, "remove_song", req)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return servererrors.ErrorRecordNotFound
		}
		p.logger.Errorf("store: %s: %s", "RemoveSong", err)
		return servererrors.ErrorInternal
	}
	return nil
}
