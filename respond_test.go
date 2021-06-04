package respond

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNotFound(test *testing.T) {

	statusCode, result := Default.NotFound()

	assert.Equal(test, 404, statusCode)
	assert.Equal(test, map[string]interface{}{
		"message": "Oops... The requested page not found!",
		"status":  "failed",
		"error":   5404,
	}, result)
}

func TestSucceed(test *testing.T) {

	statusCode, result := Default.Succeed(map[string]interface{}{
		"data": "Test",
	})

	assert.Equal(test, 200, statusCode)
	assert.Equal(test, map[string]interface{}{
		"status": "success",
		"result": map[string]interface{}{
			"data": "Test",
		},
	}, result)
}

func TestInsertSucceeded(test *testing.T) {

	statusCode, result := Default.InsertSucceeded()

	assert.Equal(test, 200, statusCode)
	assert.Equal(test, map[string]interface{}{
		"message": "The requested parameter is added successfully!",
		"status":  "success",
	}, result)
}

func TestInsertFailed(test *testing.T) {

	statusCode, result := Default.InsertFailed()

	assert.Equal(test, 448, statusCode)
	assert.Equal(test, map[string]interface{}{
		"message": "The requested parameter is not added!",
		"status":  "failed",
	}, result)
}

func TestDeleteSucceeded(test *testing.T) {

	statusCode, result := Default.DeleteSucceeded()

	assert.Equal(test, 200, statusCode)
	assert.Equal(test, map[string]interface{}{
		"message": "The requested parameter is deleted successfully!",
		"status":  "success",
	}, result)
}

func TestDeleteFailed(test *testing.T) {

	statusCode, result := Default.DeleteFailed()

	assert.Equal(test, 447, statusCode)
	assert.Equal(test, map[string]interface{}{
		"message": "The requested parameter is not deleted!",
		"status":  "failed",
	}, result)
}

func TestUpdateSucceeded(test *testing.T) {

	statusCode, result := Default.UpdateSucceeded()

	assert.Equal(test, 200, statusCode)
	assert.Equal(test, map[string]interface{}{
		"message": "The requested parameter is updated successfully!",
		"status":  "success",
	}, result)
}

func TestUpdateFailed(test *testing.T) {

	statusCode, result := Default.UpdateFailed()

	assert.Equal(test, 449, statusCode)
	assert.Equal(test, map[string]interface{}{
		"message": "The requested parameter is not updated!",
		"status":  "failed",
	}, result)
}

func TestWrongParameters(test *testing.T) {

	statusCode, result := Default.WrongParameters()

	assert.Equal(test, 406, statusCode)
	assert.Equal(test, map[string]interface{}{
		"message": "Oops... The parameters you entered are wrong!",
		"status":  "failed",
		"error":   5406,
	}, result)
}

func TestMethodNotAllowed(test *testing.T) {

	statusCode, result := Default.MethodNotAllowed()

	assert.Equal(test, 405, statusCode)
	assert.Equal(test, map[string]interface{}{
		"message": "Oops... The method you requested is not allowed!",
		"status":  "failed",
		"error":   5405,
	}, result)
}

func TestValidationErrors(test *testing.T) {

	statusCode, result := Default.ValidationErrors(map[string]interface{}{
		"data": "Test",
	})

	assert.Equal(test, 420, statusCode)
	assert.Equal(test, map[string]interface{}{
		"status": "failed",
		"result": map[string]interface{}{
			"data": "Test",
		},
	}, result)
}

func TestRequestFieldNotfound(test *testing.T) {

	statusCode, result := Default.RequestFieldNotfound()

	assert.Equal(test, 446, statusCode)
	assert.Equal(test, map[string]interface{}{
		"status":  "failed",
		"error":   1001,
		"message": "Oops... Requested field is not found!",
	}, result)
}

func TestRequestFieldDuplicated(test *testing.T) {

	statusCode, result := Default.RequestFieldDuplicated()

	assert.Equal(test, 400, statusCode)
	assert.Equal(test, map[string]interface{}{
		"status":  "failed",
		"error":   1004,
		"message": "Failed because of duplicate",
	}, result)
}

func TestDefaultWithLang(test *testing.T) {

	statusCode, result := DefaultWithLang("fa").NotFound()

	assert.Equal(test, 404, statusCode)
	assert.Equal(test, map[string]interface{}{
		"status":  "نا موفق",
		"error":   5404,
		"message": ".صفحه درخواست شده پیدا نمیشود",
	}, result)
}
