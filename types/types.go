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
	Img_Url  string
	Username string
	Email    string
	Content  string
	Created  string
	Metric   string
}

type UserInfo struct {
	Bio           string
	Gender        string
	Phone         string
	RelStatus     string
	TopArtist     string
	LookingFor    string
	InstaUsername string
	Email         string
	Img_Url       string
}
type Poke struct {
	Username string
	Email    string
	Img_Url  string
}
