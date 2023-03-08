package models

type Comment struct {
	Id       int
	UserId   int
	PostId   int
	Likes    int
	Dislikes int
	Text     string
	Author   string
	Date     string
}
