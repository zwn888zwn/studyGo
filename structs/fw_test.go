package structs

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)



func TestProcessCustomFWCfg(t *testing.T) {
	var err error
	t.Helper()
	var jsonStr string = "[{\"id\":14,\"protocol\":\"tcp\",\"priority\":25,\"direction\":\"INPUT\",\"behavior\":\"ACCEPT\",\"ports\":\"40000\",\"sourceIp\":\"192.168.0.1/24\"},{\"id\":17,\"protocol\":\"udp\",\"priority\":11,\"direction\":\"OUTPUT\",\"behavior\":\"REFUSE\",\"ports\":\"40020:40030\"},{\"id\":18,\"protocol\":\"icmp\",\"priority\":1,\"arg1\":\"8\",\"direction\":\"INPUT\",\"behavior\":\"DROP\",\"sourceIp\":\"0/0\"}]"

	var customFWList customFWSlice

	if err = json.Unmarshal([]byte(jsonStr), &customFWList); err == nil {
		println(customFWList[1].Behavior)
	}

	for _, item := range customFWList {
		var buildRuleString strings.Builder
		appendProtocol(&buildRuleString, item) //fixme 要把对象引用传过去
		appendSourceIp(&buildRuleString, item)
		appendPorts(&buildRuleString, item)
		appendBehavior(&buildRuleString, item)

		//appendToChain(buildRuleString, item)
		fmt.Printf("build str: %s \n",buildRuleString.String())
	}
}
