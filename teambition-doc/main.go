package main

import (
	"log"

	"github.com/kexin8/teambition-sdk-go/teambition-doc/domain"
)

const (
	OK_ICON = ":heavy_check_mark:"
	X_ICON  = ":x:"
)

var (
	GenNodes = map[string]bool{"API": true}

	FinishApis = Set("获取企业信息","搜索企业自定义字段分类","根据自定义字段分类统计自定义字段数","搜索企业任务类型","获取企业的管理员（含拥有者）","获取企业拥有者","获取企业成员列表","获取企业部门列表","获取部门成员列表","获取用户加入的企业部门列表")
)

func main() {
	n, err := domain.DocNode()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	for _, node := range n {
		if _, ok := GenNodes[node.Name]; !ok {
			continue
		}

		// log.Println(node.Name)
		// node.Children
		children(node.Children)
	}

}

func children(nodes []domain.Node) {
	for _, n := range nodes {
		if n.ObjectType == "folder" {

			level := "###"
			for i := 0; i < len(n.AncestorIds); i++ {
				level += "#"
			}
			println(level, n.Name)
			children(n.Children)
		} else {
			if _, ok := FinishApis[n.Name]; ok {
				println("- ", n.Name, ":heavy_check_mark:")
				continue
			}
			println("- ", n.Name, ":x:")
		}

	}
}

func Set(v ...string) map[string]bool {
	m := make(map[string]bool)
	for _, s := range v {
		m[s] = true
	}
	return m
}
