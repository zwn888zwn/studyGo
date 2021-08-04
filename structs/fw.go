package structs

import "strings"

type customFWStruct struct {
	Id        int    `json:"id"` //序号
	Name      string `json:"name"`
	Priority  int    `json:"priority"`
	Direction string `json:"direction"` //方向
	Behavior  string `json:"behavior"`  //行为
	Protocol  string `json:"protocol"`
	Ports     string `json:"ports"` //端口字符串 123 or 123:124
	SourceIp  string `json:"sourceIp"`
	Arg1 string `json:"arg1"`
	Arg2 string `json:"arg2"`
	Arg3 string `json:"arg3"`
}
type customFWSlice []customFWStruct

func (s customFWSlice) Len() int           { return len(s) }
func (s customFWSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s customFWSlice) Less(i, j int) bool { return s[i].Priority < s[j].Priority }

func appendProtocol(buildRuleString *strings.Builder, item customFWStruct) {
	switch item.Protocol {
	case TCP:
		buildRuleString.WriteString(" -p tcp ")
	case UDP:
		buildRuleString.WriteString(" -p udp ")
	case ICMP:
		buildRuleString.WriteString(" -p icmp ")
		if len(item.Arg1)>0{
			buildRuleString.WriteString(" --icmp-type " + item.Arg1)
		}
	case ALL:
		buildRuleString.WriteString(" -p all ")
	}
}

func appendSourceIp(ruleString *strings.Builder, item customFWStruct) {
	if len(item.SourceIp) > 0 {
		ruleString.WriteString(" -s " + item.SourceIp)
	}
}

func appendPorts(ruleString *strings.Builder, item customFWStruct) {
	if len(item.Ports) > 0 {
		switch item.Direction {
		case INPUT:
			ruleString.WriteString(" --dport " + item.Ports)
		case OUTPUT:
			ruleString.WriteString(" --sport " + item.Ports)
		}
	}
}

func appendBehavior(ruleString *strings.Builder, item customFWStruct) {
	switch item.Protocol {
	case ACCEPT:
		ruleString.WriteString(" -j " + ACCEPT)
	case REFUSE:
		ruleString.WriteString(" -j " + REFUSE)
	case DROP:
		ruleString.WriteString(" -j " + DROP)
	}
}

//protocol
const (
	TCP  = "tcp"
	UDP  = "udp"
	ICMP = "icmp"
	ALL  = "all"
)

//behavior
const (
	ACCEPT = "ACCEPT"
	REFUSE = "REFUSE"
	DROP   = "DROP"
)

//direction
const (
	INPUT  = "INPUT"  //下行
	OUTPUT = "OUTPUT" //上行
)
