package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"html/template"
	
	_ "RPC-HolaMundo/docs" // Importar la documentación Swagger generada
	httpSwagger "github.com/swaggo/http-swagger"    // Importar el servidor HTTP Swagger
)

// Args define la estructura para los argumentos de la llamada RPC
type Args struct {
	Name string
}

// Greeter implementa el servicio RPC
type Greeter struct{}

// SayHello es el método RPC que retorna un mensaje de saludo
// @Description Metodo que retorna un mensaje de saludo usando RPC
// @Accept  json
// @Produce  json
// @Param name query string true "Nombre del usuario"
// @Success 200 {string} string "Saludo de bienvenida"
// @Failure 400 {string} string "Error: el nombre no puede estar vacío"
// @Router /hello [get]
func (g *Greeter) SayHello(args *Args, reply *string) error {
	if args.Name == "" {
		return errors.New("name cannot be empty")
	}
	*reply = fmt.Sprintf("Hello, %s! Welcome to the RPC world!", args.Name)
	return nil
}

func main() {
	// Registra el servicio RPC
	greeter := new(Greeter)
	rpc.Register(greeter)

	// Exponer el servidor RPC sobre TCP en el puerto 1234
	go func() {
		listener, err := net.Listen("tcp", ":1234")
		if err != nil {
			fmt.Println("Error al iniciar el servidor RPC:", err)
			return
		}
		fmt.Println("Servidor RPC escuchando en el puerto 1234...")
		rpc.Accept(listener)
	}()

	// Servir una página web con un formulario
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Si es una solicitud POST (cuando el usuario envía el formulario)
		if r.Method == "POST" {
			r.ParseForm()
			name := r.FormValue("name")

			// Conectar con el servidor RPC
			client, err := rpc.Dial("tcp", "localhost:1234")
			if err != nil {
				fmt.Println("Error al conectar al servidor RPC:", err)
				http.Error(w, "Error al conectar con el servidor RPC", http.StatusInternalServerError)
				return
			}
			defer client.Close()

			// Realizar la llamada RPC
			var reply string
			err = client.Call("Greeter.SayHello", &Args{Name: name}, &reply)
			if err != nil {
				http.Error(w, "Error al llamar al método RPC", http.StatusInternalServerError)
				return
			}

			// Crear una plantilla HTML con la respuesta del RPC
			tmpl := template.Must(template.New("response").Parse(`
				<html>
				<body>
					<h1>RPC Hello World</h1>
					<p>{{.}}</p>
					<form method="post">
						<label for="name">Enter your name:</label>
						<input type="text" id="name" name="name">
						<input type="submit" value="Submit">
					</form>
				</body>
				</html>
			`))

			// Ejecutar la plantilla con la respuesta de RPC
			tmpl.Execute(w, reply)

		} else {
			// Si es una solicitud GET, mostrar el formulario vacío
			tmpl := template.Must(template.New("form").Parse(`
				<html>
				<body>
					<h1>RPC Hello World</h1>
					<form method="post">
						<label for="name">Enter your name:</label>
						<input type="text" id="name" name="name">
						<input type="submit" value="Submit">
					</form>
				</body>
				</html>
			`))

			// Ejecutar la plantilla vacía
			tmpl.Execute(w, nil)
		}
	})

	// Ruta para Swagger
	http.Handle("/swagger/", httpSwagger.WrapHandler)

	// Iniciar el servidor HTTP en el puerto 8080
	http.ListenAndServe(":8080", nil)
}