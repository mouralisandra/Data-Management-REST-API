package document

type Document struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Content   string `json:"content"`
}
