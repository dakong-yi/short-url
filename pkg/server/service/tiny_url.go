package service

import (
	"context"
	"fmt"
	"short-url/pkg/connect"
	"short-url/pkg/grpc/model"
	"short-url/pkg/pb"
)

type tinyUrlService struct {
	tinyUrlModel *model.TinyUrl
}

func NewTinyUrlService() *tinyUrlService {
	return &tinyUrlService{
		tinyUrlModel: new(model.TinyUrl),
	}
}

func (s *tinyUrlService) GerOriginUrlByTinyUrl(ctx context.Context, tinyUrl string) (string, error) {
	conn := connect.GetClient()
	client := pb.NewTinyUrlClient(conn)
	url, err := client.GetOriginUrl(ctx, &pb.ShortUrl{
		TinyUrl: tinyUrl,
	})
	if err != nil {
		return "", err
	}
	return url.OriginUrl, err
}

func (s *tinyUrlService) GerTinyUrl(ctx context.Context, url string) (string, error) {
	conn := connect.GetClient()
	client := pb.NewTinyUrlClient(conn)
	fmt.Println(client)
	tinyUrl, err := client.GetTinyUrl(ctx, &pb.OriGinUrl{OriginUrl: url})
	if err != nil {
		return "", err
	}
	defer conn.Close()
	return tinyUrl.TinyUrl, err
}
