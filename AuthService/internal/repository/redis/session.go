package repository

import (
	"AuthService/dto"
	"AuthService/entity"
	"encoding/json"
	"fmt"
	"sync"
	"github.com/gomodule/redigo/redis"
)

//SessionManager struct
type SessionManager struct {
	redisConn redis.Conn
	mu        sync.Mutex
}

//NewSessionManager creates session manager
func NewSessionManager(conn redis.Conn) *SessionManager {
	return &SessionManager{
		redisConn: conn,
	}
}

//Create creates session in db
func (sm *SessionManager) Create(cookie *entity.Cookie) error {
	dataSerialized, _ := json.Marshal(cookie.UserID)
	mkey := "sessions:" + cookie.SessionToken
	sm.mu.Lock()
	result, err := redis.String(sm.redisConn.Do("SET", mkey, dataSerialized, "EX", 540000))
	sm.mu.Unlock()
	if err != nil || result != "OK" {

		return entity.ErrInternalServerError
	}

	return nil
}

//Check checks session in db
func (sm *SessionManager) Check(sessionToken string) (*entity.Cookie, error) {
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

//Delete deletes session in db
func (sm *SessionManager) Delete(cookie *dto.DBDeleteCookie) error {
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

//Expire updates session in db
func (sm *SessionManager) Expire(cookie *entity.Cookie) error {
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

//CreateCsrf creates csrf in db
func (sm *SessionManager) CreateCsrf(sessionToken string, csrfToken string) error {
	dataSerialized, _ := json.Marshal(csrfToken)
	mkey := "csrf:" + sessionToken
	sm.mu.Lock()
	result, err := redis.String(sm.redisConn.Do("SET", mkey, dataSerialized, "EX", 540000))
	sm.mu.Unlock()
	if err != nil || result != "OK" {
		return entity.ErrInternalServerError
	}
	fmt.Println(result)
	return nil
}

//GetCsrf gets csrf from db
func (sm *SessionManager) GetCsrf(sessionToken string) (string, error) {
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
