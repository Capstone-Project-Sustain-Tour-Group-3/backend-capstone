package tests

import (
	"capstone/dto"
	"capstone/errorHandlers"
	"capstone/mocks/repositories"
	"capstone/usecases"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetAllData(t *testing.T) {
	mockRepo := new(repositories.MockDashboardRepository)
	uc := usecases.NewDashboardUsecase(mockRepo)

	testCases := []struct {
		name           string
		mockSetup      func()
		expectedResult dto.DashboardResponse
		expectedError  error
	}{
		{
			name: "Success get all data",
			mockSetup: func() {
				mockRepo.On("GetTotalAdmin").Return(int64(5), nil).Once()
				mockRepo.On("GetTotalUser").Return(int64(50), nil).Once()
				mockRepo.On("GetTotalRoute").Return(int64(10), nil).Once()
				mockRepo.On("GetTotalDestination").Return(int64(20), nil).Once()
				mockRepo.On("GetTotalVidContent").Return(int64(15), nil).Once()
				mockRepo.On("GetTotalDestinationWithCategory").Return([]dto.TotalDestinationationWithCategory{{Name: "Alam", Total: 10}}, nil).Once()
				mockRepo.On("GetMonthlyUsers").Return([]dto.MonthlyUser{{Bulan: "Jun", TotalPengguna: 5}}, nil).Once()
				mockRepo.On("GetRouteUser").Return([]dto.RouteUser{{NamaRute: "Route 1", JumlahDestinasi: 3}}, nil).Once()
			},
			expectedResult: dto.DashboardResponse{
				TotalData: dto.TotalData{
					TotalAdmin: 5,
					TotalUser:  50,
					TotalRute:  10,
				},
				PertumbuhanPengguna: []dto.MonthlyUser{{Bulan: "Jun", TotalPengguna: 5}},
				TotalKontenVideo: dto.TotalContentVid{
					TotalContent:   15,
					TotalDestinasi: 20,
				},
				KategoriDestinasi: []dto.TotalDestinationationWithCategory{{Name: "Alam", Total: 10}},
				RuteUser:          []dto.RouteUser{{NamaRute: "Route 1", JumlahDestinasi: 3}},
			},
			expectedError: nil,
		},
		{
			name: "Error on GetTotalAdmin",
			mockSetup: func() {
				mockRepo.On("GetTotalAdmin").Return(int64(0), &errorHandlers.InternalServerError{Message: "DB error"}).Once()
			},
			expectedResult: dto.DashboardResponse{},
			expectedError:  &errorHandlers.InternalServerError{Message: "DB error"},
		},
		{
			name: "Error on GetTotalUser",
			mockSetup: func() {
				mockRepo.On("GetTotalAdmin").Return(int64(5), nil).Once()
				mockRepo.On("GetTotalUser").Return(int64(0), &errorHandlers.InternalServerError{Message: "DB error"}).Once()
			},
			expectedResult: dto.DashboardResponse{},
			expectedError:  &errorHandlers.InternalServerError{Message: "DB error"},
		},
		{
			name: "Error on GetTotalRoute",
			mockSetup: func() {
				mockRepo.On("GetTotalAdmin").Return(int64(5), nil).Once()
				mockRepo.On("GetTotalUser").Return(int64(50), nil).Once()
				mockRepo.On("GetTotalRoute").Return(int64(0), &errorHandlers.InternalServerError{Message: "DB error"}).Once()
			},
			expectedResult: dto.DashboardResponse{},
			expectedError:  &errorHandlers.InternalServerError{Message: "DB error"},
		},
		{
			name: "Error on GetTotalDestination",
			mockSetup: func() {
				mockRepo.On("GetTotalAdmin").Return(int64(5), nil).Once()
				mockRepo.On("GetTotalUser").Return(int64(50), nil).Once()
				mockRepo.On("GetTotalRoute").Return(int64(10), nil).Once()
				mockRepo.On("GetTotalDestination").Return(int64(0), &errorHandlers.InternalServerError{Message: "DB error"}).Once()
			},
			expectedResult: dto.DashboardResponse{},
			expectedError:  &errorHandlers.InternalServerError{Message: "DB error"},
		},
		{
			name: "Error on GetTotalVidContent",
			mockSetup: func() {
				mockRepo.On("GetTotalAdmin").Return(int64(5), nil).Once()
				mockRepo.On("GetTotalUser").Return(int64(50), nil).Once()
				mockRepo.On("GetTotalRoute").Return(int64(10), nil).Once()
				mockRepo.On("GetTotalDestination").Return(int64(20), nil).Once()
				mockRepo.On("GetTotalVidContent").Return(int64(0), &errorHandlers.InternalServerError{Message: "DB error"}).Once()
			},
			expectedResult: dto.DashboardResponse{},
			expectedError:  &errorHandlers.InternalServerError{Message: "DB error"},
		},
		{
			name: "Error on GetTotalDestinationWithCategory",
			mockSetup: func() {
				mockRepo.On("GetTotalAdmin").Return(int64(5), nil).Once()
				mockRepo.On("GetTotalUser").Return(int64(50), nil).Once()
				mockRepo.On("GetTotalRoute").Return(int64(10), nil).Once()
				mockRepo.On("GetTotalDestination").Return(int64(20), nil).Once()
				mockRepo.On("GetTotalVidContent").Return(int64(15), nil).Once()
				mockRepo.On("GetTotalDestinationWithCategory").Return(nil, &errorHandlers.InternalServerError{Message: "DB error"}).Once()
			},
			expectedResult: dto.DashboardResponse{},
			expectedError:  &errorHandlers.InternalServerError{Message: "DB error"},
		},
		{
			name: "Error on GetMonthlyUsers",
			mockSetup: func() {
				mockRepo.On("GetTotalAdmin").Return(int64(5), nil).Once()
				mockRepo.On("GetTotalUser").Return(int64(50), nil).Once()
				mockRepo.On("GetTotalRoute").Return(int64(10), nil).Once()
				mockRepo.On("GetTotalDestination").Return(int64(20), nil).Once()
				mockRepo.On("GetTotalVidContent").Return(int64(15), nil).Once()
				mockRepo.On("GetTotalDestinationWithCategory").Return([]dto.TotalDestinationationWithCategory{{Name: "Alam", Total: 10}}, nil).Once()
				mockRepo.On("GetMonthlyUsers").Return(nil, &errorHandlers.InternalServerError{Message: "DB error"}).Once()
			},
			expectedResult: dto.DashboardResponse{},
			expectedError:  &errorHandlers.InternalServerError{Message: "DB error"},
		},
		{
			name: "Error on GetRouteUser",
			mockSetup: func() {
				mockRepo.On("GetTotalAdmin").Return(int64(5), nil).Once()
				mockRepo.On("GetTotalUser").Return(int64(50), nil).Once()
				mockRepo.On("GetTotalRoute").Return(int64(10), nil).Once()
				mockRepo.On("GetTotalDestination").Return(int64(20), nil).Once()
				mockRepo.On("GetTotalVidContent").Return(int64(15), nil).Once()
				mockRepo.On("GetTotalDestinationWithCategory").Return([]dto.TotalDestinationationWithCategory{{Name: "Alam", Total: 10}}, nil).Once()
				mockRepo.On("GetMonthlyUsers").Return([]dto.MonthlyUser{{Bulan: "Jun", TotalPengguna: 5}}, nil).Once()
				mockRepo.On("GetRouteUser").Return(nil, &errorHandlers.InternalServerError{Message: "DB error"}).Once()
			},
			expectedResult: dto.DashboardResponse{},
			expectedError:  &errorHandlers.InternalServerError{Message: "DB error"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockSetup()

			result, err := uc.GetAllData()

			if tc.expectedError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedResult, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
