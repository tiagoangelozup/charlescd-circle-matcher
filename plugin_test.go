package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/proxytest"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func Test_httpHeaders_OnHttpRequestHeaders(t *testing.T) {
	//data := base64.RawStdEncoding.EncodeToString([]byte(`{"exp":1648565060,"iat":1648564760,"jti":"8a10d8c8-55a1-4d00-9a96-915aa4120288","iss":"http://keycloak.lvh.me/auth/realms/Kurtis","aud":"account","sub":"7e37ee4c-e8ef-415f-8aa5-54ecf727bdaa","typ":"Bearer","azp":"demo-frontend","session_state":"d170c91e-b4d9-496c-b1de-922fcb615b3b","acr":"1","allowed-origins":["http://0.0.0.0:8000","http://127.0.0.1:8000","http://localhost:8000"],"realm_access":{"roles":["offline_access","uma_authorization","default-roles-kurtis"]},"resource_access":{"account":{"roles":["manage-account","manage-account-links","view-profile"]}},"scope":"profile email","sid":"d170c91e-b4d9-496c-b1de-922fcb615b3b","email_verified":true,"name":"Tiago Angelo","preferred_username":"tiago","age":32,"given_name":"Tiago","family_name":"Angelo","email":"tiago@gmail.com"}`))
	//_ = data
	// Setup configurations.
	config := `{"rings":[{"id":"9d22edc0-db79-412e-9e4d-d420ec5826d0","match":{"any":[{"key":"request.auth.claims","operator":"GreaterThan","value":30}]}}]}`
	opt := proxytest.NewEmulatorOption().
		WithPluginConfiguration([]byte(config)).
		WithVMContext(&vm{})
	host, reset := proxytest.NewHostEmulator(opt)
	defer reset()

	// Call OnPluginStart.
	require.Equal(t, types.OnPluginStartStatusOK, host.StartPlugin())

	// Initialize http context.
	id := host.InitializeHttpContext()

	// Call OnHttpResponseHeaders.
	hs := [][2]string{{"X-CharlesCD-User", "eyJleHAiOjE2NDg1NjUwNjAsImlhdCI6MTY0ODU2NDc2MCwianRpIjoiOGExMGQ4YzgtNTVhMS00ZDAwLTlhOTYtOTE1YWE0MTIwMjg4IiwiaXNzIjoiaHR0cDovL2tleWNsb2FrLmx2aC5tZS9hdXRoL3JlYWxtcy9LdXJ0aXMiLCJhdWQiOiJhY2NvdW50Iiwic3ViIjoiN2UzN2VlNGMtZThlZi00MTVmLThhYTUtNTRlY2Y3MjdiZGFhIiwidHlwIjoiQmVhcmVyIiwiYXpwIjoiZGVtby1mcm9udGVuZCIsInNlc3Npb25fc3RhdGUiOiJkMTcwYzkxZS1iNGQ5LTQ5NmMtYjFkZS05MjJmY2I2MTViM2IiLCJhY3IiOiIxIiwiYWxsb3dlZC1vcmlnaW5zIjpbImh0dHA6Ly8wLjAuMC4wOjgwMDAiLCJodHRwOi8vMTI3LjAuMC4xOjgwMDAiLCJodHRwOi8vbG9jYWxob3N0OjgwMDAiXSwicmVhbG1fYWNjZXNzIjp7InJvbGVzIjpbIm9mZmxpbmVfYWNjZXNzIiwidW1hX2F1dGhvcml6YXRpb24iLCJkZWZhdWx0LXJvbGVzLWt1cnRpcyJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoicHJvZmlsZSBlbWFpbCIsInNpZCI6ImQxNzBjOTFlLWI0ZDktNDk2Yy1iMWRlLTkyMmZjYjYxNWIzYiIsImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJuYW1lIjoiVGlhZ28gQW5nZWxvIiwicHJlZmVycmVkX3VzZXJuYW1lIjoidGlhZ28iLCJhZ2UiOjMyLCJnaXZlbl9uYW1lIjoiVGlhZ28iLCJmYW1pbHlfbmFtZSI6IkFuZ2VsbyIsImVtYWlsIjoidGlhZ29AZ21haWwuY29tIn0"}}
	action := host.CallOnRequestHeaders(id, hs, false)
	require.Equal(t, types.ActionContinue, action)

	//Check headers.
	resultHeaders := host.GetCurrentRequestHeaders(id)
	var found bool
	for _, val := range resultHeaders {
		if val[0] == "X-CharlesCD-Ring" {
			require.Equal(t, "9d22edc0-db79-412e-9e4d-d420ec5826d0", val[1])
			found = true
		}
	}
	require.True(t, found)

	// Call OnHttpStreamDone.
	host.CompleteHttpContext(id)
}
