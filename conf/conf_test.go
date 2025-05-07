package conf

import (
	"os"
	"testing"
)

func TestGetPath(t *testing.T) {
	dir, _ := os.Getwd()
	t.Log(dir)
}

func TestInit(t *testing.T) {
	Init()
	t.Log(Server)
	t.Log(MySQL)
}
