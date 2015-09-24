package soap

import (
	"encoding/xml"
)

type ResponseEnvelope struct {
	XMLName          xml.Name     `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	ResponseBodyBody ResponseBody `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type ResponseBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Fault   Fault    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
}

type Fault struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Fault"`
	Code    string   `xml:"faultcode,omitempty"`
	String  string   `xml:"faultstring,omitempty"`
	Actor   string   `xml:"faultactor,omitempty"`
	Detail  string   `xml:"detail,omitempty"`
}
