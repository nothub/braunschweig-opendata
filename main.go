package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type PackageData struct {
	Success bool `json:"success"`
	Result  []struct {
		Id                string `json:"id"`
		Name              string `json:"name"`
		Title             string `json:"title"`
		Author            string `json:"author"`
		AuthorEmail       string `json:"author_email"`
		Maintainer        string `json:"maintainer"`
		MaintainerEmail   string `json:"maintainer_email"`
		LicenseTitle      string `json:"license_title"`
		LicenseId         string `json:"license_id"`
		Notes             string `json:"notes"`
		Url               string `json:"url"`
		State             string `json:"state"`
		Private           bool   `json:"private"`
		RevisionTimestamp string `json:"revision_timestamp"`
		MetadataCreated   string `json:"metadata_created"`
		MetadataModified  string `json:"metadata_modified"`
		CreatorUserId     string `json:"creator_user_id"`
		Type              string `json:"type"`
		Resources         []struct {
			Id                string `json:"id"`
			RevisionId        string `json:"revision_id"`
			Url               string `json:"url"`
			Description       string `json:"description"`
			Format            string `json:"format"`
			State             string `json:"state"`
			RevisionTimestamp string `json:"revision_timestamp"`
			Name              string `json:"name"`
			Mimetype          string `json:"mimetype"`
			Size              string `json:"size"`
			Created           string `json:"created"`
			ResourceGroupId   string `json:"resource_group_id"`
			LastModified      string `json:"last_modified"`
		} `json:"resources"`
		Tags []struct {
			Id           string `json:"id"`
			VocabularyId string `json:"vocabulary_id"`
			Name         string `json:"name"`
		} `json:"tags"`
		Groups []struct {
			Description     string `json:"description"`
			Id              string `json:"id"`
			ImageDisplayUrl string `json:"image_display_url"`
			Title           string `json:"title"`
			Name            string `json:"name"`
		} `json:"groups"`
	} `json:"result"`
}

func main() {

	res, err := http.Get("https://opendata.braunschweig.de/api/3/action/package_list")
	if err != nil {
		log.Fatalln(err.Error())
	}

	var listData struct {
		Success bool     `json:"success"`
		Result  []string `json:"result"`
	}

	err = json.NewDecoder(res.Body).Decode(&listData)
	if err != nil {
		log.Fatalln(err.Error())
	}

	if !listData.Success {
		log.Fatalln("unable to fetch list of packages")
	}

	for _, name := range listData.Result {

		res, err := http.Get("https://opendata.braunschweig.de/api/3/action/package_show?id=" + name)
		if err != nil {
			log.Fatalln(err.Error())
		}

		var packageData PackageData
		err = json.NewDecoder(res.Body).Decode(&packageData)
		if err != nil {
			log.Fatalln(err.Error())
		}

		if !packageData.Success {
			log.Fatalln("unable to fetch package info for", name)
		}

		for _, pkg := range packageData.Result {
			log.Println(pkg.Id)
			log.Println(pkg.Name)
			log.Println(pkg.Title)
			log.Println(pkg.Author)
			log.Println(pkg.Maintainer)
			log.Println(pkg.LicenseId)
			log.Println(pkg.LicenseTitle)
			log.Println(pkg.RevisionTimestamp)
			log.Println(pkg.Type)
			for _, r := range pkg.Resources {
				log.Println(r.Id)
				log.Println(r.Name)
				log.Println(r.Url)
				log.Println(r.Format)
				log.Println(r.Mimetype)
				log.Println(r.Size)
				log.Println(r.LastModified)
			}
			log.Println("----------")
		}

	}

}
