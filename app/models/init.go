package models

import (
	"fmt"
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/revel"
	"github.com/revel/config"
)

var (
	engineRead *xorm.Engine
	engineWrite *xorm.Engine
	Smtp   SmtpType
)

type SmtpType struct {
	Username string
	Password string
	Host     string
	Address  string
	From     string
}

func init() {
	// 这函数表示在revel加载之前运行自定义（Init）方法
	revel.OnAppStart(Init)
}

func Init() {
	//获取配置my.conf文件中的值
	c, err := config.ReadDefault(revel.BasePath + "/conf/my.conf")
	if err != nil {
		revel.ERROR.Panicln(err)
	}
	driverRead, _ := c.String("databaseRead", "db.driver.read")
	dbnameRead, _ := c.String("databaseRead", "db.dbname.read")
	userRead, _ := c.String("databaseRead", "db.user.read")
	passwordRead, _ := c.String("databaseRead", "db.password.read")
	hostRead, _ := c.String("databaseRead", "db.host.read")

	// 用于读取操作数据库的数据
	paramsRead := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", userRead, passwordRead, hostRead, dbnameRead)

	engineRead, err = xorm.NewEngine(driverRead, paramsRead)
	if err != nil {
		revel.ERROR.Panicln(err)
	}

	driverWrite, _ := c.String("databaseWrite", "db.driver.write")
	dbnameWrite, _ := c.String("databaseWrite", "db.dbname.write")
	userWrite, _ := c.String("databaseWrite", "db.user.write")
	passwordWrite, _ := c.String("databaseWrite", "db.password.write")
	hostWrite, _ := c.String("databaseWrite", "db.host.write")
	paramsWrite := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", userWrite, passwordWrite, hostWrite, dbnameWrite)

	// 用于把数据写入数据库
	engineWrite, err = xorm.NewEngine(driverWrite, paramsWrite)
	if err != nil {
		revel.ERROR.Panicln(err)
	}
}

func GetEngineRead() *xorm.Engine {
	return engineRead
}

func GetEngineWrite() *xorm.Engine {
	return engineWrite
}
