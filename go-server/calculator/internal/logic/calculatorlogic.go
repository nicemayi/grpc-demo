package logic

import (
	"context"

	"server/calculator/internal/svc"
	"server/calculator/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type CalculatorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCalculatorLogic(ctx context.Context, svcCtx *svc.ServiceContext) CalculatorLogic {
	return CalculatorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CalculatorLogic) Calculator(req types.Request) (*types.Response, error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
