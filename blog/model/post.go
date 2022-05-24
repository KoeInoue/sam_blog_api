package model

type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Content  string `json:"content"`
	Tag      string `json:"tag"`
	Locale   string `json:"locale"`
	PostedAt string `json:"postedAt"`
}

func NewPost() *Post {
	p := new(Post)
	return p
}
