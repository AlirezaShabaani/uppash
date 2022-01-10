package config

import (
	"github.com/spf13/viper"
	"log"
)

type Configs struct {
	MinioEndPoint  string
	MinioAccessKey string
	MinioSecretKey string
	MinioUseSSL bool
	NatsUrl string
}
func LoadConfig()(conf Configs)  {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalf("can't load config data err :%v", err.Error())
	}
	conf.MinioEndPoint = v.GetString("MINIO_ENDPOINT")
	conf.MinioAccessKey = v.GetString("MINIO_ACCESSKEY")
	conf.MinioSecretKey = v.GetString("MINIO_SECRETKEY")
	conf.MinioUseSSL = v.GetBool("MINIO_USESSL")
	conf.NatsUrl = v.GetString("NATS_URL")
	return
}
