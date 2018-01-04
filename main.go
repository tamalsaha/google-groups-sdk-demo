package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2/google"
	gdir "google.golang.org/api/admin/directory/v1"
	"github.com/tamalsaha/go-oneliners"
)

func main() {
	filename := "/home/tamal/Downloads/tigerworks-kube-ada29bda5a1d.json"
	sa, err := ioutil.ReadFile(filename)
	if err != nil {
		oneliners.FILE(err)
		log.Fatalln(err)
	}
	fmt.Println(string(sa))

	cfg, err := google.JWTConfigFromJSON(sa, gdir.AdminDirectoryGroupReadonlyScope)
	if err != nil {
		oneliners.FILE(err)
		log.Fatalln(err)
	}
	// https://admin.google.com/ManageOauthClients

	// ref: https://developers.google.com/admin-sdk/directory/v1/guides/delegation
	// Note: Only users with access to the Admin APIs can access the Admin SDK Directory API, therefore your service account needs to impersonate one of those users to access the Admin SDK Directory API.
	cfg.Subject = "tamal@appscode.com"
	client := cfg.Client(context.Background())

	svc, err := gdir.New(client)
	if err != nil {
		oneliners.FILE(err)
		log.Fatalln(err)
	}

	r2, err := svc.Groups.List().UserKey("xyz@appscode.com").Domain("appscode.com").Do()
	if err != nil {
		oneliners.FILE(err)
		log.Fatalln(err)
	}
	for _, group := range r2.Groups {
		fmt.Println(group.Email)
	}
}
