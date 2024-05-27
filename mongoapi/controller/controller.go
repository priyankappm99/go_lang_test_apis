package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/priyankappm99/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectionstring = "mongodb+srv://yourdb:your@cluster0.zbzdsfv.mongodb.net/netflix?retryWrites=true&w=majority&appName=Cluster0"
	dbname           = "netflix"
	colname          = "watchlist"
	productColname   = "products"
)

var (
	collection        *mongo.Collection
	productCollection *mongo.Collection
)

func init() {
	clientOption := options.Client().ApplyURI(connectionstring)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongodb connection success")

	collection = client.Database(dbname).Collection(colname)
	productCollection = client.Database(dbname).Collection(productColname)

	fmt.Println("Collection instance is ready")
}

// Netflix-related functions

func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id:", inserted.InsertedID)
}

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count: ", deleteCount)
}

func deleteAllMovies() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of movies deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}
	defer cur.Close(context.Background())
	return movies
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	count := deleteAllMovies()
	json.NewEncoder(w).Encode(count)
}

// Product-related functions

func insertOneProduct(product model.Product) {
	inserted, err := productCollection.InsertOne(context.Background(), product)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 product in db with id:", inserted.InsertedID)
}

func getAllProducts() []primitive.M {
	cur, err := productCollection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var products []primitive.M
	for cur.Next(context.Background()) {
		var product bson.M
		err := cur.Decode(&product)
		if err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	defer cur.Close(context.Background())
	return products
}

func getOneProduct(productId string) (model.Product, error) {
	var product model.Product
	id, _ := primitive.ObjectIDFromHex(productId)
	filter := bson.M{"_id": id}
	err := productCollection.FindOne(context.Background(), filter).Decode(&product)
	return product, err
}

func GetMyAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allProducts := getAllProducts()
	json.NewEncoder(w).Encode(allProducts)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	product, err := getOneProduct(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product not found")
		return
	}
	json.NewEncoder(w).Encode(product)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var product model.Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	insertOneProduct(product)
	json.NewEncoder(w).Encode(product)
}
