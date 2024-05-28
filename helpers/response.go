package helpers

import "capstone/dto"

type ResponseWithData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithPaginate struct {
	Status     string   `json:"status"`
	Message    string   `json:"message"`
	Data       any      `json:"data"`
	Pagination Paginate `json:"pagination"`
}

type Paginate struct {
	Total       int64 `json:"total"`
	PerPage     int   `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	LastPage    int   `json:"last_page"`
}

type ResponseWithSort struct {
	Status     string   `json:"status"`
	Message    string   `json:"message"`
	Data       any      `json:"data"`
	Pagination Paginate `json:"pagination"`
	Sort       Sort     `json:"sort"`
}

type Sort struct {
	SortBy   string `json:"sort_by"`
	SortType string `json:"sort_type"`
}

type ResponseWithoutData struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func generateResponseWithData(status, message string, data any) ResponseWithData {
	return ResponseWithData{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func generateResponseWithPaginate(status, message string, data any, paginate Paginate) ResponseWithPaginate {
	return ResponseWithPaginate{
		Status:     status,
		Message:    message,
		Data:       data,
		Pagination: paginate,
	}
}

func generateResponseWithSort(status, message string, data any, paginate Paginate, sort Sort) ResponseWithSort {
	return ResponseWithSort{
		Status:     status,
		Message:    message,
		Data:       data,
		Pagination: paginate,
		Sort:       sort,
	}
}

func generateResponseWithoutData(status, message string) ResponseWithoutData {
	return ResponseWithoutData{
		Status:  status,
		Message: message,
	}
}

func Response(param dto.ResponseParams) any {
	var status string
	if param.StatusCode >= 200 && param.StatusCode < 300 {
		status = "Success"
	} else {
		status = "Failed"
	}

	if param.Data != nil {
		if param.IsPaginate {
			paginate := Paginate{
				Total:       param.Total,
				PerPage:     param.PerPage,
				CurrentPage: param.CurrentPage,
				LastPage:    param.LastPage,
			}
			if param.IsSort {
				sort := Sort{
					SortBy:   param.SortBy,
					SortType: param.SortType,
				}
				return generateResponseWithSort(status, param.Message, param.Data, paginate, sort)
			}
			return generateResponseWithPaginate(status, param.Message, param.Data, paginate)
		}
		return generateResponseWithData(status, param.Message, param.Data)
	}
	return generateResponseWithoutData(status, param.Message)
}
