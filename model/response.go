package model

import "github.com/SemmiDev/Go-Scheduled/entity"

type Result struct {
	ID          string                  `json:"id"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	IsDone      bool                    `json:"is_done"`
	IsOver      bool                    `json:"is_over"`
	Matrix      entity.EisenHowerMatrix `json:"matrix"`
	StartAt     string                  `json:"start_at"`
	Due         string                  `json:"due"`
	Duration    int64                   `json:"duration"`
	CreatedAt   string                  `json:"created_at"`
	UpdatedAt   string                  `json:"updated_at"`
}

type TimeManagementMatrix struct {
	Q1 []*Result `json:"q_1"`
	Q2 []*Result `json:"q_2"`
	Q3 []*Result `json:"q_3"`
	Q4 []*Result `json:"q_4"`
}

func MatrixFilter(assignments []*entity.Task) *TimeManagementMatrix {
	var matrix TimeManagementMatrix
	for _, v := range assignments {
		switch v.Matrix {
		case entity.UrgentAndImportant:
			matrix.Q1 = append(matrix.Q1, TaskToResult(v))
		case entity.ImportantButNotUrgent:
			matrix.Q2 = append(matrix.Q2, TaskToResult(v))
		case entity.UrgentButNotImportant:
			matrix.Q3 = append(matrix.Q3, TaskToResult(v))
		case entity.NotUrgentAndNotImportant:
			matrix.Q4 = append(matrix.Q4, TaskToResult(v))
		}
	}

	return &matrix
}
