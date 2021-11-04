package model

import (
	"errors"
	"github.com/SemmiDev/Go-Scheduled/entity"
	"github.com/google/uuid"
	"time"
)

var ErrParseTime = errors.New("can't parse time request")

func CreateTaskReqToTask(t *CreateTaskReq) (*entity.Task, error) {
	r := new(entity.Task)

	// load location
	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowInLoc := time.Now().In(loc).Unix()

	// start at
	start, err := time.Parse(time.RFC3339, t.StartAt)
	if err != nil {
		return nil, ErrParseTime
	}
	startInLoc := start.In(loc).Unix()

	// ends at
	due, err := time.Parse(time.RFC3339, t.Due)
	if err != nil {
		return nil, ErrParseTime
	}
	dueInLoc := due.In(loc).Unix()

	// duration
	duration := (dueInLoc - startInLoc) / 60

	switch t.Matrix {
	case "Urgent And Important":
		r.Matrix = entity.UrgentAndImportant
	case "Important But Not Urgent":
		r.Matrix = entity.ImportantButNotUrgent
	case "Urgent But Not Important":
		r.Matrix = entity.UrgentButNotImportant
	case "Not Urgent And Not Important":
		r.Matrix = entity.NotUrgentAndNotImportant
	default:
		panic(errors.New("matrix not recognize"))
	}

	r.ID = uuid.New().String()
	r.Title = t.Title
	r.Description = t.Description
	r.IsDone = false
	r.IsOver = false
	r.Duration = duration
	r.StartAt = entity.Epoch(startInLoc)
	r.Due = entity.Epoch(dueInLoc)
	r.CreatedAt = entity.Epoch(nowInLoc)
	r.UpdatedAt = entity.Epoch(nowInLoc)

	return r, nil
}

func UpdateTaskReqToTask(t *UpdateTaskReq) (*entity.Task, error) {
	r := new(entity.Task)

	// load location
	loc, _ := time.LoadLocation("Asia/Jakarta")
	nowInLoc := time.Now().In(loc).Unix()

	// start at
	start, err := time.Parse(time.RFC3339, t.StartAt)
	if err != nil {
		return nil, ErrParseTime
	}
	startInLoc := start.In(loc).Unix()

	// ends at
	due, err := time.Parse(time.RFC3339, t.Due)
	if err != nil {
		return nil, ErrParseTime
	}
	dueInLoc := due.In(loc).Unix()

	// duration
	duration := (dueInLoc - startInLoc) / 60

	switch t.Matrix {
	case "Urgent And Important":
		r.Matrix = entity.UrgentAndImportant
	case "Important But Not Urgent":
		r.Matrix = entity.ImportantButNotUrgent
	case "Urgent But Not Important":
		r.Matrix = entity.UrgentButNotImportant
	case "Not Urgent And Not Important":
		r.Matrix = entity.NotUrgentAndNotImportant
	default:
		panic(errors.New("matrix not recognize"))
	}

	r.ID = t.ID
	r.Title = t.Title
	r.Description = t.Description
	r.StartAt = entity.Epoch(startInLoc)
	r.Due = entity.Epoch(dueInLoc)
	r.Duration = duration
	r.UpdatedAt = entity.Epoch(nowInLoc)

	return r, nil
}

func TaskToResult(t *entity.Task) *Result {
	r := new(Result)

	// load location
	loc, _ := time.LoadLocation("Asia/Jakarta")
	if int64(t.Due)-time.Now().In(loc).Unix() < 0 {
		r.IsOver = true
	}

	createdAt := time.Unix(int64(t.CreatedAt), 0).Format(time.RFC3339)
	updatedAt := time.Unix(int64(t.UpdatedAt), 0).Format(time.RFC3339)
	startAt := time.Unix(int64(t.StartAt), 0).Format(time.RFC3339)
	due := time.Unix(int64(t.Due), 0).Format(time.RFC3339)

	r.ID = t.ID
	r.Title = t.Title
	r.Description = t.Description
	r.Matrix = t.Matrix
	r.IsDone = t.IsDone
	r.StartAt = startAt
	r.Duration = t.Duration
	r.Due = due
	r.CreatedAt = createdAt
	r.UpdatedAt = updatedAt

	return r
}
