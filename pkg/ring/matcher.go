package ring

import (
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/config"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/json"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/log"
	"strings"
)

type matcher interface {
	MatchAny(j *json.Object) bool
	MatchAll(j *json.Object) bool
}

func matcherOfRule(rule *config.Rule) matcher {
	log.Debugf("getting a matcher for a rule operator %q", rule.Operator)
	switch rule.Operator {
	case "GreaterThan":
		return &greaterThanMatcher{key: rule.Key, values: rule.NumberValues()}
	}
	return new(noMatcher)
}

type noMatcher struct{}

func (g *noMatcher) MatchAny(j *json.Object) bool {
	return false
}

func (g *noMatcher) MatchAll(j *json.Object) bool {
	return false
}

type greaterThanMatcher struct {
	key    string
	values []float64
}

func (m *greaterThanMatcher) MatchAny(j *json.Object) bool {
	if !strings.HasPrefix(m.key, keyPrefix) {
		return false
	}
	key := strings.Replace(m.key, keyPrefix, "", 1)
	log.Debugf("searching key %q", key)
	jval, err := j.GetValue(key)
	if err != nil {
		log.Warnf("error getting value for key %q: %v", key, err)
		return false
	}
	log.Debugf("found %q=%q", key, jval)
	num, err := jval.AsFloat()
	if err != nil {
		log.Warnf("error unmarshalling rule value %q: %v", jval, err)
		return false
	}
	for _, mval := range m.values {
		if num > mval {
			log.Debugf(`matched operation {"reason":"%g > %g"}`, num, mval)
			return true
		} else {
			log.Debugf(`unmatched operation {"reason":"%g <= %g"}`, num, mval)
		}
	}
	return false
}

func (m *greaterThanMatcher) MatchAll(j *json.Object) bool {
	return false
}
