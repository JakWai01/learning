// /*
// JSONEchoServer
// */

// package main

// type Person struct {
// 	Name Name
// 	Email []Email
// }

// type Name struct {
// 	Family string
// 	Personal string
// }

// type Email struct {
// 	Kind string
// 	Address string
// }

// func (p person) String() string {
// 	s := p.Name.Personal + " " + p.Name.Family

// 	for _, v :=  range.p.Email {
// 		s += "\n" + v.Kind + ": " + v.Address
// 	}

// 	return s
// }

// func main() {

// 	service := "0.0.0.0:1200"
// 	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
// 	checkError()

// 	listener, err := net.ListenTCP("tcp", netAddr)
// 	checkError(err)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			continue
// 		}

// 		encoder := json.newEncoder(conn)
// 		decoder := json.newDecoder(conn)

// 		for n := 0; n < 10; n++ {
// 			var person Person
// 			decoder.Decode(&person)
// 			fmt.Println(person.String())
// 			encoder.Encode(person)
// 		}

// 		conn.Close()
// 	}
// }

// func checkError(err error) {

// 	log.Fatal(err)
// 	os.Exit(1)
// }