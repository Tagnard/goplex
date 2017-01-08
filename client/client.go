package agent

import (
	"net/http"
	"fmt"
	"log"
	"crypto/tls"
	"io/ioutil"
	"encoding/xml"
	"errors"

	"github.com/tagnard/goplex/internal/models"
	"github.com/tagnard/goplex/internal/filters"
)

type Client struct {
	Host      string
	Port      int
	AutoToken string
	HTTPS     bool

	webclient *http.Client
}

func NewClient(host string, port int, authtoken string, https bool, allowinsecure bool) *Client {
	tr := &http.Transport{}
	if allowinsecure {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	} else {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: false}
	}

	return &Client{
		Host: host,
		Port: port,
		AutoToken: authtoken,
		HTTPS: https,
		webclient: &http.Client{Transport: tr},
	}
}

func (pc *Client) Call(path string) (error, *models.MediaContainer) {
	uri := ""
	if pc.HTTPS {
		uri = fmt.Sprintf("https://%v:%v%v", pc.Host, pc.Port, path)
	} else {
		uri = fmt.Sprintf("http://%v:%v%v", pc.Host, pc.Port, path)
	}

	log.Println(uri)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return err, nil
	}

	req.Header.Add("X-Plex-Token", pc.AutoToken)
	resp, err := pc.webclient.Do(req)
	if err != nil {
		return err, nil
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err, nil

	}

	mc := &models.MediaContainer{}

	err = xml.Unmarshal(body, mc)
	if err != nil {
		return err, nil
	}

	return nil, mc
}

func (pc *Client) Callf(format string, a ...interface{}) (error, *models.MediaContainer) {
	return pc.Call(fmt.Sprintf(format, a...))
}

func (pc *Client) Sessions() (error, *models.MediaContainer) {
	err, mc := pc.Call("/status/sessions")
	return err, mc
}

func (pc *Client) SessionMetrics() (error, models.SessionMetrics) {
	err, sessions := pc.Sessions()
	if err != nil {
		return errors.New("Error fetching sessions"), models.SessionMetrics{}
	}

	metrics := models.SessionMetrics{}
	metrics.Total = sessions.Size

	for _, video := range sessions.Video {
		if video.Player.State == "playing" {
			metrics.Active++
		}

		if video.Player.State == "paused" {
			metrics.Inactive++
		}
	}
	return nil, metrics
}

func (pc *Client) Sections() (error, *models.MediaContainer) {
	err, mc := pc.Call("/library/sections/")
	return err, mc
}

func (pc *Client) Section(id int, filter filters.Filter) (error, *models.MediaContainer) {
	err, mc := pc.Callf("/library/sections/%d/%v", id, filter)
	return err, mc
}

func (pc *Client) SectionMetrics(id int) (error, *models.SectionMetrics) {
	err, all := pc.Callf("/library/sections/%d/%v", id, filters.All)
	if err != nil {
		return err, nil
	}

	err, unwatched := pc.Callf("/library/sections/%d/%v", id, filters.Unwatched)
	if err != nil {
		return err, nil
	}
	return err, models.NewSectionMetrics(id, all.Size, unwatched.Size)
}
