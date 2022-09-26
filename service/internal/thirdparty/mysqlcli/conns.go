package mysqlcli

import (
	"context"
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	TRANSACTION_CTX_KEY       = "TRANSACTION_CTX_KEY"
	TRANSACTION_CTX_ISSET_KEY = "TRANSACTION_CTX_ISSET_KEY"
)

var (
	once sync.Once
	self *DBClient
)

type IMySQLClient interface {
	Session() *gorm.DB
	Close() error
}

type DBClient struct {
	in     digIn
	Client *gorm.DB
}

func initWithConfig(in digIn) IMySQLClient {
	connect := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		in.OpsConf.GetMySQLServiceConfig().Username,
		in.OpsConf.GetMySQLServiceConfig().Password,
		in.OpsConf.GetMySQLServiceConfig().Address,
		in.OpsConf.GetMySQLServiceConfig().Database,
	)

	var err error
	db, err := gorm.Open(mysql.Open(connect))
	if err != nil {
		panic(fmt.Sprintf("conn: %s err: %v", connect, err))
	}

	if in.AppConf.GetMySQLConfig().LogMode {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	in.SysLogger.Info(context.Background(), fmt.Sprintf("Database [%s] Connect success", in.OpsConf.GetMySQLServiceConfig().Database))
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(in.AppConf.GetMySQLConfig().MaxIdle)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(in.AppConf.GetMySQLConfig().MaxOpen)
	// SetConnMaxLifetime sets the maximum amount of timeUtil a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Duration(in.AppConf.GetMySQLConfig().ConnMaxLifeSec) * time.Second)

	return &DBClient{in: in, Client: db}
}

// Session creates an original gorm.DB session.
func (c *DBClient) Session() *gorm.DB {
	db := c.Client.Session(&gorm.Session{})
	return db
}

// Session creates an original gorm.DB session.
func (c *DBClient) Close() error {
	client, err := c.Client.DB()
	if err != nil {
		return err
	}

	if err := client.Close(); err != nil {
		return err
	}

	return nil
}

//NewTestSession only for unit-test
func NewMockSession() *gorm.DB {
	temp := DBClient{}
	return temp.Client.Session(&gorm.Session{})
}

type mockDBClient struct{}

func (m mockDBClient) Session() *gorm.DB {
	return &gorm.DB{}
}

func (m mockDBClient) Close() error {
	return nil
}

func NewMockClient() *mockDBClient {
	return &mockDBClient{}
}
