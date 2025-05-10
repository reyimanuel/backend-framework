package repository

import (
	"backend/contract"
	"backend/model"

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

func (t *TeamRepository) GetAllMember() ([]model.Team, error) {
	var teams []model.Team
	if err := t.db.Find(&teams).Error; err != nil {
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
	return t.db.Where("id = ?", id).Updates(&team).Error
}

func (t *TeamRepository) DeleteMember(id uint64) error {
	var team model.Team
	if err := t.db.Where("id = ?", id).Delete(&team).Error; err != nil {
		return err
	}
	return nil
}

func (t *TeamRepository) GetMemberByDivision(division string) ([]model.Team, error) {
	var teams []model.Team
	if err := t.db.Where("division = ?", division).Find(&teams).Error; err != nil {
		return nil, err
	}
	return teams, nil
}
