package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

const (
	cidPath            = "./.containerid"
	msgRefleshInterval = 500 // millisec
)

func main() {

	// For saving messages which received from containers
	rowLogFile, err := os.OpenFile("./log/"+time.Now().Format("2006-01-02-15-04-05-rows.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer rowLogFile.Close()
	rowLogger := log.New(rowLogFile, "", log.Ltime)

	// For saving messages which show to the user
	msgLogFile, err := os.OpenFile("./log/"+time.Now().Format("2006-01-02-15-04-05-msgs.log"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	defer msgLogFile.Close()
	msgLogger := log.New(msgLogFile, "", log.Ltime)

	// Read container's id
	byteCID, err := ioutil.ReadFile(cidPath)
	if err != nil {
		log.Fatalln(err)
	}

	cid := string(byteCID)
	cLogFile := "/var/lib/docker/containers/" + cid + "/" + cid + "-json.log"

	// Channels
	msgChan := make(chan string)
	errChan := make(chan error)

	// Start tailing container's log file
	go tailFile(msgChan, errChan, cLogFile)

	for {
		select {

		// firing when recived message
		case logMsg := <-msgChan:

			row, err := parseLog(logMsg)
			if err != nil {
				fmt.Println(err)
				continue
			}
			rowLogger.Println(row)

			msg, hasMsg := generateMsg(row)
			if hasMsg {
				msgLogger.Println(msg)
				fmt.Println(msg)
			}

		// firing when recived error
		case err := <-errChan:
			log.Fatalln(err)

		// firing when received nothing
		default:
			time.Sleep(msgRefleshInterval * time.Millisecond)
		}
	}

}
