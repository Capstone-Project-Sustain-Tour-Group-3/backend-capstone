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
	Sort       string   `json:"sort"`
}

type ResponseWithFilter struct {
	Status     string   `json:"status"`
	Message    string   `json:"message"`
	Data       any      `json:"data"`
	Pagination Paginate `json:"pagination"`
	Sort       string   `json:"sort"`
	Filter     string   `json:"filter"`
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

func generateResponseWithSort(status, message string, data any, paginate Paginate, sort string) ResponseWithSort {
	return ResponseWithSort{
		Status:     status,
		Message:    message,
		Data:       data,
		Pagination: paginate,
		Sort:       sort,
	}
}

func generateResponseWithFilter(status, message string, data any, paginate Paginate, sort, filter string) ResponseWithFilter {
	return ResponseWithFilter{
		Status:     status,
		Message:    message,
		Data:       data,
		Pagination: paginate,
		Sort:       sort,
		Filter:     filter,
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

	if param.Data == nil {
		return generateResponseWithoutData(status, param.Message)
	}

	if param.IsPaginate {
		paginate := Paginate{
			Total:       param.Total,
			PerPage:     param.PerPage,
			CurrentPage: param.CurrentPage,
			LastPage:    param.LastPage,
		}
		return handlePaginateResponse(status, param, paginate)
	}

	return generateResponseWithData(status, param.Message, param.Data)
}

func handlePaginateResponse(status string, param dto.ResponseParams, paginate Paginate) any {
	if param.IsSort {
		return handleSortResponse(status, param, paginate)
	}
	return generateResponseWithPaginate(status, param.Message, param.Data, paginate)
}

func handleSortResponse(status string, param dto.ResponseParams, paginate Paginate) any {
	sort := param.Sort
	if param.IsFilter {
		filter := param.Filter
		return generateResponseWithFilter(status, param.Message, param.Data, paginate, sort, filter)
	}
	return generateResponseWithSort(status, param.Message, param.Data, paginate, sort)
}
