package main

import (
	"explorerapp/counters"
	"explorerapp/hotelsrepo"
	"explorerapp/infoserver"
)

func main() {
	CounterInterfaceDemo()
	RepoDemoForInterface()
	LoggerUsage()
	ServerLoggerOperations()

}
func ServerLoggerOperations() {
	logger := infoserver.NewLogger()
	server := infoserver.NewServer("Info server", &logger)
	server.Start()
	infoserver.ServerOperations(&server)
	server.Logger.Error("Recording a error")
	server.DrainToConsole()
}

func LoggerUsage() {
	logger := infoserver.NewLogger()
	logger.Info("Inside main")
	logger.Info("Invoking task 123")
	logger.Debug("Debugging task 567")
	logger.Error("Error caught in task 111")
	logger.DrainToConsole()
}

func RepoDemoForInterface() {
	fileRepo := hotelsrepo.NewHotelRepoFile("data/golddata.json")
	hotelsrepo.HotelOperations(&fileRepo)
	apiRepo := hotelsrepo.NewApiHotelRepo("http://localhost:3000/inventory")
	hotelsrepo.HotelOperations(&apiRepo)
}

func CounterInterfaceDemo() {
	sc := counters.NewSimpleCounter()
	counters.CounterOperations(&sc)

	stCounter := counters.NewStepCounter(2, 4)
	counters.CounterOperations(&stCounter)
}
