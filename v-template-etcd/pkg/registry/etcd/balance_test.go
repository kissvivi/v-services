package etcd

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestRandomBalance_Select(t *testing.T) {

	balance := Balance{}

	service := balance.GetService()
	for {
		node, err := balance.SelectRandom(context.TODO(), service.Nodes)
		if err != nil {
			t.Errorf("err%v", err)
		}
		fmt.Println(node)

		time.Sleep(time.Second * 3)
	}
}

func TestBalance_SelectList(t *testing.T) {
	balance := Balance{}

	service := balance.GetService()
	for {
		node, err := balance.SelectList(context.TODO(), service.Nodes)
		if err != nil {
			t.Errorf("err%v", err)
		}

		fmt.Println(fmt.Sprintf("%v:%v", node.IP, node.Port))

		time.Sleep(time.Second * 3)
	}
}
