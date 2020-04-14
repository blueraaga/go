package main

import ("fmt"
		"net/http")


func index_handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Go is neat! Awesome!!")
}

func about_handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "So now you want to know about me! Sneaky!!")
}


func work_handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "<h1>With tags and vars</h1>")
	
	fmt.Fprintf(w,
		`<p>Same fprintf</p>
		<p>can do multiline</p>
		<p>texts also.</p>`)

	fmt.Fprintf(w, "I can also print %d which was passed as a %s.", 5, "variable")
}


func main() {
	fmt.Println("Starting server. Listening at port 8000...")
	// request redirector to handler function
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about", about_handler)
	http.HandleFunc("/work", work_handler)
	// Server listner
	http.ListenAndServe(":8000", nil)
}