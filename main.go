package main

import (
	"bufio"
	"os"
	"strings"
)

func write_wrapper(output_file *os.File, airport_code string, rwy_name string, sid_name string, fixs string) {
	output_file.WriteString(airport_code)
	output_file.WriteString(":")
	output_file.WriteString(rwy_name)
	output_file.WriteString(":")
	output_file.WriteString(sid_name)
	output_file.WriteString(":")
	output_file.WriteString(fixs)
	output_file.WriteString("\n")
}

var airport_rwys map[string][]string

func init_airport_rwys() {
	airport_rwys = make(map[string][]string)
	airport_rwys["RJAH"] = []string{"03L", "21R", "03R", "21L"}
	airport_rwys["RJAW"] = []string{"07", "25"}
	airport_rwys["RJBB"] = []string{"06R", "24L", "06L", "24R"}
	airport_rwys["RJBE"] = []string{"09", "27"}
	airport_rwys["RJCM"] = []string{"18", "36"}
	airport_rwys["RJCN"] = []string{"08", "26"}
	airport_rwys["RJCW"] = []string{"08", "26"}
	airport_rwys["RJDA"] = []string{"13", "31"}
	airport_rwys["RJDT"] = []string{"14", "32"}
	airport_rwys["RJDU"] = []string{"18", "36"}
	airport_rwys["RJEO"] = []string{"13", "31"}
	airport_rwys["RJFA"] = []string{"12", "30"}
	airport_rwys["RJFC"] = []string{"14", "32"}
	airport_rwys["RJFF"] = []string{"16", "34"}
	airport_rwys["RJFM"] = []string{"09", "27"}
	airport_rwys["RJFO"] = []string{"09", "19"}
	airport_rwys["RJFS"] = []string{"11", "29"}
	airport_rwys["RJFT"] = []string{"07", "25"}
	airport_rwys["RJFU"] = []string{"14", "32"}
	airport_rwys["RJFY"] = []string{"08R", "26L", "06L", "26R"}
	airport_rwys["RJKN"] = []string{"01", "19"}
	airport_rwys["RJNA"] = []string{"16", "34"}
	airport_rwys["RJNO"] = []string{"08", "26"}
	airport_rwys["RJNS"] = []string{"12", "30"}
	airport_rwys["RJNT"] = []string{"02", "20"}
	airport_rwys["RJNW"] = []string{"07", "25"}
	airport_rwys["RJNY"] = []string{"09", "27"}
	airport_rwys["RJOA"] = []string{"10", "28"}
	airport_rwys["RJOK"] = []string{"14", "32"}
	airport_rwys["RJOM"] = []string{"14", "32"}
	airport_rwys["RJOO"] = []string{"14R", "32L", "14L", "32R"}
	airport_rwys["RJOS"] = []string{"11", "29"}
	airport_rwys["RJOT"] = []string{"08", "26"}
	airport_rwys["RJOW"] = []string{"11", "29"}
	airport_rwys["RJSA"] = []string{"06", "24"}
	airport_rwys["RJSF"] = []string{"01", "19"}
	airport_rwys["RJSK"] = []string{"10", "28"}
	airport_rwys["RJSS"] = []string{"09", "27", "12", "30"}
	airport_rwys["RJTK"] = []string{"02", "20"}
	airport_rwys["RJTT"] = []string{"16L", "34R", "16R", "34L", "04", "22", "05", "23"}
	airport_rwys["RJTU"] = []string{"01", "19"}
	airport_rwys["ROAH"] = []string{"18L", "36R", "18R", "36L"}
	airport_rwys["ROIG"] = []string{"04", "22"}
	airport_rwys["RORY"] = []string{"14", "32"}
}

func all_rwy_star(output_file *os.File, airport_code string, star_name string, fixs string) {
	var i = 0
	for i < len(airport_rwys[airport_code]) {
		output_file.WriteString("STAR:")
		write_wrapper(output_file, airport_code, airport_rwys[airport_code][i], star_name, fixs)
		i++
	}
}

func main() {
	sid_dat, sid_open_err := os.Open("data/PSSSID.dat")
	if sid_open_err != nil {
		println("data/PSSSID.datを開けませんでした")
		os.Exit(1)
	}
	star_dat, star_open_err := os.Open("data/PSSSTAR.dat")
	if star_open_err != nil {
		println("data/PSSSTAR.datを開けませんでした。")
		os.Exit(1)
	}
	output_file, output_create_err := os.Create("output.txt")
	if output_create_err != nil {
		println("output.txtの作成に失敗しました。")
		os.Exit(1)
	}
	println("running...")
	init_airport_rwys()
	sid_scanner := bufio.NewScanner(sid_dat)
	for sid_scanner.Scan() {
		line := sid_scanner.Text()
		if strings.HasPrefix(line, "[RP") {
			break
		}
		if strings.HasPrefix(line, "[RJ") || strings.HasPrefix(line, "[RO") {
			airport_code := line[1:5]
			var rwy_name string
			if strings.HasPrefix(line[13:], "ALL") {
				rwy_name = line[13 : 13+strings.IndexRune(line[13:], ' ')]
			} else {
				rwy_name = line[15 : 15+strings.IndexRune(line[15:], ' ')]
			}
			sid_name := line[6 : 6+strings.IndexRune(line[6:], '/')]
			sid_name = strings.TrimSpace(sid_name)
			var fixs string
			for {
				sid_scanner.Scan()
				line = sid_scanner.Text()
				if len(line) == 0 {
					break
				}
				if len(strings.TrimSpace(line[14:19])) != 0 {
					fixs += strings.TrimSpace(line[14:19]) + ","
				}
			}
			fixs = strings.ReplaceAll(fixs, ",", " ")
			output_file.WriteString("SID:")
			write_wrapper(output_file, airport_code, rwy_name, sid_name, fixs)
		}
	}
	star_scanner := bufio.NewScanner(star_dat)
	for star_scanner.Scan() {
		line := star_scanner.Text()
		if strings.HasPrefix(line, "[RP") {
			break
		}
		if strings.HasPrefix(line, "[RJ") || strings.HasPrefix(line, "[RO") {
			airport_code := line[1:5]
			var rwy_name string
			if strings.HasPrefix(line[13:], "ALL") {
				rwy_name = line[13 : 13+strings.IndexRune(line[13:], ' ')]
			} else {
				rwy_name = line[15 : 15+strings.IndexRune(line[15:], ' ')]
			}
			star_name := line[6 : 6+strings.IndexRune(line[6:], '/')]
			star_name = strings.TrimSpace(star_name)
			var fixs string
			for {
				star_scanner.Scan()
				line = star_scanner.Text()
				if len(line) == 0 {
					break
				}
				if len(strings.TrimSpace(line[14:19])) != 0 {
					fixs += strings.TrimSpace(line[14:19]) + ","
				}
			}
			fixs = strings.ReplaceAll(fixs, ",", " ")
			if !strings.HasPrefix(rwy_name, "ALL") {
				fixs = strings.ReplaceAll(fixs, ",", " ")
				output_file.WriteString("STAR:")
				write_wrapper(output_file, airport_code, rwy_name, star_name, fixs)
			} else {
				all_rwy_star(output_file, airport_code, star_name, fixs)
			}
		}
	}
}
