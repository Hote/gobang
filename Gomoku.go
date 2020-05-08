package main

import (
	"fmt"
	"net/http"
	"strconv"
	//	"strings"
)

func main() {
	http.HandleFunc("/", indcx)
	http.ListenAndServe(":8080", nil)
}

func indcx(w http.ResponseWriter, r *http.Request) {
	num, err := strconv.Atoi(r.FormValue("num"))
	if err != nil {
		num := 0
	}

	//default position is
	var pos string = r.FormValue("pos")

	//black
	var b_pos string = r.FormValue("b_pos")

	//white
	var w_pos string = r.FormValue("w_pos")

	if num > 0 {
		if num%2 == 0 {
			if w_pos != "" {
				w_pos += ","
			}
			w_pos += w_pos
		} else {
			if b_pos != "" {
				b_pos += ","
			}
			b_pos += pos

		}
	}

	// web page initiation
	html := "<!DOCTYPE html><html lang=\"zh-tw\">"

	html += SecHead()

	html += SecBody(num, b_pos, w_pos)

	html += "</html>"
	fmt.Fprint(w, html)
}

func SecHead() string {
	head := "<head>"
	head += "<title> Amos Go practice</title>"
	head += "<meta charset=\"utf-8\">"
	return head
}

func SecBody(num int, b_pos string, w_pos string) string {
	body := "<body>"
	body += "<div class=\"container\">"

	return body

}
