package main

import (
	"os"

	"github.com/greficsmurf/go-ios/restapi/api"
	_ "github.com/greficsmurf/go-ios/restapi/docs"
	log "github.com/sirupsen/logrus"
)

// @title           Go-iOS API
// @version         0.01
// @description     Exposes go-ios features as REST API calls.
// @termsOfService  https://github.com/greficsmurf/go-ios

// @contact.name   Daniel Paulus
// @contact.url    https://github.com/greficsmurf/go-ios

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth
func main() {
	log.WithFields(log.Fields{"args": os.Args, "version": api.GetVersion()}).Infof("starting go-iOS-API")
	api.Main()
}
