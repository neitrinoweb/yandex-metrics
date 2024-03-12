package handlers

import (
	// "encoding/json"
	// "io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"yandex-metrics/internal/server/storage"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

// HTTP testing explained:
//

func TestUpdateMetric(t *testing.T) {

	type args struct {
		method string
		url    string
	}

	type want struct {
		code        int
		response    string ""
		contentType string ""
	}

	// FYI anonymous struct declaration
	test := []struct {
		name string
		code int
		want want
		args args
	}{
		{
			name: "GET 405",
			want: want{
				code: 405,
			},
			args: args{
				method: "GET",
				url:    "http://localhost/example",
			},
		},
		// {
		// 	name: "GET 404",
		// 	want: want{
		// 		code: 404,
		// 	},
		// 	args: args{
		// 		method: "GET",
		// 		url:    "http://localhost/not_exist",
		// 	},
		// },
		{
			name: "GET 200",
			want: want{
				code: 200,
			},
			args: args{
				method: "POST",
				url:    "http://localhost/update/gauge/timer/1",
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			storageMemory := storage.InitInMemoryMetricStorage()

			request := httptest.NewRequest(tt.args.method, tt.args.url, nil)
			w := httptest.NewRecorder()
			h := http.HandlerFunc(UpdateMetric(storageMemory))
			h(w, request)

			result := w.Result()

			assert.Equal(t, tt.want.code, result.StatusCode)
		})
	}

}
