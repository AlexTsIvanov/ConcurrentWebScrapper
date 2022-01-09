package links

import (
	"github.com/PuerkitoBio/goquery"
)

func uniqueHref(checkHref string, dupHref map[string]struct{}) bool {
	if _, ok := dupHref[checkHref]; ok {
		return false
	} else {
		dupHref[checkHref] = struct{}{}
		return true
	}
}

func uniqueImg(checkImg string, dupImg map[string]struct{}) bool {
	if _, ok := dupImg[checkImg]; ok {
		return false
	} else {
		dupImg[checkImg] = struct{}{}
		return true
	}
}
func GetHrefLinks(urlLink string, dupHref map[string]struct{}) ([]string, error) {
	var urlLinks []string
	doc, err := goquery.NewDocument(urlLink)
	if err != nil {
		return nil, err
	}

	doc.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		if uniqueHref(href, dupHref) {
			urlLinks = append(urlLinks, href)
		}

	})
	return urlLinks, nil
}

func GetImgLinks(urlLink string, dupImg map[string]struct{}) ([]string, error) {
	var urlLinks []string
	doc, err := goquery.NewDocument(urlLink)
	if err != nil {
		return nil, err
	}

	doc.Find("img[src]").Each(func(index int, item *goquery.Selection) {
		src, _ := item.Attr("src")
		if uniqueImg(src, dupImg) {
			urlLinks = append(urlLinks, src)
		}
	})
	return urlLinks, nil
}
