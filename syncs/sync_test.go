package syncs

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex1(t *testing.T) {
	var mu sync.Mutex
	mu.Lock()
	go func() {
		fmt.Println("hello world")
		mu.Unlock()
	}()
	mu.Lock()
}

/**
根据Go语言内存模型规范，对于从无缓存通道进行的"接收"，发生在对该通道进行的"发送"完成之前。
因此，后台线程<-done接收操作完成之后，main线程的done <- 1发送操作才可能完成
*/
func TestMutex2(t *testing.T) {
	done := make(chan int)
	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("hello")
		<-done
	}()
	fmt.Println("wait")
	done <- -1
	fmt.Println("world")
}

//从2、3看出，无缓存通道双方都会阻塞
func TestMutex3(t *testing.T) {
	done := make(chan int)
	go func() {
		time.Sleep(time.Duration(2) * time.Second)
		fmt.Println("hello")
		done <- -1
	}()
	fmt.Println("wait")
	<-done
	fmt.Println("world")
}
func TestPrimeFilter(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	ch := generateNatural(ctx)
	for i := 0; i < 100; i++ {
		prime := <-ch
		fmt.Printf("%v:%v\n\n", i+1, prime)
		ch = PrimeFilter(ctx, ch, prime)
	}
	cancel()
}
