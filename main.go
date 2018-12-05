package main

import (
    "fmt"
    "flag"
    "os"
    "github.com/BurntSushi/toml"
    log "github.com/Sirupsen/logrus"
    "github.com/gin-gonic/gin"
    "github.com/jiujuanfeng/yunheblog/db"
    "github.com/jiujuanfeng/yunheblog/config"
    "github.com/jiujuanfeng/yunheblog/router"
)

func initLog() {
    switch config.BlogConfig.Log.Level {
    case "debug":
        log.SetLevel(log.DebugLevel)
        log.Infof("Set log level: %s", "debug")
    case "info":
        log.SetLevel(log.InfoLevel)
        log.Infof("Set log level: %s", "info")
    case "warn":
        log.SetLevel(log.WarnLevel)
        log.Infof("Set log level: %s", "warn")
    case "error":
        log.SetLevel(log.ErrorLevel)
        log.Infof("Set log level: %s", "error")
    case "fatal":
        log.SetLevel(log.FatalLevel)
        log.Infof("Set log level: %s", "fatal")
    case "panic":
        log.SetLevel(log.PanicLevel)
        log.Infof("Set log level: %s", "panic")
    default:
        log.SetLevel(log.DebugLevel)
        log.Infof("Set log level: %s", "debug")
    }

    log.Infof("Set log formatter: %s", config.BlogConfig.Log.Formatter)
    if config.BlogConfig.Log.Formatter == "json" {
        log.SetFormatter(&log.JSONFormatter{})
    } else if config.BlogConfig.Log.Formatter == "text" {
        log.SetFormatter(&log.TextFormatter{})
    }

    log.Infof("Open log file: %s", config.BlogConfig.Log.File)
    logFile, err := os.OpenFile(config.BlogConfig.Log.File, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        log.Errorf("Failed to open log file: %s, err: %s", config.BlogConfig.Log.File, err.Error())
        os.Exit(1)
    }

    log.Infof("Set log output: %s", config.BlogConfig.Log.File)
    log.SetOutput(logFile)

}

func main() {
    var configFile, mode string
    flag.StringVar(&configFile, "c", "./config/conf.toml", "config file")
    flag.StringVar(&mode, "m", "localhost", "localhost to read file")

    log.Infof("Read config file: %s ", configFile)
    if _, err := toml.DecodeFile(configFile, &config.BlogConfig); err != nil {
        log.Errorf("Failed to load config file: %s err: %s ", configFile, err.Error())
    }

    initLog()
    db.InitDB()
    defer db.DB.Close()

    app := gin.New()
    router.Route(app)

    port := config.BlogConfig.Server.Port
    log.Infof("server port: %d ", port)
    app.Run(":" + fmt.Sprintf("%d", port))
}
