package domain

type Article struct {
	ID      int
	Title   string
	Content string
	Author  User
}
