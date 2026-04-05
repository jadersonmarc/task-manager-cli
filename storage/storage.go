package storage

import (
	"encoding/json"
	"os"

	"github.com/jadersonmarc/task-manager-cli/task"
)

const fileName = "tasks.json"

func Load() ([]task.Task, error) {
	var tasks []task.Task

	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func Save(tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, data, 0644)
}
