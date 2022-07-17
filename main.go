package main

import (
  "github.com/volo-h/minors/database"
  "github.com/volo-h/minors/routes"

  "github.com/gofiber/fiber/v2"
  "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
  // var n int
  // n = 5
  // or
  // n := 5

  // var name *string = new(string)
  // *name = "waw"
  // fmt.Println(*name, name)

  // name := "popsa"
  // fmt.Println(name, "\t\t" ,&name)

  // name = "new value"
  // fmt.Println(name, "\t", &name)

  database.Connect()

  // db, err = gorm.Open(mysql.Open(dsn: "root:rootroot@/go_admin"), &gorm.Config{})

  app := fiber.New()

  app.Use(cors.New(cors.Config{
   AllowCredentials: true,
  }))

  routes.Setup(app)

  app.Listen(":8000")
}
