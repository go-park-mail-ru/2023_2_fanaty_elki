package repository

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"server/internal/domain/entity"
	"github.com/gomodule/redigo/redis"
)

var (
	letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}


const sessKeyLen = 10

type SessionManager struct {
	redisConn redis.Conn
}

func NewSessionManager(conn redis.Conn) *SessionManager {
	return &SessionManager{
		redisConn: conn,
	}
}

func (sm *SessionManager) Create(in *entity.Session) (*entity.SessionID, error) {
	id := entity.SessionID{ID: randStringRunes(sessKeyLen)}
	dataSerialized, _ := json.Marshal(in)
	mkey := "sessions:" + id.ID
	result, err := redis.String(sm.redisConn.Do("SET", mkey, dataSerialized, "EX", 86400))
	if err != nil {
		return nil, err
	}
	if result != "OK" {
		return nil, fmt.Errorf("result not OK")
	}
	return &id, nil
}

func (sm *SessionManager) Check(in *entity.SessionID) (*entity.Session, error) {
	mkey := "sessions:" + in.ID
	data, err := redis.Bytes(sm.redisConn.Do("GET", mkey))
	if err != nil {
		return nil, err
	}
	sess := &entity.Session{}
	err = json.Unmarshal(data, sess)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func (sm *SessionManager) Delete(in *entity.SessionID) error {
	mkey := "sessions:" + in.ID
	_, err := redis.Int(sm.redisConn.Do("DEL", mkey))
	if err != nil {
		return err
	}
	return nil
}
