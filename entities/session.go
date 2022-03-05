package entities

type Service struct{}

type Session struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	Duration  string `json:"duration"`
	Frequency string `json:"frequency"`
	Notes     []Note `json:"notes"`
	// Reminders []Reminder `json:"reminders"`
}

type Note struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type Reminder struct {
	StartTime string `json:"start_time"`
}
