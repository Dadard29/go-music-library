package models

type AlbumDto struct {
	AlbumName string
	TitleList []string
	Artist string
	ImageUrl string
}

type ArtistDto struct {
	ArtistName string
	AlbumList []AlbumDto
}
