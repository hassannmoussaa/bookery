package config

import (
	"bufio"
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/hassannmoussaa/pill.go/clean"
)

var cfg *Config

type Config struct {
	Environment       string `json:"environment"`
	AppName           string `json:"app_name"`
	AppVersion        string `json:"app_version"`
	DomainName        string `json:"domain_name"`
	APIDomainName     string `json:"api_domain_name"`
	OnlyHTTPS         bool   `json:"only_https"`
	FacebookAppID     string `json:"facebook_app_id"`
	CSRFEncryptionKey string `json:"csrf_encryption_key"`

	//Database Configurations
	PostgresDBConfig *PostgresDBConfig `json:"postgres_db_config"`
	RethinkDBConfig  *RethinkDBConfig  `json:"rethink_db_config"`

	//Paths
	RootPath          string       `json:"root_path"`
	ResourcesDirPath  string       `json:"resources_dir_path"`
	JWTPrivateKeyPath string       `json:"jwt_private_key_path"`
	JWTPublicKeyPath  string       `json:"jwt_public_key_path"`
	Uploader          *uploader    `json:"uploader"`
	StaticFiles       *staticFiles `json:"static_files"`
	HTMLTemplatesPath string       `json:"html_templates_path"`
	TxtDirPath        string       `json:"txt_dir_path"`

	//Servers
	WebServerAddress string `json:"web_server_address"`
	APIServerAddress string `json:"api_server_address"`
	ServerAddress    string `json:"server_address"`
	APIPath          string `json:"api_path"`

	//ACCOUNT KIT
	AccountKit         *accountKit         `json:"account_kit"`
	GoogleCloudStorage *googleCloudStorage `json:"google_cloud_storage"`

	//Hosts
	WebHost       string `json:"web_host"`
	RealtimeHost  string `json:"realtime_host"`
	WebsocketHost string `json:"websocket_host"`
	APIHost       string `json:"api_host"`

	DefaultMetaImageURL string `json:"default_meta_image_url"`

	//EMAIL
	SMTPHost      string `json:"smtp_host"`
	SMTPPort      int    `json:"smtp_port"`
	Email         string `json:"email"`
	EmailPassword string `json:"email_password"`

	ImagesSizes map[string]map[string][]uint `json:"images_sizes"`

	//API Server
	APIServerUsername string `json:"api_server_username"`
	APIServerPassword string `json:"api_server_password"`

	//CORS
	CORSAllowedOrigins []string `json:"cors_allowed_origins"`

	//GOOGLE MAPSS API KEY
	GMapsAPIKey string `json:"google_maps_api_key"`

	FCMServerKey string `json:"fcm_server_key"`
}

type accountKit struct {
	AppSecret            string `json:"app_secret"`
	Version              string `json:"version"`
	MeEndpointBaseURL    string `json:"me_endpoint_base_url"`
	TokenExchangeBaseURL string `json:"token_exchange_base_url"`
}

type uploader struct {
	Path    string `json:"path"`
	URLPath string `json:"url_path"`
	ToCloud bool   `json:"to_cloud"`
}

func Uploader() *uploader {
	return cfg.Uploader
}

type staticFiles struct {
	Path      string `json:"path"`
	URLPath   string `json:"url_path"`
	FromCloud bool   `json:"from_cloud"`
}

func StaticFiles() *staticFiles {
	return cfg.StaticFiles
}

type RethinkDBConfig struct {
	Host     string `json:"host,omitempty"`
	Database string `json:"database,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	AuthKey  string `json:"auth_key,omitempty"`
	Cert     string `json:"cert,omitempty"`
}

func (this *RethinkDBConfig) GetCertPath() string {
	if this.Cert != "" {
		return path.Join(ResourcesDirPath(), this.Cert)
	}
	return ""
}

type PostgresDBConfig struct {
	Host           string `json:"host,omitempty"`
	Database       string `json:"database,omitempty"`
	User           string `json:"user,omitempty"`
	Password       string `json:"password,omitempty"`
	Port           uint16 `json:"port,omitempty"`
	MaxConnections int    `json:"max_connections,omitempty"`
	Cert           string `json:"cert,omitempty"`
}

func (this *PostgresDBConfig) GetCertPath() string {
	if this.Cert != "" {
		return path.Join(ResourcesDirPath(), this.Cert)
	}
	return ""
}

func GetRethinkDBConfig() *RethinkDBConfig {
	return cfg.RethinkDBConfig
}

func GetPostgresDBConfig() *PostgresDBConfig {
	return cfg.PostgresDBConfig
}

type googleCloudStorage struct {
	ProjectID     string `json:"project_id,omitempty"`
	ProjectNumber string `json:"project_number,omitempty"`
	JsonKeyFile   string `json:"json_key_file,omitempty"`
}

func (this *googleCloudStorage) GetJsonKeyPath() string {
	if this.JsonKeyFile != "" {
		return ResourcesDirPath() + "/" + this.JsonKeyFile
	}
	return ""
}

func GoogleCloudStorage() *googleCloudStorage {
	return cfg.GoogleCloudStorage
}

func Init(environment string, rootPath string) {
	cfg = &Config{}
	cfg.AccountKit = &accountKit{}
	cfg.RethinkDBConfig = &RethinkDBConfig{}
	cfg.PostgresDBConfig = &PostgresDBConfig{}
	cfg.GoogleCloudStorage = &googleCloudStorage{}
	cfg.Uploader = &uploader{}
	cfg.StaticFiles = &staticFiles{}

	//parse config file
	configFile, err := os.Open(path.Join(rootPath, environment+".json"))
	defer configFile.Close()
	if err == nil {
		configFileInfo, _ := configFile.Stat()
		size := configFileInfo.Size()
		data := make([]byte, size)
		reader := bufio.NewReader(configFile)
		readLen, _ := reader.Read(data)
		data = data[:readLen]
		err = json.Unmarshal(data, cfg)
		if err != nil {
			clean.Error(err)
			os.Exit(1)
		}
	} else {
		clean.Error(err)
		os.Exit(1)
	}
	environment = strings.ToLower(strings.TrimSpace(environment))
	if strings.HasPrefix(environment, "development") {
		cfg.Environment = "development"
	} else {
		cfg.Environment = "production"
	}
	cfg.RootPath = rootPath
}

func Environment() string {
	return cfg.Environment
}
func AppName() string {
	return cfg.AppName
}
func AppVersion() string {
	return cfg.AppVersion
}
func DomainName() string {
	return cfg.DomainName
}
func APIDomainName() string {
	return cfg.APIDomainName
}
func APIPath() string {
	return cfg.APIPath
}
func OnlyHTTPS() bool {
	return cfg.OnlyHTTPS
}
func CSRFEncryptionKey() string {
	return cfg.CSRFEncryptionKey
}
func FacebookAppID() string {
	return cfg.FacebookAppID
}

//ACCOUNT KIT
func AccountKit() *accountKit {
	return cfg.AccountKit
}

//Paths
func RootPath() string {
	return cfg.RootPath
}
func ResourcesDirPath() string {
	return path.Join(RootPath(), cfg.ResourcesDirPath)
}
func JWTPrivateKeyPath() string {
	return path.Join(ResourcesDirPath(), cfg.JWTPrivateKeyPath)
}
func JWTPublicKeyPath() string {
	return path.Join(ResourcesDirPath(), cfg.JWTPublicKeyPath)
}
func HTMLTemplatesPath() string {
	return path.Join(ResourcesDirPath(), cfg.HTMLTemplatesPath)
}

func TxtDirPath() string {
	return path.Join(ResourcesDirPath(), cfg.TxtDirPath)
}

//Servers
func WebServerAddress() string {
	if cfg.WebServerAddress != "" {
		return cfg.WebServerAddress
	}
	return ServerAddress()
}
func APIServerAddress() string {
	if cfg.APIServerAddress != "" {
		return cfg.APIServerAddress
	}
	return ServerAddress()
}
func ServerAddress() string {
	return cfg.ServerAddress
}

//Hosts
func WebHost() string {
	return cfg.WebHost
}
func RealtimeHost() string {
	return cfg.RealtimeHost
}
func WebsocketHost() string {
	return cfg.WebsocketHost
}
func APIHost() string {
	return cfg.APIHost
}

func DefaultMetaImageURL() string {
	return cfg.DefaultMetaImageURL
}

//EMAIL
func SMTPHost() string {
	return cfg.SMTPHost
}
func SMTPPort() int {
	return cfg.SMTPPort
}
func Email() string {
	return cfg.Email
}
func EmailPassword() string {
	return cfg.EmailPassword
}

func ImagesSizes() map[string]map[string][]uint {
	return cfg.ImagesSizes
}

func Get() *Config {
	return cfg
}

//API Server
func APIServerUsername() string {
	return cfg.APIServerUsername
}
func APIServerPassword() string {
	return cfg.APIServerPassword
}

//CORS
func CORSAllowedOrigins() []string {
	return cfg.CORSAllowedOrigins
}

//GOOGLE MAPS API KEY
func GMapsAPIKey() string {
	return cfg.GMapsAPIKey
}

func FCMServerKey() string {
	return cfg.FCMServerKey
}
