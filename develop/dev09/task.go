package main

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/gocolly/colly"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	if err := Wget(os.Args[1]); err != nil {
		fmt.Fprintln(os.Stderr, "error wget: ", err)
		os.Exit(1)
	}
}

func CreateFolder(folderName string) error {
	_, err := os.Stat(folderName)
	if os.IsExist(err) {
		return nil
	} else if os.IsNotExist(err) {
		if err := os.MkdirAll(folderName, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func Wget(urlStr string) error {
	rawURL, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return err
	}
	err = CreateFolder(rawURL.Host)
	if err != nil {
		return err
	}
	c := colly.NewCollector(
		colly.AllowedDomains(rawURL.Host),
	)
	c.OnHTML("a[href]", func(el *colly.HTMLElement) {
		ul := el.Request.AbsoluteURL(el.Attr("href"))
		_ = c.Visit(ul)
	})

	c.OnHTML("link[href]", func(el *colly.HTMLElement) {
		ul := el.Request.AbsoluteURL(el.Attr("href"))
		_ = c.Visit(ul)
	})

	c.OnHTML("script[src]", func(el *colly.HTMLElement) {
		ul := el.Request.AbsoluteURL(el.Attr("src"))
		_ = c.Visit(ul)
	})

	c.OnResponse(func(r *colly.Response) {
		pageName := r.Request.URL.Host + r.Request.URL.Path

		// нет расширения - создаем папку и сохраняем содержимое в index.html
		if path.Ext(pageName) == "" {
			err = CreateFolder(pageName)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}
			if pageName[len(pageName)-1] != '/' {
				pageName += "/"
			}
			pageName += "index.html"
		} else { // иначе сохраняем в папке с именем "до слеша перед именем файла"
			last := strings.LastIndexByte(pageName, '/')
			if last > 0 {
				if err := CreateFolder(pageName[:last]); err != nil {
					fmt.Fprintln(os.Stderr, err)
					os.Exit(1)
				}
			}
		}

		if err = r.Save(pageName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}
		fmt.Println("saved:", pageName)
	})

	if err = c.Visit(rawURL.String()); err != nil {
		return err
	}
	return nil
}
