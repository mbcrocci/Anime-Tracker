package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

var db *mgo.Collection
var animeList []Anime

type Anime struct {
	Id      bson.ObjectId `bson:"_id"`
	Title   string        `bson:"title"`
	Episode int           `bson:"episode"`
}

func (a *Anime) Increment() {
	a.Episode += 1
}

func (a *Anime) ChangeTitle(title string) {
	a.Title = title
}

func (a *Anime) ChangeEpisode(episode int) {
	a.Episode = episode
}

func addAnime(title, episode string) error {
	ep, err := strconv.Atoi(episode)
	if err != nil {
		return errors.New("Can't conver " + episode + "to string")
	}
	err = db.Insert(Anime{
		Id:      bson.NewObjectId(),
		Title:   title,
		Episode: ep,
	})
	if err != nil {
		return err
	}
	return nil
}

func Increment(title string) error {
	for _, anime := range animeList {
		if anime.Title == title {
			anime.Increment()

			err := db.Update(
				bson.M{"title": anime.Title},
				bson.M{"$set": bson.M{"episode": anime.Episode}})
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("Can't find anime: " + title)
}

func Remove(title string) error {
	for _, anime := range animeList {
		if anime.Title == title {
			if err := db.RemoveId(anime.Id); err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("Can't find anime: " + title)
}

func main() {
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		fmt.Printf("Can't connect to mongo\n")
		os.Exit(1)
	}
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	// actualy a collection
	db = session.DB("test").C("anime")

	// Populate the animeList
	if err := db.Find(nil).All(&animeList); err != nil {
		fmt.Printf("Can't find any animes")
	}

	// Run the cli app (cli.go)
	RunCli()
}
