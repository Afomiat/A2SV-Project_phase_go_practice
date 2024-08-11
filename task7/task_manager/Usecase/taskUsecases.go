package Usecase

import (
	"errors"

	"task1.go/task7/task_manager/Domain"
)

type TaskUsecase struct {
	TaskRepo Domain.TaskRepository
}

func NewTaskUsecase(taskRepo Domain.TaskRepository) *TaskUsecase {
	return &TaskUsecase{TaskRepo: taskRepo}
}

func (uc *TaskUsecase) AddTask(task Domain.Task) error {
	if task.Title == "" || task.Status == "" {
		return errors.New("missing required fields")
	}
	_, err := uc.TaskRepo.AddTask(task)
	return err
}

func (uc *TaskUsecase) GetTasks(Role_ string, username string) ([]Domain.Task, error) {
	return uc.TaskRepo.GetTasks(Role_, username)
}

func (uc *TaskUsecase) GetTaskByID(id string, creter string, Role_ string) (Domain.Task, error) {
	return uc.TaskRepo.GetTaskByID(id, creter, Role_)
}


func (uc *TaskUsecase) DeleteTask(id string) error {
	_, err := uc.TaskRepo.DeleteTask(id)
	return err
}

func (uc *TaskUsecase) UpdateTask(id string, task Domain.Task) error {
	if task.Title == "" && task.Status == "" {
		return errors.New("missing required fields")
	}
	_, err := uc.TaskRepo.UpdateTask(id, task)
	return err
}