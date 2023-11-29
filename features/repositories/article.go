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
	GetTrendingArticle(nameFilter string, page, pageSize int) ([]response.Article, int, error)
	GetArticle(id string) (response.Article, error)
	UpdateArticle(id string, input request.Article) (response.Article, error)
	DeleteArticle(id string) (response.Article, error)
	GetTotalLikes(SourceId uint, SourceStr string) (uint, error) 
	GetTotalLikesLastTwoWeeks(SourceId uint, SourceStr string) (uint, error)
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

func (ar *articleRepository) GetTrendingArticle(nameFilter string, page, pageSize int) ([]response.Article, int, error) {
    var allArticle []models.Article
    var resAllArticle []response.Article

    query := ar.db.Preload("Admin").Preload("Like")

    if nameFilter != "" {
        query = query.Where("title LIKE ?", "%"+nameFilter+"%")
    }

    offset := (page - 1) * pageSize

    query = query.Limit(pageSize).Offset(offset)
    
    twoWeeksAgo := time.Now().Add(-2 * 7 * 24 * time.Hour)
    err := query.Where("created_at >= ?", twoWeeksAgo).Find(&allArticle).Error
    if err != nil {
        return nil, 0, errors.ERR_GET_DATA
    }

    sort.Slice(allArticle, func(i, j int) bool {
        return allArticle[j].TotalLike > allArticle[i].TotalLike
    })

	for i := 0; i < len(allArticle); i++ {
		articleVm := domain.ConvertFromModelToArticleRes(allArticle[i])
	
		totalLikes, err := ar.GetTotalLikes(allArticle[i].ID, "article")
		if err != nil {
			return nil, 0, err
		}
	
		totalLikesLastTwoWeeks, err := ar.GetTotalLikesLastTwoWeeks(allArticle[i].ID, "article")
		if err != nil {
			return nil, 0, err
		}
	
		articleVm.TotalLike = totalLikes
		articleVm.TotalLike = totalLikesLastTwoWeeks

		if totalLikes == 0 {
			articleVm.Like = nil
		}

		var filteredLikes []response.Like
   		 for _, like := range articleVm.Like {
        if like.SourceStr == "article" {
            filteredLikes = append(filteredLikes, like)
        	}
    	}
   		articleVm.Like = filteredLikes
	
		resAllArticle = append(resAllArticle, *articleVm)
	}	

    var allItems int64
    query.Count(&allItems)

		// sort by like
		sort.Slice(resAllArticle, func(i, j int) bool {
			return resAllArticle[i].TotalLike > resAllArticle[j].TotalLike
		})

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

func (ar *articleRepository) GetTotalLikes(SourceId uint, SourceStr string) (uint, error) {
    var count int64
    if err := ar.db.Model(&models.Like{}).Where("source_id = ? AND source_str = ?", SourceId, SourceStr).Count(&count).Error; err != nil {
        return 0, err
    }
    return uint(count), nil
}

func (ar *articleRepository) GetTotalLikesLastTwoWeeks(SourceId uint, SourceStr string) (uint, error) {
    var count int64

    if err := ar.db.Model(&models.Like{}).
        Where("source_id = ? AND source_str = ?", SourceId, SourceStr).
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