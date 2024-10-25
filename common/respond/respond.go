package respond

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type APIModel[T interface{}] struct {
	Data   T               `json:"data"`
	Errors []ErrorAPIModel `json:"errors"`
}

type ErrorAPIModel struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func Success(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	model := APIModel[interface{}]{Data: data}
	js, err := json.Marshal(model)
	if err != nil {
		Error(w, context.Background(), WithDesc(CodeInternalServerError, "Unknown error"), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(js); err != nil {
		log.Fatalf("%v", err)
	}
}

func Error(w http.ResponseWriter, ctx context.Context, errModel APIError, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	model := APIModel[interface{}]{Data: nil, Errors: []ErrorAPIModel{ErrorAPIModel{
		Code:    errModel.Code,
		Message: errModel.Desc,
	}}}
	js, err := json.Marshal(model)
	if err != nil {
		Error(w, ctx, WithDesc(CodeInternalServerError, "Unknown error"), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(js); err != nil {
		log.Fatalf("%v", err)
	}
	log.Printf("An error has occured: %v", string(js))
}
