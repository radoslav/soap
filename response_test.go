package soap

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"testing"
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
	// {"XMLName":{"Space":"http://schemas.xmlsoap.org/soap/envelope/","Local":"Envelope"},"ResponseBodyBody":{"XMLName":{"Space":"http://schemas.xmlsoap.org/soap/envelope/","Local":"Body"},"Fault":{"XMLName":{"Space":"http://schemas.xmlsoap.org/soap/envelope/","Local":"Fault"},"Code":"env:Client","String":"Rejected (from client)","Actor":"","Detail":""}}}
	
}

func TestResponseFault(t *testing.T) {
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

	err := CheckFault(soapTest)
	if err == nil {
		t.Error("Expected fault")
	}

}

func TestResponseNotFault(t *testing.T) {
	s := `<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"><soap:Body><ns2:GetMetadataInstanceAllValueResponse xmlns:ns2="http://www.imgw.pl/c26/ws/metadata/universalmetadataservice"><value xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:ns4="http://www.imgw.pl/c26/ws/metadata/metadatatypes" xsi:type="ns4:MetadataCollectionType" partialId="station[349190600]"><collection partialId=".core"><values partialId=".nazwa"><value xmlns:ns5="http://www.imgw.pl/c26/ws/metadata/coretypes" xsi:type="ns5:NameStationValueType" value="BIELSKO-BIAĹA" validFrom="1999-06-30T00:00:00.000Z"/></values></collection></value></ns2:GetMetadataInstanceAllValueResponse></soap:Body></soap:Envelope>`
	soapTest := []byte(s)

	err := CheckFault(soapTest)
	if err != nil {
		t.Error("Expected positive")
	}

}
