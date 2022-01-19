package service

import (
	"context"
	"demo/params/article_param"
	"fmt"
)

// ArticleService 定义文章service接口，规定本service要提供的方法
type ArticleService interface {
	Create	(ctx context.Context, req *article_param.CreateReq) (*article_param.CreateResp, error)
	Detail (ctx context.Context, req *article_param.DetailReq) (*article_param.DetailResp, error)
}

// NewArticleService new service
func NewArticleService() ArticleService {
	var svc = &articleService{}
	{
		// middleware
	}
	return svc
}

// 定义文章service结构体，并实现文章service接口的所有方法
type articleService struct {}

func (s *articleService) Create	(ctx context.Context, req *article_param.CreateReq) (*article_param.CreateResp, error) {
	fmt.Printf("req:%#v\n", req)

	// mock：insert 根据传入参数req插库生成Id
	id := 1
	return &article_param.CreateResp{
		Id: int64(id),
	}, nil
}

func (s *articleService) Detail (ctx context.Context, req *article_param.DetailReq) (*article_param.DetailResp, error) {
	fmt.Printf("req:%#v\n", req)

	// mock：GetArticleById
	detail := struct {
		Id int64
		Title string
		Content string
		CateId int64
		UserId int64
	}{
		Id: 1,
		Title: "learn go",
		Content: "learn go content",
		CateId: 1,
		UserId: 1,
	}

	return &article_param.DetailResp{
		Id: detail.Id,
		Title: detail.Title,
		Content: detail.Content,
		CateId: detail.CateId,
		UserId: detail.UserId,
	}, nil
}