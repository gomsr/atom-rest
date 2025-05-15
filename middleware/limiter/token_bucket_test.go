package limiter

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"golang.org/x/time/rate"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestXRateError(t *testing.T) {
	limit := rate.NewLimiter(rate.Every(time.Minute), 1)
	testSuccessThenFailure(
		t,
		NewErrorLimiter(limit),
		ErrLimited.Error())
}

func TestXRateDelay(t *testing.T) {
	limit := rate.NewLimiter(rate.Every(time.Second), 1)
	testSuccessThenFailure(
		t,
		NewDelayLimiter(limit),
		"exceed context deadline")
}

func testSuccessThenFailure(t *testing.T, e gin.HandlerFunc, failContains string) {
	r := gin.Default()
	r.Use(e)
	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, "ok")
	})

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// First request should succeed.
	req, _ := http.NewRequest("GET", "/test", nil)
	req.WithContext(ctx)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	// Next request should fail.
	req2, _ := http.NewRequest("GET", "/test", nil)
	req2.WithContext(ctx)
	w2 := httptest.NewRecorder()
	r.ServeHTTP(w2, req2)
	assert.Contains(t, w2.Body.String(), "ok")
}
