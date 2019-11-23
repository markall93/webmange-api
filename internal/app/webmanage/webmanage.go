package webmanage

import (
	"github.com/gin-gonic/gin"
	webconfig "github.com/mengjayxc/webmanage-api/internal/app/webmanage/config"
	"github.com/mengjayxc/webmanage-api/internal/app/webmanage/middleware"
	"github.com/mengjayxc/webmanage-api/internal/app/webmanage/routers"
	"github.com/mengjayxc/webmanage-api/internal/pkg/config"
	"github.com/mengjayxc/webmanage-api/internal/pkg/models"
	"github.com/mengjayxc/webmanage-api/internal/pkg/models/sys"
	"github.com/mengjayxc/webmanage-api/internal/pkg/models/web"
	"github.com/mengjayxc/webmanage-api/pkg/convert"
	"github.com/mengjayxc/webmanage-api/pkg/logger"
	"net/http"
	"os"
	"time"
)

// create all tables
var createTableStatements = map[string]interface{} {
	"admins":       sys.Admins{},
	"accounts":     web.Accounts{},
	"user_infos":   web.UserInfos{},
}

func Init(configPath string) {
	configDir := os.ExpandEnv("$PWD/config/config.yaml")
	if configPath == ""{
		configPath = configDir
	}
	// 加载配置
	config, err := webconfig.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}

	logPath := os.ExpandEnv("$PWD/log/web.log")
	logger.InitLog("debug", logPath)

	initDB(config)
	createTables(createTableStatements)
	initWeb(config)
	logger.Debug(config.Web.Domain + " 站点已启动...")
	
}

func initDB(config *config.Config){
	models.InitDB(config)
	logger.Debug("initDB completed...")
}

func initWeb(config *config.Config) {
	gin.SetMode(gin.DebugMode) //调试环境/生产环境
	app := gin.New()
	app.NoRoute(middleware.NoRouteHandler())

	// 注册路由
	routers.RegisterRouter(app)
	go initHTTPServer(config,app)
	//app.Run("172.168.0.44:8087")
}

func createTables(tableStmts map[string]interface{}){
	models.CreateTables(tableStmts)
}

// InitHTTPServer 初始化http服务
func initHTTPServer(config *config.Config,handler http.Handler) {
	srv := &http.Server{
		Addr:         config.Web.Host + ":"+convert.ToString(config.Web.Port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	srv.ListenAndServe()
}






