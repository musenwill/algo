package src

import (
	"fmt"

	"github.com/urfave/cli"
)

func channel(c *cli.Context) error {
	sendToNilChan()
	receiveFromNilChan()
	sendToNormalChan()
	receiveFromNormalChan()
	sendToClosedChan()
	receiveFromClosedChan()
	return nil
}

func sendToNilChan() {
	var cc chan int

	select {
	case cc <- 1:
		fmt.Println("can send to a nil chan")
	default:
		fmt.Println("send to a nil chan will block")
	}
}

func receiveFromNilChan() {
	var cc chan int

	select {
	case <-cc:
		fmt.Println("can receive from a nil chan")
	default:
		fmt.Println("receive from a nil chan will block")
	}
}

func sendToNormalChan() {
	cc := make(chan int, 1)

	select {
	case cc <- 1:
		fmt.Println("can send to a normal chan")
	default:
		fmt.Println("send to a normal chan will block")
	}
	close(cc)
}

func receiveFromNormalChan() {
	cc := make(chan int, 1)
	cc <- 1

	select {
	case <-cc:
		fmt.Println("can receive from a normal chan")
	default:
		fmt.Println("receive from a nromal chan will block")
	}
}

func sendToClosedChan() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("send to a closed chan will panic")
		}
	}()

	cc := make(chan int, 1)
	close(cc)

	select {
	case cc <- 1:
		fmt.Println("can send to a normal chan")
	default:
		fmt.Println("send to a normal chan will block")
	}
}

func receiveFromClosedChan() {
	cc := make(chan int, 1)
	close(cc)

	select {
	case r := <-cc:
		if r == 0 {
			fmt.Println("can receive nil value from a closed chan")
		} else {
			fmt.Println("can receive none nil value from a closed chan")
		}
	default:
		fmt.Println("receive from a nromal chan will block")
	}
}
