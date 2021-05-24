package service

import (
	"GinStudy/dao"
	"GinStudy/model"
	"fmt"
	"github.com/go-redis/redis"
)
import "GinStudy/cache"

//得到一篇文章的详情
func GetOneArticle(articleId uint64) (*model.Article, error) {
	//从 缓存 中获取 文章信息
	article, err := cache.GetOneArticleCache(articleId)
	//缓存中不存在
	if err == redis.Nil || err != nil {
		//从数据库中查询
		oneArticle, err := dao.SelectOneArticle(articleId)
		if err != nil {
			fmt.Println("dao.SelectOneArticle", err)
			return nil, err
		} else {
			err := cache.SetOneArticleCache(articleId, oneArticle)
			if err != nil {
				fmt.Println("cache.SetOneArticleCache", err)
				return nil, err
			} else {
				return oneArticle, err
			}
		}
	} else {
		return article, nil
	}
}
