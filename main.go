package main

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				log.Println("event:", event)
			}
		}
	}()

	err = watcher.Add("tmp/")
	if err != nil {
		log.Fatal(err)
	}

	<-make(chan struct{})
}
