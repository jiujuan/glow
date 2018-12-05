package config

type Config struct {
    Mysql  mysqlConfig  `toml:"mysql"`
    Server serverConfig `toml:"server"`
    Log    logConfig    `toml:"log"`
}

type mysqlConfig struct {
    Username string
    Password string
    DB       string
    Host     string
    Port     int
    Protocol string
    Charset  string
    MaxIdleConn int
    MaxOpenConn int
}

type serverConfig struct {
    Env  string
    Port int
}

type logConfig struct {
    File      string
    Formatter string
    Level     string
}

var BlogConfig Config
