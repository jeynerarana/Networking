package main

import (
	"fmt"
	"log"
	"net"
//	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"strings"
)

/* Returns a Header response, Note: This is just like the GET response,
but no files would be sent */
func resolveHead(head []byte)([]byte){//DONE
	/*Fields function seperates file by spaces, hence the second argument 
	will be used(aka index 1) this is where the file name is located */
	file := strings.Fields(string(head))
	server := "Server: cihttpd\r\n"
	ContentLength := "Content-Length: "
	current_time := time.Now()
	print_time := current_time.Format(time.ANSIC)
	LastModified := "Last-Modified: "+string(print_time)+"\r\n"
	headerOk := "HTTP/1.1 200 OK\r\n"
	headerNotFound := "HTTP/1.1 404 Not Found\r\n"
	//reading the file that the client requested
	if string(file[1]) != "/" {
		readFile, err := ioutil.ReadFile("www"+string(file[1]))
		//file not found err 404
		if err != nil {
			fmt.Println(string(readFile)+" File not Found")
			lensize := strconv.Itoa(int(200))
			ContentLength += lensize + "\r\n\r\n"
			headerNotFound += server + LastModified + ContentLength
			return []byte(headerNotFound)
		}
		//now we look for length of file and last date Modified
		fi, err := os.Stat("www"+string(file[1]))
		if err != nil {
			log.Fatal(err)
		}
		lensize := strconv.Itoa(int(fi.Size()))
		ContentLength += lensize + "\r\n\r\n"
		headerOk += server + LastModified + ContentLength
	}else{ //we use this to accept a "/" as a GET request and send the index.html file
		index :="www/index.html" 
		readFile, err := ioutil.ReadFile(index)
		//file not found
		if err != nil {
		    fmt.Println(string(readFile)+" File not found")
		}
		//now we look for length of file and last date Modified
		fi, err := os.Stat("www"+string(file[1]))
		if err != nil {
		    log.Fatal(err)
		}
		lensize := strconv.Itoa(int(fi.Size()))
		ContentLength += lensize + "\r\n\r\n"
		headerOk += server + LastModified + ContentLength
	}
	//return the final outcome
	return []byte(headerOk)
}
//Returns a response to GET request, as well as Initiating Neccesarry Headers
func resolveGet(get []byte)([]byte){//DONE
	/*Fields function seperates file by spaces, hence the second argument 
	will be used(aka index 1) this is where the file name is located */
	file := strings.Fields(string(get))
	server := "Server: cihttpd\r\n"
	ContentLength := "Content-Length: "
	current_time := time.Now()
	print_time := current_time.Format(time.ANSIC)
	LastModified := "Last-Modified: "+string(print_time)+"\r\n"
	headerOk := "HTTP/1.1 200 OK\r\n"
	headerNotFound := "HTTP/1.1 404 Not Found\r\n"
	//reading the file that the client requested
	if string(file[1]) != "/" {
		readFile, err := ioutil.ReadFile("www"+string(file[1]))
		//file not found err 404
		if err != nil {
	    		fmt.Println("error file not Found")
			readErrFile , _ := ioutil.ReadFile("www/404.html")
			lensize := strconv.Itoa(int(200))
			ContentLength += lensize + "\r\n\r\n"
			headerNotFound += server + LastModified + ContentLength + string(readErrFile)
			return []byte(headerNotFound)
		}
		//now we look for length of file and last date Modified
		fi, err := os.Stat("www"+string(file[1]))
		if err != nil {
    			log.Fatal(err)
		}
		lensize := strconv.Itoa(int(fi.Size()))
		ContentLength += lensize + "\r\n\r\n"
		headerOk += server + LastModified + ContentLength + string(readFile)
	}else{ //we use this to accept a "/" as a GET request and send the index.html file
		index :="www/index.html" 
		readFile, err := ioutil.ReadFile(index)
		//file not found
		if err != nil {
		    fmt.Println("Index File not found")
		}
		//now we look for length of file and last date Modified
		fi, err := os.Stat("www"+string(file[1]))
		if err != nil {
    		    log.Fatal(err)
		}
		lensize := strconv.Itoa(int(fi.Size()))
		ContentLength += lensize + "\r\n\r\n"
		headerOk += server + LastModified + ContentLength + string(readFile)
	}
	//return the final outcome
	return []byte(headerOk)
}
/* Checks what request is being recieved (HEAD/GET) calls function accordingly,
returns the []byte of the reponse needed */
func resolveRequest(request []byte)([]byte){//DONE
	outputData := make([]byte,512)
	//checks for the type of request
	//this can be Modified to read other requests
	if string(request[:3]) == "GET"{
		outputData = resolveGet(request)
	}else{//HEAD response
		outputData = resolveHead(request)
	}
	return outputData
}
/* Reads data, shows errors if there is error reading/writing
calls resolveRequest function to identify the request type. Returns final ouput in []bytes */
func handleClient(conn net.Conn) {//DONE
	data := make([]byte, 512)
	_, err :=conn.Read(data)
	if err != nil {log.Fatal(err)}
	//helper function to resolveRequest
	finalOut := resolveRequest(data)
	conn.Write([]byte(finalOut))
	if err != nil {log.Fatal(err)}
	//line below is used for debugging purposes, outputs the response of server
	//fmt.Printf(string(finalOut))
	conn.Close()
}
func main() {
	ln, err := net.Listen("tcp", ":8080")
	//listener error
	if err != nil {
	log.Fatal("Could not Listen")
	}
	//waits for loop to finish then closes connection
	defer ln.Close()
	for {
		conn,err :=ln.Accept()
		if err != nil {log.Fatal("Error Accepting request")}
		//handles  sending/recieving data from client
		go handleClient(conn)
	}
}
