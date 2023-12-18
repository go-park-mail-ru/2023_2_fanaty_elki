package repository

import (
	"AuthService/dto"
	"AuthService/entity"
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"sync"
)

type sessionManager struct {
	redisConn redis.Conn
	mu        sync.Mutex
}

func NewSessionManager(conn redis.Conn) *sessionManager {
	return &sessionManager{
		redisConn: conn,
	}
}

func (sm *sessionManager) Create(cookie *entity.Cookie) error {
	dataSerialized, _ := json.Marshal(cookie.UserID)
	mkey := "sessions:" + cookie.SessionToken
	result, err := redis.String(sm.redisConn.Do("SET", mkey, dataSerialized, "EX", 540000))
	if err != nil || result != "OK" {

		return entity.ErrInternalServerError
	}

	return nil
}

func (sm *sessionManager) Check(sessionToken string) (*entity.Cookie, error) {
	mkey := "sessions:" + sessionToken
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

func (sm *sessionManager) Delete(cookie *dto.DBDeleteCookie) error {
	mkey := "sessions:" + cookie.SessionToken
	_, err := redis.Int(sm.redisConn.Do("DEL", mkey))
	if err != nil {
		if err != redis.ErrNil {
			return entity.ErrInternalServerError
		}
		return nil
	}

	return nil
}

func (sm *sessionManager) Expire(cookie *entity.Cookie) error {
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

func (sm *sessionManager) CreateCsrf(sessionToken string, csrfToken string) error {
	dataSerialized, _ := json.Marshal(csrfToken)
	mkey := "csrf:" + sessionToken
	result, err := redis.String(sm.redisConn.Do("SET", mkey, dataSerialized, "EX", 540000))
	if err != nil || result != "OK" {
		return entity.ErrInternalServerError
	}
	fmt.Println(result)
	return nil
}

func (sm *sessionManager) GetCsrf(sessionToken string) (string, error) {
	mkey := "csrf:" + sessionToken
	sm.mu.Lock()
	data, err := redis.Bytes(sm.redisConn.Do("GET", mkey))
	sm.mu.Unlock()
	if err != nil {
		if err != redis.ErrNil {
			return "", entity.ErrInternalServerError
		}
		return "", nil
	}

	var csrfToken string
	err = json.Unmarshal(data, &csrfToken)
	if err != nil {
		return "", entity.ErrInternalServerError
	}
	return csrfToken, nil
}
