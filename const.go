package httputil

const (
	AuthenticationSchemeBearer = "Bearer"
	ErrorBadJson               = "Bad JSON format"
	ErrorBodyRead              = "Error occured reading the body"
	HeaderAllow                = "Allow"
	HeaderAuthenticate         = "WWW-Authenticate"
	HeaderAuthorization        = "Authorization"
	HeaderCacheControl         = "Cache-Control"
	HeaderContentType          = "Content-Type"
	HeaderForwardedHost        = "x-forwarded-host"
	HeaderForwardedProtocol    = "x-forwarded-proto"
	HeaderHost                 = "Host"
	HeaderLocation             = "Location"
	MimeTypeJson               = "application/json"
	MimeTypeUrlEncodedForm     = "application/x-www-form-urlencoded"
)
