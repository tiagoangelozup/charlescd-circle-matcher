package ring

import (
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/config"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/json"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/log"
	"strings"
)

type matcher interface {
	Match(j *json.Object) bool
}

func matcherOfRule(rule *config.Rule) matcher {
	log.Debugf("getting a matcher for a rule operator %q", rule.Operator)
	switch rule.Operator {
	case "GreaterThan":
		return newGreaterThanMatcher(rule.Key, rule.Values)
	}
	return new(noMatcher)
}

type noMatcher struct{}

func (g *noMatcher) Match(j *json.Object) bool {
	return false
}

type greaterThanMatcher struct {
	key    string
	values []float64
}

func newGreaterThanMatcher(key string, values []interface{}) *greaterThanMatcher {
	v := make([]float64, 0)
	for _, value := range values {
		switch t := value.(type) {
		case float64:
			v = append(v, t)
		}
	}
	return &greaterThanMatcher{key: key, values: v}
}

func (m *greaterThanMatcher) Match(j *json.Object) bool {
	if !strings.HasPrefix(m.key, keyPrefix) {
		return false
	}
	key := strings.Replace(m.key, keyPrefix, "", 1)
	log.Debugf("searching key %q", key)
	jval, err := j.GetValue(key)
	if err != nil {
		return false
	}
	num, err := jval.AsFloat()
	if err != nil {
		// TODO: log warning
		return false
	}
	result := len(m.values) > 0
	for _, mval := range m.values {
		if num <= mval {
			return false
		}
	}
	return result
}
