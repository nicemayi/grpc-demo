package handler

import (
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"server/calculator/internal/logic"
	"server/calculator/internal/svc"
	"server/calculator/internal/types"
)

func CalculatorHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCalculatorLogic(r.Context(), ctx)
		resp, err := l.Calculator(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
