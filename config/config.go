package config

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
var Configinfo  = ConfigInfo{
	dbProtocol:"postgres",
	dbUsername:"dbuser",
	dbPassword:"docker",
	dbIp:"174.137.53.55",
	dbPort:"5432",
	dbName:"testdb",
	sslMode:"disable",
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
