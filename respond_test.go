package josh

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNotFound(test *testing.T) {

	statusCode, result := Respond{}.NotFound()

	assert.Equal(test, 404, statusCode)
	assert.Equal(test, map[string]interface{} {
		"message": "Oops... The requested page not found!",
		"status": "Failed!",
		"error": 5404,
	}, result)
}

func TestSucceed(test *testing.T) {

	statusCode, result := Respond{}.Succeed(map[string]interface{} {
		"data": "Test",
	})

	assert.Equal(test, 200, statusCode)
	assert.Equal(test, map[string]interface{} {
		"status": "Success.",
		"result": map[string]interface{} {
			"data": "Test",
		},
	}, result)
}

func TestInsertSucceeded(test *testing.T) {

	statusCode, result := Respond{}.InsertSucceeded()

	assert.Equal(test, 200, statusCode)
	assert.Equal(test, map[string]interface{} {
		"message": "The requested parameter is Added successfully!",
		"status": "Success.",
	}, result)
}

func TestInsertFailed(test *testing.T) {

	statusCode, result := Respond{}.InsertFailed()

	assert.Equal(test, 448, statusCode)
	assert.Equal(test, map[string]interface{} {
		"message": "The requested parameter is not Added!",
		"status": "Failed!",
	}, result)
}

func TestDeleteSucceeded(test *testing.T) {

	statusCode, result := Respond{}.DeleteSucceeded()

	assert.Equal(test, 200, statusCode)
	assert.Equal(test, map[string]interface{} {
		"message": "The requested parameter is deleted successfully!",
		"status": "Success.",
	}, result)
}

func TestDeleteFailed(test *testing.T) {

	statusCode, result := Respond{}.DeleteFailed()

	assert.Equal(test, 447, statusCode)
	assert.Equal(test, map[string]interface{} {
		"message": "The requested parameter is not deleted!",
		"status": "Failed!",
	}, result)
}

func TestUpdateSucceeded(test *testing.T) {

	statusCode, result := Respond{}.UpdateSucceeded()

	assert.Equal(test, 200, statusCode)
	assert.Equal(test, map[string]interface{} {
		"message": "The requested parameter is updated successfully!",
		"status": "Success.",
	}, result)
}

func TestUpdateFailed(test *testing.T) {

	statusCode, result := Respond{}.UpdateFailed()

	assert.Equal(test, 449, statusCode)
	assert.Equal(test, map[string]interface{} {
		"message": "The requested parameter is not updated!",
		"status": "Failed!",
	}, result)
}

func TestWrongParameters(test *testing.T) {

	statusCode, result := Respond{}.WrongParameters()

	assert.Equal(test, 406, statusCode)
	assert.Equal(test, map[string]interface{} {
		"message": "Oops... The parameters you entered are wrong!",
		"status": "Failed!",
		"error": 5406,
	}, result)
}

func TestMethodNotAllowed(test *testing.T) {

	statusCode, result := Respond{}.MethodNotAllowed()

	assert.Equal(test, 405, statusCode)
	assert.Equal(test, map[string]interface{} {
		"message": "Oops... The method you requested is not allowed!",
		"status": "Failed!",
		"error": 5405,
	}, result)
}

func TestValidationErrors(test *testing.T) {

	statusCode, result := Respond{}.ValidationErrors(map[string]interface{} {
		"data": "Test",
	})

	assert.Equal(test, 420, statusCode)
	assert.Equal(test, map[string]interface{} {
		"status": "Failed!",
		"result": map[string]interface{} {
			"data": "Test",
		},
	}, result)
}

func TestRequestFieldNotfound(test *testing.T) {

	statusCode, result := Respond{}.RequestFieldNotfound()

	assert.Equal(test, 446, statusCode)
	assert.Equal(test, map[string]interface{} {
		"status": "Failed!",
		"error": 1001,
		"message": "Oops... Requested field is not found!",
	}, result)
}

func TestRequestFieldDuplicated(test *testing.T) {

	statusCode, result := Respond{}.RequestFieldDuplicated()

	assert.Equal(test, 400, statusCode)
	assert.Equal(test, map[string]interface{} {
		"status": "Failed!",
		"error": 1004,
		"message": "Failed because of duplicate",
	}, result)
}