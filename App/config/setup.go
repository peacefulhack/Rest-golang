package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/redis/go-redis/v9"
	"time"
)

func NewMysql(conf *Config) (*sql.DB, error) {
	cfg := conf.Mysql
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database))
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(30 * time.Second)
	db.SetConnMaxIdleTime(20 * time.Second)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func NewRedis(conf *Config) *redis.Client {
	cfg := conf.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	return client
}
