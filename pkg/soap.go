package pkg

import "encoding/xml"

// SOAPEnvelope envelope
type SOAPEnvelope struct {
	XMLName xml.Name    `xml:"soap:Envelope"`
	Header  *SOAPHeader `xml:",omitempty"`
	Body    SOAPBody    `xml:",omitempty"`
	Soap	string 		`xml:"xmlns:soap,attr"`
	Xsi		string 		`xml:"xmlns:xsi,attr"`
	Xsd		string 		`xml:"xmlns:xsd,attr"`
}

// SOAPHeader header
type SOAPHeader struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Content interface{} `xml:",omitempty"`
}

// SOAPBody body
type SOAPBody struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ soap:Body"`
	Content interface{} `xml:",omitempty"`
}

