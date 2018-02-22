package api

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"

	geoip "github.com/oschwald/geoip2-golang"
)

var db *geoip.Reader

type GeoInfo struct {
	ISOCountryCode string `json:"isoCountryCode"`
	TimeZone       string `json:"timeZone"`
}

func init() {
	var err error

	// Open GeoLite2 DB
	db, err = geoip.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatalln(err)
	}
}

func GetIPGeoInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := net.ParseIP(vars["ip"])

	record, err := db.City(ip)
	if err != nil {
		log.Println("Unalble to parse ip info")
		log.Println(err)

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ipGeoInfo := &GeoInfo{
		ISOCountryCode: record.Country.IsoCode,
		TimeZone:       record.Location.TimeZone,
	}

	response, err := json.Marshal(ipGeoInfo)
	if err != nil {
		log.Println("Failed to encode info")
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func GetIPCityInfo(w http.ResponseWriter, r *http.Request) {

}

func GetIPCountryInfo(w http.ResponseWriter, r *http.Request) {

}
