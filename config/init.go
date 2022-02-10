package config

func init() {
	Server = (&server{}).Load("config.ini").Init()
	Mysql = (&mysql{}).Load("config.ini").Init()
}
