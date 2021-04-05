package database

import "github.com/HideBa/go-cleanarch/app/domain"

type ArticleRepository struct {
	SqlHandler
}

func (r *ArticleRepository) Store(a domain.Article) (int, error) {
	result, err := r.Execute("INSERT INTO articles (title, content, author_id) VALUES (?,?,?)", a.Title, a.Content, a.Author.ID)
	if err != nil {
		return 0, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	id := int(id64)
	return id, nil
}

func (r *ArticleRepository) FindById(identifier int) (domain.Article, error) {
	article := domain.Article{}
	author := domain.User{}
	row, err := r.Query("SELECT id, title, content, users.id, users.first_name, users.last_name from articles INNER JOIN users ON articles.author_id = users.id WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return article, err
	}
	var id int
	var title string
	var content string
	var authorId int
	var authorFirstName string
	var authorLastName string
	row.Next()
	if err = row.Scan(&id, &title, &content, &authorId, &authorFirstName, &authorLastName); err != nil {
		return domain.Article{}, err
	}
	article.ID = id
	article.Title = title
	article.Content = content
	article.Author = author
	article.Author.ID = authorId
	article.Author.FirstName = authorFirstName
	article.Author.LastName = authorLastName
	return article, nil
}

func (r *ArticleRepository) FindByUser(userId int) (articles []domain.Article, err error) {
	rows, err := r.Query("SELECT id, title, content, users.id, users.first_name, users.last_name from articles INNER JOIN users ON articles.author_id = users.id WHERE users.id = ?", userId)
	defer rows.Close()
	if err != nil {
		return articles, err
	}
	for rows.Next() {
		var id int
		var title string
		var content string
		var author domain.User
		if err := rows.Scan(&id, &title, &content, &author.ID, &author.FirstName, &author.LastName); err != nil {
			continue
		}
		article := domain.Article{
			ID:      id,
			Title:   title,
			Content: content,
			Author:  author,
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (r *ArticleRepository) FindAll() (articles []domain.Article, err error) {
	rows, err := r.Query("SELECT id, title, content, author_id, users.id, users.first_name, users.last_name from articles INNER JOIN users ON articles.author_id = users.id")
	defer rows.Close()
	if err != nil {
		return articles, err
	}
	for rows.Next() {
		var id int
		var title string
		var content string
		var author domain.User
		if err := rows.Scan(&id, &title, &content, &author.ID, &author.FirstName, &author.LastName); err != nil {
			continue
		}
		article := domain.Article{
			ID:      id,
			Title:   title,
			Content: content,
			Author:  author,
		}
		articles = append(articles, article)
	}
	return articles, nil
}
