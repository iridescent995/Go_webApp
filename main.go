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



type News struct{
	Titles [] string `xml:"url>n:news>n:title"`
	Keywords []string `xml:"url>n:news>n:keywords"`
	Locations []string `xml:"url>loc"`
}



func main() {
  var s Sitemapindex
  var n News


  //Obtaining sitemap data from washington post
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  
  xml.Unmarshal(bytes,&s)
  //fmt.Println(s.Locations)

  for _, Location := range(s.Locations){
  	resp, _ := http.Get(Location)
  	bytes, _ := ioutil.ReadAll(resp.Body)
  	xml.Unmarshal(bytes,&n)
  	fmt.Println(n.Locations)
  }

  
  //resp.Body.Close()
}