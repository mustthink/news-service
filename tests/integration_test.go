package tests

import (
	"bytes"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	// postsEndpoint is the endpoint for the news service.
	postsEndpoint = "/posts"
)

// TestIntegrationWithOrder is a test function. Order of cases is important for this test.
func TestIntegrationWithOrder(t *testing.T) {
	var setURL = os.Getenv("SET_URL")
	if setURL == "" {
		setURL = "http://localhost:8080"
	}

	tests := []struct {
		name       string
		endpoint   string
		method     string
		body       []byte
		wantStatus int
	}{
		{
			name:     "Post news",
			endpoint: postsEndpoint,
			method:   http.MethodPost,
			body: []byte(`
				{
					"title": "test title",
					"content": "test content",
					"author_id": 1,
					"topic_id": 1
				}`),
			wantStatus: http.StatusOK,
		},
		{
			name:       "Get news",
			endpoint:   postsEndpoint + "/1",
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
		},
		{
			name:       "Get all news",
			endpoint:   postsEndpoint,
			method:     http.MethodGet,
			wantStatus: http.StatusOK,
		},
		{
			name:     "Update news",
			endpoint: postsEndpoint + "/1",
			method:   http.MethodPut,
			body: []byte(`
				{
					"title": "test title",
					"content": "test content 2",
					"author_id": 1,
					"topic_id": 1
				}`),
			wantStatus: http.StatusOK,
		},
		{
			name:       "Delete news",
			endpoint:   postsEndpoint + "/1",
			method:     http.MethodDelete,
			wantStatus: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req, err := http.NewRequest(test.method, setURL+test.endpoint, bytes.NewReader(test.body))
			require.NoError(t, err)

			resp, err := http.DefaultClient.Do(req)
			require.NoError(t, err)

			assert.Equal(t, test.wantStatus, resp.StatusCode)
		})
	}
}
