package xjutils

import (
	"context"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"

	"go.etcd.io/etcd/clientv3"
)

type ClientDis struct {
	client     *clientv3.Client
	serverList map[string]string
	lock       sync.Mutex
}

func NewClientDis(addr []string) (*ClientDis, error) {
	conf := clientv3.Config{
		Endpoints:   addr,
		DialTimeout: 5 * time.Second,
	}
	if client, err := clientv3.New(conf); err == nil {
		return &ClientDis{
			client:     client,
			serverList: make(map[string]string),
		}, nil
	} else {
		return nil, err
	}
}

func (this *ClientDis) GetService(prefix string) ([]string, error) {
	prefix = "services/" + prefix
	resp, err := this.client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}
	addrs := this.extractAddrs(resp)

	go this.watcher(prefix)
	return addrs, nil
}

func (this *ClientDis) watcher(prefix string) {
	rch := this.client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range rch {
		for _, ev := range wresp.Events {
			switch ev.Type {
			case clientv3.EventTypePut:
				this.SetServiceList(string(ev.Kv.Key), string(ev.Kv.Value))
			case clientv3.EventTypeDelete:
				this.DelServiceList(string(ev.Kv.Key))
			}
		}
	}
}

func (this *ClientDis) extractAddrs(resp *clientv3.GetResponse) []string {
	addrs := make([]string, 0)
	if resp == nil || resp.Kvs == nil {
		return addrs
	}
	for i := range resp.Kvs {
		if v := resp.Kvs[i].Value; v != nil {
			this.SetServiceList(string(resp.Kvs[i].Key), string(resp.Kvs[i].Value))
			addrs = append(addrs, string(v))
		}
	}
	return addrs
}

func (this *ClientDis) SetServiceList(key, val string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.serverList[key] = string(val)
	log.Println("set data key :", key, "val:", val)
}

func (this *ClientDis) DelServiceList(key string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	delete(this.serverList, key)
	log.Println("del data key:", key)
}

func (this *ClientDis) SerList2Array() []string {
	this.lock.Lock()
	defer this.lock.Unlock()
	addrs := make([]string, 0)

	for _, v := range this.serverList {
		addrs = append(addrs, v)
	}
	return addrs
}

func (this *ClientDis) GetServerList() map[string][]*url.URL {
	//var ServerList map[string][]*url.URL
	ServerList := make(map[string][]*url.URL)
	for k, _ := range this.serverList {
		strs := strings.Split(k, "/")
		serverurl := &url.URL{
			Scheme: "http",
			Host:   strs[2],
		}
		key := strs[1]
		var urls []*url.URL
		urls = ServerList[key]
		urls = append(urls, serverurl)
		ServerList[key] = urls
	}
	return ServerList
}

//func main () {
//	//cli,_ := NewClientDis([]string{"192.168.147.151:2379"})
//	cli,_ := NewClientDis([]string{"127.0.0.1:2379"})
//	cli.GetService("config")
//	//fmt.Println(cli.serverList)
//	select {}
//}
