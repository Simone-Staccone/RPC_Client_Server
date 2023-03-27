package main

import (
	"bufio"
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
	"time"
)

func formatString(s string) string {
	for {
		if strings.Contains(s, "\n") {
			s = strings.TrimSuffix(s, "\n")
		} else if strings.Contains(s, "\r") {
			s = strings.TrimSuffix(s, "\r")
		} else {
			s = strings.TrimSuffix(s, " ")
			break
		}
	}
	return s
}

func addResource(resource *Resource, client *rpc.Client) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("shell> Insert the name of the resource you want to add")
	fmt.Print("shell/usr> ")
	resource.Value, _ = reader.ReadString('\n')
	resource.Value = formatString(resource.Value)

	ret := new(Reply)
	err := client.Call("Ret.AddResource", resource, ret)
	if err != nil {
		log.Printf("%v\n", err)
		log.Fatal("Error in client add resource")
	}

	fmt.Printf("shell> %+v\n", ret.RET)
}

func lookUpResource(resource *Resource, client *rpc.Client) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("shell> Insert the name of the resource you want to retrieve")
	fmt.Print("shell/usr> ")
	resource.Value, _ = reader.ReadString('\n')
	resource.Value = formatString(resource.Value)

	ret := new(Reply)
	err := client.Call("Ret.LookUpResource", resource, ret)
	if err != nil {
		log.Printf("%v\n", err)
		log.Fatal("Error in client lookup for resources")
	}

	if ret.RET == -1 {
		fmt.Println("shell> No Resource was found with this name")
	} else {
		fmt.Printf("shell> Found %v with id: %v\n", resource.Value, ret.RET)
	}
	fmt.Println("shell> Press ENTER to continue...")
	fmt.Print("shell/usr> ")
	_, _ = reader.ReadString('\n')
}

func connectionShell() int {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\033[H\033[2J") //Magic code to clean the screen
	for {
		fmt.Println("************************************")
		fmt.Println("*             MY SHELL             *")
		fmt.Println("************************************")
		fmt.Println("* 1 - Connect to RPC server        *")
		fmt.Println("* 0 - Exit                         *")
		fmt.Println("************************************")

		fmt.Print("shell/usr> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		switch text[0] {
		case '1':
			return 1
		case '0':
			fmt.Println("shell> Exiting...")
			return 0
		default:
			fmt.Println("shell> Insert only one number")
			time.Sleep(1000000000)
		}
	}
}

func mainShell(client *rpc.Client) int {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\033[H\033[2J") //Magic code to clean the screen

		fmt.Println("************************************")
		fmt.Println("*             MY SHELL             *")
		fmt.Println("************************************")
		fmt.Println("* 1 - Add node                     *")
		fmt.Println("* 2 - Add resource                 *")
		fmt.Println("* 3 - Lookup Resource              *")
		fmt.Println("* 4 - Get network list             *")
		fmt.Println("* 0 - Exit                         *")
		fmt.Println("************************************")

		fmt.Print("shell/usr> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		switch text[0] {
		case '1':
			/*TODO*/
		case '2':
			addResource(new(Resource), client)
		case '3':
			lookUpResource(new(Resource), client)
		case '4':
			/*TODO*/
		case '0':
			fmt.Printf("shell> Exiting...\n")
			return 0
		default:
			fmt.Println("shell> Insert only one number")
			time.Sleep(1000000000)
		}
	}
}

func main() {
	address := "localhost:8080"

	fmt.Printf("Starting client...\n")

	//ret := new(Reply)

	flag := connectionShell()

	//Open connection
	if flag == 1 {
		client, err := rpc.Dial("tcp", address)
		if err != nil {
			log.Printf("%v", client)
			log.Fatal("Error,", err)
		}

		mainShell(client)

	}

}
