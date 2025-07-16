package todo

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Status      string `json:"status"`
}

func GetDummyTodos() []Todo {
	return []Todo{
		{ID: 1, Title: "Learn Go", Status: "In Progress"},
		{ID: 2, Title: "Build a REST API", Status: "Completed"},
		{ID: 3, Title: "Write Tests", Status: "Pending"},
	}
}