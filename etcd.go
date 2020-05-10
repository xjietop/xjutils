package xjutils

import (
	"fmt"
	"time"

	//"github.com/etcd-io/etcd/clientv3"
	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
)

func EtcdPut(endpoints []string, key, value string) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return err
	}

	fmt.Println("connect succ")
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, key, value)
	cancel()
	if err != nil {
		fmt.Println("put failed, err:", err)
		return err
	}
	return nil
}

func EtcdGet(endpoints []string, key string) map[string]string {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("connect failed, err:", err)
		return nil
	}

	fmt.Println("connect succ")
	defer cli.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key, clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return nil
	}
	m := make(map[string]string)
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
		key := fmt.Sprintf("%s", ev.Key)
		value := fmt.Sprintf("%s", ev.Value)
		m[key] = value
	}
	return m
}
