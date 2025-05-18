package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/playwright-community/playwright-go"
)

func main() {
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("could not start playwright: %v", err)
	}
	defer pw.Stop()
	// 启动浏览器
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false), // 可以改为 false 进行可视化调试
	})
	if err != nil {
		log.Fatalf("could not launch browser: %v", err)
	}
	defer browser.Close()

	// 打开新页面
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("could not create page: %v", err)
	}
	if _, err = page.Goto("https://scrapingclub.com/exercise/list_infinite_scroll/"); err != nil {
		log.Fatalf("could not goto: %v", err)
	}
	//_, err := page.Content()
	//if err != nil {
	//	log.Fatalf("could not get entries: %v", err)
	//	return
	//}
	//fmt.Println(html)
	productHTMLElements, err := page.Locator(".post").All()
	if err != nil {
		log.Fatalf("Could not get the product node: %v", err)
	}
	var products []Product
	for _, productHTMLElement := range productHTMLElements {
		// select the name and price nodes
		// and extract the data of interest from them
		name, err := productHTMLElement.Locator("h4").First().TextContent()
		price, err := productHTMLElement.Locator("h5").First().TextContent()
		if err != nil {
			log.Fatal("Could not apply the scraping logic:", err)
		}

		// add the scraped data to the list
		product := Product{}
		product.name = strings.TrimSpace(name)
		product.price = strings.TrimSpace(price)
		products = append(products, product)
	}
	for i, obj := range products {
		fmt.Println("no.", i, "= ", obj)
	}
}

type Product struct {
	name, price string
}
