// This Go application takes a user inputted phone number and sends it a text message via Twilio API
// ---------------------------------------------------------------------------------------------------

package main

import (
	"fmt"
 	"net/http" 
 	"net/url"
 	"strings"
)

func main() {

	// Set variables
 	accountSid := "" // Your Twilio accountSid
 	authToken := "" // Your Twilio AuthToken
 	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// Intake phone number
    fmt.Printf("\nEnter target phone number:   ")
	var input string
    fmt.Scanln(&input)
    fmt.Print("Number inputted: ",input,"\n\n")

 	// Build message 
 	v := url.Values{}
 	v.Set("To",input)
 	v.Set("From","12223334444") // Your Twilio number
 	v.Set("Body","This is a successful live message!!") // Your Text Body
 	//	v.Set("MediaUrl","") // Your Media (Optional)
  	rb := *strings.NewReader(v.Encode())

	// Create Client
	client := &http.Client{}
 	req, _ := http.NewRequest("POST", urlStr, &rb)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Make request
	resp, _ := client.Do(req)
	fmt.Println(resp.Status)
	
}


	


