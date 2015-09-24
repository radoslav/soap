package soap

import (
	"encoding/json"
	"encoding/xml"
	//		"fmt"
	"os"
)

func ExampleResponse() {
	s := `
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
   <env:Body>
      <env:Fault>
         <faultcode>env:Client</faultcode>
         <faultstring>Rejected (from client)</faultstring>
      </env:Fault>
   </env:Body>
</env:Envelope>`

	soapTest := []byte(s)

	env := ResponseEnvelope{}

	xml.Unmarshal(soapTest, &env)

	json, _ := json.Marshal(env)
	os.Stdout.Write(json)

	// Output:
	// {"XMLName":{"Space":"http://schemas.xmlsoap.org/soap/envelope/","Local":"Envelope"},"ResponseBodyBody":{"XMLName":{"Space":"http://schemas.xmlsoap.org/soap/envelope/","Local":"Body"},"Fault":{"XMLName":{"Space":"http://schemas.xmlsoap.org/soap/envelope/","Local":"Fault"},"Code":"env:Client","String":"Rejected (from client)","Actor":null,"Detail":null}}}

}
