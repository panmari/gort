package renderer

import (
	"encoding/gob"
	"films"
	"fmt"
	"log"
	"net"
	"scenes"
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
	registerTypes()

	for {
		conn, err := ln.Accept() // this blocks until connection or error
		if err != nil {
			log.Println("Could not accept connection: ", err)
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
	registerTypes()
	encoder := gob.NewEncoder(conn)

	err = encoder.Encode(scene)
	if err != nil {
		log.Fatal("Failed to encode: ", err)
	}
	log.Println("Waiting for server to finish rendering...")

	// Expect a BoxFilterFilm as answer.
	// TODO(smoser) Find out the type of the film by reflection of scene.
	dec := gob.NewDecoder(conn)
	f := new(films.BoxFilterFilm)
	dec.Decode(f)
	// Tonemapper is a function pointer, which can not be transmitted by gob.
	// Set it here again
	f.Tonemapper = scene.Film.GetTonemapper()
	f.WriteToPng(scene.Filename)
	conn.Close()
	fmt.Println("Done, wrote to ", scene.Filename)
}
