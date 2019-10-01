package main

import (
	"fmt"
	//"reflect"

	//"github.com/bitly/go-simplejson"
	"github.com/verdverm/frisby"
)

func main(){
	fmt.Println("Frisby!")
	var json = `{"id":11,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}`
	frisby.Create("Test GET Go homepage").
	Get("http://localhost:8080/api/entries").
	Send().
	ExpectStatus(200)
	//ExpectContent("[{"id":11,"first_name":"xyz","last_name":"pqr","email_address":"xyz@pqr.com","phone_number":"1234567890"}]")
	frisby.Global.PrintReport()
	frisby.Create("Test GET Go homepage").
	Get("http://localhost:8080/api/entry?id=11").
	Send().
	ExpectStatus(200).
	ExpectContent(json)
	frisby.Global.PrintReport()
}
	
	