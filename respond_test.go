package respond

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getExpectedMap(r io.Reader) (map[string]interface{}, error) {
	expected := map[string]interface{}{}
	if err := json.NewDecoder(r).Decode(&expected); err != nil {
		return nil, err
	}
	return expected, nil
}

func TestNotFound(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).NotFound()

	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"message": "Oops... The requested page not found!",
		"status":  "failed",
		"error":   float64(5404),
	}, expected)

}

func TestSucceed(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).Succeed(map[string]interface{}{
		"data": "Test",
	})

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"status": "success",
		"result": map[string]interface{}{
			"data": "Test",
		},
	}, expected)
}

func TestInsertSucceeded(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).InsertSucceeded()

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"message": "The requested parameter is added successfully!",
		"status":  "success",
	}, expected)
}

func TestInsertFailed(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).InsertFailed()

	assert.Equal(t, 448, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"message": "The requested parameter is not added!",
		"status":  "failed",
	}, expected)
}

func TestDeleteSucceeded(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).DeleteSucceeded()

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"message": "The requested parameter is deleted successfully!",
		"status":  "success",
	}, expected)
}

func TestDeleteFailed(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).DeleteFailed()

	assert.Equal(t, 447, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"message": "The requested parameter is not deleted!",
		"status":  "failed",
	}, expected)
}

func TestUpdateSucceeded(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).UpdateSucceeded()

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"message": "The requested parameter is updated successfully!",
		"status":  "success",
	}, expected)
}

func TestUpdateFailed(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).UpdateFailed()

	assert.Equal(t, 449, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"message": "The requested parameter is not updated!",
		"status":  "failed",
	}, expected)
}

func TestWrongParameters(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).WrongParameters()

	assert.Equal(t, 406, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"message": "Oops... The parameters you entered are wrong!",
		"status":  "failed",
		"error":   float64(5406),
	}, expected)
}

func TestMethodNotAllowed(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).MethodNotAllowed()

	assert.Equal(t, 405, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"message": "Oops... The method you requested is not allowed!",
		"status":  "failed",
		"error":   float64(5405),
	}, expected)
}

func TestValidationErrors(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).ValidationErrors(map[string]interface{}{
		"data": "Test",
	})

	assert.Equal(t, 420, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"status": "failed",
		"result": map[string]interface{}{
			"data": "Test",
		},
	}, expected)
}

func TestRequestFieldNotfound(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).RequestFieldNotfound()

	assert.Equal(t, 446, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"status":  "failed",
		"error":   float64(1001),
		"message": "Oops... Requested field is not found!",
	}, expected)
}

func TestRequestFieldDuplicated(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).RequestFieldDuplicated()

	assert.Equal(t, http.StatusBadRequest, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"status":  "failed",
		"error":   float64(1004),
		"message": "Failed because of duplicate",
	}, expected)
}

func TestDefaultWithLang(t *testing.T) {

	t.Parallel()

	recorder := httptest.NewRecorder()
	NewWithWriter(recorder).Language("fa").NotFound()

	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"status":  "نا موفق",
		"error":   float64(5404),
		"message": ".صفحه درخواست شده پیدا نمیشود",
	}, expected)
}

func TestCustomLanguage(t *testing.T) {

	t.Parallel()

	ruMessages := map[string]interface{}{
		"success": "успех",
		"failed":  "не смогли",
		"errors": map[string]map[string]interface{}{
			"5404": {
				"message": "Упс ... Запрошенная страница не найдена!",
				"type":    "ошибка",
			},
		},
	}

	var (
		recorder = httptest.NewRecorder()
		respond  = NewWithWriter(recorder)
	)

	respond.Messages().AddLanguageTranslation("ru", ruMessages)
	respond.Language("ru").NotFound()

	assert.Equal(t, http.StatusNotFound, recorder.Result().StatusCode)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))

	expected, err := getExpectedMap(recorder.Body)
	assert.NoError(t, err)

	assert.Equal(t, map[string]interface{}{
		"status":  "не смогли",
		"error":   float64(5404),
		"message": "Упс ... Запрошенная страница не найдена!",
	}, expected)

}
