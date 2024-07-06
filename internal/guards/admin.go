package guards

// import (
// 	"context"
// 	"errors"
// 	"gomodule/pkg/ctn"
// 	"gomodule/pkg/pst"
// 	"gomodule/pkg/resp"
// 	"gomodule/pkg/zlog"
// 	"net/http"
//
// 	"strings"
// 	"time"
//
// 	"github.com/gofrs/uuid/v5"
// )
//
// var (
// 	ErrUnauthorized  = errors.New("error you are not authorized")
// 	ErrToken         = errors.New("error token not match")
// 	ErrHeaderEmpty   = errors.New("error header is empty")
// 	ErrHeaderInvalid = errors.New("error invalid header format")
// 	ErrTokenEmpty    = errors.New("error token is empty")
// )
//
// func Admin(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
// 		defer cancel()
//
// 		xApiKey := r.Header.Get(ctn.X_API_KEY)
// 		if len(xApiKey) <= 0 {
// 			zlog.Error(ErrHeaderEmpty)
// 			resp.Error(ErrHeaderEmpty, http.StatusUnauthorized, w)
// 			return
// 		}
// 		apiKeyParts := strings.Split(xApiKey, ",")
// 		userID, errID := uuid.FromString(apiKeyParts[0])
// 		if errID != nil {
// 			zlog.Error(errID)
// 			resp.Error(errID, http.StatusUnauthorized, w)
// 			return
// 		}
// 		publicKey := apiKeyParts[1]
// 		repo, err := service_admin_auth.New(ctx).GetToken(userID)
//
// 		if err != nil {
// 			zlog.Error(err)
// 			resp.Error(ErrUnauthorized, http.StatusUnauthorized, w)
// 			return
// 		}
//
// 		errVerify := pst.New().Claim(pst.Payload{
// 			JTI:       userID,
// 			Signed:    repo.Token,
// 			PublicKey: publicKey,
// 		})
// 		if errVerify != nil {
// 			zlog.Error(errVerify)
// 			resp.Error(ErrUnauthorized, http.StatusUnauthorized, w)
// 			return
// 		}
//
// 		adminCtx := context.WithValue(r.Context(), ctn.USER_ID, userID)
// 		next.ServeHTTP(w, r.WithContext(adminCtx))
// 	})
// }
