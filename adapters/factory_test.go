package adapters

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Factory(t *testing.T) {
	testCases := []struct {
		name         string
		adapterName  string
		depContainer *DepContainer
		wantErr      bool
	}{
		{
			name:        "redis - should pass",
			adapterName: "redis",
			depContainer: &DepContainer{
				RedisURL: os.Getenv("REDIS_URL"),
			},
		},
		{
			name:        "non existent adapter - should err",
			adapterName: "redis_",
			wantErr:     true,
		},
		{
			name:    "empty adapter - should err",
			wantErr: true,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			_, err := Factory(tCase.adapterName, tCase.depContainer)

			assert.Equal(t, tCase.wantErr, err != nil)
		})
	}
}
