package ring

import (
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/config"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/http"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/json"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/log"
)

const keyPrefix = "request.auth.claims."

type Service struct {
	rings config.Rings
}

func NewService(rings config.Rings) *Service {
	return &Service{rings: rings}
}

func (s *Service) FindRings(req http.Request) ([]string, error) {
	encoded, ok, err := req.GetHeader("X-CharlesCD-User")
	if err != nil {
		return nil, err
	}
	if !ok {
		log.Debugf("user not found")
		return nil, nil
	}
	j, err := json.FromBase64(encoded)
	if err != nil {
		return nil, err
	}
	log.Debugf("checking %d ring(s) rules against user %q", len(s.rings), j)
	results := make([]string, 0)
	for _, ring := range s.rings {
		for _, rule := range ring.Match.Any {
			if matcherOfRule(rule).MatchAny(j) {
				results = append(results, ring.ID)
			}
		}
	}
	return results, nil
}
