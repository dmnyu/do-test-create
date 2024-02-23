package main

import (
	"fmt"

	"github.com/nyudlts/go-aspace"
)

var (
	client *aspace.ASClient
	ao     aspace.ArchivalObject
	repoID = 6
	aoID   = 980443
)

func main() {
	var err error
	client, err = aspace.NewClient("go-aspace.yml", "dev", 20)
	if err != nil {
		panic(err)
	}

	ao, err = client.GetArchivalObject(6, aoID)
	if err != nil {
		panic(err)
	}

	createDO()

}

func createDO() {

	fmt.Println(ao)

	do := new(aspace.DigitalObject)
	do.Title = ao.Title
	do.DigitalObjectID = ao.ComponentId
	do.Publish = true

	fv1 := new(aspace.FileVersion)
	fv1.FileURI = "https://hdl.handle.net/2333.1/material-request-placeholder"
	fv1.Publish = true
	fv1.UseStatement = "electronic-records-reading-room"
	fv1.XLinkActuateAttribute = "onLoad"
	fv1.XLinkShowAttribute = "new"

	fv2 := new(aspace.FileVersion)
	fv2.FileURI = "https://rstar.url.com"
	fv2.Publish = false
	fv2.UseStatement = "electronic-records-master"

	do.FileVersions = append(do.FileVersions, *fv1, *fv2)

	responseBody, err := client.CreateDigitalObject(6, *do)
	if err != nil {
		panic(err)
	}

	fmt.Println(responseBody)
}
