package models

type AlbumDto struct {
	Name string `json:"name"`
	TitleList []string `json:"title_list"`
	Artist string `json:"artist"`
	ImageUrl string `json:"image_url"`
}

type ArtistDto struct {
	Name string `json:"name"`
	AlbumList []AlbumDto `json:"album_list"`
}
