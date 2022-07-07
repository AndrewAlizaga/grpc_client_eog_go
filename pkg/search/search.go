package search

import (
	"context"
	"log"
	"time"

	searchv1 "github.com/AndrewAlizaga/eog_protos/pkg/proto/search"
	"github.com/AndrewAlizaga/grpc_client_eog_go/utils"
)

func PostSearch(search *searchv1.Search)  (*searchv1.SearchResponse, error){

	//Req body 
	now := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()


	//Conx
	searchServiceConnnection, err := utils.GetSearchConnection()

	if err != nil {
		log.Fatal("Fail obtaining searchServiceConnection")
		return nil, err
	}

	searchResult, err := searchServiceConnnection.PostSearch(ctx, search)

	if err != nil {
		log.Fatal("Fail obtaining searchServiceConnection")
		return searchResult, err
	}

	log.Println("Time for CreateAgent", time.Since(now))


	return searchResult, nil



}
