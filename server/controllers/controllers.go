package controllers

import (
	"context"
	"encoding/json"
	"go_mefo-music-server/server/db"
	"go_mefo-music-server/server/models"
	res "go_mefo-music-server/utils"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func AddSong(w http.ResponseWriter, r *http.Request) {
	collection := db.ConnectDatabase()
	decoder := json.NewDecoder(r.Body)
	var song models.Song
	err := decoder.Decode(&song)
	if err != nil {
		panic(err)
	}
	_, error := collection.InsertOne(context.TODO(), bson.M{"title": song.Title, "src": song.Src})
	if error != nil {
		res.JSON(w, 501, "Create song failed")
	}
	res.JSON(w, 200, "Song created successfully")
}

func GetAllSongs(w http.ResponseWriter, r *http.Request) {
	collection := db.ConnectDatabase()

	var result []bson.M
	data, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		res.JSON(w, 500, "Internal Server Error")
		return
	}

	defer data.Close(context.Background())
	for data.Next(context.Background()) {
		var elem bson.M
		err := data.Decode(&elem)

		if err != nil {
			res.JSON(w, 500, "Internal Server Error")
			return
		}

		result = append(result, elem)
	}

	res.JSON(w, 200, result)
}
