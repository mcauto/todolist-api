package database

import (
	"fmt"
	"testing"
	"todo-list-back/config"
)

func TestConnectDB(t *testing.T) {
	url := config.Conf.Database.GetURL()
	fmt.Println(url)
	_, err := ConnectDB(url)
	if err != nil {
		t.Error(err)
	}
}

func TestConnectDBFail(t *testing.T) {
	url := ""
	_, err := ConnectDB(url)
	if err == nil {
		t.Errorf("상상도 하지 못한 에러")
	}
}
