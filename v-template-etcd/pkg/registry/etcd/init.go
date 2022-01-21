package etcd

import (
	"context"
	"fmt"
	"time"
	"v-template-etcd/pkg/registry"
)

func InitRegister(service *registry.Service) {
	//初始化注册中心  注册etcd 服务中心
	registryInst, err := registry.InitRegistry(context.TODO(), "etcd",
		registry.WithAddrs([]string{"127.0.0.1:2379"}),
		registry.WithTimeout(time.Second),
		registry.WithRegistryPath("/etcd-Service/"),
		registry.WithHeartBeat(5),
	)
	if err != nil {
		fmt.Println(fmt.Sprintf("init registry failed, err:%v", err))
		return
	}
	registryInst.Register(context.TODO(), service)
}

const envEtcd = "ETCD_HOST"

func EtcdRegister() {
	etcd := "10.10.10.2"
	fmt.Println("etcd:", etcd)
	url := fmt.Sprintf("%s:2379", etcd)
	//初始化注册中心  注册etcd 服务中心
	registryInst, err := registry.InitRegistry(context.TODO(), "etcd",
		registry.WithAddrs([]string{url}),
		registry.WithTimeout(time.Second),
		registry.WithRegistryPath("/etcd-Service/"),
		registry.WithHeartBeat(5),
	)
	if err != nil {
		fmt.Println(fmt.Sprintf("init registry failed, err:%v", err))
		return
	}

	service := &registry.Service{
		Name: "test-service",
	}

	service.Nodes = append(service.Nodes, &registry.Node{
		IP:   "10.10.10.11",
		Port: 8081,
	},
		&registry.Node{
			IP:   "10.10.10.12",
			Port: 8082,
		},
		&registry.Node{
			IP:   "10.10.10.13",
			Port: 8083,
		},
	)
	registryInst.Register(context.TODO(), service)
}
