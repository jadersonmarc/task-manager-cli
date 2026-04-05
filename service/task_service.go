package service

import (
	"errors"
	"strconv"
	"time"

	"github.com/jadersonmarc/task-manager-cli/storage"
	"github.com/jadersonmarc/task-manager-cli/task"
)

func Add(description string) error {
	tasks, err := storage.Load()
	if err != nil {
		return err
	}

	now := time.Now()

	newTask := task.Task{
		ID:          len(tasks) + 1,
		Description: description,
		Status:      task.StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	tasks = append(tasks, newTask)

	return storage.Save(tasks)
}

func UpdateStatus(idStr string, status string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.New("ID inválido")
	}

	tasks, err := storage.Load()
	if err != nil {
		return err
	}

	found := false

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return errors.New("tarefa não encontrada")
	}

	return storage.Save(tasks)
}

func Delete(idStr string) error {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return errors.New("ID inválido")
	}

	tasks, err := storage.Load()
	if err != nil {
		return err
	}

	var updated []task.Task
	found := false

	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		updated = append(updated, t)
	}

	if !found {
		return errors.New("tarefa não encontrada")
	}

	return storage.Save(updated)
}

func List(filter func(task.Task) bool) ([]task.Task, error) {
	tasks, err := storage.Load()
	if err != nil {
		return nil, err
	}

	var result []task.Task

	for _, t := range tasks {
		if filter(t) {
			result = append(result, t)
		}
	}

	return result, nil
}
