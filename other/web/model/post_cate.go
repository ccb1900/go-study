package model

type PostCate struct {
	Model
	Name  string
	Order int
}

func NewPostCate() *PostCate {
	return &PostCate{}
}
