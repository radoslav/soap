package soap

import (
	"bytes"
	"crypto/tls"
//	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
)

func Request(url string, soapRequest []byte, soapAction string) ([]byte, error) {
	buffer := new(bytes.Buffer)
	buffer.Write(soapRequest)

	req, err := http.NewRequest("POST", url, buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Add("SOAPAction", soapAction)
	req.Header.Set("User-Agent", "github.com/radoslav/soap/0.1")
	req.Close = true

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{Transport: tr}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rawbody, err := ioutil.ReadAll(res.Body)
	if len(rawbody) == 0 {
		return nil, errors.New("Empty response")
	}

	soapResponse, err := SoapFomMTOM(rawbody)
	if err != nil {
		return nil, err
	}

	// test for fault
//	xmlEnvelope := ResponseEnvelope{}
//
//	err = xml.Unmarshal(soapResponse, xmlEnvelope)
//	if err != nil {
//		return nil, err
//	}
//	
//	fault := xmlEnvelope.ResponseBodyBody.Fault
//	if fault != nil {
//		var sFault string = fault.Code 
//		return nil, errors.New(sFault) 
//	}

	return soapResponse, nil
}

func SoapFomMTOM(soap []byte) ([]byte, error) {
	reg := regexp.MustCompile(`(?ims)<[env:|soap:].+Envelope>`)
	s := reg.FindString(string(soap))
	if s == "" {
		return nil, errors.New("Response without soap envelope")
	}

	return []byte(s), nil
}
