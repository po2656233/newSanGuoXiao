package redis

/*
# -*- coding: utf-8 -*-
# @Time : 2020/4/30 10:12
# @Author : Pitter
# @File : redisHandle.go
# @Software: GoLand
*/

import (
	"errors"
	"sanguoxiao/internal/component/jettengame/base"
	"sanguoxiao/internal/component/jettengame/conf"
	"sync"
	"time"

	"github.com/go-redis/redis"
	"github.com/po2656233/goleaf/log"
)

type RedisMan struct {
	//sync.RWMutex
	//user     string
	password string
	address  string
	dbNum    int
	DB       *RedisClient //本地使用
	Client   *redis.Client
}
type RedisClient struct {
	essence *redis.Client
}

var once sync.Once
var redisObj *RedisMan

func RedisHandle() *RedisClient {
	once.Do(func() {
		redisObj = &RedisMan{
			address:  conf.Server.RedisAddr,
			password: conf.Server.RedisPSW,
			dbNum:    conf.Server.RedisDBNum,
			DB:       nil,
		}
		redisObj.Init()
	})

	return redisObj.DB
}
func (self *RedisMan) Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     self.address,
		Password: self.password, // no password set
		DB:       self.dbNum,    // use default DB
	})
	pong, err := client.Ping().Result()
	log.Debug(":%v :%v", pong, err)
	self.DB = &RedisClient{
		essence: client,
	}
	if nil != err {
		//self.DB.essence = nil //允许不使用redis
		log.Fatal("cann't to connect redis:%v  error:%v", self.address, err)
	}
}

func (self *RedisClient) Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	if self.essence == nil {
		return nil
	}
	return self.essence.Set(key, value, expiration)
}

func (self *RedisClient) Get(key string) *redis.StringCmd {
	if self.essence == nil {
		return nil
	}
	return self.essence.Get(key)
}
func (self *RedisClient) IncrBy(key string, value int64) *redis.IntCmd {
	if self.essence == nil {
		return nil
	}
	return self.essence.IncrBy(key, value)
}

func (self *RedisClient) ZAdd(key string, score float64, value interface{}) *redis.IntCmd {
	if self.essence == nil {
		return nil
	}
	return self.essence.ZAdd(key, redis.Z{
		Score:  score,
		Member: value,
	})
}
func (self *RedisClient) ZIncrBy(key, member string) (float64, error) {
	if self.essence == nil {
		return 0, errors.New(base.StatusText[base.Redis02])
	}
	return self.essence.ZIncrBy(key, 1, member).Result()
}
func (self *RedisClient) ZScore(key, member string) (float64, error) {
	if self.essence == nil {
		return 0, errors.New(base.StatusText[base.Redis02])
	}
	return self.essence.ZScore(key, member).Result()
}

func (self *RedisClient) ZMaxScore(key string) (float64, interface{}, error) {
	if self.essence == nil {
		return 0, nil, errors.New(base.StatusText[base.Redis02])
	}
	vals, err := self.essence.ZRevRangeWithScores(key, 0, 0).Result()
	if err != nil {
		log.Error("ZMaxScore key:%v error:%v", key, err)
		return 0, nil, err
	}
	if 0 == len(vals) {
		return 0, nil, errors.New(base.StatusText[base.Redis01])
	}
	return vals[0].Score, vals[0].Member, nil
}
func (self *RedisClient) ZRangeAll(key string) ([]redis.Z, error) {
	if self.essence == nil {
		return nil, errors.New(base.StatusText[base.Redis02])
	}
	return self.essence.ZRangeWithScores(key, 0, -1).Result()
}
