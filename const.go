package httputil

const (
	HeaderContentType = "Content-Type"
	HeaderLocation = "Location"
	HeaderForwardedProtocol = "x-forwarded-proto"
	HeaderForwardedHost = "x-forwarded-host"
	HeaderHost = "Host"
	MimeTypeUrlEncodedForm = "application/x-www-form-urlencoded"
	MimeTypeJson = "application/json"
	ErrorBodyRead = "Error occured reading the body"
	ErrorBadJson = "Bad JSON format"
)