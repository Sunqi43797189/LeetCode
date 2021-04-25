package link_list

import (
	"fmt"
	"leet_code_go/helper"
	"testing"
)

func TestIsLinkListCircle(t *testing.T) {
	head := helper.CircleLinkList()
	res, node := IsLinkListCircle(head)
	fmt.Println(res, node.Value)
}
