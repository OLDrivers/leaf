package leaf

import (
	"os"
	"os/signal"

	"github.com/OLDrivers/leaf/cluster"
	"github.com/OLDrivers/leaf/console"
	"github.com/OLDrivers/leaf/log"
	"github.com/OLDrivers/leaf/module"
)

func Run(mods ...module.Module) {
	log.Info("Leaf %v starting up", version)

	// module
	for i := 0; i < len(mods); i++ {
		module.Register(mods[i])
	}
	module.Init()

	// cluster
	cluster.Init()

	// console
	console.Init()

	// close
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	log.Info("Leaf closing down (signal: %v)", sig)
	console.Destroy()
	cluster.Destroy()
	module.Destroy()
}
