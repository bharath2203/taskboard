package ddb

import "taskboard/internal/api"

type DbProvider interface {
	GetBoards(userId string) []*api.Board
	GetBoardById(id string) *api.Board
	GetTasks(listId string) []*api.Task
	GetLists(boardId string) []*api.List
	GetTaskById(id string) *api.Task
}
