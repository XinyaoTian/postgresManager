package config

import "os"

// 用于配置数据库访问的信息 struct
type ConfigInfo struct {
	dbProtocol string
	dbUsername string
	dbPassword string
	dbIp string
	dbPort string
	dbName string
	sslMode string

}

// 具体的配置项信息
//var Configinfo  = ConfigInfo{
//	dbProtocol:"postgres",
//	dbUsername:"dbuser",
//	dbPassword:"docker",
//	dbIp:"174.137.53.55",
//	dbPort:"5432",
//	dbName:"testdb",
//	sslMode:"disable",
//}

// 具体的配置项信息，从 os 的环境变量中获取，方便 docker 的参数传入
var Configinfo  = ConfigInfo{
	dbProtocol: os.Getenv("POSTGRES_PROTOCOL"),
	dbUsername: os.Getenv("POSTGRES_USERNAME"),
	dbPassword: os.Getenv("POSTGRES_PASSWORD"),
	dbIp: os.Getenv("POSTGRES_IPADDR"),
	dbPort: os.Getenv("POSTGRES_PORT"),
	dbName: os.Getenv("POSTGRES_DBNAME"),
	sslMode: os.Getenv("POSTGRES_SSLMODE"),
}


func (c ConfigInfo) GetDbProtocol() string  {
	return c.dbProtocol
}

func (c ConfigInfo) GetUsername() string  {
	return c.dbUsername
}

func (c ConfigInfo) GetDbPassword() string  {
	return c.dbPassword
}

func (c ConfigInfo) GetDbIp() string  {
	return c.dbIp
}

func (c ConfigInfo) GetDbPort() string  {
	return c.dbPort
}

func (c ConfigInfo) GetDbName() string  {
	return c.dbName
}

func (c ConfigInfo) GetSslMode() string  {
	return c.sslMode
}
