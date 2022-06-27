package service

import (
	"context"
	"errors"
	"log"
	"short-url/pkg/grpc/model"
	"sync"

	"gorm.io/gorm"
)

var SenderNumClient *urlSenderNum

type urlSenderNum struct {
	StartNum int64
	EndNum   int64
	CurrNum  int64
	Next     *urlSenderNum
	Lock     *sync.Mutex
}

func NewUrlSenderNum(ctx context.Context) (*urlSenderNum, error) {
	if SenderNumClient != nil {
		return SenderNumClient, nil
	}
	m := &model.UrlSenderNum{}
	u := &urlSenderNum{Lock: &sync.Mutex{}}
	err := m.GetMaxNum(ctx)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		m.StartNum = 1
		m.EndNum = 1000
		m.Version = 1
		err = m.Create(ctx)
		u.StartNum = m.StartNum
		u.EndNum = m.EndNum
		u.CurrNum = m.StartNum
		return u, err
	}
	if err != nil {
		log.Println(err)
		return u, err
	}
	u.StartNum = m.EndNum + 1
	u.CurrNum = u.StartNum
	u.EndNum = u.StartNum + 999
	version := m.Version
	m.Version++
	m.StartNum = u.StartNum
	m.EndNum = u.EndNum
	err = m.UpdateByVersion(ctx, version)
	return u, err
}

func (u *urlSenderNum) GetCurrNum(ctx context.Context) int64 {
	u.Lock.Lock()
	defer u.Lock.Unlock()
	num := u.CurrNum
	u.IncrCurrNum(ctx)
	return num
}

func (u *urlSenderNum) IncrCurrNum(ctx context.Context) {
	if u.CurrNum+1 >= u.EndNum {
		if u.Next != nil {
			u = u.Next
		} else {
			u, _ = NewUrlSenderNum(ctx)
		}
	} else {
		u.CurrNum++
	}

	if (u.CurrNum-u.StartNum)/(u.EndNum-u.CurrNum)*10 > 8 {
		go func() {
			next, _ := NewUrlSenderNum(ctx)
			u.SetNext(ctx, next)
		}()
	}
}
func (u *urlSenderNum) SetNext(ctx context.Context, next *urlSenderNum) {
	u.Next = next
}
