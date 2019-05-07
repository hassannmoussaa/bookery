package tmpls

import "github.com/hassannmoussaa/pill.go/templates"

var (
	uploadsURLPath     string
	staticFilesURLPath string
)

func init() {
	templates.AddTmplFunc("GetStaticFileURLPath", GetStaticFileURLPath)
	templates.AddTmplFunc("GetUploadURLPath", GetUploadURLPath)
}

func Init(templatesPath string, StaticFilesURLPath string, UploadsURLPath string) {
	staticFilesURLPath = StaticFilesURLPath
	uploadsURLPath = UploadsURLPath
	templates.Setup(templatesPath)
}
