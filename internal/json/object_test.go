//go:build proxytest
// +build proxytest

package json_test

import (
	"encoding/base64"
	"github.com/stretchr/testify/require"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/json"
	"testing"
)

func TestObject_Get(t *testing.T) {
	raw := `{"exp":1648565060,"iat":1648564760,"jti":"8a10d8c8-55a1-4d00-9a96-915aa4120288","iss":"http://keycloak.lvh.me/auth/realms/Kurtis","aud":"account","sub":"7e37ee4c-e8ef-415f-8aa5-54ecf727bdaa","typ":"Bearer","azp":"demo-frontend","session_state":"d170c91e-b4d9-496c-b1de-922fcb615b3b","acr":"1","allowed-origins":["http://0.0.0.0:8000","http://127.0.0.1:8000","http://localhost:8000"],"realm_access":{"roles":["offline_access","uma_authorization","default-roles-kurtis"]},"resource_access":{"account":{"roles":["manage-account","manage-account-links","view-profile"]}},"scope":"profile email","sid":"d170c91e-b4d9-496c-b1de-922fcb615b3b","email_verified":true,"name":"Tiago Angelo","preferred_username":"tiago","age":32,"given_name":"Tiago","family_name":"Angelo","email":"tiago@gmail.com"}`
	j, err := json.FromBase64(base64.RawStdEncoding.EncodeToString([]byte(raw)))
	require.NoError(t, err)

	_, err = j.GetValue("resource_access.account.roles[0]")
	require.NoError(t, err)
}

func TestObject_GetString(t *testing.T) {
	raw := `{"exp":1648565060,"iat":1648564760,"jti":"8a10d8c8-55a1-4d00-9a96-915aa4120288","iss":"http://keycloak.lvh.me/auth/realms/Kurtis","aud":"account","sub":"7e37ee4c-e8ef-415f-8aa5-54ecf727bdaa","typ":"Bearer","azp":"demo-frontend","session_state":"d170c91e-b4d9-496c-b1de-922fcb615b3b","acr":"1","allowed-origins":["http://0.0.0.0:8000","http://127.0.0.1:8000","http://localhost:8000"],"realm_access":{"roles":["offline_access","uma_authorization","default-roles-kurtis"]},"resource_access":{"account":{"roles":["manage-account","manage-account-links","view-profile"]}},"scope":"profile email","sid":"d170c91e-b4d9-496c-b1de-922fcb615b3b","email_verified":true,"name":"Tiago Angelo","preferred_username":"tiago","age":32,"given_name":"Tiago","family_name":"Angelo","email":"tiago@gmail.com"}`
	j, err := json.FromBase64(base64.RawStdEncoding.EncodeToString([]byte(raw)))
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
