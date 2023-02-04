package dto

import "gin_serve/app/model"

type TodoCreateDTO struct {
	Status  model.Todo_Status_Type `json:"status" form:"status" binding:""`
	Content string                 `json:"content" form:"content" binding:"required"`
}

type TodoUpdateDTO struct {
	Status  model.Todo_Status_Type `json:"status" form:"status" binding:""`
	Content string                 `json:"content" form:"content" binding:"required"`
}
