package router

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/http"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/log"
)

type RingService interface {
	FindRings(req http.Request) ([]string, error)
}

type JWT struct {
	types.DefaultHttpContext
	contextID uint32
	svc       RingService
}

func NewJWT(svc RingService) *JWT {
	return &JWT{svc: svc}
}

func (h *JWT) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	log.Debug("request intercepted")
	req := http.NewRequest(numHeaders, endOfStream)
	rings, err := h.svc.FindRings(req)
	if err != nil {
		log.Criticalf(err, "error finding rings based on JWT claim")
		return types.ActionContinue
	}
	for _, ring := range rings {
		err = req.AddHeader("X-CharlesCD-Ring", ring)
		if err != nil {
			log.Criticalf(err, "error adding ring %q to request header", ring)
			return types.ActionContinue
		}
		log.Debugf("CharlesCD ring %q was added to request header", ring)
	}
	return types.ActionContinue
}

func (h *JWT) OnHttpResponseHeaders(numHeaders int, endOfStream bool) types.Action {
	res := http.NewResponse(numHeaders, endOfStream)
	if err := res.AddHeader("hello", "kurtis"); err != nil {
		log.Criticalf(err, "error routing based on JWT claims")
		return types.ActionContinue
	}
	return types.ActionContinue
}
