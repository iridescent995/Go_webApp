package main

import (
"encoding/xml"
  "fmt"
  "io/ioutil"
  "net/http")


//slice of location of sitemaps depending on different topics from main site map
type Sitemapindex struct{
  Locations []string `xml:"sitemap>loc"`
}


//extracting data like titles keywords locations from each xml sitemap
type News struct {
  Titles []string `xml:"url>news>title"`
  Keywords []string `xml:"url>news>keywords"`
  Locations []string `xml:"url>loc"`
}

//putting news in a map
type NewsMap struct {
  Keyword string
  Location string
}



func main() {

  var s Sitemapindex
  var n News
  news_map := make(map[string]NewsMap)

  //Obtaining sitemap data from washington post
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  
  xml.Unmarshal(bytes,&s)
  //fmt.Println(s.Locations)

  for _, Location := range(s.Locations){
  	resp, _ := http.Get(Location)
  	bytes, _ := ioutil.ReadAll(resp.Body)
  	xml.Unmarshal(bytes,&n)
  	//fmt.Println(n.Locations)

    //storing data in a map
    for idx, _ := range n.Keywords {
      news_map[n.Titles[idx]] = NewsMap{n.Keywords[idx], n.Locations[idx]}
    }

  }

  //fmt.Println("reached")

  for idx, data := range news_map {
    fmt.Println("\n\n\n",idx)
    fmt.Println("\n",data.Keyword)
    fmt.Println("\n",data.Location)
  }


  //resp.Body.Close()
}