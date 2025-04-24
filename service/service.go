package service

import "backend/contract"

func New(repo *contract.Repository) *contract.Service {
	return &contract.Service{
		Auth: ImplAuthService(repo),
		// Add your service methods here

	}
}
