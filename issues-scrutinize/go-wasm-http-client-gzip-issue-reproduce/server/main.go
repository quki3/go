/* SL=(standard library)*/
package main

import (
        "bytes"         // SL manipulation of byte slices 
                        //used to hold binary data in memory
        "compress/gzip" // SL reading and writing compressed
                        //file as specified in RFC 1952
        "fmt"           // SL implement formatted I/O 
                        //analogous to C's prinf and scanf
        "log"           // SL simple logging package
        "net/http"      // SL client and server implementation
)

func httpNoGzip(w http.ResponseWriter, req *http.Request) {             // this por now I don't know clear what do this?
        fmt.Fprintf(w, "hello, world")
}
func httpGzip(w http.ResponseWriter, req *http.Request) {               //func used for the Handlers
        var b bytes.Buffer //declare var b type
                           //bytes.Buffer
        gw := gzip.NewWriter(&b) //create new gzip writer
                                 //this will write compressed
                                 //data to the bytes.Buffer
                                 //(&b) specify where the compressed
                                 //data should be writen
        gw.Write([]byte("hello, world")) //write byte slice to 
                                         //the underlying buffer b
        gw.Close() //ensure that any remaining data is flushed

        w.Header().Set("Content-Encoding", "gzip") //this header indicates
                                                   //taht the response
                                                   //body is compressed
                                                   //using the gzip
        w.Write(b.Bytes()) //this write the commpressed data and 
                           //sends as the response body using methods
                           //of http.ResponseWriter
}

func main() {
        http.Handle("/", http.FileServer(http.Dir(".")))
        http.HandleFunc("/gzip", httpGzip)
        http.HandleFunc("/nogzip", httpNoGzip)
        log.Fatalln(http.ListenAndServe(":8080", nil))
}

