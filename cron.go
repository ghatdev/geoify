package main

import (
	"archive/tar"
	"compress/gzip"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	cron "gopkg.in/robfig/cron.v2"
)

func updateDB() {
	resp, err := http.Get("http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz")
	if err != nil {
		log.Fatalln("DB update failed!")
	}

	defer resp.Body.Close()

	gzf, err := gzip.NewReader(resp.Body)
	defer gzf.Close()

	tarf := tar.NewReader(gzf)
	i := 0
	for {
		h, err := tarf.Next()
		if err != nil {
			break
		}

		if h.Typeflag == tar.TypeDir {
			continue
		}

		name := h.Name[strings.LastIndex(h.Name, "/")+1:]

		if name == "GeoLite2-City.mmdb" {
			buffer, err := ioutil.ReadAll(tarf)
			if err != nil {
				break
			}

			err = ioutil.WriteFile(name, buffer, 0644)
			if err != nil {
				break
			}

			return
		}

		i++
		if i > 4 {
			break
		}
	}

	log.Fatalln("DB update failed!2")
}

func addCron() {
	c := cron.New()
	c.AddFunc("@daily", updateDB)
	c.Start()
}
