package resp

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	errUnauthorized = "error you are not authorized"
	keyData         = "data"
	keyMessage      = "message"
	keyStatusCode   = "statusCode"
	keySuccess      = "success"
	keyStatus       = "status"
	contentType     = "Content-Type"
	appJson         = "application/json"
)

func Success(data interface{}, message string, statusCode int, w http.ResponseWriter) {
	response := map[string]interface{}{
		keyData:    data,
		keyMessage: message,
		keyStatus:  true,
	}
	w.Header().Set(contentType, appJson)
	w.WriteHeader(statusCode)
	errJson(response, w)
}

func SuccessText(message string, statusCode int, w http.ResponseWriter) {
	response := fmt.Sprintf("%s", message)
	w.Header().Set(contentType, "text/plain")
	w.WriteHeader(statusCode)
	_, err := w.Write([]byte(response))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SuccessCreateOrUpdate(data interface{}, message string, w http.ResponseWriter) {
	w.Header().Set(contentType, appJson)
	w.WriteHeader(http.StatusCreated)
	errJson(successResponse(data, message, http.StatusCreated), w)
}

func UnAuthorized(w http.ResponseWriter, err error) {
	w.Header().Set(contentType, appJson)
	w.WriteHeader(http.StatusUnauthorized)
	errJson(errResponse(errUnauthorized, err, http.StatusUnauthorized), w)
}

func Error(err error, statusCode int, w http.ResponseWriter) {
	response := map[string]interface{}{
		keyData:    nil,
		keyMessage: err.Error(),
		keyStatus:  false,
	}
	w.Header().Set(contentType, appJson)
	w.WriteHeader(statusCode)
	errJson(response, w)
}

func errJson(response map[string]interface{}, w http.ResponseWriter) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func errResponse(msg string, err error, statusCode int) map[string]interface{} {
	return map[string]interface{}{
		keyData: nil,
		keyMessage: map[string]any{
			"errorMessage": msg,
			"error":        err,
		},
		"statusCode": statusCode,
		"success":    false,
	}
}

func successResponse(data interface{}, message string, statusCode int) map[string]interface{} {
	return map[string]interface{}{
		keyData:       data,
		keyMessage:    message,
		keyStatusCode: statusCode,
		keySuccess:    true,
	}
}
