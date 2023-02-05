package api

type Board struct {
	Id    string  `json:"id"`
	Lists []*List `json:"lists"`
}

type List struct {
	Id    string  `json:"id"`
	Tasks []*Task `json:"tasks"`
}

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
