//copied from https://github.com/weaveworks/common/blob/master/http/client/client_test.go
// because it is not included in dskit

package internal

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTimedClient_operationName(t *testing.T) {
	r, err := http.NewRequest("GET", "https://weave.test", nil)
	assert.NoError(t, err)

	r = r.WithContext(context.WithValue(t.Context(), OperationNameContextKey, "opp"))
	c := NewTimedClient(http.DefaultClient, nil)

	assert.Equal(t, "opp", c.operationName(r))
}

func TestTimedClient_operationName_Default(t *testing.T) {
	r, err := http.NewRequest("GET", "https://weave.test/you/know/me", nil)
	assert.NoError(t, err)

	r = r.WithContext(t.Context())
	c := NewTimedClient(http.DefaultClient, nil)

	assert.Equal(t, "/you/know/me", c.operationName(r))
}
