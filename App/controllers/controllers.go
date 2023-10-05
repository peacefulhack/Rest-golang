package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"restArchitecture/mikail/App/clients"
	"restArchitecture/mikail/App/config"
	"restArchitecture/mikail/App/models"
	"restArchitecture/mikail/App/utils/enum"
	"strings"
	"time"
)

type ProductControllers interface {
	InsertProductData(req *models.InsertProductRequest) error
	SortProductData(key, sorter string) ([]*models.SortProductResponse, error)
	getRedishandler(redisKey string) ([]*models.SortProductResponse, bool, error)
}

type ProductControllerRepo struct {
	client clients.ProductClientRepo
	redis  clients.RedisClientRepo
	err    error
}

func NewControllers(env string) *ProductControllerRepo {
	repo := &ProductControllerRepo{}
	repo.init(env)
	return repo
}

func (m *ProductControllerRepo) init(env string) {
	cfg, err := config.StartConfig(env)
	if err != nil {
		m.err = err
		return
	}
	db, err := config.NewMysql(cfg)
	if err != nil {
		m.err = err
		return
	}
	rd := config.NewRedis(cfg)
	productClient := clients.NewProductClient(db)
	redisClient := clients.NewRedisClient(rd)
	m.client = productClient
	m.redis = redisClient
}

func (m *ProductControllerRepo) setMock(repo clients.ProductClientRepo, redis clients.RedisClientRepo, err error) {
	m.client = repo
	m.redis = redis
	m.err = err
}

func (m *ProductControllerRepo) InsertProductData(req *models.InsertProductRequest) error {
	if m.err != nil {
		return m.err
	}
	err := m.client.InsertProductList(&models.ProductList{
		ProductName:     req.ProductName,
		ProductDesc:     req.ProductDesc,
		ProductPrice:    req.ProductPrice,
		ProductQuantity: req.ProductQuantity,
		CreatedDate:     time.Now().Format(enum.DbTimeFormat),
	})
	if err != nil {
		return err
	}
	return nil
}

func (m *ProductControllerRepo) SortProductData(key, order string) ([]*models.SortProductResponse, error) {
	if m.err != nil {
		return nil, m.err
	}
	key = strings.ToLower(key)
	order = strings.ToLower(order)
	if order != "asc" && order != "desc" {
		return nil, errors.New("parameter not valid")
	}
	redisKey := "sort-" + key + "-" + order
	val, isExist, err := m.getRedishandler(redisKey)
	if err != nil {
		fmt.Println(err)
	}
	if isExist {
		return val, nil
	}
	res := []*models.SortProductResponse{}
	if key == "new" {
		res, err = m.client.SortDateProductList(order)
		if err != nil {
			return nil, err
		}
	} else if key == "price" {
		res, err = m.client.SortPriceProductList(order)
		if err != nil {
			return nil, err
		}
	} else if key == "name" {
		res, err = m.client.SortNameProductList(order)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("parameter not valid")
	}
	jsonData, err := json.Marshal(res)
	_ = m.redis.Set(redisKey, jsonData, 1*time.Hour)
	return res, nil
}

func (m *ProductControllerRepo) getRedishandler(redisKey string) ([]*models.SortProductResponse, bool, error) {
	val, isExist, err := m.redis.Get(redisKey)
	if err != nil {
		return nil, isExist, err
	}
	if isExist {
		var res []*models.SortProductResponse
		err = json.Unmarshal([]byte(val), &res)
		if err != nil {
			err := m.redis.Del(redisKey)
			if err != nil {
				return nil, false, err
			}
			return nil, false, err
		}
		return res, true, nil
	}
	return nil, false, nil
}
