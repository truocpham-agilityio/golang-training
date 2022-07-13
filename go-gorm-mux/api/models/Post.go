package models

import (
	"errors"
	"go-gorm-mux/api/utils/pagination"
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

// Post represents a post model.
type Post struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title     string    `gorm:"size:255;not null;unique" json:"title"`
	Content   string    `gorm:"size:255;not null;" json:"content"`
	Author    User      `json:"author"`
	AuthorID  uint32    `sql:"type:int REFERENCES users(id)" json:"author_id"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Prepare represents the preparation of the post model data.
func (p *Post) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Content = html.EscapeString(strings.TrimSpace(p.Content))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

// Validate represents the validation of the post model data.
func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("required Title")
	}

	if p.Content == "" {
		return errors.New("required Content")
	}

	if p.AuthorID < 1 {
		return errors.New("required Author")
	}

	return nil
}

// SavePost represents the saving of the post model data.
func (p *Post) SavePost(db *gorm.DB) (*Post, error) {
	err := db.Debug().Create(&p).Error
	if err != nil {
		return &Post{}, err
	}

	if p.ID != 0 {
		err := db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}

	return p, nil
}

// FindAllPosts represents the finding of all post model data.
func (p *Post) FindAllPosts(db *gorm.DB, pagination *pagination.Pagination) (*[]Post, error) {
	var err error
	posts := []Post{}
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := db.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuilder.Model(&Post{}).Find(&posts)

	var totalRows int64
	db.Model(&Post{}).Count(&totalRows)
	pagination.TotalRows = totalRows

	err = result.Error
	if err != nil {
		return nil, err
	}

	if len(posts) > 0 {
		for i := range posts {
			err := db.Debug().Model(&User{}).Where("id = ?", posts[i].AuthorID).Take(&posts[i].Author).Error
			if err != nil {
				return nil, err
			}
		}
	}

	return &posts, nil
}

// FindPostByID represents the finding of the post model data by ID.
func (p *Post) FindPostByID(db *gorm.DB, pid uint64) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

// UpdatePost represents the updating of the post model data.
func (p *Post) UpdatePost(db *gorm.DB) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Where("id = ?", p.ID).Updates(Post{Title: p.Title, Content: p.Content, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Post{}, err
	}

	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}

	return p, nil
}

// DeletePost represents the deleting of the post model data.
func (p *Post) DeletePost(db *gorm.DB, pid uint64, uid uint32) (int64, error) {
	db = db.Debug().Model(&Post{}).Where("id = ? and author_id = ?", pid, uid).Take(&Post{}).Delete(&Post{})

	if db.Error != nil {
		if errors.Is(db.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("Post not found")
		}
		return 0, db.Error
	}

	return db.RowsAffected, nil
}
