package api

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	geoip "github.com/oschwald/geoip2-golang"
)

var db *geoip.Reader

type GeoInfo struct {
	CityName        string `json:"cityName"`
	SubdivisionName string `json:"subdivisionName"`
	CountryName     string `json:"countryName"`
	ISOCountryCode  string `json:"isoCountryCode"`
	TimeZone        string `json:"timeZone"`
}

func OpenDB() {
	var err error

	// Open GeoLite2 DB
	db, err = geoip.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatalln(err)
	}
}

func queryIP(ip net.IP) *GeoInfo {
	record, err := db.City(ip)
	if err != nil {
		log.Println("Unalble to parse ip info")
		log.Println(err)

		return nil
	}

	ipGeoInfo := &GeoInfo{
		CityName:        record.City.Names["en"],
		SubdivisionName: record.Subdivisions[0].Names["en"],
		CountryName:     record.Country.Names["en"],
		ISOCountryCode:  record.Country.IsoCode,
		TimeZone:        record.Location.TimeZone,
	}

	return ipGeoInfo
}

func GetMyIPGeoInfo(w http.ResponseWriter, r *http.Request) {
	ip := net.ParseIP(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0])

	ipGeoInfo := queryIP(ip)
	if ipGeoInfo == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
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

func GetIPGeoInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ip := net.ParseIP(vars["ip"])

	if ip == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ipGeoInfo := queryIP(ip)
	if ipGeoInfo == nil {
		w.WriteHeader(http.StatusBadRequest)
		return
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
