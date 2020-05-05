package repositories

import (
	"github.com/Dadard29/go-music-library/api"
	"github.com/Dadard29/go-music-library/models"
)

func MusicListGet(token string) []models.MusicEntity {
	var res = make([]models.MusicEntity, 0)
	api.Api.Database.Orm.Where(&models.MusicEntity{
		Token: token,
	}).Find(&res)

	return res
}

func AlbumListGet(token string) []models.MusicEntity {
	var res = make([]models.MusicEntity, 0)
	api.Api.Database.Orm.Table("music").Select("DISTINCT album").
		Where(&models.MusicEntity{}).Scan(&res)

	return res
}

func ArtistListGet(token string) []models.MusicEntity {
	var res = make([]models.MusicEntity, 0)
	api.Api.Database.Orm.Table("music").Select("DISTINCT artist").Scan(&res)

	return res
}

