package middleware

import (
	"context"
	"log"
	"net/http"
	"usus-sehat/cmd/user/repo"
	"usus-sehat/internal/domain/entity"
	"usus-sehat/internal/domain/model"

	"github.com/gorilla/securecookie"
)

type Middleware interface {
	Authentication(next http.Handler) http.Handler
	Authorization(next http.Handler) http.Handler
	// SetCSRFKey(next http.Handler) http.Handler
}

type middleware struct {
	sc *securecookie.SecureCookie
	ur repo.UserRepo
}

// // SetCSRFKey implements Middleware.
// func (md *middleware) SetCSRFKey(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		session, err := md.ss.Get(r, "CSRF")

// 		if err != nil {
// 			log.Println("[warn] An error occured", err.Error())
// 			http.Error(w, "Something went wrong", http.StatusInternalServerError)
// 			return
// 		}

// 		session.Values["CSRFKey"] = securecookie.GenerateRandomKey(32)
// 		session.Save(r, w)

// 		next.ServeHTTP(w, r)
// 	})
// }

type userKey string

const Key userKey = "user"

// Authentication implements Middleware.
func (md *middleware) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, e := r.Cookie("token")

		if e != nil {
			log.Println("[warn]", e.Error())
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		data := &model.TokenResponse{}

		e = md.sc.Decode("token", cookie.Value, &data)

		if e != nil {
			log.Println("[warn]", e.Error())
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		u := &entity.User{}

		err := u.ValidateToken(data.Token)

		if err != nil {
			log.Println("[warn]", err.Message())
			http.Error(w, err.Error(), err.Status())
			return
		}

		user, err := md.ur.FetchById(int(u.ID))

		if err != nil {
			http.Error(w, err.Error(), err.Status())
			return
		}

		r = r.WithContext(context.WithValue(context.Background(), Key, user))

		next.ServeHTTP(w, r)
	})
}

// Authorization implements Middleware.
func (md *middleware) Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie, e := r.Cookie("token")

		if e != nil {
			log.Println("[warn]", e.Error())
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		data := &model.TokenResponse{}

		e = md.sc.Decode("token", cookie.Value, &data)

		if e != nil {
			http.Error(w, e.Error(), http.StatusInternalServerError)
			return
		}

		u := &entity.User{}

		err := u.ValidateToken(data.Token)

		if err != nil {
			log.Println("[warn]", err.Message())
			http.Error(w, err.Error(), err.Status())
			return
		}

		if u.Role != "Admin" {
			http.Error(w, "You're not authorized to access this endpoint", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func NewMiddleware(sc *securecookie.SecureCookie, ur repo.UserRepo) Middleware {
	return &middleware{
		sc: sc,
		ur: ur,
	}
}
