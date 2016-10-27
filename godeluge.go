package godeluge

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

//NewDeluge creates a new deluge instance
func NewDeluge(url string, password string) (*Deluge, error) {
	var deluge = Deluge{URL: url, Password: password}
	var err error

	err = deluge.login()

	return &deluge, err
}

//GetTorrentStatus returns the current status of a torrent
func (deluge Deluge) GetTorrentStatus(hash string) (TorrentStatus, error) {
	result, err := deluge.sendCommand("web.get_torrent_status", []interface{}{strings.ToLower(hash), STATUSTYPES})
	var i TorrentStatus
	if err != nil {
		return i, err
	}

	err1 := json.Unmarshal(result, &i)
	if err1 != nil {
		return TorrentStatus{}, err1
	}

	if (i == TorrentStatus{}) {
		return i, errors.New("Torrent could not be found in Deluge")
	}

	spew.Dump(i)

	return i, nil
}

//RemoveTorrent removes a torrent from Deluge
func (deluge Deluge) RemoveTorrent(magnet string) error {
	result, err := deluge.sendCommand("core.remove_torrent", []interface{}{strings.ToLower(magnet), true})

	if err != nil {
		return err
	}

	var i bool
	err1 := json.Unmarshal(result, &i)

	if err1 != nil {
		return err1
	}

	if !i {
		return errors.New("Error removing torrent")
	}

	return nil
}

//AddTorrent adds a torrent to Deluge
func (deluge Deluge) AddTorrent(magnet string) error {
	params := []interface{}{[]interface{}{map[string]interface{}{"path": magnet, "options": nil}}}
	result, err := deluge.sendCommand("web.add_torrents", params)
	if err != nil {
		return err
	}

	var i bool
	err1 := json.Unmarshal(result, &i)

	if err1 != nil {
		return err1
	}
	if !i {
		return errors.New("Incorrect Magnet")
	}

	return nil
}
