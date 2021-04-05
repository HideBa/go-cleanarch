package controllers

import (
	"strconv"

	"github.com/HideBa/go-cleanarch/app/domain"
	"github.com/HideBa/go-cleanarch/app/interfaces/database"
	"github.com/HideBa/go-cleanarch/app/usecase"
)

type ArticleController struct {
	Interactor usecase.ArticleInteractor
}

func NewArticleController(sqlHandler database.SqlHandler) *ArticleController {
	return &ArticleController{
		Interactor: usecase.ArticleInteractor{
			ArticleRepository: &database.ArticleRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *ArticleController) Create(c Context) {
	a := domain.Article{}
	c.Bind(&a)
	err := controller.Interactor.Add(a)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(201, nil)
}

func (controller *ArticleController) Index(c Context) {
	articles, err := controller.Interactor.Articles()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, articles)
}

func (controller *ArticleController) Show(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article, err := controller.Interactor.ArticleById(id)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, article)
}
