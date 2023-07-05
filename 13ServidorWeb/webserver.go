package servidorweb

import "net/http"

func MiWebServer() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)

}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./13ServidorWeb/index.html")
}
