package cmd

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"

	"g.hz.netease.com/horizon/core/http/api/v1/group"
	"g.hz.netease.com/horizon/core/http/health"
	"g.hz.netease.com/horizon/core/middleware/user"
	"g.hz.netease.com/horizon/lib/orm"
	"g.hz.netease.com/horizon/server/middleware"
	logmiddle "g.hz.netease.com/horizon/server/middleware/log"
	"g.hz.netease.com/horizon/server/middleware/requestid"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"

	ormmiddle "g.hz.netease.com/horizon/server/middleware/orm"
)

// Flags defines agent CLI flags.
type Flags struct {
	ConfigFile string
}

// ParseFlags parses agent CLI flags.
func ParseFlags() *Flags {
	var flags Flags

	flag.StringVar(
		&flags.ConfigFile, "config", "", "configuration file path")

	flag.Parse()
	return &flags
}

// Run runs the agent.
func Run(flags *Flags) {
	var config Config
	// load config
	data, err := ioutil.ReadFile(flags.ConfigFile)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(err)
	}

	// init db
	mySQLDB, err := orm.NewMySQLDB(&orm.MySQL{
		Host:     config.DBConfig.Host,
		Port:     config.DBConfig.Port,
		Username: config.DBConfig.Username,
		Password: config.DBConfig.Password,
		Database: config.DBConfig.Database,
	})
	if err != nil {
		panic(err)
	}

	// init controller
	var controller = group.NewController()

	// init server
	log.Printf("Server started")
	r := gin.New()

	// use middleware
	r.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health"),
		gin.Recovery(),
		requestid.Middleware(),
		logmiddle.Middleware(),
		ormmiddle.Middleware(mySQLDB),
		user.Middleware(config.OIDCConfig,
			middleware.MethodAndPathSkipper(http.MethodGet, regexp.MustCompile("^/health"))),
	)

	gin.ForceConsoleColor()

	// register routes
	health.RegisterRoutes(r)
	group.RegisterRoutes(r, controller)

	log.Fatal(r.Run(fmt.Sprintf(":%d", config.ServerConfig.Port)))
}
