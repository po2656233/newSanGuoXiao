package conf

import (
	"encoding/json"
	"github.com/po2656233/goleaf/log"
	"os"
)

var Server struct {
	//CarryMainSubID bool `json:"carryMainSubID"` //传输的时候 是否需要携带mainID 或 subID
	MaxConnNum  int `json:"maxConnNum"`
	ConsolePort int
	RedisDBNum  int    `json:"redisDBNum"`
	LogLevel    string `json:"logLevel"`
	LogPath     string `json:"logPath"`
	KCPAddr     string `json:"kcpAddr"`
	TCPAddr     string `json:"tcpAddr"`
	WSAddr      string `json:"wsAddr"`
	RedisAddr   string `json:"redisAddr"`
	RedisPSW    string `json:"redisPSW"`
	CertFile    string
	KeyFile     string

	ProfilePath string `json:"profilePath"`

	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBAddress  string `json:"dbAddress"`
	DBPort     string `json:"dbPort"`
	DBName     string `json:"dbName"`

	CenterWeb  string `json:"centerWeb"`
	CenterUser string `json:"centerUser"`
	CenterPSW  string `json:"centerPSW"`
}

func InitJson() {
	data, err := os.ReadFile(ServerJsonPath)
	if err != nil {
		log.Fatal("%v", err)
	}
	err = json.Unmarshal(data, &Server)
	if err != nil {
		log.Fatal("%v", err)
	}
}

func updateJsonConf(data string) (err error) {
	temp := Server
	err = json.Unmarshal([]byte(data), &temp)
	if err != nil {
		log.Error("updateJsonConf data:[%v] err:%v", data, err)
		return
	}
	Server = temp
	return
}
