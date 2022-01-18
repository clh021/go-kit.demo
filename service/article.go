package service

import (
	"context"
	"demo/params/article_param"
)

type articleService struct {}

func (s *articleService) Create	(ctx context.Context, req *article_param.CreateResp) (*article_param.CreateResp, error) {
	// mock：insert 根据传入参数req插库生成Id
	id := 1
	return &article_param.CreateResp{
		Id: int64(id),
	}, nil
}

func (s *articleService) Detail (ctx context.Context, req *article_param.DetailReq) (*article_param.DetailResp, error) {
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