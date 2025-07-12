package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSerializeAndParse(t *testing.T) {
	content := "testcontent"
	r := UDPRequest{
		endpoint:    "/helloWorld",
		reqId:       "uniqueid",
		contentSize: uint16(len(content)),
		content:     content,
	}

	serializedReq, err := r.Serialize()
	if err != nil {
		t.Errorf("Got an error while serializing: %d", err)
	}

	reqParsed, err := ParseUDPMessage(serializedReq)
	if err != nil {
		t.Errorf("Got an error while serializing: %d", err)
	}

	if !reflect.DeepEqual(r, reqParsed) {
		t.Errorf("Parsed request is not the same as serialized requiest:\n %s", fmt.Sprintf("\tOriginal: \n\t%v \n\tParsed: \n\t%v", r, reqParsed))
	}

	fmt.Println("Success")
}
