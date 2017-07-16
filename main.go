package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoring = 2
const delay = 2

func main() {

	showIntroduction()

	for {
		showMenu()               //Function that shows the menu
		command := readCommand() //Menu reading
		chooseCommand(command)   //According to chosen menu direct
	}

}

func showIntroduction() {

	name := "Abraao"
	print("Hi ", name)
	println(" your software has version 1.1")
}

func showMenu() {

	fmt.Println("1- Start Monitoring ")
	fmt.Println("2- Show logs ")
	fmt.Println("0- Exit the software")
}

func readCommand() int {

	var command int
	fmt.Scan(&command)
	fmt.Println(" command chose is ", command)
	return command
}

func chooseCommand(command int) {

	switch command {
	case 1:
		fmt.Println("Monitoring...")
		startMonitoring()
	case 2:
		fmt.Println("showing logs..." + "\n")
		showLogs()
	case 0:
		fmt.Println(" bye!")
		os.Exit(0)
	default:
		fmt.Println(" Unrecongized command! ")
		os.Exit(-1)
	}
}

func startMonitoring() {
	sites := readSitesFiles()

	for i := 0; i < monitoring; i++ {
		for i, site := range sites {
			fmt.Println("testando site", i, ":", site)
			verificationSites(site)
		}

	}
	time.Sleep(delay * time.Second) //sleep in 5 seconds
	fmt.Println("")
}

func verificationSites(site string) {
	reply, err := http.Get(site)
	check(err)

	if reply.StatusCode == 200 {
		fmt.Println("site: ", site, " want load with sucess")

		saveLog(site, true)
	} else {
		fmt.Println("site:", site,
			"was load with failure \n status code is", reply.StatusCode)

		saveLog(site, false)
	}
}

func readSitesFiles() []string {

	var sites []string
	file, err := os.Open("input/sites.txt")
	check(err)

	reads := bufio.NewReader(file)
	for {
		line, err := reads.ReadString('\n')

		if len(line) > 1 {
			line = strings.TrimSpace(line)
			sites = append(sites, line)
		}

		if err == io.EOF {
			break
		}

	}

	fmt.Println(sites)

	file.Close()
	return sites
}

func saveLog(site string, status bool) {

	file, err := os.OpenFile("output/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	check(err)
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " -Online: " +
		strconv.FormatBool(status) + "\n")
	file.Close()
}

func showLogs() {

	file, err := ioutil.ReadFile("output/log.txt")
	check(err)

	fmt.Println(string(file))

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
