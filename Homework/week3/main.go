package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// WithContext 使用 WithCancel 创建一个可以取消的 context 将 cancel 赋值给 Group 保存起来，然后再将 context 返回回去
	// 注意这里有一个坑，在后面的代码中不要把这个 ctx 当做父 context 传给下游，因为 errgroup取消了，这个 context 就没用了，会导致下游复用时出错
	group, errCtx := errgroup.WithContext(ctx)
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	// 服务器监听 8080 端口
	srv := &http.Server{Addr: ":8080", Handler: mux}

	// Go 方法是一个封装，相当于 go 关键字的加强，会启动一个协程，然后利用 waitgroup 来控制是否结束。
	// 如果有一个非 nil 的 error 出现就会保存起来并且如果有 cancel 就会调用 cancel 取消掉，使 ctx 返回
	group.Go(func() error {
		err := srv.ListenAndServe()
		if err != nil {
			log.Println("http failed, will exit:", err.Error())
		}
		return err
	})

	// g2
	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("http server stop")
		return srv.Shutdown(errCtx)
	})

	// subscribe SIGINT/SIGTERM signals
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-stopChan:
				cancel()
			}
		}
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}

	fmt.Println("Finish")
}

func index(w http.ResponseWriter, req *http.Request) {
	delay := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(delay))

	userAddr := req.RemoteAddr
	if strings.Contains(userAddr, ":") {
		fmt.Println("IP", net.ParseIP(strings.Split(userAddr, ":")[0]), "Response Code", http.StatusOK)
	} else {
		fmt.Println("IP", net.ParseIP(userAddr), "Response Code", http.StatusOK)
	}
	w.Write([]byte(fmt.Sprintf("<h1>%d</h1>", delay)))
}
