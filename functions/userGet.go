package functions

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/HamelBarrer/api-go/connections"
	"go.mongodb.org/mongo-driver/bson"
)

// UserGet es la funcion que me permite obtener los usuarios
func UserGet(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	collection := connections.Connection().Database("testing").Collection("users")

	var users []bson.M

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = cursor.All(ctx, &users); err != nil {
		log.Fatal(err.Error())
	}

	json.NewEncoder(w).Encode(users)
}
