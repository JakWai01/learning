// /*
// JSONEchoClient
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

// func (p Person) String() string {
// 	s := p.Name.Personal + " " + p.Name.Family

// 	for _, v := range p.Email {
// 		s += "\n" + v.Kind + ": " + v.Address
// 	}
// 	return s
// }

// func main() {
// 	person := Person{
// 		Name: Name{Family: "Waibel", Personal: "Jakob"}, Email: []Email{Email{Kind: "work", Address: "mail@jakobwaibel.com"}, Email{Kind: "home", Address: "asdasd"}}

// 		if len(os.Args) != 2 {
// 			fmt.Println("Usage: ", os.Args[0], "host:port")
// 			os.Exit(1)
// 		}

// 		service := os.Args[1]

// 		conn, err := net.Dial("tcp", service)
// 		checkError(err)

// 		encoder := json.NewEncoder(conn)
// 		decoder := json.NewDecoder(conn)

// 		for n := 0; n < 10; n++ {
// 			encoder.Encode(person)
// 			var newPerson Person
// 			decoder.Decode(&newPerson)
// 			fmt.Println(newPerson.String())
// 		}

// 		os.Exit(0)
// 	}
// }

// func checkError(err) {

// 	log.Fatal(err)
// 	os.Exit(1)
// }

// func readFully(conn net.Conn) (byte[], error) {
// 	defer conn.Close()

// 	result := bytes.NewBuffer(nil)
// 	var buf [512]byte
// 	for {
// 		n, err :=  conn.Read(buf[0:])
// 		result.Write(buf[0:n])
// 		if err != nill {
// 			if err == io.EOF {
// 				break
// 			}
// 		}

// 		return nil, err
// 	}

// 	return result.Bytes(), nil
// }