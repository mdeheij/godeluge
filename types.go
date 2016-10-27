package godeluge

import (
	"encoding/json"
)

//STATUSTYPES is an enum of all possible types for torrent status requests to Deluge
var STATUSTYPES = [...]string{"queue", "name", "total_size", "state", "progress", "num_seeds", "total_seeds", "num_peers", "total_peers", "download_payload_rate", "upload_payload_rate", "eta", "ratio", "distributed_copies", "is_auto_managed", "time_added", "tracker_host", "save_path", "total_done", "total_uploaded", "max_download_speed", "max_upload_speed", "seeds_peers_ratio"}

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

// type oldTorrentStatus struct {
// 	Name                string
// 	Progress            float64
// 	ETA                 float64
// 	Queue               interface{}
// 	TotalSize           interface{} `json: "total_size"`
// 	State               interface{}
// 	NumberSeeds         interface{} `json: "num_seeds"`
// 	TotalSeeds          interface{} `json: "total_seeds"`
// 	NumPeers            interface{} `json: "num_peers"`
// 	TotalPeers          interface{} `json: "total_peers"`
// 	DownloadPayloadRate interface{} `json: "download_payload_rate"`
// 	UploadPayloadRate   interface{} `json: "upload_payload_rate"`
// 	Ratio               interface{}
// 	DistributedCopies   interface{} `json: "distributed_copies"`
// 	IsAutoManaged       interface{} `json: "is_auto_managed"`
// 	TimeAdded           interface{} `json: "time_added"`
// 	TrackerHost         interface{} `json: "tracker_host"`
// 	SavePath            interface{} `json: "save_path"`
// 	TotalDone           interface{} `json: "total_done"`
// 	TotalUploaded       interface{} `json: "total_uploaded"`
// 	MaxDownloadSpeed    interface{} `json: "max_download_speed"`
// 	MaxUploadSpeed      interface{} `json: "max_upload_speed"`
// 	SeedsPeerRatio      interface{} `json: "seeds_peers_ratio"`
// }

//TorrentStatus is the response of GetTorrentStatus and contains information about the requested torrent
type TorrentStatus struct {
	Name                string  `json:"name"`
	Progress            float64 `json:"progress"`
	ETA                 float64 `json:"eta"`
	State               string  `json:"state"`
	NumPeers            int     `json:"num_peers"`
	NumSeeds            int     `json:"num_seeds"`
	TotalPeers          int     `json:"total_peers"`
	TotalSeeds          int     `json:"total_seeds"`
	SeedsPeersRatio     float64 `json:"seeds_peers_ratio"`
	MaxDownloadSpeed    int     `json:"max_download_speed"`
	MaxUploadSpeed      int     `json:"max_upload_speed"`
	TimeAdded           float64 `json:"time_added"`
	TotalUploaded       int     `json:"total_uploaded"`
	TotalDone           int64   `json:"total_done"`
	TotalSize           int64   `json:"total_size"`
	DistributedCopies   float64 `json:"distributed_copies"`
	TrackerHost         string  `json:"tracker_host"`
	SavePath            string  `json:"save_path"`
	IsAutoManaged       bool    `json:"is_auto_managed"`
	Queue               int     `json:"queue"`
	Ratio               float64 `json:"ratio"`
	DownloadPayloadRate int     `json:"download_payload_rate"`
	UploadPayloadRate   int     `json:"upload_payload_rate"`
}
