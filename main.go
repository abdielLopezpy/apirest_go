package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Este es el tipo de dato que usaremos para almacenar los elementos de nuestra lista
type Item struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// Esta es la lista que usaremos para almacenar nuestros elementos
var items []Item

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/items", handleItems)
	http.HandleFunc("/items/", handleItem)
	fmt.Println("Servidor iniciado en el puerto 8080")
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Bienvenido a nuestra API")
}

// Esta función maneja las solicitudes a la ruta /items
func handleItems(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Obtener todos los elementos
		json.NewEncoder(w).Encode(items)
	case "POST":
		// Crear un nuevo elemento
		var item Item
		json.NewDecoder(r.Body).Decode(&item)
		items = append(items, item)
		json.NewEncoder(w).Encode(item)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

// Esta función maneja las solicitudes a la ruta /items/{id}
func handleItem(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID del elemento de la ruta
	idStr := r.URL.Path[len("/items/"):]
	if idStr == "" {
		http.Error(w, "Falta el ID del elemento", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":
		// Obtener el elemento con el ID especificado
		for _, item := range items {
			if item.ID == id {
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		http.Error(w, "No se encontró el elemento", http.StatusNotFound)
	case "PUT":
		// Actualizar el elemento con el ID especificado
		var item Item
		json.NewDecoder(r.Body).Decode(&item)
		if item.ID != id {
			http.Error(w, "ID del elemento no coincide", http.StatusBadRequest)
			return
		}
		for i, oldItem := range items {
			if oldItem.ID == id {
				items[i] = item
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		http.Error(w, "No se encontró el elemento", http.StatusNotFound)
	case "DELETE":
		// Borrar el elemento con el ID especificado
		for i, item := range items {
			if item.ID == id {
				items = append(items[:i], items[i+1:]...)
				json.NewEncoder(w).Encode(item)
				return
			}
		}
		http.Error(w, "No se encontró el elemento", http.StatusNotFound)
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
