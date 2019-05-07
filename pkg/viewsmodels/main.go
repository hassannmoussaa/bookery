package viewsmodels

var webHost, apiHost, webDomain, environment, appVersion, staticFilesURLPath, uploadsURLPath, defaultPageImage string

func Init(WebHost, APIHost, WebDomain, StaticFilesURLPath, UploadsURLPath, Environment, AppVersion, DefaultPageImage string) {
	webHost = WebHost
	apiHost = APIHost
	webDomain = WebDomain
	defaultPageImage = DefaultPageImage
	environment = Environment
	appVersion = AppVersion
	staticFilesURLPath = StaticFilesURLPath
	uploadsURLPath = UploadsURLPath
	initGlobal()
}

type CSRFToken struct {
	Value       string
	HiddenField string
}