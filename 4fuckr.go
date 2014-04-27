package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Test: http://4fuckr.com/image_5195262.htm")
	pic := GetPicUrlFromPicPage("http://4fuckr.com/image_5195262.htm")
	fmt.Println("Got: " + pic)
}

func GetPicUrlFromPicPage(PicPageUrl string) string {
	src, _ := HttpGet(PicPageUrl)

	// now search for pic url in source:
	// <img src="http://images.4fuckr.com/d6c5dea0c460/6/c/1/7/4/6c17421939.gif" alt="e62aaad2a4.gif" title="e62aaad2a4.gif" border="0" class="scaled" id="mypic"/>
	re := regexp.MustCompile("http://images.4fuckr.com/[a-f0-9]{12}/[a-f0-9]{1}/[a-f0-9]{1}/[a-f0-9]{1}/[a-f0-9]{1}/[a-f0-9]{1}/[a-f0-9]{10}.(gif|jpg|jpeg|png)")
	match := re.FindString(src)

	return match
}

func GetPicPagesFromPage(PageUrl string) []string {
	return make([]string, 42)
}

// Site Cound is max 299 due to the fact
// that _non_ premium members can not
// access pages > 299
// Members however can "reach" them, but
// only to $number (need to be done) :default
// BUT!!! if you traverse over each pic_id link
// then you can get to all picutres!!!
// Sooooooo, meeeehhhh.....
func GetPageCountFromSite() int {
	return 0
}
