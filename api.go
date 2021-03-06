package godeluge

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"
)

func (deluge *Deluge) sendCommand(method string, params interface{}) (json.RawMessage, error) {
	reader, writer := io.Pipe()
	var err error

	go func() {
		defer writer.Close()

		var request = Request{Method: method, Params: params, ID: deluge.ID}

		err = json.NewEncoder(writer).Encode(&request)
	}()

	req, err1 := http.NewRequest("POST", deluge.URL, reader)
	if method != "auth.login" {
		req.Header.Add("Cookie", deluge.Session)
	}
	resp, err11 := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}
	if err1 != nil {
		return nil, err1
	}
	if err11 != nil {
		return nil, err11
	}

	defer resp.Body.Close()

	var r Response
	err2 := json.NewDecoder(resp.Body).Decode(&r)
	if err2 != nil {
		return nil, err2
	}

	c := resp.Header.Get("Set-Cookie")
	if c != "" {
		deluge.Session = strings.Split(c, ";")[0]
	}

	var err3 error
	if r.Error.Message != "" {
		if r.Error.Message == "Not authenticated" {
			deluge.login()
			return deluge.sendCommand(method, params)
		}
		err3 = errors.New("error:" + r.Error.Message)
	}

	return r.Result, err3
}

func (deluge *Deluge) login() error {
	result, err := deluge.sendCommand("auth.login", []string{deluge.Password})

	if err != nil {
		return err
	}

	var test bool
	json.Unmarshal(result, &test)

	if !test {
		err = errors.New("Password incorrect")
	}

	return err
}
