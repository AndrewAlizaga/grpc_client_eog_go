package utils

import (
	"flag"
	"log"

	searchv1 "github.com/AndrewAlizaga/eog_protos/pkg/proto/search"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	tls                = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containing the CA root cert file")
	serverAddr         = flag.String("addr", "https://eog-engine.herokuapp.com:50051", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "x.test.example.com", "The server name used to verify the hostname returned by the TLS handshake")
	searchConfigurationConnnection searchv1.SearchServiceClient
)

//var Conn *grpc.ClientConn


func init(){}


func GetSearchConnection() (searchv1.SearchServiceClient, error){

	//var err error

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if searchConfigurationConnnection == nil {
		conn, err := grpc.Dial(*serverAddr, opts...)

		if err != nil {
			log.Fatalf("fail to dial: %v", err)
			return nil, err
		}

		//defer conn.Close()
		searchConfigurationConnnection = searchv1.NewSearchServiceClient(conn)
	}

	return searchConfigurationConnnection, nil
	


}