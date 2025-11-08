package main

import (
	"github.com/HanThamarat/GO-Bucket-Service/packages/conf"
	"github.com/HanThamarat/GO-Bucket-Service/packages/initialize"
	"github.com/HanThamarat/GO-Bucket-Service/server"
)

// @title THE AJ EXPLORER BUCKET
// @version 1.0
// @description BUCKET | Doc by Swagger.
// @contact.name Developer Team
// @contact.url https://technexify.site
// @contact.email technexify@outlook.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3004
// @BasePath /api
// @schemes http https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Use "Bearer <API_KEY>"
func main() {
	config := conf.GetConfig();
	initialize.FolderInitialize();
	server.NewFiberServer(config).Start();
}