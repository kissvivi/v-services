package config

import (
	"github.com/spf13/viper"
)

//var (
//	GVA_DB              *gorm.DB
//	GVA_INFLUXDB        influxdb.Client
//	GVA_MQTT            mqtt.Client //mqtt客户端
//	GVA_REDIS           *redis.Client
//	GVA_CONFIG          config.Server
//	GVA_VP              *viper.Viper
//	GVA_WEBSOCKET       *websocket.Conn
//	GVA_WEBSOCKET_ALARM *websocket.Conn
//	//GVA_LOG    *oplogging.Logger
//	GVA_LOG   *zap.Logger
//	GVA_Timer timer.Timer = timer.NewTimerTask()
//)

type Config struct {
	Application Application `json:"application"`
	Database    Database    `json:"database"`
}

type Application struct {
	Name       string     `json:"name"`
	Version    string     `json:"version"`
	HttpServer HttpServer `json:"httpserver"`
}

type HttpServer struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	Mode         string `json:"mode"`
	Readtimeout  int    `json:"readtimeout"` //单位分钟
	Writetimeout int    `json:"writetimeout"`
}

type Mysql struct {
	//Host     string `json:"host"`
	//Port     string `json:"port"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type Redis struct {
	//Host     string `json:"host"`
	//Port     string `json:"port"`
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type Database struct {
	Mysql Mysql `json:"mysql"`
	Redis Redis `json:"redis"`

	//Type     string `json:"type"`
	//Host     string `json:"host"`
	//Port     string `json:"port"`
	//Username string `json:"username"`
	//Password string `json:"password"`
	//Dbname   string `json:"dbname"`
}

func InitConfig() (*Config, error) {
	v := viper.New()
	v.AddConfigPath("config")
	v.SetConfigName("setting")
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var cfg Config
	err = v.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

//func Gorm(cfg *Database) *gorm.DB {
//	//todo 读写分离
//	switch cfg.Type {
//	case "mysql":
//		return DSN(cfg)
//	default:
//		return DSN(cfg)
//	}
//}
//
//func DSN(cfg *Database) *gorm.DB {
//	var dsn string
//	dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
//		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Dbname)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		fmt.Println("database connect failed")
//	}
//	fmt.Println("数据库success")
//	return db
//}
