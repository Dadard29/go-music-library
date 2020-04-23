package managers

import (
	"github.com/Dadard29/go-music-library/models"
	"github.com/Dadard29/go-music-library/repositories"
)

func LibraryManagerCreate(b models.MusicDto, token string) (models.MusicDto, string, error) {
	var f models.MusicDto

	m, err := repositories.LibraryCreate(models.MusicEntity{
		Title:       b.Title,
		Artist:      b.Artist,
		Album:       b.Album,
		PublishedAt: b.PublishedAt,
		Genre:       b.Genre,
		ImageUrl:    b.ImageUrl,
		Token:       token,
	})
	if err != nil {
		return f, "error creating music", err
	}

	return m.ToDto(), "music added to library", nil
}

func LibraryManagerGet(p models.MusicParam) (models.MusicDto, string, error) {
	var f models.MusicDto

	m, err := repositories.LibraryGet(p)
	if err != nil {
		return f, "error getting music", err
	}

	return m.ToDto(), "music retrieved", nil
}

func LibrarManagerDelete(p models.MusicParam) (models.MusicDto, string, error) {
	var f models.MusicDto

	m, err := repositories.LibraryDelete(p)
	if err != nil {
		return f, "error deleting music", err
	}

	return m.ToDto(), "music deleted", nil
}
