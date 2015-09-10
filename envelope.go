package soap

import (
	"encoding/xml"
)

type Envelope struct {
	XMLName      xml.Name `xml:"soapenv:Envelope"`
	XmlnsSoapenv string   `xml:"xmlns:soapenv,attr"`
	XmlnsUniv    string   `xml:"xmlns:univ,attr"`

	Header *Header
	Body   *Body
}

type Body struct {
	XMLName xml.Name `xml:"soapenv:Body"`
	Payload interface{}
}

type Header struct {
	XMLName      xml.Name `xml:"soapenv:Header"`
	WsseSecurity *WsseSecurity
}

type WsseSecurity struct {
	MustUnderstand string   `xml:"soapenv:mustUnderstand,attr"`
	XMLName        xml.Name `xml:"wsse:Security"`
	XmlnsWsse      string   `xml:"xmlns:wsse,attr"`
	XmlnsWsu       string   `xml:"xmlns:wsu,attr"`

	UsernameToken *UsernameToken
}

type UsernameToken struct {
	XMLName  xml.Name `xml:"wsse:UsernameToken"`
	WsuId    string   `xml:"wsu:Id,attr,omitempty"`
	Username *Username
	Password *Password
	Nonce    *Nonce
	Created  *Created
}

type Username struct {
	XMLName xml.Name `xml:"wsse:Username"`
	Value   string   `xml:",chardata"`
}
type Password struct {
	XMLName xml.Name `xml:"wsse:Password"`
	Type    string   `xml:"Type,attr"`
	Value   string   `xml:",chardata"`
}
type Nonce struct {
	XMLName      xml.Name `xml:"wsse:Nonce,omitempty"`
	EncodingType string   `xml:"EncodingType,attr,omitempty"`
	Value        string   `xml:",chardata"`
}
type Created struct {
	XMLName xml.Name `xml:"wsu:Created,omitempty"`
	Value   string   `xml:",chardata"`
}
