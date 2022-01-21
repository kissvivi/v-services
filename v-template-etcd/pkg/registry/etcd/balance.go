package etcd

import (
	"context"
	"fmt"
	"math/rand"
	"time"
	"v-template-etcd/pkg/registry"
)

type Balance struct {
	index int
}

func (r *Balance) GetService() *registry.Service {
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
		return nil
	}

	//for {
	service, err := registryInst.GetService(context.TODO(), "test-service")
	if err != nil {
		fmt.Println(fmt.Sprintf("init registry failed, err:%v", err))
		return nil
	}

	return service

	//for _, node := range service.Nodes {
	//	fmt.Printf("service:%s, node:%#v\n", service.Name, node)
	//}
	//time.Sleep(time.Second * 5)
	//}
}

//随机
func (r *Balance) SelectRandom(ctx context.Context, nodes []*registry.Node) (node *registry.Node, err error) {

	if len(nodes) == 0 {
		//err = ErrNotHaveNodes
		return
	}

	index := rand.Intn(len(nodes))
	node = nodes[index]

	return node, nil
}

//轮训

func (r *Balance) SelectList(ctx context.Context, nodes []*registry.Node) (node *registry.Node, err error) {

	if len(nodes) == 0 {
		//err = ErrNotHaveNodes
		return
	}

	r.index = (r.index + 1) % len(nodes)
	node = nodes[r.index]

	return node, nil
}

//权重
