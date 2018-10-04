package lib1

import "github.com/hyangah/exp/internal/common"

var Ver = "2"

func Name() string {
	return common.Name("lib1", Ver)
}


