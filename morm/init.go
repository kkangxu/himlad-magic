package morm

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

// DBSetting 数据库连接字符串属性
type DBSetting struct {
	Dialect         string // 数据库类型
	DbName          string // 数据库名称
	Host            string // IP地址
	User            string // 用户名
	Password        string // 密码
	Port            int    // 端口
	MaxOpenConns    int    // 最大连接数
	MaxIdleConns    int    // 最大IDLE链接数
	ConnMaxLifetime int    // 分钟
}

var globalDB *gorm.DB

// GetGlobalDB 获取全局DB
func GetGlobalDB(ctx context.Context) *gorm.DB {
	return globalDB.WithContext(ctx)
}

// GetDB 获取全局DB
func GetDB() *gorm.DB {
	return globalDB
}

// InitInstance 注册多个数据库连接
func InitInstance(confs []*DBSetting, plugins ...gorm.Plugin) error {
	if len(confs) == 0 {
		return fmt.Errorf("config is not corrrect, check it")
	}
	for _, c := range confs {
		fixConf(c)
	}
	// fist database
	var err error
	globalDB, err = connectDB(confs[0])
	if err != nil {
		return err
	}
	confs = confs[1:]

	// init resolver
	dr := &dbresolver.DBResolver{}
	for _, conf := range confs {
		dl, err := getDialector(conf)
		if err != nil {
			return err
		}
		dr = dr.Register(
			dbresolver.Config{
				Sources: []gorm.Dialector{dl},
				Policy:  dbresolver.RandomPolicy{},
			},
			conf.DbName,
		)
		dr.SetConnMaxLifetime(time.Duration(240) * time.Second)
		dr.SetMaxIdleConns(conf.MaxIdleConns)
		dr.SetMaxOpenConns(conf.MaxOpenConns)
		dr.SetConnMaxLifetime(time.Minute * time.Duration(conf.ConnMaxLifetime))
	}
	plugins = append(plugins, dr)

	// use plugin
	for _, p := range plugins {
		if err := globalDB.Use(p); err != nil {
			return err
		}
	}

	return nil
}

func fixConf(conf *DBSetting) {
	if conf.MaxIdleConns <= 0 {
		conf.MaxIdleConns = 5
	}
	if conf.MaxOpenConns <= 0 {
		conf.MaxOpenConns = 20
	}
	if conf.ConnMaxLifetime <= 0 {
		conf.ConnMaxLifetime = 5
	}
}

func getDialector(conf *DBSetting) (gorm.Dialector, error) {
	p := strconv.Itoa(conf.Port)
	switch conf.Dialect {
	case "mysql":
		return mysql.Open(
			fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4,utf8&parseTime=true&loc=Local",
				conf.User, conf.Password, conf.Host, p, conf.DbName)), nil
	case "mssql":
		return sqlserver.Open(
			fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&connection+timeout=600&encrypt=disable",
				conf.User, conf.Password, conf.Host, p, conf.DbName)), nil
	}
	return nil, fmt.Errorf("unsupported db：%s", conf.Dialect)
}

func connectDB(conf *DBSetting) (*gorm.DB, error) {
	dl, err := getDialector(conf)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(dl, &gorm.Config{
		AllowGlobalUpdate: true,
		PrepareStmt:       true,
		Logger:            &OrmLogger{loglevel: logger.Info},
	})
	if err != nil {
		return nil, fmt.Errorf("connnet db failed, err: %s", err.Error())
	}
	db.Set("gorm:table_options", "CHARSET=utf8mb4")
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Duration(240) * time.Second)
	sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
	sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Minute * time.Duration(conf.ConnMaxLifetime))
	db.Use(otelgorm.NewPlugin())
	return db, nil
}

// MockDB 模拟数据库
func MockDB() sqlmock.Sqlmock {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}))
	// 替换全局db
	globalDB = gormDB
	return mock
}
