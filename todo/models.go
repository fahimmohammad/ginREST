package todo

// Todo - struct
type Todo struct {
	ID        string `json:"id"  bson:"_id"`
	Title     string `json:"title" bson:"title"`
	Completed bool   `json:"completed" bson:"completed"`
}
