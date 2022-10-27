package test

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"testing"
)

func TestGenerateUUID(t *testing.T) {
	s := uuid.NewV4().String()
	for i := 0; i < 20; i++ {
		fmt.Println(len(s), s)
	}
}
