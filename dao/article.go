package dao

import (
	"GinStudy/global"
	"GinStudy/model"
)

func SelectOneArticle(articleId uint64) (*model.Article, error) {
	fields := []string{"articleId", "subject", "url"}
	article := model.Article{}
	err := global.DB.Select(fields).Where("articleId=?", articleId).First(&article).Error
	if err != nil {
		return nil, err
	} else {
		return &article, nil
	}
}
