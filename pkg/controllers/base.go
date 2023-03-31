package controllers

import (
	"bytes"
	"context"
	models "diwa/kredit-plus/pkg/models"
	"diwa/kredit-plus/pkg/usecase"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/validator"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
)

type ctrl struct {
	uc usecase.UC
}

var (
	validate       = validator.New()
	decoder        = schema.NewDecoder()
	SUCCESS        = "SUCCESS"
	INTERNALSERVER = "INTERNAL SERVER ERROR"
	INVALIDREQUEST = "INVALID REQUEST"
	NOTFOUND       = "DATA NOT FOUND"
)

type Controllers interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	GetDataUser(w http.ResponseWriter, r *http.Request)
	FindByIdLogin(w http.ResponseWriter, r *http.Request)
	CreateTransaction(w http.ResponseWriter, r *http.Request)
}

func NewControllers(us usecase.UC) Controllers {
	return &ctrl{uc: us}
}

func Response(w http.ResponseWriter, ctx context.Context, code int, status bool, message string, res, pagination interface{}) {
	resservice := models.Responseservice{}
	resservice.Status = code
	if status {
		resservice.Data = res
		resservice.Pagination = pagination
	} else {
		resservice.ErrorMessage = "Error"
	}

	var input []byte

	resservice.Message = message
	switch res.(type) {
	case string:
		input = []byte(res.(string))
	case []byte:
		input = res.([]byte)
	default:
		input, _ = JSONMarshal(res)
	}

	if ctx == nil {
		ctx = context.TODO()
	}

	Logger(ctx, string(input), code)
	origin := "*"

	v, ok := ctx.Value(models.LogKey).(*models.Data)
	if ok {
		words := strings.Fields(v.RequestHeader)
		for i := 0; i < len(words); i++ {
			if words[i] == "Origin:" {
				origin = words[i+1]
				break
			}
		}

	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("Strict-Transport-Security", "max-age=15552000; includeSubDomains")
	w.Header().Set("X-DNS-Prefetch-Control", "off")
	w.Header().Set("Vary", "X-HTTP-Method-Override")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resservice)
}

// Logger
func Logger(ctx context.Context, response string, statuscode int) {
	var level string

	v, ok := ctx.Value(models.LogKey).(*models.Data)
	if ok {
		t := time.Since(v.TimeStart)
		if statuscode >= 200 && statuscode < 400 {
			level = "INFO"
		} else if statuscode >= 400 && statuscode < 500 {
			level = "WARM"
		} else {
			level = "ERROR"
		}

		v.StatusCode = statuscode
		v.Response = response
		v.ExecTime = t.Seconds()

		if statuscode == 0 {
			v.StatusCode = 200
		}

		Output(v, level)
	}
}

// UTCFormatter ...
type UTCFormatter struct {
	logrus.Formatter
}

// Output for output to terminal
func Output(out *models.Data, level string) {
	logrus.SetFormatter(UTCFormatter{&logrus.JSONFormatter{}})
	if level == "ERROR" {
		logrus.WithField("data", out).Error("apps")
	} else if level == "INFO" {
		logrus.WithField("data", out).Info("apps")
	} else if level == "WARN" {
		logrus.WithField("data", out).Warn("apps")
	}
}

// JSONMarshal is func
func JSONMarshal(t interface{}) ([]byte, error) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	err := encoder.Encode(t)
	return buffer.Bytes(), err
}
