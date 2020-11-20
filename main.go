package main

func main() {
	wd, err := getWeatherData("Berlin")

	printReport(wd, err)
}
