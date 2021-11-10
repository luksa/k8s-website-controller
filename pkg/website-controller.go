package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"log"
	"github.com/luksa/website-controller/pkg/v1"
	"io/ioutil"
	"strings"
)

func main() {
	log.Println("website-controller started.")
	for {
		resp, err := http.Get("http://localhost:8001/apis/extensions.example.com/v1/websites?watch=true")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		for {
			var event v1.WebsiteWatchEvent
			if err := decoder.Decode(&event); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}

			log.Printf("Received watch event: %s: %s: %s\n", event.Type, event.Object.Metadata.Name, event.Object.Spec.GitRepo)

			if event.Type == "ADDED" {
				createWebsite(event.Object)
			} else if event.Type == "DELETED" {
				deleteWebsite(event.Object)
			}
		}
	}

}

func createWebsite(website v1.Website) {
	createResource(website, "api/v1", "services", "service-template.json")
	createResource(website, "apis/apps/v1", "deployments", "deployment-template.json")
}

func deleteWebsite(website v1.Website) {
	deleteResource(website, "api/v1", "services", getName(website))
	deleteResource(website, "apis/apps/v1", "deployments", getName(website))
}

func createResource(webserver v1.Website, apiGroup string, kind string, filename string) {
	log.Printf("Creating %s with name %s in namespace %s", kind, getName(webserver), webserver.Metadata.Namespace)
	templateBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	template := strings.Replace(string(templateBytes), "[NAME]", getName(webserver), -1)
	template = strings.Replace(template, "[GIT-REPO]", webserver.Spec.GitRepo, -1)

	resp, err := http.Post(fmt.Sprintf("http://localhost:8001/%s/namespaces/%s/%s/", apiGroup, webserver.Metadata.Namespace, kind), "application/json", strings.NewReader(template))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("response Status:", resp.Status)
}

func deleteResource(webserver v1.Website, apiGroup string, kind string, name string) {
	log.Printf("Deleting %s with name %s in namespace %s", kind, name, webserver.Metadata.Namespace)
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:8001/%s/namespaces/%s/%s/%s", apiGroup, webserver.Metadata.Namespace, kind, name), nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("response Status:", resp.Status)

}

func getName(website v1.Website) string {
	return website.Metadata.Name + "-website"
}
