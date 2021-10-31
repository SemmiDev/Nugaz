package model

import (
	"errors"
	"github.com/SemmiDev/Go-Scheduled/entity"
)

type CreateTaskReq struct {
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Matrix      entity.EisenHowerMatrix `json:"matrix"`
	StartAt     string                  `json:"start_at"`
	Due         string                  `json:"due"`
}

var (
	EmptyTitle       = errors.New("title is empty")
	EmptyDescription = errors.New("description is empty")
	EmptyMatrix      = errors.New("matrix is empty")
	EmptyDue         = errors.New("due is empty")
	EmptyStartAt     = errors.New("start_at is empty")
)

func (c *CreateTaskReq) Validate() error {
	if c.Title == "" {
		return EmptyTitle
	}
	if c.Description == "" {
		return EmptyDescription
	}
	if c.Matrix == "" {
		return EmptyMatrix

	}
	if c.StartAt == "" {
		return EmptyStartAt
	}
	if c.Due == "" {
		return EmptyDue
	}

	switch string(c.Matrix) {
	case "Urgent And Important":
		c.Matrix = entity.UrgentAndImportant
	case "Important But Not Urgent":
		c.Matrix = entity.ImportantButNotUrgent
	case "Urgent But Not Important":
		c.Matrix = entity.UrgentButNotImportant
	case "Not Urgent And Not Important":
		c.Matrix = entity.NotUrgentAndNotImportant
	default:
		return errors.New("matrix not recognize")
	}
	return nil
}

type UpdateTaskReq struct {
	ID          string
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Matrix      entity.EisenHowerMatrix `json:"matrix"`
	StartAt     string                  `json:"start_at"`
	Due         string                  `json:"due"`
}

func (c *UpdateTaskReq) Validate() error {
	if c.Title == "" {
		return EmptyTitle
	}
	if c.Description == "" {
		return EmptyDescription
	}
	if c.Matrix == "" {
		return EmptyMatrix

	}
	if c.StartAt == "" {
		return EmptyStartAt
	}
	if c.Due == "" {
		return EmptyDue
	}

	switch string(c.Matrix) {
	case "Urgent And Important":
		c.Matrix = entity.UrgentAndImportant
	case "Important But Not Urgent":
		c.Matrix = entity.ImportantButNotUrgent
	case "Urgent But Not Important":
		c.Matrix = entity.UrgentButNotImportant
	case "Not Urgent And Not Important":
		c.Matrix = entity.NotUrgentAndNotImportant
	default:
		return errors.New("matrix not recognize")
	}
	return nil
}
