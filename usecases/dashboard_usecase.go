package usecases

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/repositories"
)

type DashboardUsecase interface {
	GetAllData() (dto.DashboardResponse, error)
}

type dashboardUsecase struct {
	repository repositories.DashboardRepository
}

func NewDashboardUsecase(repository repositories.DashboardRepository) *dashboardUsecase {
	return &dashboardUsecase{repository: repository}
}

func (uc *dashboardUsecase) GetAllData() (dto.DashboardResponse, error) {
	totalAdmin, err := uc.repository.GetTotalAdmin()
	if err != nil {
		return dto.DashboardResponse{}, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	totalUser, err := uc.repository.GetTotalUser()
	if err != nil {
		return dto.DashboardResponse{}, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	totalRoute, err := uc.repository.GetTotalRoute()
	if err != nil {
		return dto.DashboardResponse{}, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	totalDestination, err := uc.repository.GetTotalDestination()
	if err != nil {
		return dto.DashboardResponse{}, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	totalVidContent, err := uc.repository.GetTotalVidContent()
	if err != nil {
		return dto.DashboardResponse{}, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	totalDestinationWithCategory, err := uc.repository.GetTotalDestinationWithCategory()
	if err != nil {
		return dto.DashboardResponse{}, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	monthlyUsers, err := uc.repository.GetMonthlyUsers()
	if err != nil {
		return dto.DashboardResponse{}, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	routeUser, err := uc.repository.GetRouteUser()
	if err != nil {
		return dto.DashboardResponse{}, &errorHandlers.InternalServerError{Message: err.Error()}
	}

	response := dto.DashboardResponse{
		TotalData: dto.TotalData{
			TotalAdmin: totalAdmin,
			TotalUser:  totalUser,
			TotalRute:  totalRoute,
		},
		PertumbuhanPengguna: monthlyUsers,
		TotalKontenVideo: dto.TotalContentVid{
			TotalContent:   totalVidContent,
			TotalDestinasi: totalDestination,
		},
		KategoriDestinasi: totalDestinationWithCategory,
		RuteUser:          routeUser,
	}

	return response, nil
}
