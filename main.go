package main

import (
    "flag"
    "os"
    
    "github.com/sirupsen/logrus"
    
    "github.com/jmshin92/scraper/router"
)

var (
    FlagPort string
    FlagHelp bool
)

func init() {
    parseFlag()
}

func parseFlag() {
    flag.StringVar(&FlagPort, "p", "9090", "port")
    flag.BoolVar(&FlagHelp, "h", false, "help")
    flag.Parse()
    
    
    if FlagHelp {
        flag.Usage()
        os.Exit(0)
    }
    
    if len(FlagPort) == 0 {
        logrus.Error("Port is nil")
        flag.Usage()
        os.Exit(-1)
    } else {
        FlagPort = ":" + FlagPort
    }
}

func main() {
    r := router.Init()
    logrus.Info(FlagPort)
    if err := r.Run(FlagPort); err != nil {
        logrus.Fatal(err)
    }
}