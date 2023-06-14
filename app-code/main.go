// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Hello World")
// }

// func main() {
// 	http.HandleFunc("/", handler)
// 	http.ListenAndServe(":8000", nil)
// }
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	message := `
	<!DOCTYPE html>
	<html>
		<head>
			<style>
				body {
					background-color: #000;
					color: #FFF;
					font-size: 24px;
					text-align: center;
					padding-top: 200px;
				}
			</style>
		</head>
		<body>
			Hello World! ITI's DevOps Track says Hi
		</body>
	</html>
	`

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, message)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}
