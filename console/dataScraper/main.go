package main

import (
	"fmt"
	"log"
	"strings"
	"net/url"
	"net/http"

	"github.com/tidwall/gjson"
	"github.com/PuerkitoBio/goquery"
)

type DataSet struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Length string `json:"length"`
}

type DataAccessStrSet struct {
	BaseAccessStr   string
	IdAccessStr     string
	TitleAccessStr  string
	LengthAccessStr string
}

func getVideoId(uri string)(string, error) {
	videoId := ""
	param, err := url.Parse(uri)
	if err != nil {
		return videoId, err
	}
	if len(param.Query()["v"][0]) > 0 {
		videoId = param.Query()["v"][0]
	}
	return videoId, nil
}

func Search(param string)([]DataSet, error){
	var err error
	var res []DataSet
	if strings.HasPrefix(param, "http://") {
		return res, nil
	}
	videoId := ""
	if strings.HasPrefix(param, "https://www.youtube.com/watch?v=") {
		videoId, err = getVideoId(param)
		if err != nil {
			return res, err
		}
		param = "https://www.youtube.com/results?search_query='" + videoId + "'"
	} else if !strings.HasPrefix(param, "https://www.youtube.com/") {
		param = "https://www.youtube.com/results?search_query='" + param
	}
	req, err := http.NewRequest("GET", param, nil)
	if err != nil {
		return res, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return res, err
	}
	var initialData string
	doc.Find("body script").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(strings.Replace(s.Text(), "\n", "", -1))
		if len(text) > 0 && strings.HasPrefix(text, "// scraper_data_begin") {
			text = text[21:]
		}
		if len(text) > 0 && strings.HasPrefix(text, "window[\"ytInitialData\"]") {
			initialData = text[26:]
		} else if len(text) > 0 && strings.HasPrefix(text, "var ytInitialData") {
			initialData = text[20:]
		}
	})
	if len(initialData) > 0 {
		jsonStr := strings.Split(initialData, "};")[0] + "}"
		dataAccessStrSet := GetAccessStrSet(param)
		result := gjson.Get(jsonStr, dataAccessStrSet.BaseAccessStr)
		result.ForEach(func(key, value gjson.Result) bool {
			var data DataSet
			v := value.String()
			data.Id = gjson.Get(v, dataAccessStrSet.IdAccessStr).String()
			if len(data.Id) > 0 && (len(videoId) < 1 || strings.Compare(videoId, data.Id) == 0) {
				data.Title = gjson.Get(v, dataAccessStrSet.TitleAccessStr).String()
				data.Length = gjson.Get(v, dataAccessStrSet.LengthAccessStr).String()
				res = append(res, data)
			}
			return true
		})
	}

	return res, nil
}

func GetChannelAccessStrSet() DataAccessStrSet {
	return DataAccessStrSet {
		BaseAccessStr: "contents.twoColumnBrowseResultsRenderer.tabs.0.tabRenderer.content.sectionListRenderer.contents.1.itemSectionRenderer.contents.0.shelfRenderer.content.horizontalListRenderer.items",
		IdAccessStr: "gridVideoRenderer.videoId",
		TitleAccessStr: "gridVideoRenderer.title.simpleText",
		LengthAccessStr: "gridVideoRenderer.thumbnailOverlays.0.thumbnailOverlayTimeStatusRenderer.text.simpleText",
	}
}

func GetTrendAccessStrSet() DataAccessStrSet {
	return DataAccessStrSet {
		BaseAccessStr: "contents.twoColumnBrowseResultsRenderer.tabs.0.tabRenderer.content.sectionListRenderer.contents.0.itemSectionRenderer.contents.0.shelfRenderer.content.expandedShelfContentsRenderer.items",
		IdAccessStr: "videoRenderer.videoId",
		TitleAccessStr: "videoRenderer.title.runs.0.text",
		LengthAccessStr: "videoRenderer.lengthText.simpleText",
	}
}

func GetPlaylistAccessStrSet() DataAccessStrSet {
	return DataAccessStrSet {
		BaseAccessStr: "contents.twoColumnBrowseResultsRenderer.tabs.0.tabRenderer.content.sectionListRenderer.contents.0.itemSectionRenderer.contents.0.playlistVideoListRenderer.contents",
		IdAccessStr: "playlistVideoRenderer.videoId",
		TitleAccessStr: "playlistVideoRenderer.title.runs.0.text",
		LengthAccessStr: "playlistVideoRenderer.lengthText.simpleText",
	}
}

func GetSearchAccessStrSet() DataAccessStrSet {
	return DataAccessStrSet {
		BaseAccessStr: "contents.twoColumnSearchResultsRenderer.primaryContents.sectionListRenderer.contents.0.itemSectionRenderer.contents",
		IdAccessStr: "videoRenderer.videoId",
		TitleAccessStr: "videoRenderer.title.runs.0.text",
		LengthAccessStr: "videoRenderer.lengthText.simpleText",
	}
}

func GetTopAccessStrSet() DataAccessStrSet {
	return DataAccessStrSet {
		BaseAccessStr: "contents.twoColumnBrowseResultsRenderer.tabs.0.tabRenderer.content.richGridRenderer.contents",
		IdAccessStr: "richItemRenderer.content.videoRenderer.videoId",
		TitleAccessStr: "richItemRenderer.content.videoRenderer.title.runs.0.text",
		LengthAccessStr: "richItemRenderer.content.videoRenderer.lengthText.simpleText",
	}
}

func GetAccessStrSet(param string) DataAccessStrSet {
	if strings.HasPrefix(param, "https://www.youtube.com/c/") ||
		strings.HasPrefix(param, "https://www.youtube.com/user/") ||
		strings.HasPrefix(param, "https://www.youtube.com/channel/") {
		return GetChannelAccessStrSet()
	} else if strings.HasPrefix(param, "https://www.youtube.com/feed/trending") {
		return GetTrendAccessStrSet()
	} else if strings.HasPrefix(param, "https://www.youtube.com/results?search_query") {
		return GetSearchAccessStrSet()
	} else if strings.HasPrefix(param, "https://www.youtube.com/playlist") {
		return GetPlaylistAccessStrSet()
	} else {
		return GetTopAccessStrSet()
	}
}

func main() {
	var res []DataSet
	var err error

	list := []string{
		"https://www.youtube.com/feed/trending",
		"https://www.youtube.com/",
	}

	for _, v := range list {
		res, err = Search(v)
		if err != nil {
			log.Fatal(err)
		} else {
			for _, x := range res {
				fmt.Printf("%s, %s, %s\n", x.Id, x.Title, x.Length)
			}
		}
		fmt.Println("==========")
	}
}
