package filters

import (
	"github.com/revel/revel"
	"github.com/elirenato/jwt"
	"regexp"
	"net/http"
	"github.com/elirenato/stringUtils"
	"strings"
	"fmt"
	"bytes"
	"golang.org/x/net/context"
	"time"
)

type JWTToken struct {
	jwt.Payload
	UserID int64 `json:"user_id"`
	Authorities []string `json:"authorities"`
}

type JWTTokenHeader string

const (
	authorizationHeader = "Authorization"
	jwtTokenHeader = JWTTokenHeader("jwtTokenHeader")
	bearerPrefix = "Bearer "
)

var (
	anonymousPaths *regexp.Regexp
	jwtSecretKey JWTTokenHeader
)

func init() {
	revel.OnAppStart(func() {
		anonymousPaths = regexp.MustCompile(revel.Config.StringDefault("auth.jwt.anonymous","/token"))
		jwtSecretKey = JWTTokenHeader(revel.Config.StringDefault("auth.jwt.secret","my-secret-key-in-production"))
    })
}

func parseFromRequest(r *http.Request) (*JWTToken, error) {
 	authHeader := r.Header.Get(authorizationHeader)
	if stringUtils.IsNotBlank(authHeader) && strings.HasPrefix(authHeader, bearerPrefix) {
		authHeader = strings.TrimPrefix(authHeader, bearerPrefix)
		jwtToken := &JWTToken{}
		v := jwt.NewHSValidator(jwt.HS512)
		v.Key = []byte(jwtSecretKey)
		err := jwt.NewDecoder(bytes.NewBufferString(authHeader), v).Decode(jwtToken)
		if err != nil {
			return nil, err
		}
		return jwtToken, nil
	}
	return nil, fmt.Errorf("Authorization header not found")
}

func setResponseError(c *revel.Controller) {
	c.Response.Status = http.StatusUnauthorized
	c.Result = c.RenderJSON (map[string]string{
		"id":      "unauthorized",
		"message": "Invalid or token is not provided",
	})
}

/*
Filter AuthFilter is Revel Filter for JWT Auth Token verification
Register it in the revel.Filters in <APP_PATH>/app/init.go

Add jwt.AuthFilter anywhere deemed appropriate, it must be register after revel.PanicFilter

	revel.Filters = []revel.Filter{
		revel.PanicFilter,
		...
		jwt.AuthFilter,		// JWT Auth Token verification for Request Paths
		...
	}

Note: If everything looks good then Claims map made available via c.Args
and can be accessed using c.Args[jwt.TOKEN_CLAIMS_KEY]
*/
func JWTFilter(c *revel.Controller, fc []revel.Filter) {
	if !anonymousPaths.MatchString(c.Request.URL.Path) {
		fmt.Println("1")
		var token *JWTToken
		token, err := parseFromRequest(c.Request.Request)
		if err != nil {
			revel.ERROR.Println(err)
			setResponseError(c)
			return
		}
		context.WithValue(context.Background(),jwtTokenHeader, token)
	}
	fc[0](c, fc[1:]) //execute the next filter
}

func CreateToken(userID int64, authorities[]string) (string, error) {
	expirationTime := time.Now().AddDate(0,0,1)
	token := JWTToken{}
	token.UserID = userID
	token.Authorities = authorities
	token.ExpirationTime = &expirationTime
	tokenBuffer := bytes.NewBuffer(nil)
	v := jwt.NewHSValidator(jwt.HS512)
	v.Key = []byte(jwtSecretKey)
	err := jwt.NewEncoder(tokenBuffer, v).Encode(token)
	if err != nil {
		return "", err
	}		
	return tokenBuffer.String(), nil
}