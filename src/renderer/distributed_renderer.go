package renderer

import (
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"scenes"
//	"films"
)

func handleConnection(conn net.Conn) {
	log.Println("Received connection.")
	dec := gob.NewDecoder(conn)
	s := new(scenes.Scene)
	log.Print("Decoding scene...")
	err := dec.Decode(s)
	if err != nil {
		log.Fatal("Failed: ", err)
	} 
	log.Println("Received scene ", s.Filename)
	log.Println("Start rendering...")
	StartRendering(s, false)
	log.Println("Finished rendering.")
	encoder := gob.NewEncoder(conn)
    encoder.Encode(s.Film)
    log.Println("Sent result back.")
}

func StartServer() {
	log.Println("Starting server...")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	log.Println("Listening for requests...")
	for {
		conn, err := ln.Accept() // this blocks until connection or error
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn) // a goroutine handles conn so that the loop can accept other connections
	}
}

func RenderOnServer(scene *scenes.Scene) {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        log.Fatal("Connection error: ", err)
    }
    log.Println("Connected, sending scene...")
    encoder := gob.NewEncoder(conn)
    encoder.Encode(scene)
    log.Println("Waiting for server to finish rendering...")
    
    // Expect a BoxFilterFilm as answer.
//    dec := gob.NewDecoder(conn)
//	f := new(films.BoxFilterFilm)
//	dec.Decode(f)
//	f.WriteToPng(scene.Filename)
    conn.Close()
    fmt.Println("Done, wrote to ", scene.Filename);
}
