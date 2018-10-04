package lib1

import "github.com/hyangah/exp/internal/common"

var Ver = "1"

func Name() string {
	return common.Name("lib1", Ver)
}


