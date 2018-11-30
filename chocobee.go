//Build chocolatey packages using docker containers
//for windows 2016/2019
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

//Environment variables are declared on the top before any function
var nugeturl string = os.Getenv("NUGET_URL") // https://chocolatey.org/api/v2
var docker_registry string = os.Getenv("DOCKER_REGISTRY")
var docker_image string = os.Getenv("DOCKER_IMAGE")
var docker_image_tag string = os.Getenv("DOCKER_IMAGE_TAG")

//var artifactory_apikey string = os.Getenv("ARTIFACTORY_APIKEY")

var waitGroup sync.WaitGroup

//Make sure all global variables are set
func validateVariables() {
	if strings.Compare("", nugeturl) == 0 {
		fmt.Println("Please set NUGET_URL value")
		os.Exit(1)
	}

	if strings.Compare("", docker_registry) == 0 {
		fmt.Println("Please set DOCKER_REGISTRY value")
		//os.Exit(1)
	}

	if strings.Compare("", docker_image) == 0 {
		fmt.Println("Please set DOCKER_IMAGE value")
		//os.Exit(1)
	}

	if strings.Compare("", docker_image_tag) == 0 {
		fmt.Println("Please set DOCKER_IMAGE_TAG value")
		//os.Exit(1)
	}

}

func runContainer(container int) {
	defer waitGroup.Done()
	cmd := exec.Command("cmd", "/c", "echo", "Hello World")

	cmdReader, err := cmd.StdoutPipe()

	err = cmd.Start()

	if err != nil {
		os.Stderr.WriteString(err.Error())
	} else {

		scanner := bufio.NewScanner(cmdReader)
		go func() {
			for scanner.Scan() {
				//Scanner if flushing the StdoutPipe was messages are read
				fmt.Printf("%s\n", scanner.Text())
			}

		}()

		cmd.Wait()
	}
}

//Scans the package directory for nuspec files and
//verifies that only non-existing packages are built-tested-deployed
//against a target chocolatey server
func buildPackages() {
	list_packages, _ := filepath.Glob("packages/*/*/*.nuspec")

	fmt.Printf("I found %d packages\n", len(list_packages))

	for i := 0; i < len(list_packages); i++ {
		//TODO: use filepath.FromSlash() to make it os independent
		//      this program will most likely be used on windows
		name := strings.Split(list_packages[i], "\\")[1]
		version := strings.Split(list_packages[i], "\\")[2]

		if isDeployed(name, version, nugeturl) {
			fmt.Println("Going to skip the deployment...")
		} else {
			fmt.Println("Going to run the deployment...")
		}
	}
}

func isDeployed(name, version, source string) bool {
	log.Printf("INFO: choco search %s --version %s --source %s ", name, version, source)
	out, err := exec.Command("choco", "search", name, "--version", version, "--source", source).Output()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s \n", out)
	if strings.Contains(string(out), "1 packages found") {
		fmt.Printf("I found a package %s \n", name)
		return true
	} else {
		return false
	}
}

func main() {
	validateVariables()
	buildPackages()
	//	fmt.Printf("%#v\n", waitGroup)
	//	for i := 0; i < 10; i++ {
	//		waitGroup.Add(1)
	//		go runcontainer(i)
	//}

	//fmt.Printf("%#v\n", waitGroup)
	//waitGroup.Wait()
	fmt.Println("\nExiting...")
}
