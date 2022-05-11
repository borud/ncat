package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strings"

	"github.com/adrianmo/go-nmea"
)

var gpsDevice = flag.String("dev", "/dev/ttyUSB1", "path to GPS serial device")

func main() {
	flag.Parse()

	gps, err := os.Open(*gpsDevice)
	if err != nil {
		log.Fatalf("unable to open serial device: %v", err)
	}
	defer gps.Close()

	scanner := bufio.NewScanner(gps)
	for scanner.Scan() {
		line := scanner.Text()

		s, err := nmea.Parse(line)
		if err != nil {
			continue
		}

		switch dd := s.(type) {
		case nmea.RMC:
			log.Printf("time=[%s %s] validity=[%s] lat=[%s] lon=[%s] speed=[%f.2] course=[%f.2] variation=[%f]",
				dd.Date,
				dd.Time,
				dd.Validity,
				nmea.FormatGPS(dd.Latitude),
				nmea.FormatGPS(dd.Longitude),
				dd.Speed,
				dd.Course,
				dd.Variation,
			)

		case nmea.GSV:
			log.Printf("talker=[%s] systemID=[%d] numberSVsInView=[%d]", dd.TalkerID(), dd.SystemID, dd.NumberSVsInView)

		case nmea.GSA:
			log.Printf("talker=[%s] Satelites=[%+v]", dd.TalkerID(), strings.Join(dd.SV, ","))

		}
	}
}
