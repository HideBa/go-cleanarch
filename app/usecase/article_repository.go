package usecase

import "github.com/HideBa/go-cleanarch/app/domain"

type ArticleRepository interface {
	Store(domain.Article) (int, error)
	FindById(id int) (domain.Article, error)
	FindByUser(id int) ([]domain.Article, error)
	FindAll() ([]domain.Article, error)
}
