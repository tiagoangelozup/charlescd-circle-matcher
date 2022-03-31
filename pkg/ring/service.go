package ring

import (
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/config"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/http"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/json"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/logger"
)

const keyPrefix = "request.auth.claims."

type ServiceLogger logger.Interface

type Service struct {
	log   ServiceLogger
	rings []*config.Ring
}

func NewService(log ServiceLogger, rings []*config.Ring) *Service {
	return &Service{log: log, rings: rings}
}

func (s *Service) FindRings(req http.Request) ([]string, error) {
	encoded, err := req.GetHeader("X-CharlesCD-User")
	if err != nil {
		return nil, err
	}
	j, err := json.FromBase64(encoded)
	if err != nil {
		return nil, err
	}
	results := make([]string, 0)
	for _, ring := range s.rings {
		for _, rule := range ring.Match.Any {
			if matcherOfRule(rule).Match(j) {
				results = append(results, ring.ID)
			}
		}
	}
	return results, nil
}
