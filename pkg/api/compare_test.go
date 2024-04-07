package api_test

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetRoute(t *testing.T) {
	r := gin.Default()

	//r.GET("/get/:name", api.Get)
	//r.GET("/g/:name", api.Get)

	//t.Run("if keys is not set, get returns 0", func(t *testing.T) {
	//	w := httptest.NewRecorder()
	//	req, _ := http.NewRequest(http.MethodGet, "/get/test?index=0", nil)
	//	r.ServeHTTP(w, req)
	//
	//	snaps.MatchSnapshot(t, w.Body.String())
	//	assert.Equal(t, http.StatusOK, w.Code)
	//})
}
