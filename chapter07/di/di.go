package main

import (
	"errors"
	"fmt"
	"net/http"
)

// ログを記録する関数
func LogOutput(message string) {
	fmt.Println(message)
}

// データの保存場所
type SmpleDataStore struct {
	userData map[string]string
}

func (sds SmpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

// SimpleDataStoreのインスタンスを作成するファクトリ関数
func NewSimpleDataStore() SmpleDataStore {
	return SmpleDataStore{
		userData: map[string]string{
			"1": "Alice",
			"2": "Bob",
			"3": "Charlie",
		},
	}
}

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

type Logger interface {
	Log(message string)
}

// LogOutputがLoggerインターフェースを満たすようにアダプタを作成
type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("SayHello(" + userID + ")")
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("User not found")
	}
	return name + "さんこんにちは", nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("SayGoodbye(" + userID + ")")
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("User not found")
	}
	return name + "さんさようなら", nil
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

type Logic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	l     Logger
	logic Logic // Controller型の変数はLogicインターフェースを満たす任意の型を保持することができる
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.l.Log("SayHello内： ")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(l Logger, logic Logic) Controller {
	return Controller{
		l:     l,
		logic: logic,
	}
}

func main() {
	l := LoggerAdapter(LogOutput)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)

	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
