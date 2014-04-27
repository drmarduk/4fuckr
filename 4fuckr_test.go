package main

import (
	"testing"
)

func TestGetPageCountFromsite(t *testing.T) {

}

// http://4fuckr.com/image_5195262.htm
// -> http://images.4fuckr.com/029e8d750a87/6/c/1/7/4/6c17421939.gif
func TestGetPicUrlFromPicPage(t *testing.T) {
	var sInput string = "http://4fuckr.com/image_5195262.htm"
	var sExpect string = "http://images.4fuckr.com/029e8d750a87/6/c/1/7/4/6c17421939.gif"
	var sOutput string = GetPicUrlFromPicPage(sInput)
	if sOutput != sExpect {
		t.Error("Expected " + sExpect + " got " + sOutput)
	}
}
