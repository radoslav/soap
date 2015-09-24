package soap

import (
	"os"
)

func ExampleSoapFomMTOM() {
	s := []byte("dasdasd---dasdad<><soap:envelope>adsdasfasdfsadfsaf<adsfas><asdfasdf></soap:envelope>--sadadad<>asdasd")
	out, _ := SoapFomMTOM(s)

	os.Stdout.Write(out)
	// Output:
	// <soap:envelope>adsdasfasdfsadfsaf<adsfas><asdfasdf></soap:envelope>
}

func ExampleSoapFomMTOM2() {
	s := `
	HTTP/1.1 500 Internal Server Error
Content-Type: text/xml; charset="utf-8"
X-Backside-Transport: FAIL FAIL
Connection: close
	
<env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
   <env:Body>
      <env:Fault>
         <faultcode>env:Client</faultcode>
         <faultstring>Rejected (from client)</faultstring>
      </env:Fault>
   </env:Body>
</env:Envelope>

<><>
	`

	out, _ := SoapFomMTOM([]byte(s))

	os.Stdout.Write(out)
	// Output:
	// <env:Envelope xmlns:env="http://schemas.xmlsoap.org/soap/envelope/">
	//    <env:Body>
	//       <env:Fault>
	//          <faultcode>env:Client</faultcode>
	//          <faultstring>Rejected (from client)</faultstring>
	//       </env:Fault>
	//    </env:Body>
	// </env:Envelope>
}
