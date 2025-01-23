package blog

import "gorm.io/gorm"

type BlogPost struct {
	gorm.Model
	Category      string `gorm:"index"`
	Title         string
	Summary       string
	Content       string
	Duration_time string
	Writer        string `gorm:"index"`
	WriterRole    string
	Image         string
	ImageAlt      string
	WriterPicture string
	PictureAlt    string
	Link          string
}

func (BlogPost) TableName() string {
	return "tbl_blog_posts"
}
