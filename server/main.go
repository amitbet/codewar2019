package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type fileHandler struct {
	root    string
	headers map[string]string
}

type MatchLogger struct {
	fname string
}

func NewFileHandler(root string, hlist map[string]string) http.Handler {
	return &fileHandler{root, hlist}
}

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	upath := r.URL.Path
	if !strings.HasPrefix(upath, "/") {
		upath = "/" + upath
		r.URL.Path = upath
	}
	fName := f.root + path.Clean(upath)
	//println("serving: ", fName)

	for key, value := range f.headers {
		w.Header().Set(key, value)
	}

	http.ServeFile(w, r, fName)
}

var rec = &GameRecord{}

func GetFilenameFromDate() string {
	// Use layout string for time format.
	const layout = "01-02-2006[15.04.05]"
	// Place now in the string.
	t := time.Now()
	return t.Format(layout)
}

func (ml *MatchLogger) MatchLoggerHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := io.ReadAll(r.Body)
	var m Match
	err := json.Unmarshal(b, &m)
	if err != nil {
		fmt.Printf("MatchLoggerHandler, a problem in unmarshaling json: %v", err)
	}
	fmt.Printf("%s", m)
	rec.AddMatch(&m)

	err = os.WriteFile(ml.fname, []byte(rec.Csv()), 0644)
	if err != nil {
		fmt.Printf("MatchLoggerHandler, problem writing  csv: %v", err)
	}
}

// Helper function to parse headers from a comma-separated string
func parseHeaders(headerStr string) map[string]string {
	headers := make(map[string]string)
	if headerStr == "" {
		return headers
	}

	headerPairs := strings.Split(headerStr, ",")
	for _, pair := range headerPairs {
		parts := strings.SplitN(pair, ":", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			headers[key] = value
		}
	}

	return headers
}

func main() {
	csvFileName := "gameLog-" + GetFilenameFromDate() + ".csv"
	ml := MatchLogger{fname: csvFileName}

	var portFlag = flag.Int("p", 3000, "help message for flag n")
	headers := flag.String("H", "", "Custom headers (comma-separated)")

	flag.Parse()

	// Register the handler function
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// Set custom headers

	// Split the headers string into individual headers
	headerList := parseHeaders(*headers)

	port := *portFlag
	fs := NewFileHandler("./", headerList)
	http.Handle("/", fs)

	r := mux.NewRouter()
	r.HandleFunc("/log-match", ml.MatchLoggerHandler).Methods("POST")
	http.Handle("/log-match", r)

	log.Println("Listening on " + strconv.Itoa(port))
	err := http.ListenAndServe(":"+strconv.Itoa(port), nil)
	if err != nil {
		log.Fatal(err)
	}

}
