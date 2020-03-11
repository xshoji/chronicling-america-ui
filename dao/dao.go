package dao

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/xshoji/chronicling-america-ui/httputil"
	"github.com/xshoji/chronicling-america-ui/jsonutil"
	"net/url"
)

const API_URL_BASE = "https://chroniclingamerica.loc.gov"
const API_PATH_SEARCH = "/search/titles/results"

func Search(terms string, page string) []map[string]interface{} {
	// set default parameter
	if page == "" {
		page = "1"
	}

	// get request
	queryString := fmt.Sprintf(`?format=json&terms=%v&page=%v`, url.QueryEscape(terms), url.QueryEscape(page))
	urlFull := API_URL_BASE + API_PATH_SEARCH + queryString
	responseBodyObject := httputil.DoGet(urlFull)
	log.Info("responseBodyObject")
	log.Info(responseBodyObject)

	// response -> map
	itemsObject := jsonutil.Get(responseBodyObject, "items")
	log.Info("itemsObject")
	log.Info(itemsObject)

	if itemsObject == nil {
		return []map[string]interface{}{}
	}
	resultMap := make([]map[string]interface{}, 0)
	if items, ok := itemsObject.([]interface{}); ok {
		for _, item := range items {
			resultMap = append(resultMap, jsonutil.ToMap(item, GetKeysSearchResponse()))
		}
	}
	log.Info("resultMap")
	log.Info(resultMap)
	return resultMap
}

func GetKeysSearchResponse() []string {
	return []string{
		"essay.0",
		"place_of_publication",
		"start_year",
		"publisher",
		"county.0",
		"edition",
		"frequency",
		"url",
		"id",
		"subject.0",
		"city.0",
		"language.0",
		"title",
		"holding_type.0",
		"end_year",
		"alt_title.0",
		"note.0",
		"lccn",
		"state.0",
		"place.0",
		"country",
		"type",
		"title_normal",
		"oclc",
	}
}
