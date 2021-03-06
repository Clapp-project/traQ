package oauth2

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/traPtitech/traQ/model"
	"github.com/traPtitech/traQ/rbac"
	"github.com/traPtitech/traQ/repository"
	"github.com/traPtitech/traQ/router/consts"
	"github.com/traPtitech/traQ/router/extension"
	"github.com/traPtitech/traQ/router/middlewares"
	"go.uber.org/zap"
)

const (
	grantTypeAuthorizationCode = "authorization_code"
	grantTypePassword          = "password"
	grantTypeClientCredentials = "client_credentials"
	grantTypeRefreshToken      = "refresh_token"

	errInvalidRequest          = "invalid_request"
	errUnauthorizedClient      = "unauthorized_client"
	errAccessDenied            = "access_denied"
	errUnsupportedResponseType = "unsupported_response_type"
	errInvalidScope            = "invalid_scope"
	errServerError             = "server_error"
	errInvalidClient           = "invalid_client"
	errInvalidGrant            = "invalid_grant"
	errUnsupportedGrantType    = "unsupported_grant_type"
	errLoginRequired           = "login_required"
	errConsentRequired         = "consent_required"

	oauth2ContextSession = "oauth2_context"
	authScheme           = "Bearer"

	authorizationCodeExp = 60 * 5
)

type Config struct {
	RBAC   rbac.RBAC
	Repo   repository.Repository
	Logger *zap.Logger

	// AccessTokenExp アクセストークンの有効時間(秒)
	AccessTokenExp int
	// IsRefreshEnabled リフレッシュトークンを発行するかどうか
	IsRefreshEnabled bool
}

func (h *Config) Setup(e *echo.Group) {
	e.GET("/authorize", h.AuthorizationEndpointHandler)
	e.POST("/authorize/decide", h.AuthorizationDecideHandler, middlewares.UserAuthenticate(h.Repo), middlewares.BlockBot(h.Repo))
	e.POST("/authorize", h.AuthorizationEndpointHandler)
	e.POST("/token", h.TokenEndpointHandler)
	e.POST("/revoke", h.RevokeTokenEndpointHandler)
}

// splitAndValidateScope スペース区切りのスコープ文字列を分解し、検証します
func (h *Config) splitAndValidateScope(str string) (model.AccessScopes, error) {
	scopes := model.AccessScopes{}
	scopes.FromString(str)
	if err := scopes.Validate(); err != nil {
		return nil, errors.New(errInvalidScope)
	}
	return scopes, nil
}

func (h *Config) requestContextLogger(c echo.Context) *zap.Logger {
	l, ok := c.Get(consts.KeyLogger).(*zap.Logger)
	if ok {
		return l
	}
	l = h.Logger.With(zap.String("logging.googleapis.com/trace", extension.GetTraceID(c)))
	c.Set(consts.KeyLogger, l)
	return l
}
