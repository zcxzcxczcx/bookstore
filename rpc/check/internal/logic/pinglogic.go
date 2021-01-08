package logic

import (
	"context"

	"bookstore/rpc/check/check"
	"bookstore/rpc/check/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *check.Request) (*check.Response, error) {
	// todo: add your logic here and delete this line

	return &check.Response{}, nil
}

// func (l *CheckLogic) Check(in *check.CheckReq) (*check.CheckResp, error) {
//     // 手动代码开始
//     resp, err := l.svcCtx.Model.FindOne(in.Book)
//     if err != nil {
//         return nil, err
//     }

//     return &check.CheckResp{
//         Found: true,
//         Price: resp.Price,
//     }, nil
//     // 手动代码结束
// }
