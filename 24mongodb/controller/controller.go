package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/kamblepratik90/mongoDbGoLang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = "mongodb://127.0.0.1:27017/"
const dbName = "GoLangTut"
const collectionName = "watchlist"

// MOST imp
var collection *mongo.Collection

// connect with MongoDB

func init() { // this is special method /func in go IT ONLY EXECUTES ONCE AND AT VERY BEGINNING.

	// client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongo
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongo connection success...")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection instance ready...")
}

// mongo db helpers - files

// insert one

func insertOneMovie(movie model.Netflix) {
	fmt.Println("Inserting movie ")

	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted new movie with Id: ", inserted.InsertedID)
}

// update 1
func updateOneMovie(movieId string) {
	fmt.Println("Updating movie")
	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updated, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Updated movie count: ", updated.ModifiedCount)
}

// delete 1
func deleteOneMovie(movieId string) {
	fmt.Println("Deleting movie")

	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}

	deletedOne, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted one movie: ", deletedOne.DeletedCount)
}

// delete all movies
func deleteAllMovie() int64 {
	fmt.Println("Deleting all movie")

	filter := bson.M{} // delete all

	deletedAll, err := collection.DeleteMany(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Deleted all movie: ", deletedAll.DeletedCount)
	return deletedAll.DeletedCount
}

// get all movies
func geyAllMovies() []primitive.M {
	fmt.Println("Getting all movies")

	filter := bson.M{} // for all

	cursor, err := collection.Find(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cursor.Close(context.Background())
	return movies

	// fmt.Println("Find all movie: ", findAll)
}

// actual controllers - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/x-www-form-urlencode")
	// w.Header().Add("Content-Type", "application/json")
	allMovies := geyAllMovies()

	err := json.NewEncoder(w).Encode(allMovies)

	if err != nil {
		log.Fatal(err)
	}
}

// create movies

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	// error can be handled for validation / input parsing
	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)

}

// mark is watched

func MarkMovieWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode(params["id"])
}

// Delete a movie

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

// Delete all movie

func DeleteALLMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := deleteAllMovie()
	json.NewEncoder(w).Encode(count)
}

// serve home route
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>This is the homepage ... root... API</h1>"))
}
