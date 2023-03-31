package route

import (
	"bytes"
	"context"
	"diwa/kredit-plus/pkg/controllers"
	"diwa/kredit-plus/pkg/models"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"reflect"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/rs/cors"
)

const SECRET_KEY = "1234567"

type Claims struct {
	jwt.StandardClaims
	UserId int
}

const TOKEN_EXP = int64(time.Hour * 3)

type route struct {
	ctrl controllers.Controllers
}

type Route interface {
	Router() http.Handler
}

func NewRoute(ctrl controllers.Controllers) Route {
	return &route{ctrl: ctrl}
}

func (c *route) Router() http.Handler {
	router := chi.NewRouter()
	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,PUT,PATCH,POST,DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,x-api-key,X-API-KEY")
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedHeaders:   []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token", "Authorization"},
			AllowedMethods:   []string{"GET", "PATCH", "POST", "PUT", "OPTIONS", "DELETE"},
			Debug:            true,
			AllowCredentials: true,
		})
		cors.Handler(corsMiddle())
		if r.Method == "OPTIONS" {
			w.WriteHeader(204)
			w.Header().Set("Access-Control-Allow-Origin", "*")
			json.NewEncoder(w).Encode("PREFLIGHT OK ")
			return
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	})

	// user router
	router.Group(func(r chi.Router) {
		r.Use(Loggers)
		r.Post("/login", c.ctrl.Login)
		r.Post("/register", c.ctrl.Register)
		r.Get("/users/all", c.ctrl.GetDataUser)

	})

	router.Group(func(r chi.Router) {
		r.Use(Loggers)
		r.Get("/users/me", c.ctrl.FindByIdLogin)
	})

	// credit transaction router
	router.Group(func(r chi.Router) {
		r.Use(Loggers)
		r.Post("/credit-transaction", c.ctrl.CreateTransaction)
	})

	return router
}

func corsMiddle() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, request *http.Request) {
		if request.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
	})
}

func Loggers(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		loc, _ := time.LoadLocation("Asia/Jakarta")
		start := time.Now().In(loc)
		r = StartRecord(r, start, "")
		f.ServeHTTP(w, r)
	})
}

func StartRecord(req *http.Request, start time.Time, id string) *http.Request {
	ctx := req.Context()

	v := new(models.Data)
	v.RequestID = uuid.New().String()

	v.Host = req.Host
	v.Endpoint = req.URL.Path
	v.TimeStart = start
	v.Device = "Web-Base"

	v.RequestMethod = req.Method
	v.RequestHeader = DumpRequest(req)
	v.UserId = id

	ctx = context.WithValue(ctx, models.LogKey, v)

	return req.WithContext(ctx)
}

func DumpRequest(req *http.Request) string {
	header, err := httputil.DumpRequest(req, true)
	if err != nil {
		return "cannot dump request"
	}

	trim := bytes.ReplaceAll(header, []byte("\r\n"), []byte("   "))
	return string(trim)
}

func GetToken(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		authorization := r.Header.Get("Authorization")
		if reflect.ValueOf(authorization).IsZero() {
			log.Println("Unauthorized")
			controllers.Response(w, nil, 401, true, "Unauthorized", nil, nil)
			return
		}
		aut := ExtractToken(authorization)

		claim := &Claims{}
		_, err := jwt.ParseWithClaims(aut, claim, func(authorization *jwt.Token) (interface{}, error) {
			// verify iss claim
			return []byte(SECRET_KEY), nil
		})
		if err != nil {
			controllers.Response(w, nil, 401, true, "Unauthorized", nil, nil)
			return
		}

		loc, _ := time.LoadLocation("Asia/Jakarta")
		start := time.Now().In(loc)
		r = StartRecord(r, start, claim.Id)
		f.ServeHTTP(w, r)
	})
}

func ExtractToken(bearToken string) string {
	// normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}
