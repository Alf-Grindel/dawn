package middleware

import (
	"context"
	"github.com/alf-grindel/dawn/internal/model/data"
	"github.com/alf-grindel/dawn/pkg/constants"
	"github.com/alf-grindel/dawn/pkg/errno"
	"github.com/alf-grindel/dawn/pkg/resp"
	"github.com/gorilla/sessions"
	"net/http"
)

func AuthMiddleware(store *sessions.CookieStore) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, constants.UserLoginState)
			if session.Values["login"] != true {
				resp.WriteJson(rw, errno.NotLoginErr, nil)
				return
			}

			user, ok := session.Values["user"].(data.User)
			if !ok {
				resp.WriteJson(rw, errno.SystemErr, nil)
				return
			}

			if user.UserRole != "admin" {
				resp.WriteJson(rw, errno.NoAuthErr, nil)
				return
			}
			ctx := context.WithValue(r.Context(), "user", user)
			r = r.WithContext(ctx)
			next.ServeHTTP(rw, r)
		})
	}
}
