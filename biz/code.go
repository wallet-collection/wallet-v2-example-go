package biz

import (
	"errors"
	"strconv"
	"wallet-example/redis"
)

const (
	// CodeTemplateSceneWithdraw 提现
	CodeTemplateSceneWithdraw = 11
)

type CodeBiz struct {
	rdb *redis.DB
}

func NewCodeBiz(rdb *redis.DB) *CodeBiz {
	return &CodeBiz{
		rdb: rdb,
	}
}

func (c *CodeBiz) Verify(scene int, to string, code string, areaCodeId int64) error {

	sceneStr := strconv.Itoa(scene)

	if areaCodeId > 0 {
		sceneStr += strconv.FormatInt(areaCodeId, 10)
	}

	key := "code:" + sceneStr + ":" + to

	val, err := c.rdb.Get(key)

	_ = c.rdb.Del(key)

	if err != nil {
		return err
	}

	if code != val {
		return errors.New("code error")
	}

	return nil
}
