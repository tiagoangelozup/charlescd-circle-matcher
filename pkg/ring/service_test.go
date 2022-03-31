//go:build proxytest
// +build proxytest

package ring_test

import (
	"encoding/base64"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/config"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/http"
	"github.com/tiagoangelozup/charlescd-circle-matcher/internal/logger"
	"github.com/tiagoangelozup/charlescd-circle-matcher/pkg/ring"
	"reflect"
	"testing"
)

type httpRequestMock struct{}

func (h *httpRequestMock) GetHeader(key string) (string, error) {
	if key == "X-CharlesCD-User" {
		j := `{"exp":1648565060,"iat":1648564760,"jti":"8a10d8c8-55a1-4d00-9a96-915aa4120288","iss":"http://keycloak.lvh.me/auth/realms/Kurtis","aud":"account","sub":"7e37ee4c-e8ef-415f-8aa5-54ecf727bdaa","typ":"Bearer","azp":"demo-frontend","session_state":"d170c91e-b4d9-496c-b1de-922fcb615b3b","acr":"1","allowed-origins":["http://0.0.0.0:8000","http://127.0.0.1:8000","http://localhost:8000"],"realm_access":{"roles":["offline_access","uma_authorization","default-roles-kurtis"]},"resource_access":{"account":{"roles":["manage-account","manage-account-links","view-profile"]}},"scope":"profile email","sid":"d170c91e-b4d9-496c-b1de-922fcb615b3b","email_verified":true,"name":"Tiago Angelo","preferred_username":"tiago","age":32,"given_name":"Tiago","family_name":"Angelo","email":"tiago@gmail.com"}`
		return base64.RawStdEncoding.EncodeToString([]byte(j)), nil
	}
	return "", nil
}

func TestService_FindRings(t *testing.T) {
	type fields struct{ rings []*config.Ring }
	type args struct{ req http.Request }
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "should find a single ring when the expression matches",
			args: args{req: new(httpRequestMock)},
			fields: fields{rings: []*config.Ring{{
				ID: "321",
				Match: &config.Match{Any: []*config.Rule{{
					Key:      "request.auth.claims.age",
					Operator: "GreaterThan",
					Values:   []interface{}{31},
				}}},
			}}},
			want: []string{"321"},
		},
		{name: "should result in empty array when no rings", args: args{req: new(httpRequestMock)}, want: []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := ring.NewService(new(logger.Local), tt.fields.rings)
			got, err := s.FindRings(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindRings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindRings() got = %v, want %v", got, tt.want)
			}
		})
	}
}
