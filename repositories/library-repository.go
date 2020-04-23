package repositories

import (
	"errors"
	"github.com/Dadard29/go-music-library/api"
	"github.com/Dadard29/go-music-library/models"
)

func LibraryGet(m models.MusicParam) (models.MusicEntity, error) {
	var f models.MusicEntity
	var r models.MusicEntity
	e := m.ToEntity()
	api.Api.Database.Orm.Where(&e).First(&r)

	if r.Title != m.Title || r.Artist != m.Artist || r.Token != m.Token {
		return f, errors.New("entity music not found")
	}
	return r, nil
}

func LibraryExists(m models.MusicParam) bool {
	_, err := LibraryGet(m)
	return err == nil
}

func LibraryCreate(m models.MusicEntity) (models.MusicEntity, error) {
	var f models.MusicEntity

	if LibraryExists(models.MusicParam{
		Title:  m.Title,
		Artist: m.Artist,
		Token: m.Token,
	}) {
		return f, errors.New("entity music already exists")
	}

	api.Api.Database.Orm.Create(&m)

	if !LibraryExists(models.MusicParam{
		Title:  m.Title,
		Artist: m.Artist,
		Token: m.Token,
	}) {
		return f, errors.New("unexpected error while creating music entity")
	}


	return m, nil
}

func LibraryDelete(m models.MusicParam) (models.MusicEntity, error) {
	var f models.MusicEntity

	mu, err := LibraryGet(m)
	if err != nil {
		return f, errors.New("entity music not found")
	}

	api.Api.Database.Orm.Delete(&mu)

	if LibraryExists(m) {
		return f, errors.New("unexpected error while deleting music entity")
	}

	return mu, nil
}
