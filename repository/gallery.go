package repository

import (
	"backend/contract"
	"backend/model"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type GalleryRepository struct {
	db *gorm.DB
}

func ImplGalleryRepository(db *gorm.DB) contract.GalleryRepository {
	return &GalleryRepository{
		db: db,
	}
}

func (g *GalleryRepository) GetAllGalleries(search, sort string) ([]model.Gallery, error) {
	var galleries []model.Gallery
	query := g.db.Model(&model.Gallery{})

	if search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if sort != "" {
		parts := strings.Split(sort, ":")
		if len(parts) == 2 {
			column := parts[0]
			order := strings.ToUpper(parts[1])
			if order == "ASC" || order == "DESC" {
				query = query.Order(fmt.Sprintf("%s %s", column, order))
			}
		}
	}

	if err := query.Find(&galleries).Error; err != nil {
		return nil, err
	}
	return galleries, nil
}

func (g *GalleryRepository) GetGalleryByID(id uint64) (*model.Gallery, error) {
	var gallery model.Gallery
	if err := g.db.First(&gallery, id).Error; err != nil {
		return nil, err
	}
	return &gallery, nil
}

func (g *GalleryRepository) CreateGallery(gallery *model.Gallery) (*model.Gallery, error) {
	if err := g.db.Create(gallery).Error; err != nil {
		return nil, err
	}
	return gallery, nil
}

func (g *GalleryRepository) UpdateGallery(id uint64, gallery *model.Gallery) error {
	err := g.db.Model(gallery).Where("id = ?", id).Updates(gallery).Error
	return err
}

func (g *GalleryRepository) DeleteGallery(id uint64) error {
	var gallery model.Gallery
	if err := g.db.Where("id = ?", id).Delete(&gallery).Error; err != nil {
		return err
	}
	return nil
}
