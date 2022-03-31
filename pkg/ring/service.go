package ring

import (
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/config"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/http"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/json"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/logger"
)

type Service struct {
	log   logger.Interface
	rings []*config.Ring
}

func NewService(rings []*config.Ring, loggerFactory *logger.Factory) *Service {
	log := loggerFactory.GetLogger("ring.Service")
	return &Service{rings: rings, log: log}
}

func (s *Service) FindRings(req *http.Request) ([]string, error) {
	encoded, err := req.GetHeader("X-CharlesCD-User")
	if err != nil {
		return nil, err
	}
	j, err := json.FromBase64(encoded)
	if err != nil {
		return nil, err
	}
	email, err := j.GetString("email")
	if err != nil {
		return nil, err
	}
	s.log.Debugf("found email = %s", email)
	results := make([]string, 0)
	for _, ring := range s.rings {
		for _, rule := range ring.Match.Any {
			value, err := j.GetInt(rule.Key)
			if err != nil {
				return nil, err
			}
			s.log.Debugf("%s=%s", rule.Key, value)
		}
	}
	return results, nil
}
