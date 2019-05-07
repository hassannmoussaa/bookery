package apiCtrls

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/hassannmoussaa/pill.go/clean"
	"github.com/hassannmoussaa/pill.go/helpers"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
	"github.com/valyala/fasthttp"
)

var (
	webHost, domainName string
)

func Init(WebHost, DomainName string) {
	webHost = WebHost
	domainName = DomainName
}

type APICtrl struct{}

type JSONResponse struct {
	Status  string      `json:"status,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Send(requestCtx *fasthttp.RequestCtx, status string, data interface{}, message string, options ...int) {
	var statusCode int
	if status == "fail" || status == "error" {
		statusCode = 500
	} else if status == "success" {
		statusCode = 200
	}
	if options != nil && len(options) > 0 {
		statusCode = options[0]
	}
	response := JSONResponse{Status: status, Message: message, Data: data}
	requestCtx.SetContentType("application/json")
	json, err := json.Marshal(response)
	if err != nil {
		clean.Error(err)
		ServerError(requestCtx, "JSON Parse Error [Syntax Error]")
	} else {
		requestCtx.SetStatusCode(statusCode)
		requestCtx.Response.SetBodyStream(bytes.NewReader(json), len(json))
	}
}

func Fail(requestCtx *fasthttp.RequestCtx, data interface{}, message string, options ...int) {
	Send(requestCtx, "fail", data, textualContent.OfFailureMsg(message), options...)
}

func Success(requestCtx *fasthttp.RequestCtx, data interface{}, message string, options ...int) {
	Send(requestCtx, "success", data, textualContent.OfSuccessfulMsg(message), options...)
}

func Error(requestCtx *fasthttp.RequestCtx, data interface{}, message string, options ...int) {
	Send(requestCtx, "error", data, textualContent.OfErrorMsg(message), options...)
}

func ValidationError(requestCtx *fasthttp.RequestCtx, data interface{}, message string) {
	Send(requestCtx, "error", data, textualContent.OfValidationMsg(message), 422)
}
func ServerError(requestCtx *fasthttp.RequestCtx, message string) {
	Error(requestCtx, nil, message, 500)
}

func Unauthorized(requestCtx *fasthttp.RequestCtx) {
	Fail(requestCtx, nil, "unauthorized", 401)
}

func (this *APICtrl) Send(requestCtx *fasthttp.RequestCtx, status string, data interface{}, message string, options ...int) {
	Send(requestCtx, status, data, message, options...)
}

func (this *APICtrl) Fail(requestCtx *fasthttp.RequestCtx, data interface{}, message string, options ...int) {
	Fail(requestCtx, data, message, options...)
}

func (this *APICtrl) Success(requestCtx *fasthttp.RequestCtx, data interface{}, message string, options ...int) {
	Success(requestCtx, data, message, options...)
}
func (this *APICtrl) Error(requestCtx *fasthttp.RequestCtx, data interface{}, message string, options ...int) {
	Error(requestCtx, data, message, options...)
}
func (this *APICtrl) ServerError(requestCtx *fasthttp.RequestCtx, message string) {
	ServerError(requestCtx, message)
}

func (this *APICtrl) Unauthorized(requestCtx *fasthttp.RequestCtx) {
	Unauthorized(requestCtx)
}
func (this *APICtrl) ValidationError(requestCtx *fasthttp.RequestCtx, data interface{}, message string) {
	ValidationError(requestCtx, data, message)
}

func (this *APICtrl) SetFields(requestCtx *fasthttp.RequestCtx) ([]string, bool) {
	fieldsRaw := ""
	excluded := false
	if fieldsRaw = helpers.BytesToString(requestCtx.QueryArgs().Peek("fields")); fieldsRaw == "" {
		excluded = true
		fieldsRaw = helpers.BytesToString(requestCtx.QueryArgs().Peek("excluded_fields"))
	}
	if fieldsRaw != "" {
		fieldsRaw = strings.Replace(fieldsRaw, " ", "", -1)
		fieldsRaw = strings.ToLower(fieldsRaw)
		if fieldsRaw != "" {
			fields := strings.Split(fieldsRaw, ",")
			return fields, excluded
		}
	}
	return []string{}, false
}

func IsFieldRequested(field string, excluded bool, fields ...string) bool {
	check := models.Included
	if excluded {
		check = models.Excluded
	}

	if check(field, fields...) {
		return true
	}
	return false
}
