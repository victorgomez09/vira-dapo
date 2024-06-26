package middlewares

// import (
// 	"net/http"

// 	"github.com/uptrace/bunrouter"
// 	"github.com/victorgomez09/vira-dapo/internal/auth"
// )

// func AuthMiddleware(next bunrouter.HandlerFunc) bunrouter.HandlerFunc {
// 	// you can initialize the middleware here

// 	// Return the middleware handler.
// 	return func(w http.ResponseWriter, req bunrouter.Request) error {
// 		w.Header().Set("Content-Type", "application/json")
// 		tokenString := req.Header.Get("Authorization")
// 		if tokenString == "" {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			bunrouter.JSON(w, bunrouter.H{
// 				"error": "Missing authorization header",
// 			})

// 			return nil
// 		}

// 		tokenString = tokenString[len("Bearer "):]
// 		err := auth.VerifyToken(tokenString)
// 		if err != nil {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			bunrouter.JSON(w, bunrouter.H{
// 				"error": "Invalid token",
// 			})

// 			return nil
// 		}

// 		return next(w, req)
// 	}
// }
