package blog

import (
	"nextgen/internals/db/pg"
)

func GetAllPosts() (*[]BlogPost, error) {
	var posts []BlogPost
	
	if err := pg.GDB.Find(&posts).Error; err != nil {
		return nil, err
	}
	return &posts, nil
}

func GetPostById(id int) (*BlogPost, error) {
	var post BlogPost
	
	if err := pg.GDB.Where("id = ?", id).Find(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

