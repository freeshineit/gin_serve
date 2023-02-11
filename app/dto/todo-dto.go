package dto

import "gin_serve/app/model"

type TodoDTO struct {
	Model
	ID      uint                    `json:"id" form:"id"`
	Status  *model.Todo_Status_Type `json:"status" form:"status"`
	Content string                  `json:"content" form:"content"`
	UserID  uint64                  `json:"user_id" form:"user_id"`
}

type TodoCreateDTO struct {
	// Status  *model.Todo_Status_Type `json:"status" form:"status" binding:""`
	Content string `json:"content" form:"content" binding:"required"`
}

// type TodoUpdateDTO struct {
// 	Status  model.Todo_Status_Type `json:"status" form:"status" binding:"required"`
// 	Content string                 `json:"content" form:"content" binding:"required"`
// }

type TodoUpdateStatusDTO struct {
	Status *model.Todo_Status_Type `json:"status" form:"status" binding:"required"`
}

type TodoUpdateContentDTO struct {
	Content string `json:"content" form:"content" binding:"required"`
}
