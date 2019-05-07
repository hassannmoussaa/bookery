package main

import (
	"flag"
	"fmt"

	"github.com/hassannmoussaa/pill.go/antiCSRF"
	"github.com/hassannmoussaa/pill.go/auth"
	"github.com/hassannmoussaa/pill.go/fastmux"
	"github.com/hassannmoussaa/pill.go/fastmux/util"
	"github.com/hassannmoussaa/pill.go/mailer"
	"github.com/hassannmoussaa/pill.go/uploader"
	"github.com/hassannmoussaa/pill.go/uploader/gcs"
	"github.com/hassannmoussaa/bookery/pkg/apiCtrls"
	"github.com/hassannmoussaa/bookery/pkg/config"
	"github.com/hassannmoussaa/bookery/pkg/db"
	"github.com/hassannmoussaa/bookery/pkg/hooks"
	"github.com/hassannmoussaa/bookery/pkg/middlewares"
	"github.com/hassannmoussaa/bookery/pkg/models"
	"github.com/hassannmoussaa/bookery/pkg/textualContent"
	"github.com/hassannmoussaa/bookery/pkg/tmpls"
	"github.com/hassannmoussaa/bookery/pkg/viewsmodels"
	"github.com/hassannmoussaa/bookery/pkg/webCtrls"
	"github.com/valyala/fasthttp"
)

var (
	RootPath    = flag.String("root", ".", "the root patg of project")
	Environment = flag.String("env", "production", "the project environment")
)

func init() {
	flag.Parse()
	config.Init(*Environment, *RootPath)
	pgxConn := db.ConnectToDBs(config.GetPostgresDBConfig())
	models.Init(pgxConn, config.StaticFiles().URLPath, config.Uploader().URLPath)
	util.InitXServe(config.StaticFiles().Path, config.StaticFiles().URLPath, config.StaticFiles().FromCloud, config.Uploader().Path, config.Uploader().URLPath, config.Uploader().ToCloud, config.AppVersion())
	auth.Init(config.JWTPrivateKeyPath(), config.JWTPublicKeyPath(), config.DomainName(), config.OnlyHTTPS())
	middlewares.Init(config.WebHost(), config.WebServerAddress(), config.APIServerAddress(), config.OnlyHTTPS(), config.APIServerUsername(), config.APIServerPassword(), config.CORSAllowedOrigins())
	if config.Uploader().ToCloud {
		gcs.Init(config.GoogleCloudStorage().ProjectID, config.GoogleCloudStorage().ProjectNumber, config.GoogleCloudStorage().GetJsonKeyPath())
	}
	uploader.Init(config.Uploader().Path, config.Uploader().URLPath, config.Uploader().ToCloud, config.ImagesSizes())
	hooks.Init(config.WebHost())
	textualContent.Init(config.TxtDirPath())
	tmpls.Init(config.HTMLTemplatesPath(), config.StaticFiles().URLPath, config.Uploader().URLPath)
	antiCSRF.Init(32, config.DomainName(), config.OnlyHTTPS(), config.CSRFEncryptionKey())
	mailer.Init(config.SMTPHost(), config.SMTPPort(), config.AppName(), config.Email(), config.EmailPassword(), config.Uploader().ToCloud)
	apiCtrls.Init(config.WebHost(), config.DomainName())
	viewsmodels.Init(config.WebHost(), config.APIHost(), config.DomainName(), config.StaticFiles().URLPath, config.Uploader().URLPath, config.Environment(), config.AppVersion(), config.DefaultMetaImageURL())
}

func main() {
	webRouter := webCtrls.Register()
	apiRouter := apiCtrls.Register(config.APIPath())
	globalRouter := fastmux.NewGlobalRouter(webRouter, apiRouter, config.APIDomainName(), config.APIPath())
	s := &fasthttp.Server{
		Handler:            fasthttp.CompressHandlerLevel(globalRouter.ServeHTTP, 6),
		MaxRequestBodySize: 40 * 1024 * 1024,
	}
	if config.ServerAddress() != "" {
		fmt.Printf("%s server is runing on %s \n", config.AppName(), config.ServerAddress())
		s.ListenAndServe(config.ServerAddress())
	} else {
		fmt.Printf("%s server is runing on %s \n", config.AppName(), ":8080")
		s.ListenAndServe(":8080")
	}
}
