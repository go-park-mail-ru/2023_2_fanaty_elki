package redis

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"sync"
)

type adminManager struct {
	redisConn redis.Conn
	mu        sync.Mutex
}

func NewadminManager(conn redis.Conn) *adminManager {
	return &adminManager{
		redisConn: conn,
	}
}

func (sm *adminManager) Create(cookie *entity.Cookie) error {
	dataSerialized, _ := json.Marshal(cookie.UserID)
	mkey := "admin_id:" + cookie.SessionToken
	result, err := redis.String(sm.redisConn.Do("SET", mkey, dataSerialized, "EX", 540000))
	if err != nil || result != "OK" {
		return entity.ErrInternalServerError
	}

	return nil
}

func (sm *adminManager) Check(sessionToken string) (*entity.Cookie, error) {
	mkey := "admin_id:" + sessionToken
	sm.mu.Lock()
	data, err := redis.Bytes(sm.redisConn.Do("GET", mkey))
	sm.mu.Unlock()
	if err != nil {
		if err != redis.ErrNil {
			return nil, entity.ErrInternalServerError
		}
		return nil, nil
	}

	cookie := &entity.Cookie{}
	err = json.Unmarshal(data, &cookie.UserID)
	if err != nil {
		return nil, entity.ErrInternalServerError
	}
	return cookie, nil
}

func (sm *adminManager) Delete(cookie *dto.DBDeleteCookie) error {
	mkey := "admin_id:" + cookie.SessionToken
	_, err := redis.Int(sm.redisConn.Do("DEL", mkey))
	if err != nil {
		if err != redis.ErrNil {
			return entity.ErrInternalServerError
		}
		return nil
	}

	return nil
}

func (sm *adminManager) Expire(cookie *entity.Cookie) error {
	err := sm.Delete(dto.ToDBDeleteCookie(cookie))
	if err != nil {
		return entity.ErrDeletingCookie
	}

	err = sm.Create(cookie)
	if err != nil {
		return entity.ErrCreatingCookie
	}
	return nil
}
