package soap

import (
	"encoding/xml"
	"fmt"
	"os"
)

const prefix string = ""
const indent string = "   "
const namespaceUniv = "http://www.example.pl/ws/test/universal"

func ExampleBody() {
	env := &Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    namespaceUniv,
	}

	env.Body = &Body{}

	output, err := xml.MarshalIndent(env, prefix, indent)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	// Output:
	// <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:univ="http://www.example.pl/ws/test/universal">
	//    <soapenv:Body></soapenv:Body>
	// </soapenv:Envelope>
}

func ExampleEnvelope() {
	env := &Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    namespaceUniv,
	}

	output, err := xml.MarshalIndent(env, prefix, indent)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	// Output:
	// <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:univ="http://www.example.pl/ws/test/universal"></soapenv:Envelope>
}

func ExampleHeader() {
	env := &Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    namespaceUniv,
		Header:       &Header{},
	}

	output, err := xml.MarshalIndent(env, prefix, indent)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	// Output:
	// <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:univ="http://www.example.pl/ws/test/universal">
	//    <soapenv:Header></soapenv:Header>
	// </soapenv:Envelope>
}

func ExampleWsseSecurity() {
	env := &Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    namespaceUniv,
		Header: &Header{
			WsseSecurity: &WsseSecurity{
				MustUnderstand: "1",
				XmlnsWsse:      "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
				XmlnsWsu:       "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
			},
		},
	}

	output, err := xml.MarshalIndent(env, prefix, indent)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	// Output:
	// <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:univ="http://www.example.pl/ws/test/universal">
	//    <soapenv:Header>
	//       <wsse:Security soapenv:mustUnderstand="1" xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"></wsse:Security>
	//    </soapenv:Header>
	// </soapenv:Envelope>

}

func ExampleUsernameToken() {
	env := &Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    namespaceUniv,
		Header: &Header{
			WsseSecurity: &WsseSecurity{
				MustUnderstand: "1",
				XmlnsWsse:      "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
				XmlnsWsu:       "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
				UsernameToken: &UsernameToken{
					WsuId: "UsernameToken-1",
				},
			},
		},
	}

	output, err := xml.MarshalIndent(env, prefix, indent)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	// Output:
	// <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:univ="http://www.example.pl/ws/test/universal">
	//    <soapenv:Header>
	//       <wsse:Security soapenv:mustUnderstand="1" xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
	//          <wsse:UsernameToken wsu:Id="UsernameToken-1"></wsse:UsernameToken>
	//       </wsse:Security>
	//    </soapenv:Header>
	// </soapenv:Envelope>

}

func ExampleUsername() {
	env := &Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    namespaceUniv,
		Header: &Header{
			WsseSecurity: &WsseSecurity{
				MustUnderstand: "1",
				XmlnsWsse:      "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
				XmlnsWsu:       "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
				UsernameToken: &UsernameToken{
					WsuId:    "UsernameToken-1",
					Username: &Username{},
				},
			},
		},
	}

	env.Header.WsseSecurity.UsernameToken.Username.Value = "test"

	output, err := xml.MarshalIndent(env, prefix, indent)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	// Output:
	// <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:univ="http://www.example.pl/ws/test/universal">
	//    <soapenv:Header>
	//       <wsse:Security soapenv:mustUnderstand="1" xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
	//          <wsse:UsernameToken wsu:Id="UsernameToken-1">
	//             <wsse:Username>test</wsse:Username>
	//          </wsse:UsernameToken>
	//       </wsse:Security>
	//    </soapenv:Header>
	// </soapenv:Envelope>

}

func ExamplePassword() {
	env := &Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    namespaceUniv,
		Header: &Header{
			WsseSecurity: &WsseSecurity{
				MustUnderstand: "1",
				XmlnsWsse:      "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
				XmlnsWsu:       "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
				UsernameToken: &UsernameToken{
					WsuId: "UsernameToken-1",
					Password: &Password{
						Type: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText",
					},
				},
			},
		},
	}

	env.Header.WsseSecurity.UsernameToken.Password.Value = "pass"

	output, err := xml.MarshalIndent(env, prefix, indent)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	// Output:
	// <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:univ="http://www.example.pl/ws/test/universal">
	//    <soapenv:Header>
	//       <wsse:Security soapenv:mustUnderstand="1" xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
	//          <wsse:UsernameToken wsu:Id="UsernameToken-1">
	//             <wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">pass</wsse:Password>
	//          </wsse:UsernameToken>
	//       </wsse:Security>
	//    </soapenv:Header>
	// </soapenv:Envelope>

}

func ExampleNonce() {
	env := &Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    namespaceUniv,
		Header: &Header{
			WsseSecurity: &WsseSecurity{
				MustUnderstand: "1",
				XmlnsWsse:      "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
				XmlnsWsu:       "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
				UsernameToken: &UsernameToken{
					WsuId:    "UsernameToken-1",
					Username: &Username{},
					Password: &Password{
						Type: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText",
					},
					Nonce: &Nonce{
						EncodingType: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary",
					},
				},
			},
		},
	}

	env.Header.WsseSecurity.UsernameToken.Username.Value = "test"
	env.Header.WsseSecurity.UsernameToken.Password.Value = "pass"
	env.Header.WsseSecurity.UsernameToken.Nonce.Value = "nvKKZ20LNP8wpCa4vAeQhQ=="

	output, err := xml.MarshalIndent(env, prefix, indent)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	// Output:
	// <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:univ="http://www.example.pl/ws/test/universal">
	//    <soapenv:Header>
	//       <wsse:Security soapenv:mustUnderstand="1" xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
	//          <wsse:UsernameToken wsu:Id="UsernameToken-1">
	//             <wsse:Username>test</wsse:Username>
	//             <wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">pass</wsse:Password>
	//             <wsse:Nonce EncodingType="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary">nvKKZ20LNP8wpCa4vAeQhQ==</wsse:Nonce>
	//          </wsse:UsernameToken>
	//       </wsse:Security>
	//    </soapenv:Header>
	// </soapenv:Envelope>

}

func ExampleCreated() {
	env := &Envelope{
		XmlnsSoapenv: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsUniv:    namespaceUniv,
		Header: &Header{
			WsseSecurity: &WsseSecurity{
				MustUnderstand: "1",
				XmlnsWsse:      "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
				XmlnsWsu:       "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
				UsernameToken: &UsernameToken{
					WsuId:    "UsernameToken-1",
					Username: &Username{},
					Password: &Password{
						Type: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText",
					},
					Nonce: &Nonce{
						EncodingType: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary",
					},
					Created: &Created{},
				},
			},
		},
	}

	env.Header.WsseSecurity.UsernameToken.Username.Value = "test"
	env.Header.WsseSecurity.UsernameToken.Password.Value = "pass"
	env.Header.WsseSecurity.UsernameToken.Nonce.Value = "nvKKZ20LNP8wpCa4vAeQhQ=="
	env.Header.WsseSecurity.UsernameToken.Created.Value = "2015-09-10T12:25:55.121Z"

	output, err := xml.MarshalIndent(env, prefix, indent)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	// Output:
	// <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:univ="http://www.example.pl/ws/test/universal">
	//    <soapenv:Header>
	//       <wsse:Security soapenv:mustUnderstand="1" xmlns:wsse="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd" xmlns:wsu="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
	//          <wsse:UsernameToken wsu:Id="UsernameToken-1">
	//             <wsse:Username>test</wsse:Username>
	//             <wsse:Password Type="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordText">pass</wsse:Password>
	//             <wsse:Nonce EncodingType="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary">nvKKZ20LNP8wpCa4vAeQhQ==</wsse:Nonce>
	//             <wsu:Created>2015-09-10T12:25:55.121Z</wsu:Created>
	//          </wsse:UsernameToken>
	//       </wsse:Security>
	//    </soapenv:Header>
	// </soapenv:Envelope>

}
