package dto

type ResponseParams struct {
	StatusCode  int    `json:"status_code"`
	Message     string `json:"message"`
	Data        any    `json:"data"`
	IsPaginate  bool   `json:"is_paginate"`
	Total       int64  `json:"total"`
	PerPage     int    `json:"per_page"`
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
	IsSort      bool   `json:"is_sort"`
	Sort        string `json:"sort"`
}

type ResponseError struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Errors  any    `json:"errors"`
}
