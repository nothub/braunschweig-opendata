package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type PackageData struct {
	Help    string `json:"help"`
	Success bool   `json:"success"`
	Result  []struct {
		Id                string    `json:"id"`
		Name              string    `json:"name"`
		Title             string    `json:"title"`
		Author            string    `json:"author"`
		AuthorEmail       string    `json:"author_email"`
		Maintainer        string    `json:"maintainer"`
		MaintainerEmail   string    `json:"maintainer_email"`
		LicenseTitle      string    `json:"license_title"`
		LicenseId         string    `json:"license_id"`
		Notes             string    `json:"notes"`
		Url               string    `json:"url"`
		State             string    `json:"state"`
		Private           bool      `json:"private"`
		RevisionTimestamp time.Time `json:"revision_timestamp"`
		MetadataCreated   time.Time `json:"metadata_created"`
		MetadataModified  time.Time `json:"metadata_modified"`
		CreatorUserId     string    `json:"creator_user_id"`
		Type              string    `json:"type"`
		Resources         []struct {
			Id                string    `json:"id"`
			RevisionId        string    `json:"revision_id"`
			Url               string    `json:"url"`
			Description       string    `json:"description"`
			Format            string    `json:"format"`
			State             string    `json:"state"`
			RevisionTimestamp time.Time `json:"revision_timestamp"`
			Name              string    `json:"name"`
			Mimetype          string    `json:"mimetype"`
			Size              string    `json:"size"`
			Created           time.Time `json:"created"`
			ResourceGroupId   string    `json:"resource_group_id"`
			LastModified      string    `json:"last_modified"`
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

		log.Println(packageData)

	}

}
