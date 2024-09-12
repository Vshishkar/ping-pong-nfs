package main

import (
	"log"
	"log/slog"
	"nfs/internal/coordinator"
)

func main() {
	slog.Info("Hello from coordinator! This is coord")

	c := coordinator.MakeCoordinator()
	log.Fatal(c.Start())
}
