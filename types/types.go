package types

type User struct {
	Username string
	Email    string
	Deanery  string
	Year     string
	Created  string
}

type Message struct {
	Alert string
}

type Post struct {
	Username string
	Email    string
	Content  string
	Created  string
	Metric   string
}
