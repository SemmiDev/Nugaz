package service

import (
	"context"
	"github.com/SemmiDev/Go-Scheduled/model"
	"github.com/SemmiDev/Go-Scheduled/repository"
)

type TaskService struct {
	repo repository.Querier
}

func NewTaskService(repo repository.Querier) *TaskService {
	return &TaskService{repo: repo}
}

func (as *TaskService) CreateTask(ctx context.Context, req *model.CreateTaskReq) (*model.Result, error) {
	arg, err := model.CreateTaskReqToTask(req)
	if err != nil {
		return nil, err
	}

	task, err := as.repo.Save(ctx, arg)
	if err != nil {
		return nil, err
	}

	response := model.TaskToResult(task)
	return response, nil
}

func (as *TaskService) GetTaskById(ctx context.Context, ID string) (*model.Result, error) {
	task, err := as.repo.FindById(ctx, ID)
	if err != nil {
		return nil, err
	}

	response := model.TaskToResult(task)
	return response, nil
}

func (as *TaskService) GetListTasks(ctx context.Context) (*model.TimeManagementMatrix, error) {
	listTasks, err := as.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	response := model.MatrixFilter(listTasks)
	return response, nil
}

func (as *TaskService) UpdateStatusTaskById(ctx context.Context, ID string) (*model.Result, error) {
	task, err := as.repo.UpdateIsDone(ctx, ID)
	if err != nil {
		return nil, err
	}

	response := model.TaskToResult(task)
	return response, nil
}

func (as *TaskService) UpdateTask(ctx context.Context, req *model.UpdateTaskReq) (*model.Result, error) {
	arg, err := model.UpdateTaskReqToTask(req)
	if err != nil {
		return nil, err
	}

	task, err := as.repo.Update(ctx, arg)
	if err != nil {
		return nil, err
	}

	response := model.TaskToResult(task)
	return response, nil
}

func (as *TaskService) DeleteTask(ctx context.Context, ID string) error {
	err := as.repo.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
