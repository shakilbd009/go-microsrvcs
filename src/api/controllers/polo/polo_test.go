package polo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/shakilbd009/go-microsrvcs/src/api/utils/test_utils"
	"github.com/stretchr/testify/assert"
)

func TestPolo(t *testing.T) {

	request := httptest.NewRequest(http.MethodGet, "/marco", nil)
	resp := httptest.NewRecorder()
	c := test_utils.GetMockedContext(request, resp)
	Polo(c)
	assert.EqualValues(t, http.StatusOK, resp.Code)
	assert.EqualValues(t, "polo", resp.Body.String())
}

func TestConst(t *testing.T) {
	assert.EqualValues(t, "polo", polo)
}
