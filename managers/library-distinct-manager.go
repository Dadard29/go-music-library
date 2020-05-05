package managers

import (
	"github.com/Dadard29/go-music-library/models"
	"github.com/Dadard29/go-music-library/repositories"
)

func LibraryArtistManagerGet(token string) ([]models.ArtistDto, string, error) {
	albums, msg, err := LibraryAlbumManagerGet(token)
	if err != nil {
		return nil, msg, err
	}

	artists := repositories.ArtistListGet(token)

	var res = make([]models.ArtistDto, 0)
	for _, ar := range artists {
		albumList := make([]models.AlbumDto, 0)
		for _, al := range albums {
			if ar.Artist == al.Artist {
				albumList = append(albumList, al)
			}
		}

		if len(albumList) == 0 {
			continue
		}

		res = append(res, models.ArtistDto{
			ArtistName: ar.Artist,
			AlbumList:  albumList,
		})
	}

	return res, "artist list retrieved", nil

}

func LibraryAlbumManagerGet(token string) ([]models.AlbumDto, string, error) {

	songList := repositories.MusicListGet(token)

	albumList := repositories.AlbumListGet(token)

	var res = make([]models.AlbumDto, 0)
	for _, a := range albumList {

		var artist string
		var imageUrl string
		var titleList = make([]string, 0)
		for _, s := range songList {
			if s.Album == a.Album {
				titleList = append(titleList, s.Title)
				artist = s.Artist
				imageUrl = s.ImageUrl
			}
		}

		if len(titleList) == 0 {
			continue
		}

		res = append(res, models.AlbumDto{
			AlbumName: a.Album,
			TitleList: titleList,
			Artist:    artist,
			ImageUrl:  imageUrl,
		})
	}

	return res, "album list retrieved", nil
}
