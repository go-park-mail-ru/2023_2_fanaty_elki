package repository

import (
	"flag"
	"log"
	"reflect"
	"server/internal/domain/dto"
	"server/internal/domain/entity"
	"testing"
	"time"

	"github.com/gomodule/redigo/redis"
)

var (
	redisAddr = flag.String("addr", "redis://user:@localhost:6379/1", "redis addr")
)

func TestCreateSuccess(t *testing.T) {
	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatal("can`t connect to redis", err)
	}

	repo := &sessionManager{
		redisConn: redisConn,
	}

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebdbYudb",
		MaxAge:       50 * time.Hour,
	}

	err = repo.Create(&cookie)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestCheckSuccess(t *testing.T) {
	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatal("can`t connect to redis", err)
	}

	repo := &sessionManager{
		redisConn: redisConn,
	}

	reqcookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebdbYudb",
		MaxAge:       50 * time.Hour,
	}

	cookie := entity.Cookie{
		UserID: 1,
	}

	actual, err := repo.Check(reqcookie.SessionToken)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	if !reflect.DeepEqual(actual, &cookie) {
		t.Errorf("results not match, want %v, have %v", cookie, actual)
		return
	}

}

func TestDeleteSuccess(t *testing.T) {
	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatal("can`t connect to redis", err)
	}

	repo := &sessionManager{
		redisConn: redisConn,
	}

	reqcookie := &dto.DBDeleteCookie{
		SessionToken: "TYebdbYudb",
	}

	err = repo.Delete(reqcookie)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestExpireSuccess(t *testing.T) {
	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatal("can`t connect to redis", err)
	}

	repo := &sessionManager{
		redisConn: redisConn,
	}

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	err = repo.Create(&cookie)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	reqcookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	err = repo.Expire(&reqcookie)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestCreateCsrfSuccess(t *testing.T) {
	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatal("can`t connect to redis", err)
	}

	repo := &sessionManager{
		redisConn: redisConn,
	}

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	err = repo.CreateCsrf(cookie.SessionToken, "HBBGFCCDFG")
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}

func TestGetCsrfSuccess(t *testing.T) {
	redisConn, err := redis.DialURL(*redisAddr)
	if err != nil {
		log.Fatal("can`t connect to redis", err)
	}

	repo := &sessionManager{
		redisConn: redisConn,
	}

	cookie := entity.Cookie{
		UserID:       1,
		SessionToken: "TYebbYudb",
		MaxAge:       50 * time.Hour,
	}

	_, err = repo.GetCsrf(cookie.SessionToken)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

}
