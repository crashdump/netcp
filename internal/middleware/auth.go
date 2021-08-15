package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	auth0 "github.com/auth0-community/go-auth0"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/gin-gonic/gin"
	jose "gopkg.in/square/go-jose.v2"
)

var (
	Domain   string
	ClientID string
	Audience string
	ApiKey   string
)

func AuthMiddleware() gin.HandlerFunc {
	// TODO: Store this in the database
	// Credentials below are test credentials and
	// should not be used in production.
	Domain = "netcp-dev.eu.auth0.com"
	ClientID = "HlpkvBqBPwLLSTMTpaeI54Gh5H0R73NB"
	Audience = "http://127.0.0.1:3000/srv/v1/"
	ApiKey = "15HPj35uON5V77zP1xvNsR1eLCOv4idn"

	return func(c *gin.Context) {
		ah := c.Request.Header.Get("Authorization")
		if len(strings.TrimSpace(ah)) == 0 {
			c.AbortWithStatus(401)
			return
		}

		switch at := strings.Fields(ah); at[0] {
		// Server KEY
		case "Basic":
			if ApiKey != "" && ah == fmt.Sprintf("Basic %s", ApiKey) {
				c.Next()
				return
			}
		// OIDC
		case "Bearer":
			err, status := checkJwt(c.Request)
			if err != nil {
				c.AbortWithStatus(status)
				return
			}
			c.Next()
			return
		}
		c.AbortWithStatus(401)
	}
}

// This middleware expects a HS256-compliant JSON Web Token to authenticate
// users. It MUST be used to secure all handlers related to the Web
// application. The user's auth0_id should be in the "sub" claim of this token,
// according to Auth0. The JWT must be passed in the Authorization header:
//
//   Authorization: Bearer <JWT goes here>
//
// When a new user authenticates (i.e. auth_id missing in the database), we create
// the user first by calling the Auth0 Server and get basic user information.
//
// Once the user has been found its added to the request's context.
func checkJwt(r *http.Request) (error, int) {
	auth0Endpoint := fmt.Sprintf("https://%s/", Domain)
	auth0JwksUri := fmt.Sprintf("https://%s/.well-known/jwks.json", Domain)

	client := auth0.NewJWKClient(auth0.JWKClientOptions{URI: auth0JwksUri}, nil)
	configuration := auth0.NewConfiguration(client, []string{Audience}, auth0Endpoint, jose.RS256)
	validator := auth0.NewValidator(configuration, nil)

	token, err := validator.ValidateRequest(r)
	if err != nil {
		log.Println("invalid token", err.Error())
		return err, http.StatusUnauthorized
	}

	// user's auth0_id is stored in a JWT claim (`sub`)
	claims := map[string]interface{}{}
	err = validator.Claims(r, token, &claims)
	if err != nil {
		log.Println("cannot retrieve JWT claims", err.Error())
		return err, http.StatusBadRequest
	}

	id := claims["sub"].(string)

	user := entity.User{Auth0ID: id}
	log.Println(user)

	//u, err := repo.GetUserByAuth0ID(id)
	//if err != nil {
	//	if err == sql.ErrNoRows {
	//		profile, err := getUserProfile(c.Domain, r.Header.Get("Authorization"))
	//		if err != nil {
	//			logger.Warn("cannot retrieve user profile", zap.Error(err))
	//			SendError(w, http.StatusUnauthorized, DetailUserProfileRetrievalFailed)
	//			return
	//		}
	//
	//		logger.Info(
	//			"create new authenticated user",
	//			zap.String("auth0_id", id),
	//			zap.String("login", profile["nickname"]),
	//			zap.String("avatar_url", profile["picture"]),
	//		)
	//
	//		u, err = repo.CreateNewUser(profile["sub"], profile["nickname"], profile["picture"])
	//		if err != nil {
	//			logger.Error("cannot create new user", zap.Error(err))
	//			SendError(w, http.StatusInternalServerError, DetailUserCreationFailed)
	//			return
	//		}
	//	} else {
	//		logger.Error("could not select user by ID", zap.Error(err), zap.String("auth0_id", id))
	//		SendError(w, http.StatusInternalServerError, DetailUserSelectionFailed)
	//		return
	//	}
	//}
	//
	//ctx := context.WithValue(r.ctx(), ContextCurrentUser, u)

	return nil, http.StatusOK
}
