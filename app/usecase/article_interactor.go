package usecase

import "github.com/HideBa/go-cleanarch/app/domain"

type ArticleInteractor struct {
	ArticleRepository ArticleRepository
}

func (i *ArticleInteractor) Add(a domain.Article) (err error) {
	_, err = i.ArticleRepository.Store(a)
	return
}

func (i *ArticleInteractor) Articles() (articles []domain.Article, err error) {
	articles, err = i.ArticleRepository.FindAll()
	return
}

func (i *ArticleInteractor) UserArticles(id int) (articles []domain.Article, err error) {
	articles, err = i.ArticleRepository.FindByUser(id)
	return
}

func (i *ArticleInteractor) ArticleById(id int) (article domain.Article, err error) {
	article, err = i.ArticleRepository.FindById(id)
	return
}
