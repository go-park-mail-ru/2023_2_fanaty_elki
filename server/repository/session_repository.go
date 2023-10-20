package repository

import (
	"encoding/json"
	"fmt"
	"math/rand"

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

type Session struct {
	UserID    uint
	Useragent string
}

type SessionID struct {
	ID string
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

func (sm *SessionManager) Create(in *Session) (*SessionID, error) {
	id := SessionID{randStringRunes(sessKeyLen)}
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

func (sm *SessionManager) Check(in *SessionID) (*Session, error) {
	mkey := "sessions:" + in.ID
	data, err := redis.Bytes(sm.redisConn.Do("GET", mkey))
	if err != nil {
		return nil, err
	}
	sess := &Session{}
	err = json.Unmarshal(data, sess)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

func (sm *SessionManager) Delete(in *SessionID) error {
	mkey := "sessions:" + in.ID
	_, err := redis.Int(sm.redisConn.Do("DEL", mkey))
	if err != nil {
		return err
	}
	return nil
}
