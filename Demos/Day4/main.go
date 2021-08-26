package main

func main() {
	// CounterInterfaceDemo()
	// RepoDemoForInterface()
	// LoggerUsage()
	// ServerLoggerOperations()
	SimpleClosureDemo()
	ClosureDemoWithAdder()
}

/*

func ServerLoggerOperations() {
	logger := NewLogger()
	server := NewServer("Info server", &logger)
	server.Start()
	ServerOperations(&server)
	server.Logger.Error("Recording a error")
	server.DrainToConsole()
}
func ServerOperations(server *Server) string {
	server.Info("getting data")
	info := server.GetInfo()
	server.Debug("Got the data " + info)
	return info
}
func LoggerUsage() {
	logger := NewLogger()
	logger.Info("Inside main")
	logger.Info("Invoking task 123")
	logger.Debug("Debugging task 567")
	logger.Error("Error caught in task 111")
	logger.DrainToConsole()
}

func RepoDemoForInterface() {
	fileRepo := NewHotelRepoFile("data/golddata.json")
	HotelOperations(&fileRepo)
	apiRepo := NewApiHotelRepo("http://localhost:3000/inventory")
	HotelOperations(&apiRepo)
}
func HotelOperations(hotelRepo HotelRepo) {
	hotelData, err := hotelRepo.GetAllHotels()
	if err != nil {
		panic(err)
	}
	fmt.Println(hotelData)
}
func CounterInterfaceDemo() {
	sc := NewSimpleCounter()
	CounterOperations(&sc)

	stCounter := NewStepCounter(2, 4)
	CounterOperations(&stCounter)
}
func CounterOperations(counter Counter) {
	counter.Increment()
	counter.Increment()
	counter.Increment()
	fmt.Println(counter.GetValue())
}
*/
