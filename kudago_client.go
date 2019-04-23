package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type Event struct {
	Title    string
	Place    string
	Location string
	Site     string
}

type resultResponse struct {
	Results []eventResponse `json:"results"`
}

type eventResponse struct {
	Title    string           `json:"title"`
	Place    placeResponse    `json:"place"`
	Location locationResponse `json:"location"`
	Site     string           `json:"site_url"`
}

type locationResponse struct {
	Slug string `json:"slug"`
}

type placeResponse struct {
	ID int `json:"id"`
}

func NewKudaGoClient(url string) KudaGoClient {
	return &kudago{
		Url: url,
	}
}

type KudaGoClient interface {
	GetRandomEvent(city string, user string) (*Event, error)
	GetEventForToday(city string, user string) (*Event, error)
	GetEventForTomorrow(city string, user string) (*Event, error)
}

type kudago struct {
	Url string
}

func (k *kudago) GetRandomEvent(city string, user string) (*Event, error) {
	return k.getEvent(city, user, time.Now().Unix(), time.Now().AddDate(0, 6, 0).Unix())
}
func (k *kudago) GetEventForToday(city string, user string) (*Event, error) {
	return k.getEvent(city, user, time.Now().Unix(), time.Now().AddDate(0, 0, 1).Unix())
}
func (k *kudago) GetEventForTomorrow(city string, user string) (*Event, error) {
	return k.getEvent(city, user, time.Now().AddDate(0, 0, 1).Unix(), time.Now().AddDate(0, 0, 2).Unix())
}

func (k *kudago) getEvent(city string, user string, from int64, until int64) (*Event, error) {
	url := fmt.Sprintf("%s/events?fields=location,title,place,site_url&location=%s&actual_since=%d&actual_until=%d&page_size=100", k.Url, city, from, until)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var resultResp resultResponse
	err = json.Unmarshal(responseData, &resultResp)
	if err != nil {
		return nil, err
	}
	event := resultResp.Results[rand.Intn(len(resultResp.Results))]
	e := Event{
		Title:    event.Title,
		Location: event.Location.Slug,
		Place:    string(event.Place.ID),
		Site:     event.Site,
	}
	return &e, nil
}
