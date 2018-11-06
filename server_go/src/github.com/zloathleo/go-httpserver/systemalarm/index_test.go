package systemalarm

import (
	"github.com/pkg/errors"
	"strconv"
	"testing"
	"log"
)

func TestAddSystemAlarm(t *testing.T) {
	for i := 0; i < 30; i++ {
		AddSystemAlarm(errors.New("111:" + strconv.Itoa(i)))
	}
	log.Println(GetSystemAlarm())
}
