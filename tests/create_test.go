package tests

import (
	"fmt"
	"github.com/sunfmin/hyperboot/services"
	"io/ioutil"
	"testing"
)

func TestCreateConfig(t *testing.T) {
	dir, err := ioutil.TempDir("", "hyperboot")
	if err != nil {
		t.Error(err)
	}

	fmt.Println("TestCreateConfig root dir: ", dir)
	err = services.CreateConfig(dir, "profile", "mysql")
	if err != nil {
		t.Error(err)
	}
}

func TestDbConfig(t *testing.T) {
	dir, err := ioutil.TempDir("", "hyperboot")
	if err != nil {
		t.Error(err)
	}
	fmt.Println("TestCreateConfig root dir: ", dir)
	err = services.CreateDb(dir, "profile")
	if err != nil {
		t.Error(err)
	}
}
