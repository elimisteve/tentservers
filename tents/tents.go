// Steve Phillips / elimisteve
// 2012.09.19

package tentserver

import (
	"encoding/json"
    "appengine"
    "appengine/datastore"
	"io/ioutil"
    "net/http"
	"time"
)

type TentServer struct {
	Author     string     `json:"author"`
	URL        string     `json:"url"`
	CreatedAt  time.Time  `json:"created_at"`
}

func init() {
    http.HandleFunc("/", root)
    http.HandleFunc("/tents", tents)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nothing here yet! Maybe you meant to visit /tents?\n"))
}

func tents(w http.ResponseWriter, r *http.Request) {
	// Too bad I can't import "github.com/bmizerany/pat"...
	if r.Method == "GET" {
		getTents(w, r)
	} else if r.Method == "POST" {
		postTents(w, r)
	}
}

func getTents(w http.ResponseWriter, r *http.Request) {
	// Grab all TentServer objects from DB
    c := appengine.NewContext(r)
    q := datastore.NewQuery("TentServer")
    tents := []TentServer{}
    _, err := q.GetAll(c, &tents)
	if err != nil {
		writeError(w, err)
        return
    }
	// Marshall all TentServers to JSON
	jsonStr, err := json.Marshal(&tents)
	if err != nil {
		writeError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Write(jsonStr)
}

func postTents(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)
	t := TentServer{CreatedAt: time.Now()}
	// Read POSTed body (should be JSON)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		writeError(w, err)
		return
	}
	defer r.Body.Close()
	// Unmarshal JSON into TentServer var
    if err := json.Unmarshal(body, &t); err != nil {
		writeError(w, err)
		return
	}
	// Store new TentServer
	key := datastore.NewIncompleteKey(c, "TentServer", nil)
    if _, err := datastore.Put(c, key, &t); err != nil {
		writeError(w, err)
        return
    }
	// Return new list of TentServer so user can see theirs was added
	getTents(w, r)
}

func writeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	http.Error(w, err.Error(), http.StatusInternalServerError)
}