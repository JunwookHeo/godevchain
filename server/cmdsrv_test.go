package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/JunwookHeo/godevchain/bclogger"
	"github.com/stretchr/testify/assert"
)

func TestCmdSrv(t *testing.T) {
	bclogger.Init(false)

	go StartCmdSrv()
}

func TestCmdHandler(t *testing.T) {
	bclogger.Init(false)

	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello BC", string(data))
}
