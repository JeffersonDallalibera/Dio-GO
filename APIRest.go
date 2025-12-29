package main

import (
    "encoding/json"
    "log"
    "net/http"

    "github.com/gorilla/mux"
)

// Person representa um contato
type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}

// Address representa endereço de um contato
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person

// Utilitário para resposta JSON
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(status)
    json.NewEncoder(w).Encode(payload)
}

// GetPeople retorna todos os contatos
func GetPeople(w http.ResponseWriter, r *http.Request) {
    respondJSON(w, http.StatusOK, people)
}

// GetPerson retorna um contato específico
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            respondJSON(w, http.StatusOK, item)
            return
        }
    }
    respondJSON(w, http.StatusNotFound, map[string]string{"error": "Contato não encontrado"})
}

// CreatePerson cria um novo contato
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    var person Person
    if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
        respondJSON(w, http.StatusBadRequest, map[string]string{"error": "JSON inválido"})
        return
    }

    people = append(people, person)
    respondJSON(w, http.StatusCreated, person)
}

// DeletePerson remove um contato
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            respondJSON(w, http.StatusOK, map[string]string{"message": "Contato removido"})
            return
        }
    }
    respondJSON(w, http.StatusNotFound, map[string]string{"error": "Contato não encontrado"})
}

func main() {
    router := mux.NewRouter()

    // Dados iniciais
    people = append(people, Person{
        ID:        "1",
        Firstname: "John",
        Lastname:  "Doe",
        Address:   &Address{City: "City X", State: "State X"},
    })
    people = append(people, Person{
        ID:        "2",
        Firstname: "Koko",
        Lastname:  "Doe",
        Address:   &Address{City: "City Z", State: "State Y"},
    })

    // Rotas
    router.HandleFunc("/contato", GetPeople).Methods("GET")
    router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/contato", CreatePerson).Methods("POST")
    router.HandleFunc("/contato/{id}", DeletePerson).Methods("DELETE")

    log.Println("Servidor rodando em http://localhost:8000")
    log.Fatal(http.ListenAndServe(":8000", router))
}
