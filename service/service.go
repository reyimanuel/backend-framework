package service

import "backend/contract"

func New(repo *contract.Repository) *contract.Service {
	return &contract.Service{
		// Add your service methods here
		Auth:    ImplAuthService(repo),
		Team:    ImplTeamService(repo),
		Gallery: ImplGalleryService(repo),
		Event:   ImplEventService(repo),
	}
}
