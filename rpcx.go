package xjutils

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"log"
	"strings"
)

func RpcxCall(etcdAddr []string, servicePath, method string, args interface{}) (reply interface{}, err error){
	basePath := ""
	index := strings.LastIndex(servicePath,"/")
	service := servicePath[index+1:]
	d := client.NewEtcdDiscovery(basePath, servicePath,etcdAddr, nil)
	xclient := client.NewXClient(service, client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	err = xclient.Call(context.Background(), method, args, &reply)
	if err != nil {
		log.Printf("failed to call: %v", err)
	}
	return
}

func RpcxCallStr(etcdAddr []string, servicePath, method string, args interface{}) (reply string, err error){
	basePath := ""
	index := strings.LastIndex(servicePath,"/")
	service := servicePath[index+1:]
	d := client.NewEtcdDiscovery(basePath, servicePath,etcdAddr, nil)
	xclient := client.NewXClient(service, client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	err = xclient.Call(context.Background(), method, args, &reply)
	if err != nil {
		log.Printf("failed to call: %v", err)
	}
	return
}

func RpcxCallByte(etcdAddr []string, servicePath, method string, args interface{}) (reply []byte, err error){
	basePath := ""
	index := strings.LastIndex(servicePath,"/")
	service := servicePath[index+1:]
	d := client.NewEtcdDiscovery(basePath, servicePath,etcdAddr, nil)
	xclient := client.NewXClient(service, client.Failtry, client.RandomSelect, d, client.DefaultOption)
	defer xclient.Close()

	err = xclient.Call(context.Background(), method, args, &reply)
	if err != nil {
		log.Printf("failed to call: %v", err)
	}
	return
}