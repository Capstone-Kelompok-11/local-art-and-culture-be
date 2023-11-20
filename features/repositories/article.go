package repositories

import (
	"fmt"
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"

	"gorm.io/gorm"
)

type IArticleRepository interface {
	CreateArticle(data *request.Article) (response.Article, error)
	GetAllArticle() ([]response.Article, error)
	GetArticle(id string) (response.Article, error)
	UpdateArticle(id string, input request.Article) (response.Article, error)
	DeleteArticle(id string) (response.Article, error)
	//GetAdminWithArticles(adminID uint) (*response.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) CreateArticle(data *request.Article) (response.Article, error){
	dataArticle := domain.ConvertFromArticleReqToModel(*data)
	err := ar.db.Create(&dataArticle).Error
	if err != nil {
		return response.Article{}, err
	}
	return *domain.ConvertFromModelToArticleRes(*dataArticle), nil
}

func (ar *articleRepository) GetAllArticle() ([]response.Article, error) {
	var allArticle []models.Article
	var resAllArticle []response.Article
	err := ar.db.Preload("SuperAdmin").Find(&allArticle).Error
	if err != nil {
		return nil, errors.ERR_GET_DATA
	}

	fmt.Println(allArticle)
	for i := 0; i < len(allArticle); i++ {
		articleVm := domain.ConvertFromModelToArticleRes(allArticle[i])
		resAllArticle = append(resAllArticle, *articleVm)
	}
	return resAllArticle, nil
}

func (ar *articleRepository) GetArticle(id string) (response.Article, error) {
	var articleData models.Article
	err := ar.db.Preload("SuperAdmin").First(&articleData, "id = ?", id).Error
	if err != nil {
		return response.Article{}, err
	}
	return *domain.ConvertFromModelToArticleRes(articleData), nil
}

func (ar *articleRepository) UpdateArticle(id string, input request.Article) (response.Article, error) {
	articleData := models.Article{}
	err := ar.db.First(&articleData, "id = ?", id).Error
	if err != nil {
		return response.Article{}, err
	}

	if input.Title != "" {
		articleData.Title = input.Title
	} else if input.Content != "" {
		articleData.Content = input.Content
	}

	if err = ar.db.Save(&articleData).Error; err != nil {
		return response.Article{}, err
	}
	return *domain.ConvertFromModelToArticleRes(articleData), nil
}

func (ar *articleRepository) DeleteArticle(id string) (response.Article, error) {
	articleData := models.Article{}
	res := response.Article{}
	find := ar.db.First(&articleData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToArticleRes(articleData)
	}
	err := ar.db.Delete(&articleData, "id = ?", id).Error
	if err != nil {
		return response.Article{}, err
	}
	return res, nil
}

// func (ar *articleRepository) GetAdminWithArticles(adminID uint) (*response.Article, error) {
//     var admin models.SuperAdmin
//     //var articles []models.Article

//     if err := ar.db.Preload("Article").First(&admin, admin.ID).Error; err != nil {
//         return nil, err
//     }

//     adminRes := domain.ConvertFromModelToAdminRes(admin)
//     articlesRes := make([]response.Article, len(admin.Articles))
//     for i, article := range admin.Articles {
//         articlesRes[i] = *domain.ConvertFromModelToArticleRes(article)
//     }

//     return &response.Article{
//         Id: adminRes.Id,
//     }, nil
// }
