package test

import (
	"testing"
	"time"

	searchv1 "github.com/AndrewAlizaga/eog_protos/pkg/proto/search"
	searchClient "github.com/AndrewAlizaga/grpc_client_eog_go/pkg/search"
)



var searchPost = &searchv1.Search{
	Name: "juan lopez",		
	Leads: make([]string, 0),
	Id: "",
}


func TestSearchPost(t *testing.T){

	startTime := time.Now()
	t.Log("STARTING POST TEST")
	response, err := searchClient.PostSearch(searchPost)
	t.Log("POST TEST FINISHED")

	if  err != nil {
		t.Logf("POST SEARCH FAILED WITH ERROR: ")
		t.Logf(err.Error())
		t.Log("DURATION: ", time.Since(startTime))

	} 

	if response.SearchStatus != 200 {
		t.Logf("POST SEARCH FAILED WITH INTERNAL STATUS ERROR: ")
		t.Logf(response.String())
		t.Log("DURATION: ", time.Since(startTime))
		return
	} 

	t.Log("POT SEARCH SUCCESSFULL")
	t.Log(response.String())
	t.Log("DURATION: ", time.Since(startTime))
}