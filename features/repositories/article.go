package repositories

import (
	"lokasani/entity/domain"
	"lokasani/entity/models"
	"lokasani/entity/request"
	"lokasani/entity/response"
	"lokasani/helpers/errors"
	"sort"
	"time"

	"gorm.io/gorm"
)

type IArticleRepository interface {
	CreateArticle(data *request.Article) (response.Article, error)
	GetAllArticle(nameFilter string, page, pageSize int) ([]response.Article, int, error)
	GetArticle(id string) (response.Article, error)
	UpdateArticle(id string, input request.Article) (response.Article, error)
	DeleteArticle(id string) (response.Article, error)
	GetTotalLikes(SourceId uint) (uint, error)
	GetTotalLikesLastTwoWeeks(SourceId uint) (uint, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) *articleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) CreateArticle(data *request.Article) (response.Article, error) {
	dataArticle := domain.ConvertFromArticleReqToModel(*data)
	err := ar.db.Create(&dataArticle).Error
	if err != nil {
		return response.Article{}, err
	}
	err = ar.db.Preload("Admin").First(&dataArticle, "id = ? ", dataArticle.ID).Error
	return *domain.ConvertFromModelToArticleRes(*dataArticle), nil
}

func (ar *articleRepository) GetAllArticle(nameFilter string, page, pageSize int) ([]response.Article, int, error) {
	var allArticle []models.Article
	var resAllArticle []response.Article

	query := ar.db.Preload("Admin").Preload("Like")

	if nameFilter != "" {
		query = query.Where("title LIKE ?", "%"+nameFilter+"%")
	}

	offset := (page - 1) * pageSize

	query = query.Limit(pageSize).Offset(offset)
	err := query.Find(&allArticle).Error
	if err != nil {
		return nil, 0, errors.ERR_GET_DATA
	}

	sort.Slice(allArticle, func(i, j int) bool {
		return allArticle[j].TotalLike > allArticle[i].TotalLike
	})

	for i := 0; i < len(allArticle); i++ {
		articleVm := domain.ConvertFromModelToArticleRes(allArticle[i])

		totalLikes, err := ar.GetTotalLikes(allArticle[i].ID)
		if err != nil {
			return nil, 0, err
		}

		totalLikesLastTwoWeeks, err := ar.GetTotalLikesLastTwoWeeks(allArticle[i].ID)
		if err != nil {
			return nil, 0, err
		}

		articleVm.TotalLike = totalLikes
		articleVm.TotalLike = totalLikesLastTwoWeeks

		resAllArticle = append(resAllArticle, *articleVm)
	}

	var allItems int64
	query.Count(&allItems)

	return resAllArticle, int(allItems), nil
}

func (ar *articleRepository) GetArticle(id string) (response.Article, error) {
	var articleData models.Article
	err := ar.db.Preload("Admin").Preload("Like").First(&articleData, "id = ?", id).Error
	if err != nil {
		return response.Article{}, err
	}
	return *domain.ConvertFromModelToArticleRes(articleData), nil
}

func (ar *articleRepository) GetTotalLikes(SourceId uint) (uint, error) {
	var count int64
	if err := ar.db.Model(&models.Like{}).Where("source_id = ?", SourceId).Count(&count).Error; err != nil {
		return 0, err
	}
	return uint(count), nil
}

func (ar *articleRepository) GetTotalLikesLastTwoWeeks(SourceId uint) (uint, error) {
	var count int64
	twoWeeksAgo := time.Now().Add(-2 * 7 * 24 * time.Hour)

	if err := ar.db.Model(&models.Like{}).
		Where("source_id = ?", SourceId).
		Where("created_at > ?", twoWeeksAgo).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return uint(count), nil
}

func (ar *articleRepository) UpdateArticle(id string, input request.Article) (response.Article, error) {
	articleData := models.Article{}
	err := ar.db.Preload("Admin").First(&articleData, "id = ?", id).Error
	if err != nil {
		return response.Article{}, err
	}

	if input.Title != "" {
		articleData.Title = input.Title
	}
	if input.Content != "" {
		articleData.Content = input.Content
	}
	if input.FilesId != nil {
		articleData.FilesId = input.FilesId
	}
	if input.Status != "" {
		articleData.Status = input.Status
	}

	if err = ar.db.Save(&articleData).Error; err != nil {
		return response.Article{}, err
	}
	return *domain.ConvertFromModelToArticleRes(articleData), nil
}

func (ar *articleRepository) DeleteArticle(id string) (response.Article, error) {
	articleData := models.Article{}
	res := response.Article{}
	find := ar.db.Preload("Admin").First(&articleData, "id = ?", id).Error
	if find == nil {
		res = *domain.ConvertFromModelToArticleRes(articleData)
	}
	err := ar.db.Delete(&articleData, "id = ?", id).Error
	if err != nil {
		return response.Article{}, err
	}
	return res, nil
}
