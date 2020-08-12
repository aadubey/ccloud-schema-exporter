package client

//
// context.go
// Author: Abraham Leal
//

import (
	"flag"
	"fmt"
	"os"
)

func GetFlags() {

	flag.StringVar(&SrcSRUrl, "src-sr-url", "", "Url to the Source Schema Registry Cluster")
	flag.StringVar(&SrcSRKey, "src-sr-key", "", "API KEY for the Source Schema Registry Cluster")
	flag.StringVar(&SrcSRSecret, "src-sr-secret", "", "API SECRET for the Source Schema Registry Cluster")
	flag.StringVar(&DestSRUrl, "dest-sr-url", "", "Url to the Source Schema Registry Cluster")
	flag.StringVar(&DestSRKey, "dest-sr-key", "", "API KEY for the Destination Schema Registry Cluster")
	flag.StringVar(&DestSRSecret, "dest-sr-secret", "", "API SECRET for the Destination Schema Registry Cluster")
	flag.IntVar(&httpCallTimeout, "timeout", 60, "Timeout, in second, to use for all REST call with the Schema Registries")
	versionFlag := flag.Bool("version", false, "Print the current version and exit")
	usageFlag := flag.Bool("usage", false, "Print the usage of this tool")
	deleteFlag := flag.Bool("deleteAllFromDestination", false, "Setting this will run a delete on all schemas written to the destination registry")

	flag.Parse()

	if *versionFlag {
		printVersion()
		os.Exit(0)
	}

	if *usageFlag {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *deleteFlag{
		fmt.Println("Deleting all schemas from DESTINATION registry")
		deleteAll(DestSRUrl,DestSRKey,DestSRSecret)
		os.Exit(0)
	}

}

func printVersion() {
	fmt.Printf("ccloud-schema-exporter: %s\n", Version)
}

func deleteAll(sr string, key string, secret string){
	destClient := NewSchemaRegistryClient(sr,key,secret, "dst")
	destClient.GetSubjectsWithVersions()
	destClient.DeleteAllSubjectsPermanently()
}



