package config

import (
	"github.com/jinzhu/configor"
	"log"
)

type Configuration struct {
	Payment      Payment
	MiniProgram  MiniProgram
	WeCom        WeCom
	OffiAccount  OffiAccount
	OpenPlatform OpenPlatform
	RobotChat    RobotChat
}

type Payment struct {
	AppID              string `required:"true" env:"app_id"`
	MchID              string `required:"true" env:"mch_id"`
	CertPath           string `required:"true" env:"wx_cert_path"`
	KeyPath            string `required:"true" env:"wx_key_path"`
	SerialNo           string `required:"true" env:"serial_no"`
	CertificateKeyPath string `required:"false" env:"certificate_key_path"`
	WechatPaySerial    string `required:"false" env:"wechat_pay_serial"`
	RSAPublicKeyPath   string `required:"false" env:"wx_rsa_public_key_path"`
	MchApiV3Key        string `env:"mch_api_v3_key"`
	Key                string `env:"key"`
	ResponseType       string `default:"map"`
	NotifyURL          string `env:"notify_url"`
	SubMchID           string `env:"sub_mch_id"`
	SubAppID           string `env:"sub_app_id"`
	HttpDebug          bool   `default:"false"`
	RedisAddr          string `env:"redis_addr"`
}

type MiniProgram struct {
	AppID         string `required:"true" env:"miniprogram_app_id"`
	Secret        string `required:"true" env:"miniprogram_secret"`
	RedisAddr     string `env:"redis_addr"`
	MessageToken  string `env:"message_token"`
	MessageAesKey string `env:"message_aes_key"`

	VirtualPayAppKey  string `env:"virtual_pay_app_key"`
	VirtualPayOfferID string `env:"virtual_pay_offer_id"`
}

type WeCom struct {
	CorpID          string `env:"corp_id"`
	AgentID         int    `env:"wecom_agent_id"`
	Secret          string `env:"wecom_secret"`
	MessageToken    string `env:"app_message_token"`
	MessageAesKey   string `env:"app_message_aes_key"`
	MessageCallback string `env:"app_message_callback_url"`
	OAuthCallback   string `env:"app_oauth_callback_url"`
	ContactSecret   string `env:"contact_secret"`
	ContactToken    string `env:"contact_token"`
	ContactAESKey   string `env:"contact_aes_key"`
	ContactCallback string `env:"contact_callback_url"`
	RedisAddr       string `env:"redis_addr"`
}

type OffiAccount struct {
	AppID         string `required:"true" env:"appid"`
	AppSecret     string `required:"true" env:"appsecret"`
	RedisAddr     string `env:"redis_addr"`
	MessageToken  string `env:"message_token"`
	MessageAesKey string `env:"message_aes_key"`
}

type OpenPlatform struct {
	AppID         string
	AppSecret     string
	MessageToken  string
	MessageAesKey string
	RedisAddr     string
}

type RobotChat struct {
	ArtBot  `env:"art_bot"`
	ChatBot `env:"chat_bot"`
}
type ArtBot struct {
	StableDiffusion `env:"stale_diffusion"`
}
type ChatBot struct {
	ChatGPT `env:"chat_gpt"`
}

type StableDiffusion struct {
	Token     string `env:"token"`
	BaseUrl   string `env:"base_url"`
	PrefixUri string `env:"prefix_uri"`
	Version   string `env:"version"`
	HttpDebug bool   `env:"http_debug"`
	ProxyURL  string `env:"proxy_url"`
}

type ChatGPT struct {
	OpenAPIKey   string `env:"open_api_key"`
	Organization string `env:"organization"`
	Model        string `env:"model"`
	HttpDebug    bool   `env:"http_debug"`
	BaseURL      string `env:"base_url"`
	APIType      string `env:"api_type"`
	APIVersion   string `env:"api_version"`
}

func configFiles() []string {
	return []string{"config.yml", "/etc/gotify/config.yml"}
}

// Get returns the configuration extracted from env variables or config file.
func Get() *Configuration {
	conf := new(Configuration)
	err := configor.New(&configor.Config{}).Load(conf, configFiles()...)
	if err != nil {
		log.Printf("%#v", conf)
		panic(err)
	}
	return conf
}
