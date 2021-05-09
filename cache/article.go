package cache

import (
	"GinStudy/global"
	"GinStudy/model"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

const ArticleDuration = time.Minute * 5

//redis的key为：article_3
func getArticleCacheName(articleId uint64) string {
	return "article_" + strconv.FormatUint(articleId, 10)
}

// 从缓存中得到一篇文章

func GetOneArticleCache(articleId uint64) (*model.Article, error) {
	key := getArticleCacheName(articleId)
	value, err := global.RedisDb.Get(global.Ctx, key).Result()
	// 该数据不存在与缓存中
	if err == redis.Nil || err != nil {
		return nil, err
	} else {
		article := model.Article{}
		//json--反序列化-->对象
		err := json.Unmarshal([]byte(value), &article)
		if err != nil {
			fmt.Println("json.Unmarshal[article] err：", err)
		}
		return &article, err
	}
}
func SetOneArticleCache(articleId uint64, article *model.Article) error {
	key := getArticleCacheName(articleId)
	valueJson, err := json.Marshal(article)
	if err != nil {
		fmt.Println("json.Marshal[article]", err)
		return err
	}
	err = global.RedisDb.Set(global.Ctx, key, valueJson, ArticleDuration).Err()
	if err != nil {
		fmt.Println("redis.Set[article]", err)
		return err
	}
	return nil
}
