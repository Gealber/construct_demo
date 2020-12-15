package api

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Gealber/construct_demo/accounts"
	rmongodb "github.com/Gealber/construct_demo/repository/mongodb"
	rredis "github.com/Gealber/construct_demo/repository/redis"
	"github.com/Gealber/construct_demo/serializer/jwt"
)

func signUpHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
		}
	}()

	user := decodeUserBodyReq(r)
	if user == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unable to decode body request")
		return
	}

	//here I should store the user in the DB
	mongoRepo, err := rmongodb.NewMongoRepository()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Unable to connect to MONGODB: %v", err)
		return
	}

	err = mongoRepo.Store(user)
	if err != nil {
		mongoRepo.Logger.Printf("Unable to store error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to store user to MONGODB: %v", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
		}
	}()

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method: %s Not Allowed", r.Method)
		return
	}

	user := decodeUserBodyReq(r)
	if user == nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Unable to decode body request")
		return
	}

	mongoRepo, err := rmongodb.NewMongoRepository()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Unable to connect to MONGODB: %v", err)
		return
	}

	userDB, err := mongoRepo.Find(user.Email)
	if err != nil {
		mongoRepo.Logger.Printf("Unable to find user error: %v", err)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "The user is not in the DB")
		return
	}

	if userDB.Password != user.Password {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "The user info provided is not correct")
		return
	}

	//Creating token
	tokenString, err := jwtser.Encode(user)
	if err != nil {
		log.Printf("Token Signing error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while Signing Token!")
		return
	}

	//here you should store the token on Redis
	redisRepo, err := rredis.NewRedisRepo()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to init Redis repository: %v", err)
		return
	}

	err = redisRepo.Store(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to store token on Redis repository: %v", err)
		return
	}

	response := accounts.Token{Token: tokenString}
	//send the token as a HTTP response
	jsonResponse(response, w)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err, ok := recover().(error); ok {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %v", err)
		}
	}()
	//this handler is maybe not needed
	if _, ok := r.Header["Authorization"]; !ok {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Authorization header not found")
		return
	}

	//verify if token exists on redis
	redisRepo, err := rredis.NewRedisRepo()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to init Redis repository")
		return
	}

	auth := r.Header["Authorization"]
	if len(auth) == 0 {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Authorization header not found")
		return
	}
	tokenString := strings.ReplaceAll(auth[0], "Bearer ", "")
	user, err := jwtser.Decode(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to parse tokenString")
		return
	}

	redisToken, err := redisRepo.Find(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Unable to find token on redis")
		return
	}

	// check if token match with the one stored on Redis
	if redisToken != tokenString {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Token doesn't match")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
