package hystrix

import "testing"

func TestHystrix(t *testing.T) {
	Hystrix()
}

/*
 wrk -t100 -c100 -d1s http://192.168.19.135:8080/api/ping/baidu
wrk 压测工具
*/
