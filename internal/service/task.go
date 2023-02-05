package service

import "taskboard/internal/api"

func FetchTasks() []*api.Task {
	tasks := []*api.Task{
		{
			Title: "task one",
		},
		{
			Title: "task two",
		},
	}
	return tasks
}
