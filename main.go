// You can start with "go run main.go" command in your terminal.
// You can see the result in your browser.
// http://localhost:8080/

package main

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/mmcdole/gofeed"
)

func main() {
	url := "https://medium.com/feed/t%C3%BCrk-telekom-bulut-teknolojileri"

	// Start a new fiber app
	app := fiber.New()

	// Get articles items on the root endpoint "/"
	app.Get("/", func(c *fiber.Ctx) error {
		fp := gofeed.NewParser()
		feed, _ := fp.ParseURL(url)
		items, _ := json.Marshal(feed.Items)
		err := c.SendString(string(items))
		
		return err
	})

	// Listen on PORT 300
	app.Listen(":8080")
}