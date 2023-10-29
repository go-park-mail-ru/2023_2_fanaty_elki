package postgres

import (
	"encoding/json"
	"server/internal/domain/dto"
	"server/internal/domain/entity"

	"github.com/gomodule/redigo/redis"
)

type sessionManager struct {
	redisConn redis.Conn
}

func NewSessionManager(conn redis.Conn) *sessionManager {
	return &sessionManager{
		redisConn: conn,
	}
}

func (sm *sessionManager) Create(cookie *entity.Cookie) error {
	dataSerialized, _ := json.Marshal(cookie.UserID)
	mkey := "sessions:" + cookie.SessionToken
	
	result, err := redis.String(sm.redisConn.Do("SET", mkey, dataSerialized, "EX", cookie.MaxAge.Seconds()))
	if err != nil || result != "OK" {
		return entity.ErrInternalServerError
	}

	return nil
}

func (sm *sessionManager) Check(sessionToken string) (*entity.Cookie, error) {
	mkey := "sessions:" + sessionToken
	data, err := redis.Bytes(sm.redisConn.Do("GET", mkey))
	if err != nil {
		return nil, entity.ErrInternalServerError
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
		return entity.ErrInternalServerError
	}
	
	return nil
}
