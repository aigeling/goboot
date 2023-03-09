package cache

import (
	"errors"
	"fmt"
	"testing"
)

var client *RedisClient

func init() {
	var err error
	client, err = NewRedisClient(&RedisConf{
		Host:     "192.168.19.141",
		Port:     6379,
		Password: "123456",
		DB:       0,
	})
	if err != nil {
		fmt.Printf("conn redis err :%v", err)
	}

}

func TestLPush(t *testing.T) {

	for i := 0; i < 5; i++ {
		err := client.LPush("queue_jrxd", fmt.Sprintf("value:%d", i))
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestBRPop(t *testing.T) {
	for i := 0; i < 6; i++ {
		rst, err := client.BRPop("queue_jrxd", 5)
		if err != nil {
			if errors.Is(err, NilError) {
				t.Logf("result is nil : %v", err)
			} else {
				t.Fatalf("err:  %v", err)
			}
		}
		t.Logf("rst: %s", rst)
	}
}

func TestGet(t *testing.T) {
	rst, err := client.Get("hello")
	if err != nil {
		if errors.Is(err, NilError) {
			t.Logf("result is nil : %v", err)
		} else {
			t.Fatalf("err:  %v", err)
		}
	}

	t.Logf("rst: %s", rst)
}

func TestSet(t *testing.T) {
	rst, err := client.Set("hello", "123", 0)

	if err != nil {
		t.Fatalf("err:  %v", err)
	}

	t.Logf("rst: %s", rst)
}

func TestDel(t *testing.T) {
	_, err := client.Set("hello", "123", 0)
	if err != nil {
		t.Fatalf("err:  %v", err)
	}

	rst, _ := client.Get("hello")
	t.Logf("Get rst:%s", rst)
	i, _ := client.Del("hello")
	t.Logf("Del rst:%d", i)
	rst, _ = client.Get("hello")
	t.Logf("Get rst:%s", rst)

}
