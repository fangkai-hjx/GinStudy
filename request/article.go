package request

type ArticleRequest struct {
	ID uint64 `form:"id"`
}
type ArticleListRequest struct {
	Page int `form:"page"`
}
