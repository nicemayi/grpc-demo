package logic

import (
	"context"

	"go-server/rpc/calculator"
	"go-server/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddLogic) Add(in *calculator.AddRequest) (*calculator.AddResponse, error) {
	resp := &calculator.AddResponse{}
	resp.Results = in.A + in.B
	resp.Message = "Success!"
	return resp, nil
}
