package main

import (
	"log"
	"os"
	"time"

	_ "github.com/AshaJenvasu/fiber-test/docs"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/swagger"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Slice (ตู้หนังสือ) เก็บหนังสือได้หลายเล่ม เป็นร้อยเป็นพันเล่ม
var books []Book 

//Middleware Check log
func checkMiddleware(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["role"] != "admin" {
		return fiber.ErrUnauthorized
	}

	return c.Next()
}

// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env error")
	}
	
	engine := html.New("./views", ".html")
	app:= fiber.New(fiber.Config{
	Views: engine,
} )

books = append(books, Book{ID: 1, Title: "Nanatsu no Taizai", Author: "Nakaba Suzuki"})
books = append(books, Book{ID: 2, Title: "One Piece", Author: "Eichiro Oda"})

// 🟢 Public Routes (ใครก็เข้าถึงได้ ไม่ต้องใช้ Token)
app.Get("/swagger/*", swagger.HandlerDefault)
app.Post("/login", login)
app.Get("/config", getEnv)      // ถ้าอยากให้ดู config ได้โดยไม่ต้องล็อกอิน
app.Get("/test-html", testHTML)

// 🔒 Protected Routes (ต้องใช้ Token ตั้งแต่บรรทัดนี้เป็นต้นไป)
app.Use(jwtware.New(jwtware.Config{
    SigningKey: []byte(os.Getenv("JWT_SECRET")),
}))

app.Use(checkMiddleware)

app.Get("/books", getBooks)
app.Get("/books/:id", getBookById)
app.Post("/books", createBook)
app.Put("/books/:id", updateBook)
app.Delete("/books/:id", deleteBook)
app.Post("/upload", uploadFile)

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
		secret := os.Getenv("SECRET")

		if secret == "" {
			secret = "defaultsecret"
		}
	  return c.Render("index", fiber.Map{
        "Title": "Hello, World!",
        "Name": "Asha",
    })
}

func getEnv(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"SECRET": os.Getenv("SECRET"),
		})
}

// Dummy user for example
type User struct {
  Email    string `json:"email"`
  Password string `json:"password"`
}

var memberUser =  User {
  Email:    "user@example.com",
  Password: "password123",
}


func login(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user) ; err != nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	//user, pass not match = Unauth
	if user.Email != memberUser.Email || user.Password != memberUser.Password{
		return fiber.ErrUnauthorized
	}
	// Create token
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["email"] = user.Email
    claims["role"] = "admin"
    claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

    // Generate encoded token
    t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    if err != nil {
      return c.SendStatus(fiber.StatusInternalServerError)
    }
	return c.JSON(fiber.Map{
			"message" : "Login success",
			"token" : t,
		})
	}
