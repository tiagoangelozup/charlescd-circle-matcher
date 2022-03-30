package json_test

import (
	"github.com/stretchr/testify/require"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/json"
	"testing"
)

func TestObject_GetString(t *testing.T) {
	j, err := json.FromBase64("eyJleHAiOjE2NDg1NjUwNjAsImlhdCI6MTY0ODU2NDc2MCwianRpIjoiOGExMGQ4YzgtNTVhMS00ZDAwLTlhOTYtOTE1YWE0MTIwMjg4IiwiaXNzIjoiaHR0cDovL2tleWNsb2FrLmx2aC5tZS9hdXRoL3JlYWxtcy9LdXJ0aXMiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiN2UzN2VlNGMtZThlZi00MTVmLThhYTUtNTRlY2Y3MjdiZGFhIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiZGVtby1mcm9udGVuZCIsInNlc3Npb25fc3RhdGUiOiJkMTcwYzkxZS1iNGQ5LTQ5NmMtYjFkZS05MjJmY2I2MTViM2IiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbImh0dHA6Ly8wLjAuMC4wOjgwMDAiLCJodHRwOi8vMTI3LjAuMC4xOjgwMDAiLCJodHRwOi8vbG9jYWxob3N0OjgwMDAiXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLWt1cnRpcyJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoicHJvZmlsZSBlbWFpbCIsInNpZCI6ImQxNzBjOTFlLWI0ZDktNDk2Yy1iMWRlLTkyMmZjYjYxNWIzYiIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJuYW1lIjoiVGlhZ28gQW5nZWxvIiwicHJlZmVycmVkX3VzZXJuYW1lIjoidGlhZ28iLCJnaXZlbl9uYW1lIjoiVGlhZ28iLCJmYW1pbHlfbmFtZSI6IkFuZ2VsbyIsImVtYWlsIjoidGlhZ29AZ21haWwuY29tIn0")
	require.NoError(t, err)

	val, err := j.GetString("resource_access.account.roles[1]")
	require.NoError(t, err)
	require.Equal(t, "manage-account-links", val)
}

func TestSplitKey(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "should split to 2 items", args: args{key: "realm_access.roles"}, want: 2},
		{name: "should split to 4 items", args: args{key: "resource_access.account.roles[2]"}, want: 4},
		{name: "should split to 5 item", args: args{key: "resource_access.account.roles[2][1]"}, want: 5},
		{name: "should split to 5 items", args: args{key: "realm_access.roles[33][12][1]"}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := json.SplitKey(tt.args.key)
			require.Len(t, got, tt.want)
		})
	}
}
