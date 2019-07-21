package pkg

import "encoding/xml"

type TestWebServiceResponse struct {
	XMLName xml.Name    `xml:"TestWebServiceResponse"`
	TestWebServiceResult bool `xml:"TestWebServiceResult"`
}
