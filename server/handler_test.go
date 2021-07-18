package server

import (
	"fmt"
	"testing"
)

func TestGetRootDir(t *testing.T) {
	getRootDir()
	fmt.Println(rootDir)
}
