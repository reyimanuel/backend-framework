package service

import (
	"backend/contract"
	"backend/dto"
	"backend/model"
	"backend/pkg/errs"
	"backend/pkg/helpers"
	"fmt"
	"net/http"
)

type TeamService struct {
	TeamRepository contract.TeamRepository
}

func ImplTeamService(repo *contract.Repository) contract.TeamService {
	return &TeamService{
		TeamRepository: repo.TeamRepository,
	}
}

func (t *TeamService) GetAllMember(search, status, sort string) (*dto.TeamResponse, error) {
	teams, err := t.TeamRepository.GetAllMember(search, status, sort)
	if err != nil {
		return nil, errs.NotFound("No members found")
	}

	data := []dto.TeamData{}
	for _, team := range teams {
		data = append(data, dto.TeamData{
			ID:        team.ID,
			Name:      team.Name,
			Role:      team.Role,
			Division:  team.Division,
			Year:      team.Year,
			Status:    team.Status,
			CreatedAt: team.CreatedAt,
			UpdatedAt: team.UpdatedAt,
		})
	}

	response := &dto.TeamResponse{
		StatusCode: http.StatusOK,
		Message:    "Member data retrieved successfully",
		Data:       data,
	}

	return response, nil
}

func (t *TeamService) GetMemberByID(id uint64) (*dto.TeamResponse, error) {
	team, err := t.TeamRepository.GetMemberByID(id)
	if err != nil {
		return nil, errs.NotFound("Member not found")
	}

	data := dto.TeamData{
		ID:        team.ID,
		Name:      team.Name,
		Role:      team.Role,
		Division:  team.Division,
		Year:      team.Year,
		Status:    team.Status,
		CreatedAt: team.CreatedAt,
		UpdatedAt: team.UpdatedAt,
	}

	response := &dto.TeamResponse{
		StatusCode: http.StatusOK,
		Message:    "Member data retrieved successfully",
		Data:       []dto.TeamData{data},
	}

	return response, nil
}

func (t *TeamService) CreateMember(payload *dto.TeamRequest) (*dto.TeamResponse, error) {
	validPayload := helpers.ValidateStruct(payload)
	if validPayload != nil {
		return nil, validPayload
	}

	team := &model.Team{
		Name:     payload.Name,
		Role:     payload.Role,
		Division: payload.Division,
		Year:     payload.Year,
		Status:   payload.Status,
	}

	newTeam, err := t.TeamRepository.CreateMember(team)
	if err != nil {
		return nil, fmt.Errorf("failed to create member: %w", err)
	}

	response := &dto.TeamResponse{
		StatusCode: http.StatusCreated,
		Message:    "Member data created successfully",
		Data: []dto.TeamData{
			{
				ID:        newTeam.ID,
				Name:      newTeam.Name,
				Role:      newTeam.Role,
				Division:  newTeam.Division,
				Year:      newTeam.Year,
				Status:    newTeam.Status,
				CreatedAt: newTeam.CreatedAt,
				UpdatedAt: newTeam.UpdatedAt,
			},
		},
	}

	return response, nil
}

func (t *TeamService) UpdateMember(id uint64, payload *dto.TeamRequest) (*dto.TeamResponse, error) {
	validPayload := helpers.ValidateStruct(payload)
	if validPayload != nil {
		return nil, validPayload
	}

	_, err := t.TeamRepository.GetMemberByID(id)
	if err != nil {
		return nil, errs.NotFound("Member data not found")
	}

	team := &model.Team{
		Name:     payload.Name,
		Role:     payload.Role,
		Division: payload.Division,
		Year:     payload.Year,
		Status:   payload.Status,
	}

	err = t.TeamRepository.UpdateMember(id, team)
	if err != nil {
		return nil, fmt.Errorf("failed to update member: %w", err)
	}

	response := &dto.TeamResponse{
		StatusCode: http.StatusOK,
		Message:    "Member data updated successfully",
		Data: []dto.TeamData{
			{
				ID:        team.ID,
				Name:      team.Name,
				Role:      team.Role,
				Division:  team.Division,
				Year:      team.Year,
				Status:    team.Status,
				CreatedAt: team.CreatedAt,
			},
		},
	}

	return response, nil
}

func (t *TeamService) DeleteMember(id uint64) (*dto.TeamResponse, error) {
	err := t.TeamRepository.DeleteMember(id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete member: %w", err)
	}

	response := &dto.TeamResponse{
		StatusCode: http.StatusOK,
		Message:    "Member data deleted successfully",
	}

	return response, nil
}
