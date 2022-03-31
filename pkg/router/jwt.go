package router

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/http"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/logger"
)

type RingService interface {
	FindRings(req http.Request) ([]string, error)
}

type JWT struct {
	types.DefaultHttpContext
	log logger.Interface
	svc RingService
}

func NewJWT(svc RingService, loggerFactory *logger.Factory) *JWT {
	log := loggerFactory.GetLogger("router.JWT")
	return &JWT{svc: svc, log: log}
}

func (h *JWT) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	req := http.NewRequest(numHeaders, endOfStream)
	rings, err := h.svc.FindRings(req)
	if err != nil {
		h.log.Criticalf(err, "error finding rings based on JWT claim")
		return types.ActionContinue
	}
	for _, ring := range rings {
		err = req.AddHeader("X-CharlesCD-Ring", ring)
		if err != nil {
			h.log.Criticalf(err, "error adding ring %q to request header", ring)
			return types.ActionContinue
		}
	}
	return types.ActionContinue
}

func (h *JWT) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	res := http.NewResponse(numHeaders, endOfStream)
	if err := res.AddHeader("hello", "kurtis"); err != nil {
		h.log.Criticalf(err, "error routing based on JWT claims")
		return types.ActionContinue
	}
	return types.ActionContinue
}
