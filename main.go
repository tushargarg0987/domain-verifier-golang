package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
    fmt.Println("--- DOMAIN VERIFIER ---")
    // fmt.Println(" ---   ---   - -   ---   ---   -  ")
    // fmt.Println("|   |  ---   - -   ---   ---   -  ")

    fmt.Println("Enter the domain to verify or press CTRL+C to exit")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		checkDomain(scanner.Text())
	}

	if err:= scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return
}

func checkDomain(domain string){
	var hasMx, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords,err := net.LookupMX(domain)

	if err!= nil {
        log.Fatal(err)
    }

	if len(mxRecords) >0 {
		hasMx = true
	}

	txtRecord,err := net.LookupTXT(domain)
	if err!= nil {
        log.Fatal(err)
    }

	for _, record := range txtRecord{
		if strings.HasPrefix(record, "v=spf1"){
			spfRecord = record
			hasSPF = true
		}
	}

	dmarkRecords, err := net.LookupTXT("_dmarc." + domain)
	if err!= nil {
        log.Fatal(err)
    }

	for _, record := range dmarkRecords{
		if strings.HasPrefix(record, "v=DMARC1"){
            dmarcRecord = record
            hasDMARC = true
        }
	}

	fmt.Printf("domain - %v, \nhas mx record - %v, \nhas spf - %v, \nspf record - %v, \nhas dmarc - %v, \ndmarc record - %v\n\nEnter the domain to verify\n",domain,hasMx,hasSPF,spfRecord,hasDMARC,dmarcRecord)
	return

}