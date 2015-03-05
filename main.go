package main

import (
	"errors"
	"os"
	"strconv"

	"log"

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
		return errors.New("Can't convert " + episode + "to string")
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

func Search(title string) (Anime, error) {
	for _, anime := range animeList {
		if anime.Title == title {
			return anime, nil
		}
	}
	err := errors.New("Can't find anime: " + title)
	return Anime{bson.NewObjectId(), "err", 0}, err
}

func Increment(title string) error {
	a, err := Search(title)
	if err != nil {
		return errors.New("Can't find anime: " + title)
	}

	a.Increment()

	err = db.Update(
		bson.M{"title": a.Title},
		bson.M{"$set": bson.M{"episode": a.Episode}})
	if err != nil {
		return err
	}
	return nil

}

// Searches for the title, sees what id it has, and then uses it to remove.
func Remove(title string) error {
	a, err := Search(title)
	if err != nil {
		return err
	}

	if err = db.RemoveId(a.Id); err != nil {
		return err
	}
	return nil
}

func main() {
	session, err := mgo.Dial(os.Getenv("MONGO_URL"))
	if err != nil {
		log.Println("Can't connect to mongo")
		os.Exit(1)
	}
	defer session.Close()
	session.SetSafe(&mgo.Safe{})

	// actualy a collection
	db = session.DB("test").C("anime")

	// Populate the animeList
	if err := db.Find(nil).All(&animeList); err != nil {
		log.Println("Can't find any animes")
	}

	// Run the cli app (cli.go)
	RunCli()
}
