package main

import (
	"fmt"
	"net"
	"os"
)

var rcon_password string

type Server struct {
	name     string
	ip       string
	port     string
	password string
}

/**
 * Force wallfly on the supplied server to say something
 */
func StuffWallflySay(srv Server, message string) {
	if srv.password == "" {
		srv.password = rcon_password
	}

	bytes := []byte{0xff, 0xff, 0xff, 0xff, 'r', 'c', 'o', 'n', ' '}
	bytes = append(bytes, srv.password...)
	cmd := []byte(" sv !stuff WallFly[BZZZ] say ")
	bytes = append(bytes, cmd...)
	bytes = append(bytes, message...)
	bytes = append(bytes, 0x00)

	conn, err := net.Dial("udp", fmt.Sprintf("%s:%s", srv.ip, srv.port))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Stuffing wallfly on %s\n", srv.name)
	fmt.Fprintf(conn, string(bytes))
	conn.Close()
}

/**
 * All servers share a common rcon password, pass it
 * via environment variable
 */
func main() {
	var servers = []Server{
		{"aq2ffa", "23.227.170.214", "27926", ""},
		{"aq2tdm", "23.227.170.214", "27927", ""},
		{"awaken", "23.227.170.214", "27953", ""},
		{"bang", "23.227.170.214", "27951", ""},
		{"btfffa", "23.227.170.214", "27910", ""},
		{"btftdm", "23.227.170.214", "27911", ""},
		{"classic", "23.227.170.222", "27927", ""},
		{"coop", "23.227.170.214", "27932", ""},
		{"ctf", "23.227.170.222", "27914", ""},
		{"ctf-match", "23.227.170.222", "27924", ""},
		{"dm", "23.227.170.222", "27916", ""},
		{"dm8", "23.227.170.222", "27922", ""},
		{"eom", "23.227.170.214", "27919", ""},
		{"ffa", "23.227.170.222", "27909", ""},
		{"freeze", "23.227.170.222", "27928", ""},
		{"hardcore", "23.227.170.222", "27929", ""},
		{"hell", "23.227.170.222", "27666", ""},
		{"jail", "23.227.170.222", "27921", ""},
		{"jump", "23.227.170.222", "27940", ""},
		{"lithctf", "23.227.170.214", "27923", ""},
		{"lithium", "23.227.170.214", "27924", ""},
		{"mutant", "23.227.170.222", "27910", ""},
		{"nadl", "23.227.170.222", "27990", ""},
		{"nail", "23.227.170.214", "27945", ""},
		{"pygmy", "23.227.170.222", "27918", ""},
		{"ra2", "23.227.170.222", "27915", ""},
		{"ra2-match", "23.227.170.222", "27925", ""},
		{"ra2-match2", "23.227.170.222", "27935", ""},
		{"ra2-railz", "23.227.170.222", "27926", ""},
		{"railz", "23.227.170.222", "27911", ""},
		{"railztdm", "23.227.170.214", "27914", ""},
		{"coop-rogue", "23.227.170.214", "27934", ""},
		{"rogue", "23.227.170.222", "27947", ""},
		{"rox", "23.227.170.222", "27923", ""},
		{"sheridan", "23.227.170.222", "27913", ""},
		{"tdm-noalias", "23.227.170.222", "27933", ""},
		{"tdm-novice", "23.227.170.222", "27932", ""},
		{"tourney", "23.227.170.222", "27930", ""},
		{"tourney2", "23.227.170.222", "27931", ""},
		{"ts500", "23.227.170.222", "27999", ""},
		{"tstournament01", "23.227.170.222", "27991", ""},
		{"tstournament02", "23.227.170.222", "27992", ""},
		{"tstournament04", "23.227.170.222", "27994", ""},
		{"tstournament05", "23.227.170.222", "27995", ""},
		{"tstournament06", "23.227.170.222", "27996", ""},
		{"tstournament07", "23.227.170.222", "27997", ""},
		{"tstournament08", "23.227.170.222", "27998", ""},
		{"vanilla", "23.227.170.222", "27912", ""},
		{"vanilla2", "23.227.170.222", "27920", ""},
		{"wod", "23.227.170.214", "27925", ""},
		{"xatrix", "23.227.170.222", "27919", ""},
		{"xquake", "23.227.170.222", "27917", ""},
		{"xrazy", "23.227.170.214", "27929", ""},
		{"xtra", "23.227.170.222", "27949", ""},
		{"coop-zaero", "23.227.170.214", "27933", ""},
		{"zaero", "23.227.170.222", "27948", ""},
		{"zigbot", "23.227.170.222", "27950", ""},
	}

	rcon_password = os.Getenv("RCON")

	msg := "HEY YOU! Join the TS500 and be the first to get 500 frags in a HUGE FFA! GOTO TS500 NOW"
	//msg := "This Saturday! Back-to-back-to-back MASSIVE FFA events! Games start at 13 ET / 11 PT on TS500 - practice up and GOTO TS500"

	for _, srv := range servers {
		StuffWallflySay(srv, msg)
	}
}
