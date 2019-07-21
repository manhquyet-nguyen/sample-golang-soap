package pkg

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// NewSOAPMux return SOAP server mux
func NewSOAPMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/WebService.asmx", licenseServer)
	return mux
}

const ActionTestWebService = `"TestWebService"`

func TestWebService() TestWebServiceResponse {
	return TestWebServiceResponse{
		TestWebServiceResult:true,
	}
}


func licenseServer(w http.ResponseWriter, r *http.Request) {

	soapAction := r.Header.Get("SOAPAction")
	var res interface{}
	switch soapAction {
	case ActionTestWebService:
		fmt.Println("ActionTestWebService")
		res = TestWebService()
	default:
		fmt.Println("Called: " + soapAction)
		res = nil
	}

	v := SOAPEnvelope{
		Xsi: "http://www.w3.org/2001/XMLSchema-instance",
		Xsd: "http://www.w3.org/2001/XMLSchema",
		Soap:"http://schemas.xmlsoap.org/soap/envelope/",
		Body: SOAPBody{
			Content: res,
		},
	}
	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/xml")
	x, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write([]byte(xml.Header + string(x)))
	return
}

// NewSOAPServer create i2c mock server
func NewSOAPServer(port string) *http.Server {
	mux := NewSOAPMux()
	server := &http.Server{
		Handler: mux,
		Addr:    "localhost:" + port,
	}
	return server
}
