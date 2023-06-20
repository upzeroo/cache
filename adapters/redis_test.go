package adapters

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewRedisAdapter(t *testing.T) {
	testCases := []struct {
		name         string
		depContainer *DepContainer
		wantErr      bool
	}{
		{
			name: "redis - should pass",
			depContainer: &DepContainer{
				RedisURL: os.Getenv("REDIS_URL"),
			},
		},
		{
			name: "redis - broker ulr - should err",
			depContainer: &DepContainer{
				RedisURL: "fake_url",
			},
			wantErr: true,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			_, err := NewRedisAdapter(tCase.depContainer)

			assert.Equal(t, tCase.wantErr, err != nil)
		})
	}
}

func Test_Get(t *testing.T) {
	adapter, _ := NewRedisAdapter(&DepContainer{
		RedisURL: os.Getenv("REDIS_URL"),
	})
	testKey := "test_key"
	data := struct{ Data string }{Data: "some values here ..."}
	jsonData, err := json.Marshal(&data)
	if err != nil {
		t.Fatal(err)
	}

	err = adapter.Set(testKey, data, 20*time.Second)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name    string
		key     string
		res     string
		wantErr bool
	}{
		{
			name: "get real key - must pass",
			key:  testKey,
			res:  string(jsonData),
		},
		{
			name:    "get non existent - must err",
			key:     "fake_one",
			wantErr: true,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			res, err := adapter.Get(tCase.key)

			assert.Equal(t, tCase.res, res)
			assert.Equal(t, tCase.wantErr, err != nil)
		})
	}

	adapter.Delete("fake_one")
}

func Test_Set(t *testing.T) {
	adapter, _ := NewRedisAdapter(&DepContainer{
		RedisURL: os.Getenv("REDIS_URL"),
	})

	testKey := "test_key"

	testCases := []struct {
		name    string
		key     string
		data    interface{}
		wantErr bool
	}{
		{
			name: "set - must pass",
			key:  testKey,
			data: struct{ Data string }{Data: "some values here ..."},
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			err := adapter.Set(tCase.key, tCase.data, 20*time.Second)

			assert.Equal(t, tCase.wantErr, err != nil)
		})
	}
}

func Test_Delete(t *testing.T) {
	adapter, _ := NewRedisAdapter(&DepContainer{
		RedisURL: os.Getenv("REDIS_URL"),
	})

	testKey := "test_key"

	testCases := []struct {
		name    string
		key     string
		wantErr bool
	}{
		{
			name: "delete - must pass",
			key:  testKey,
		},
	}

	for _, tCase := range testCases {
		t.Run(tCase.name, func(t *testing.T) {
			err := adapter.Delete(tCase.key)

			assert.Equal(t, tCase.wantErr, err != nil)
		})
	}
}
