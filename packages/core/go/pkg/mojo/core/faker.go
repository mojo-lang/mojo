package core

import (
	"sync"

	"github.com/mojo-lang/mojo/packages/core/go/pkg/logs"
)

type Faker interface {
	Fake() *Value
}

var fakers map[string]Faker
var fakersOnce sync.Once

func RegisterFaker(name string, faker Faker) {
	fakersOnce.Do(func() {
		fakers = make(map[string]Faker)
	})

	if _, ok := fakers[name]; ok {
		logs.Warnf("the faker %s already exist, will be override", name)
	}

	fakers[name] = faker
}
