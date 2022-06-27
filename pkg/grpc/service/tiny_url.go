package service

import (
	"context"
	"log"
	"short-url/pkg/grpc/model"
	"short-url/pkg/pb"
	"short-url/pkg/util"
	"time"
)

type TinyUrlService struct {
	senderNum *urlSenderNum
}

func (s *TinyUrlService) GetTinyUrl(ctx context.Context, url *pb.OriGinUrl) (*pb.ShortUrl, error) {
	var err error
	SenderNumClient, err = NewUrlSenderNum(ctx)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	num := SenderNumClient.GetCurrNum(ctx)
	str := util.Encode(int(num))
	m := &model.TinyUrl{
		TinyUrl:     str,
		OriginalUrl: url.OriginUrl,
		ExpireTime:  time.Now().Add(time.Hour * 24 * 30),
	}
	err = m.Create(ctx)
	return &pb.ShortUrl{TinyUrl: str}, err
}

func (s *TinyUrlService) GetOriginUrl(ctx context.Context, url *pb.ShortUrl) (*pb.OriGinUrl, error) {
	m := &model.TinyUrl{}
	err := m.GetOriginUrlByTinyUrl(ctx, url.TinyUrl)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.OriGinUrl{OriginUrl: m.OriginalUrl}, err
}
