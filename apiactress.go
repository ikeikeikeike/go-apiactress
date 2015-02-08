package apiactress

import (
	"encoding/json"
	"io"
	"net/url"
	"os"

	behavior "github.com/ikeikeikeike/gopkg/net/http"
)

const EndPoint = "http://apiactress.appspot.com/api"

var Kunrei = []string{
	"a", "i", "u", "e", "o",
	"ka", "ki", "ku", "ke", "ko",
	"sa", "si", "su", "se", "so",
	"ta", "ti", "tu", "te", "to",
	"na", "ni", "nu", "ne", "no",
	"ha", "hi", "hu", "he", "ho",
	"ma", "mi", "mu", "me", "mo",
	"ya", "yu", "yo",
	"ra", "ri", "ru", "re", "ro",
	"wa",
}

func tee(r io.Reader, debug bool) io.Reader {
	if !debug {
		return r
	}
	return io.TeeReader(r, os.Stdout)
}

type Client struct {
	*behavior.UserBehavior
	Debug bool
}

func NewClient() *Client {
	return &Client{
		UserBehavior: behavior.NewUserBehavior(),
		Debug:        false,
	}
}

type Actress struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Gyou  string `json:"gyou"`
	Thumb string `json:"thumb"`
	Yomi  string `json:"yomi"`
	Oto   string `json:"oto"`
}

type ApiActress struct {
	Count     int       `json:"count"`
	Actresses []Actress `json:"Actresses"`
}

func (c *Client) Fetch(prefix string) (*ApiActress, error) {
	res, err := c.Get(EndPoint + "/1/getdata/" + url.QueryEscape(prefix))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	r := new(ApiActress)
	err = json.NewDecoder(tee(res.Body, c.Debug)).Decode(&r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

/*
	[Deplicated] Server is poverty.
*/
func (c *Client) FetchAll() []*ApiActress {
	var all []*ApiActress
	for _, g := range Kunrei {
		if r, err := c.Fetch(g); err == nil {
			all = append(all, r)
		}
	}

	return all
}

/*
	[Deplicated] Server is poverty.
*/
func AsyncFetchAll() {}
