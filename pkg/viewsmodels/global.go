package viewsmodels

import "github.com/hassannmoussaa/bookery/pkg/textualContent"

var global *Global

type Global struct {
	WebHost            string
	APIHost            string
	WebDomain          string
	Environment        string
	AppVersion         string
	StaticFilesURLPath string
	UploadsURLPath     string
	TextualContent     *textualContent.TextualContent
}

func initGlobal() {
	global = &Global{}
	global.WebDomain = webDomain
	global.APIHost = apiHost
	global.WebHost = webHost
	global.Environment = environment
	global.AppVersion = appVersion
	global.StaticFilesURLPath = staticFilesURLPath
	global.UploadsURLPath = uploadsURLPath
	global.TextualContent = textualContent.Get()
}
