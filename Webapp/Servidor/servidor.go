package main
 
import (
	"fmt"// Imprimir en consola
/*	"io"// Ayuda a escribir en la respuesta
	"log"//Loguear si algo sale mal*/
	"net/http"// El paquete HTTP
)
 
    func manejador(w http.ResponseWriter, r *http.Request){
      fmt.Fprintf(w,"Hola, %s, ¡este es el servidor!", r.URL.Path)
    }
    func main(){
      http.HandleFunc("/", manejador)
      fmt.Println("El servidor se encuentra en ejecución")
      http.ListenAndServe(":8080", nil)
    }

/*func main() {
 
	http.HandleFunc("/hola", func(w http.ResponseWriter, peticion *http.Request) {
		io.WriteString(w, "Solicitaste hola")
	})
 
 
	http.HandleFunc("/ruta/un/poco/larga", func(w http.ResponseWriter, peticion *http.Request) {
		io.WriteString(w, "Solicitaste ruta/un/poco/larga")
	})
	direccion := ":8080" // Como cadena, no como entero; porque representa una dirección
	fmt.Println("Servidor listo escuchando en " + direccion)
	log.Fatal(http.ListenAndServe(direccion, nil))
}*/
