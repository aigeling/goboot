package global

type DBConfig struct {
	DBType   string `json:"dbType"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbName"`

	//是否启用
	Enabled bool `json:"enabled"`

	//最大连接数
	MaxOpenConns int `json:"maxOpenConns"`
	//最大空闲连接数
	MaxIdleConn int `json:"maxIdleConn"`
	//连接最大空闲时间，单位：分钟
	ConnMaxIdleTime int `json:"connMaxIdleTime"`
}

// 最大连接数
func (o *DBConfig) GetMaxOpenConns() int {
	if o.MaxOpenConns == 0 {
		o.MaxOpenConns = 20
	}
	return o.MaxOpenConns
}

// 最大空闲连接数
func (o *DBConfig) GetMaxIdleConns() int {
	if o.MaxIdleConn == 0 {
		o.MaxIdleConn = 10
	}
	return o.MaxIdleConn
}

// 连接最大空闲时间，单位：分钟
func (o *DBConfig) GetConnMaxIdleTime() int {
	if o.ConnMaxIdleTime == 0 {
		o.ConnMaxIdleTime = 10
	}
	return o.ConnMaxIdleTime
}
