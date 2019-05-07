package middlewares

var webHost, webAddress, apiAddress, apiServerUsername, apiServerPassword string
var onlyHttps bool
var corsAllowedOrigins []string

func Init(WebHost string, WebAddress string, ApiAddress string, OnlyHttps bool, ApiServerUsername string, ApiServerPassword string, CorsAllowedOrigins []string) {
	webHost = WebHost
	onlyHttps = OnlyHttps
	webAddress = WebAddress
	apiAddress = ApiAddress
	apiServerUsername = ApiServerUsername
	apiServerPassword = ApiServerPassword
	corsAllowedOrigins = CorsAllowedOrigins
}
