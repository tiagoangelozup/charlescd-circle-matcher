package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRings(t *testing.T) {
	j := `{"rings":[{"id":"9d22edc0-db79-412e-9e4d-d420ec5826d0","match":{"any":[{"key":"request.auth.claims.age","operator":"GreaterThan","values":[30,31]}]}},{"id":"cbb548ce-e412-4b16-9191-e06544beb69d","match":{"all":[{"key":"request.auth.claims.city","operator":"Equals","values":["São Carlos/SP",33]}]}}]}`
	rings, err := NewRings([]byte(j))
	require.NoError(t, err)
	require.Len(t, rings, 2)
}
