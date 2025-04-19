package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type EmailResponse struct {
	Status      string `json:"status"`
	Domain      string `json:"domain"`
	HasMx       bool   `json:"hasMx"`
	IsValid     bool   `json:"isValid"`
	HasSPF      bool   `json:"hasSPF"`
	SpfRecord   string `json:"spfRecord"`
	HasDMARC    bool   `json:"hasDMARC"`
	DmarcRecord string `json:"dmarcRecord"`
	Description string `json:"description"`
}

func main() {
	fmt.Println("starting the server....")

	r := mux.NewRouter()
	r.HandleFunc("/validate/{emailId}", EmailValidator).Methods("GET")
	fmt.Println("Application started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func EmailValidator(w http.ResponseWriter, r *http.Request) {

	hasspf := false
	hasdmarc := false
	isValid := false
	hasMx := false
	description := ""
	spfRecord := ""
	dmarcRecord := ""
	params := mux.Vars(r)

	emailId := params["emailId"]
	w.Header().Set("Content-Type", "application/json")

	if emailId == "" {
		json.NewEncoder(w).Encode(EmailResponse{Status: "Success", IsValid: false})
		return
	}

	mxr, err := net.LookupMX(emailId)
	if len(mxr) > 0 {
		hasMx = true
	}

	resp, err := net.LookupTXT(emailId)

	if err != nil {
		fmt.Println("Email is not valid")
		json.NewEncoder(w).Encode(EmailResponse{Status: "Success", IsValid: false, HasSPF: hasspf, HasDMARC: hasdmarc})
		return
	}

	for _, record := range resp {
		isValid = true
		if strings.HasPrefix(record, "v=spf1") {
			fmt.Println("SPF record found")
			hasspf = true
			spfRecord = record
			description = description + record + "\n"
			break
		}
	}

	dmarcResp, _ := net.LookupTXT("_dmarc." + emailId)

	if len(dmarcResp) > 0 {
		for _, record := range dmarcResp {
			if strings.HasPrefix(record, "v=DMARC1") {
				fmt.Println("DMARC record found")
				hasdmarc = true
				dmarcRecord = record
				description = description + record + "\n"
				break
			}
		}
	}

	json.NewEncoder(w).Encode(EmailResponse{Status: "Success", Domain: emailId, IsValid: isValid, HasMx: hasMx, HasSPF: hasspf, SpfRecord: spfRecord, HasDMARC: hasdmarc, DmarcRecord: dmarcRecord, Description: description})
}
