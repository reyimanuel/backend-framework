package repository

import (
	"backend/contract"
	"backend/model"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func ImplTeamRepository(db *gorm.DB) contract.TeamRepository {
	return &TeamRepository{
		db: db,
	}
}

func (t *TeamRepository) GetAllMember(search, status, sort string) ([]model.Team, error) {
	var teams []model.Team
	query := t.db.Model(&model.Team{})

	if search != "" {
		query = query.Where("name ILIKE ? OR role ILIKE ? OR division ILIKE ?", "%"+search+"%", "%"+search+"%", "%"+search+"%")
	}

	if status != "" {
		query = query.Where("status = ?", status)
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

	if err := query.Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}

func (t *TeamRepository) GetMemberByID(id uint64) (*model.Team, error) {
	var team model.Team
	if err := t.db.Where("id = ?", id).First(&team).Error; err != nil {
		return nil, err
	}
	return &team, nil
}

func (t *TeamRepository) CreateMember(team *model.Team) (*model.Team, error) {
	if err := t.db.Create(team).Error; err != nil {
		return nil, err
	}
	return team, nil
}

func (t *TeamRepository) UpdateMember(id uint64, team *model.Team) error {
	err := t.db.Model(team).Where("id = ?", id).Updates(team).Error
	return err
}

func (t *TeamRepository) DeleteMember(id uint64) error {
	var team model.Team
	if err := t.db.Where("id = ?", id).Delete(&team).Error; err != nil {
		return err
	}
	return nil
}
