package functions

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/HamelBarrer/api-go/connections"
	"github.com/HamelBarrer/api-go/models"
)

// UserAdd es la funcion que me permite agregar usuarios
func UserAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applications/json")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := connections.Connection().Database("testing").Collection("users")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(w).Encode(result.InsertedID)
}
