package godeluge

import (
	"encoding/json"
)

//Deluge instance
type Deluge struct {
	Session  string
	Password string
	URL      string
	ID       int32
}

//Request to deluge
type Request struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
	ID     int32       `json:"id"`
}

//Response from Deluge
type Response struct {
	ID     int32
	Result json.RawMessage
	Error  Error
}

//Error received from Deluge
type Error struct {
	Message string
	Code    int32
}

//TorrentStatus is the response of GetTorrentStatus and contains information about the requested torrent
type TorrentStatus struct {
	Name     string
	Progress float64
	ETA      float64
}
