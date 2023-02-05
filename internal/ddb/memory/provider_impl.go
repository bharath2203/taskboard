package memory

import (
	"taskboard/internal/api"
	"taskboard/internal/ddb"
)

type provider struct {
}

func NewProvider() ddb.DbProvider {
	return &provider{}
}

func (p *provider) GetBoards(userId string) []*api.Board {
	return nil
}

func (p *provider) GetBoardById(id string) *api.Board {
	return nil
}

func (p *provider) GetTasks(listId string) []*api.Task {
	return nil
}

func (p *provider) GetLists(boardId string) []*api.List {
	return nil
}

func (p *provider) GetTaskById(id string) *api.Task {
	return nil
}
