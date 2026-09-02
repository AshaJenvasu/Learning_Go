package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Slice (ตู้หนังสือ) เก็บหนังสือได้หลายเล่ม เป็นร้อยเป็นพันเล่ม
var books []Book 

func main() {
	engine := html.New("./views", ".html")
app:= fiber.New(fiber.Config{
	Views: engine,
} )

books = append(books, Book{ID: 1, Title: "Nanatsu no Taizai", Author: "Nakaba Suzuki"})
books = append(books, Book{ID: 2, Title: "One Piece", Author: "Eichiro Oda"})

app.Get("/books", getBooks)
app.Get("/books/:id", getBookById)
app.Post("/books", createBook)
app.Put("/books/:id", updateBook)
app.Delete("/books/:id", deleteBook)

app.Post("/upload", uploadFile)
app.Get("test-html", testHTML)

app.Listen(":8080")
}

func uploadFile( c *fiber.Ctx) error {
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	err = c.SaveFile(
		file, "./uploads/" + file.Filename)
		
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("File upload complete!")
}

func testHTML( c *fiber.Ctx) error {
	  return c.Render("index", fiber.Map{
        "Title": "Hello, World!",
        "Name": "Asha",
    })
}