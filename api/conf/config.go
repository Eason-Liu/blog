package conf

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	gormmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
	"time"
)

func NewDefaultConfig() *Config {
	return &Config{
		App:   newDefaultApp(),
		Mysql: newDefaultMysql(),
	}
}

type Config struct {
	App   *app   `toml:"app" json:"app"`
	Mysql *mysql `toml:"mysql" json:"mysql"`
}

func (c *Config) String() string {
	jd, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}
	return string(jd)
}

func newDefaultApp() *app {
	return &app{
		Name: "blog",
		Http: newDefaultHttp(),
	}
}

type app struct {
	Name string `toml:"name" env:"APP_NAME" json:"name"`
	Http *http  `toml:"http" json:"http"`
}

func newDefaultHttp() *http {
	return &http{
		Host: "loadlhost",
		Port: "8080",
	}
}

type http struct {
	Host string `toml:"host" json:"host" env:"HTTP_HOST"`
	Port string `toml:"port" json:"port" env:"HTTP_PORT"`
}

func newDefaultMysql() *mysql {
	return &mysql{
		Host:     "localhost",
		Port:     "3306",
		Database: "blog",
		Username: "admin",
		Password: "123456",
	}
}

type mysql struct {
	Host        string `toml:"host" json:"host" env:"MYSQL_HOST"`
	Port        string `toml:"port" json:"port" env:"MYSQL_PORT"`
	Database    string `toml:"database" json:"database" env:"MYSQL_DATABASE"`
	Username    string `toml:"username" json:"username" env:"MYSQL_USERNAME"`
	Password    string `toml:"password" json:"password" env:"MYSQL_PASSWORD"`
	MaxOpenConn int    `toml:"max_open_conn" json:"max_open_conn" env:"MYSQL_MAX_OPEN_CONN"`
	MaxIdleConn int    `toml:"max_idle_conn" json:"max_idle_conn" env:"MYSQL_MAX_IDLE_CONN"`
	MaxLifeTime int    `toml:"max_life_time" json:"max_life_time" env:"MYSQL_MAX_LIFE_TIME"`
	MaxIdleTime int    `toml:"max_idle_time" json:"max_idle_time" env:"MYSQL_MAX_IDLE_TIME"`
	lock        sync.Mutex
	dbconn      *sql.DB
	orm         *gorm.DB
}

func (m *mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&multiStatements=true",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
	)
}

func (m *mysql) GetORMDB() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.orm == nil {
		db, err := gorm.Open(gormmysql.Open(m.Dsn()))
		if err != nil {
			panic(err)
		}
		m.orm = db
	}
	return m.orm
}

//???????????? ????????????
func (m *mysql) GetDB() *sql.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.dbconn == nil {
		conn, err := m.getDB()
		if err != nil {
			panic(err)
		}
		m.dbconn = conn
	}
	return m.dbconn
}

// ??????MySQL???????????????????????????
func (m *mysql) getDB() (*sql.DB, error) {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&multiStatements=true",
		m.Username,
		m.Password,
		m.Host,
		m.Port,
		m.Database,
	)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("connect to mysql<%s> error, %s", dsn, err.Error())
	}
	// ?????????????????????
	db.SetMaxOpenConns(m.MaxOpenConn)
	db.SetMaxIdleConns(m.MaxIdleConn)
	if m.MaxLifeTime != 0 {
		db.SetConnMaxLifetime(time.Second * time.Duration(m.MaxLifeTime))
	}
	if m.MaxIdleConn != 0 {
		db.SetConnMaxIdleTime(time.Second * time.Duration(m.MaxIdleTime))
	}

	//??????Ping???????????????MySQL??????????????????
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping mysql<%s> error, %s", dsn, err.Error())
	}
	return db, nil
}

//func (m *mysql) getOrmDB() (*gorm.DB, error) {
//	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&multiStatements=true",
//		m.Username,
//		m.Password,
//		m.Host,
//		m.Port,
//		m.Database,
//	)
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		return nil, err
//	}
//}
