package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gorilla/websocket"
)


type OPBody struct {
	Content string
}


// file del frontend statici, nel caso di framework è necessario inserire il path della cartella con i distribuibili (dist di solito)
//go:embed frontend
var staticFiles embed.FS
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}


func getPrefix(path string) string {
	return strings.Split(path, "/")[1]
}
 

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// websocket per live reload
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	//... Use conn to send and receive messages.
	//defer conn.Close()
}


func indexHandler(w http.ResponseWriter, r *http.Request) {	
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	if strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
		return
	}

/* mime types: 
	text/css
	text/javascript
	text/html
*/
	prefix := getPrefix(r.URL.Path)
	var contenttype string
	switch prefix {
		case "css": contenttype = "text/css"
		case "js": contenttype = "text/javascript"
		case "json": contenttype = "application/json"
		default: contenttype = "text/html"
	}
	//fmt.Println("URL: " + r.URL.Path + " prefix: " + prefix + " contetype: " + contenttype)
	w.Header().Set("Content-Type", contenttype)
	
	var tao string
	if(r.URL.Path == "/") {
		tao = "frontend/index.html"
	} else {
		tao = "frontend" + r.URL.Path
	}
	//fmt.Println("path: " + r.URL.Path + " ricavato: " + tao)
	rawFile, _ := staticFiles.ReadFile(tao) 
	w.Write(rawFile)
}


// funzione di backend generica
func operazioniHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)
	var t OPBody
	err := decoder.Decode(&t)
	if err != nil {
		fmt.Printf("%v",err)	
		fmt.Println(err)
	}

	// check stringa QR content: supponiamo 8 cifre numeriche
	numeric := regexp.MustCompile(`[0-9]{8}`).MatchString
	if !numeric(t.Content) {
		fmt.Println("string QR content errata: " +  t.Content)
		w.Write([]byte("string QR content errata: " +  t.Content))
		return
	}

	// tanto per ritornare qualcosa
	dati := &OPBody{Content: t.Content}

	b, err := json.Marshal(*dati)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(b)
}


/* attivazione del servizio web */
func main() {
	ReadEnv()
	fmt.Printf("%s\n", "Webrouter attivo")

	//altri path da gestire
	//http.HandleFunc("/api/", )
	http.HandleFunc("/ws", wsHandler)
	http.HandleFunc("/", indexHandler)

	log.Fatal(http.ListenAndServe(porta, nil))
	// per connessione sicura:
	//log.Fatal(http.ListenAndServeTLS(porta, "file.crt", "file.key", nil))
}
