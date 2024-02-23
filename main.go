package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/nyudlts/go-aspace"
)

func main() {
	var err error
	client, err := aspace.NewClient("go-aspace.yml", "dev", 20)
	if err != nil {
		panic(err)
	}

	//create the do
	do := new(aspace.DigitalObject)
	do.Title = "DO TEST OBJECT"
	do.DigitalObjectID = uuid.New().String()
	do.Publish = true

	//create file version 1
	fv1 := new(aspace.FileVersion)
	fv1.FileURI = "https://hdl.handle.net/2333.1/material-request-placeholder"
	fv1.Publish = true
	fv1.UseStatement = "electronic-records-reading-room"
	fv1.XLinkActuateAttribute = "onLoad"
	fv1.XLinkShowAttribute = "new"

	//create fileversion 2
	fv2 := new(aspace.FileVersion)
	fv2.FileURI = "https://rstar.url.com"
	fv2.Publish = false
	fv2.UseStatement = "electronic-records-master"

	//append new file versions to do fv slice
	do.FileVersions = append(do.FileVersions, *fv1, *fv2)

	//make the request
	responseBody, err := client.CreateDigitalObject(6, *do)
	if err != nil {
		panic(err)
	}

	//print the request
	fmt.Println(responseBody)

}
