package models

type MusicDto struct {
	Title string `json:"title"`
	Artist string `json:"artist"`
	Album string `json:"album"`
	PublishedAt string `json:"published_at"`
	Genre string `json:"genre"`
	ImageUrl string `json:"image_url"`
}

type MusicParam struct {
	Title string
	Artist string
	Token string
}

func (m MusicParam) CheckSanity() bool {
	return m.Artist != "" && m.Title != "" && m.Token != ""
}

func (m MusicParam) ToEntity() MusicEntity {
	return MusicEntity{
		Title:       m.Title,
		Artist:      m.Artist,
		Album:       "",
		PublishedAt: "",
		Genre:       "",
		ImageUrl:    "",
		Token: m.Token,
	}
}

type MusicEntity struct {
	Title       string `gorm:"type:varchar(70);index:title"`
	Artist      string `gorm:"type:varchar(70);index:artist"`
	Album       string `gorm:"type:varchar(70);index:album"`
	PublishedAt string `gorm:"type:varchar(20);index:published_at"`
	Genre       string `gorm:"type:varchar(40);index:genre"`
	ImageUrl    string `gorm:"type:varchar(200);index:image_url"`

	Token string `gorm:"type:varchar(70);index:token"`
}

func (m MusicEntity) ToDto() MusicDto {
	return MusicDto{
		Title:       m.Title,
		Artist:      m.Artist,
		Album:       m.Album,
		PublishedAt: m.PublishedAt,
		Genre:       m.Genre,
		ImageUrl:    m.ImageUrl,
	}
}

func (MusicEntity) TableName() string {
	return "music"
}


