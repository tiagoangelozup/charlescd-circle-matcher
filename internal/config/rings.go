package config

import (
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/log"
)

const pathNotFoundErrMsg = "Key path not found"

type (
	Rings []*Ring
	Ring  struct {
		ID    string
		Match Match
	}
	Match struct {
		All []*Rule
		Any []*Rule
	}
	Rule struct {
		Key      string
		Operator string
		Values   []*RuleValue
	}
	RuleValue struct {
		value    []byte
		dataType jsonparser.ValueType
	}
	builder struct {
		rings Rings
	}
	converter func([]byte, jsonparser.ValueType, int, error)
)

func (r *RuleValue) IsNumber() (float64, bool, error) {
	if r.dataType != jsonparser.Number {
		return 0, false, nil
	}
	val, err := jsonparser.ParseFloat(r.value)
	if err != nil {
		return 0, true, fmt.Errorf("error unmarshalling rule value %q: %v", string(r.value), err)
	}
	return val, true, nil
}

func (r *Rule) NumberValues() []float64 {
	results := make([]float64, 0)
	for _, value := range r.Values {
		if val, ok, err := value.IsNumber(); err != nil {
			log.Warnf("%v", err)
		} else if ok {
			results = append(results, val)
		}
	}
	return results
}

func NewRings(raw PluginRawData) (Rings, error) {
	b := &builder{rings: make(Rings, 0)}
	_, err := jsonparser.ArrayEach(raw, extractRing(b), "rings")
	if err != nil && err.Error() != pathNotFoundErrMsg {
		return nil, fmt.Errorf("error unmarshalling rings from plugin configuration: %w", err)
	}
	log.Debugf("%d rings have been configured", len(b.rings))
	return b.rings, nil
}

func extractRing(b *builder) converter {
	return func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		ring := new(Ring)
		ring.ID, err = jsonparser.GetString(value, "id")
		if err != nil {
			log.Warnf("failed to extract the ring ID from %q: %v", string(value), err)
			return
		}
		_, err = jsonparser.ArrayEach(value, extractAllRules(ring), "match", "all")
		if err != nil && err.Error() != pathNotFoundErrMsg {
			log.Warnf("failed to extract rules from %q: %v", string(value), err)
			return
		}
		_, err = jsonparser.ArrayEach(value, extractAnyRules(ring), "match", "any")
		if err != nil && err.Error() != pathNotFoundErrMsg {
			log.Warnf("failed to extract rules from %q: %v", string(value), err)
			return
		}
		b.rings = append(b.rings, ring)
	}
}

func extractAllRules(r *Ring) converter {
	return func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		rule := new(Rule)
		rule.Key, err = jsonparser.GetString(value, "key")
		if err != nil && err.Error() != pathNotFoundErrMsg {
			log.Warnf("failed to extract rule key from %q: %v", string(value), err)
			return
		}
		rule.Operator, err = jsonparser.GetString(value, "operator")
		if err != nil && err.Error() != pathNotFoundErrMsg {
			log.Warnf("failed to extract rule operator from %q: %v", string(value), err)
			return
		}
		_, err = jsonparser.ArrayEach(value, extractValues(rule), "values")
		if err != nil && err.Error() != pathNotFoundErrMsg {
			log.Warnf("failed to extract rule values from %q: %v", string(value), err)
			return
		}
		r.Match.All = append(r.Match.All, rule)
	}
}

func extractAnyRules(r *Ring) converter {
	return func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		rule := new(Rule)
		rule.Key, err = jsonparser.GetString(value, "key")
		if err != nil && err.Error() != pathNotFoundErrMsg {
			log.Warnf("failed to extract rule key from %q: %v", string(value), err)
			return
		}
		rule.Operator, err = jsonparser.GetString(value, "operator")
		if err != nil && err.Error() != pathNotFoundErrMsg {
			log.Warnf("failed to extract rule operator from %q: %v", string(value), err)
			return
		}
		_, err = jsonparser.ArrayEach(value, extractValues(rule), "values")
		if err != nil && err.Error() != pathNotFoundErrMsg {
			log.Warnf("failed to extract rule values from %q: %v", string(value), err)
			return
		}
		r.Match.Any = append(r.Match.Any, rule)
	}
}

func extractValues(rule *Rule) converter {
	return func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		rule.Values = append(rule.Values, &RuleValue{value: value, dataType: dataType})
	}
}
