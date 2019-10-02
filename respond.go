package respond

import "strconv"

type Respond struct {

	// Status code of response
	statusCode int

	// Status text of response
	statusText string

	// Error code of response
	errorCode int

	// Language of response
	// this option work with en.yml and fa.yml files for responding
	lang string
}

var Default = &Respond{}

var DefaultWithLang = func(lang string) *Respond {
	return &Respond{ lang: lang }
}

// Get message type
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return *Message
func (respond Respond) Message() *Message {

	var message Message
	message.Lang = respond.lang

	if message.Lang == "" {
		message.Lang = "en"
	}

	data := message.LoadConfig()
	return data
}

// Set status code of response and set default value as 0
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return Respond
// @param code int
func (respond Respond) SetStatusCode(code int) Respond {

	if respond.statusCode == 0 {
		respond.statusCode = code
	}
	return respond
}

// Set status code of response and set default value as 0
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return Respond
// @param code int
func (respond Respond) SetErrorCode(code int) Respond {

	if respond.errorCode == 0 {
		respond.errorCode = code
	}
	return respond
}

// Set status text of response
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return Respond
// @param text string
func (respond Respond) SetStatusText(text string) Respond {

	if respond.statusText == "" {
		respond.statusText = text
	}
	return respond
}

// Pass response with result data like this array
//
//      array := map[string]interface{} {
//              "status": respond.statusText,
//              "result": result,
//      }
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @param result map[string]interface{}
// @return (statuscode int, result interface{})
func (respond Respond) RespondWithResult(result interface{}) (int, interface{}) {

	return respond.statusCode, map[string] interface{}{
		"status": respond.statusText,
		"result": result,
	}
}

// Pass response with message text as string
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @param message interface{}
// @return (statuscode int, result interface{})
func (respond Respond) RespondWithMessage(message interface{}) (int, interface{}) {

	data := map[string]interface{}{
		"status":  respond.statusText,
		"message": message,
	}

	if respond.errorCode != 0 {
		data["error"] = respond.errorCode
	}

	return respond.statusCode, data
}

// return notfound result
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (respond Respond) NotFound() (int, interface{}) {

	return respond.Error(404, 5404)
}

// return success result with data
//
//      data := map[string]interface{} {
//              "data": "somedata"
//      }
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @param data map[string]interface{}
// @return (statuscode int, result interface{})
func (respond Respond) Succeed(data interface{}) (int, interface{}) {

	return respond.SetStatusCode(200).SetStatusText(respond.Message().Success).RespondWithResult(data)
}

// Insert action is succeed
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (respond Respond) InsertSucceeded() (int, interface{}) {

	message := respond.Message().Errors["success"]

	return respond.SetStatusCode(200).
		SetStatusText(respond.Message().Success).RespondWithMessage(message["insert"])
}

// Insert action is failed
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (respond Respond) InsertFailed() (int, interface{}) {

	message := respond.Message().Errors["failed"]

	return respond.SetStatusCode(448).
		SetStatusText(respond.Message().Failed).RespondWithMessage(message["insert"])
}

// Delete action is succeed
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (respond Respond) DeleteSucceeded() (int, interface{}) {

	message := respond.Message().Errors["success"]

	return respond.SetStatusCode(200).
		SetStatusText(respond.Message().Success).RespondWithMessage(message["delete"])
}

// Delete action is failed
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (respond Respond) DeleteFailed() (int, interface{}) {

	message := respond.Message().Errors["failed"]

	return respond.SetStatusCode(447).
		SetStatusText(respond.Message().Failed).RespondWithMessage(message["delete"])
}

// Update action is succeed
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (respond Respond) UpdateSucceeded() (int, interface{}) {

	message := respond.Message().Errors["success"]

	return respond.SetStatusCode(200).
		SetStatusText(respond.Message().Success).RespondWithMessage(message["update"])
}

// Update action is failed
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (respond Respond) UpdateFailed() (int, interface{}) {

	message := respond.Message().Errors["failed"]

	return respond.SetStatusCode(449).
		SetStatusText(respond.Message().Failed).RespondWithMessage(message["update"])
}

// Wrong parameters are entered
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (respond Respond) WrongParameters() (int, interface{}) {

	return respond.Error(406, 5406)
}

// Wrong parameters are entered
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (respond Respond) MethodNotAllowed() (int, interface{}) {

	return respond.Error(405, 5405)
}

// There ara validation translations
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @param translations map[string]interface{}
// @return (statuscode int, result interface{})
func (respond Respond) ValidationErrors(errors interface{}) (int, interface{}) {

	return respond.SetStatusCode(420).
		SetStatusText(respond.Message().Failed).SetErrorCode(5420).RespondWithResult(errors)
}

// The request field is not found
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (respond Respond) RequestFieldNotfound() (int, interface{}) {

	return respond.Error(446, 1001)
}

// The request field is duplicated
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @return (statuscode int, result interface{})
func (respond Respond) RequestFieldDuplicated() (int, interface{}) {

	return respond.Error(400, 1004)
}

// The error message
//
// @author Alireza Josheghani <josheghani.dev@gmail.com>
// @since 15 Mar 2018
// @param statusCode int,errorCode string
// @return (statuscode int, result interface{})
func (respond Respond) Error(statusCode int, errorCode int) (int, interface{}) {

	message := respond.Message().Errors[strconv.Itoa(errorCode)]

	return respond.SetStatusCode(statusCode).
		SetStatusText(respond.Message().Failed).SetErrorCode(errorCode).RespondWithMessage(message["message"])
}
