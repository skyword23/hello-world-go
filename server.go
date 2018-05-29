package main 

import (
	"fmt"
	"log"
	"net/http"
	"context"
	"cloud.google.com/go/datastore"
)

type Task struct {
    Value string
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}

func retrieveHandler(w http.ResponseWriter, r *http.Request) {
	
	ctx := context.Background()
    dsClient, err := datastore.NewClient(ctx, "notional-portal-146510")
    if err != nil {
        fmt.Fprintf(w, "%s", err)
        return
    }
    
    var tasks []*Task
    query := datastore.NewQuery("Task").Order("Value")
    keys, err := dsClient.GetAll(ctx, query, &tasks)

    if err != nil {
        fmt.Fprintf(w, "%s", err)
        return
    }

    for i, _ := range keys {
        fmt.Fprintln(w, "<h1>", tasks[i].Value, "</h1>")
    }
}


func saveHandler(w http.ResponseWriter, r *http.Request) {

	ctx := context.Background()
    dsClient, err := datastore.NewClient(ctx, "notional-portal-146510")
    if err != nil {
		fmt.Fprintf(w, "%s", err)
        return
    }

    k := datastore.IncompleteKey("Task", nil)
    
    input := r.URL.Query().Get("input")

    if (len(input) <= 0){
    	fmt.Fprintf(w, "<h1>Please provide some input</h1>")
        return
    }

    e := new(Task)
    e.Value = input

    if _, err := dsClient.Put(ctx, k, e); err != nil {
		fmt.Fprintf(w, "%s", err)
        return
    }

	fmt.Fprintf(w, "<h1>Added value %q to datastore</h1>", e.Value)
}

func main () {

	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/retrieve/", retrieveHandler)
	http.HandleFunc("/", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}