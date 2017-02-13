package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"time"
)

type lang struct {
	name string
	rank int
}

func main() {
	fmt.Println("Starting application")
	slices()
	structs()
}

type FakePhoto struct {
	Id          json.RawMessage `json:"id"`
	Description string          `json:"title"`
}

type TumblrPhoto struct {
	Id          json.RawMessage `json:"id"`
	Description string          `json:"summary"`
}

type TumblrResponse struct {
	Response TumblrPosts `json:"response"`
}
type TumblrPosts struct {
	Posts []TumblrPhoto `json:"posts"`
}

type unifiedPhoto struct {
	Id          json.RawMessage
	Description string
}

func structs() {
	c := &http.Client{
		Timeout: 5 * time.Second,
	}

	fp := getFakePhoto(c)
	fmt.Printf("Fake photo: %+v\n", unifiedPhoto(fp))
	tp := getTumblrPhoto(c)
	fmt.Printf("Tumblr photo: %+v\n", unifiedPhoto(tp))

}

func getFakePhoto(c *http.Client) FakePhoto {
	var fRes []FakePhoto
	resp, err := c.Get("https://jsonplaceholder.typicode.com/photos")
	//resp, err := c.Get("https://api.flickr.com/services/rest/?method=flickr.photos.getInfo&api_key=f17639e3d18eca2dea2f321aaf3e2e84&photo_id=32070157923&format=json&nojsoncallback=1")
	if err != nil {
		fmt.Printf("Error calling Flickr's API: %v\n", err)
		return FakePhoto{}
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&fRes)
	if err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return FakePhoto{}
	}
	return fRes[0]

}

func getTumblrPhoto(c *http.Client) TumblrPhoto {
	var tPhoto TumblrResponse
	resp, err := c.Get("https://api.tumblr.com/v2/blog/pitchersandpoets.tumblr.com/posts/photo?api_key=fuiKNFp9vQFvjLNvx4sUwti4Yb5yGutBN4Xh10LXZhhRKjWlV4&tag=new+york+yankees")
	if err != nil {
		fmt.Printf("Error calling Tumblr's API: %v\n", err)
		return TumblrPhoto{}
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&tPhoto)
	if err != nil {
		fmt.Printf("Error decoding response: %v\n", err)
		return TumblrPhoto{}
	}

	return tPhoto.Response.Posts[0]
}

func slices() {
	langs := []lang{
		{"Java", 1},
		{"Javascript", 7},
		{"C", 2},
		{"Go", 14},
		{"Python", 5},
	}
	fmt.Printf("Initial slice: %v\n", langs)

	sort.Slice(langs, func(i, j int) bool {
		return langs[i].rank < langs[j].rank
	})
	fmt.Printf("Slice sorted after rank: %v\n", langs)

	sort.Sort(byName(langs))
	fmt.Printf("Slice sorted after name: %v\n", langs)

}

type byName []lang

func (langs byName) Len() int {
	return len(langs)
}

func (langs byName) Less(i, j int) bool {
	return langs[i].name < langs[j].name
}

func (langs byName) Swap(i, j int) {
	langs[i], langs[j] = langs[j], langs[i]
}
