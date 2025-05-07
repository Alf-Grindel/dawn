package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type server struct {
	Version string
	Name    string
}
type mysql struct {
	Addr     string
	Database string
	Username string
	Password string
}
type config struct {
	Server server
	MySQL  mysql
}

var (
	Server *server
	MySQL  *mysql

	runtimeViper = viper.New()
)

func Init() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dir := getPath(path)
	if dir == "" {
		log.Fatalln("config.Init: could not find config.yaml")
	}
	runtimeViper.SetConfigName("config")
	runtimeViper.SetConfigType("yml")
	runtimeViper.AddConfigPath(dir)
	if err = runtimeViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalln("config.Init: could not find config files")
		} else {
			log.Fatalln("config.Init: read config failed, ", err)
		}
	}
	configMapping()
	runtimeViper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("config: notice config changed, %v\n", in.String())
		configMapping()
	})
	runtimeViper.WatchConfig()
}

func configMapping() {
	c := &config{}
	if err := runtimeViper.Unmarshal(&c); err != nil {
		log.Fatalln("config.configMapping: config unmarshal failed, ", err)
	}
	Server = &c.Server
	MySQL = &c.MySQL
}

func getPath(path string) string {
	dir := path
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}
