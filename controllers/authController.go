package controllers

import (
  "github.com/volo-h/minors/database"
  "github.com/volo-h/minors/models"
  "github.com/volo-h/minors/util"

  "github.com/gofiber/fiber/v2"
  "strconv"
  "time"
)


func Hello(c *fiber.Ctx) error {
  return c.SendString("hello world")
}


func Register(c *fiber.Ctx) error {
  // array with keys, string and values string
  var data map[string]string

  if err := c.BodyParser(&data); err != nil {
    return err
  }

  if data["password"] != data["password_confirm"] {
    c.Status(400)
    return c.JSON(fiber.Map{
      "message": "passwords do not match",
    })
  }

  user := models.User{
    FirstName: data["first_name"],
    LastName:  data["last_name"],
    Email:     data["email"],
    RoleId:    1,
  }

  user.SetPassword(data["password"])

  database.DB.Create(&user)

  return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
  var data map[string]string

  if err := c.BodyParser(&data); err != nil {
   return err
  }

  var user models.User

  database.DB.Where("email = ?", data["email"]).First(&user)

  if user.Id == 0 {
    c.Status(404)
    return c.JSON(fiber.Map{
      "message": "not found",
    })
  }

  if err := user.ComparePassword(data["password"]); err != nil {
    c.Status(400)
    return c.JSON(fiber.Map{
      "message": "incorrect password",
    })
  }

  // if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
      // c.Status(400)
      //  return c.JSON(fiber.Map{
      //    "message": "incorrect password",
      //  })
  // }
  // return c.JSON(user)

  token, err := util.GenerateJwt(strconv.Itoa(int(user.Id)))

  if err != nil {
    return c.SendStatus(fiber.StatusInternalServerError)
  }

  // return c.JSON(token)

  cookie := fiber.Cookie{
    Name:     "jwt",
    Value:    token,
    Expires:  time.Now().Add(time.Hour * 24),
    HTTPOnly: true,
  }

  c.Cookie(&cookie)

  return c.JSON(fiber.Map{
    "message": "success",
  })
}

func User(c *fiber.Ctx) error {
  cookie := c.Cookies("jwt")

  id, _ := util.ParseJwt(cookie)

  var user models.User

  database.DB.Where("id = ?", id).First(&user)

  return c.JSON(user)
}

func Logout(c *fiber.Ctx) error {
  cookie := fiber.Cookie{
    Name:     "jwt",
    Value:    "",
    Expires:  time.Now().Add(-time.Hour),
    HTTPOnly: true,
  }

  c.Cookie(&cookie)

  return c.JSON(fiber.Map{
    "message": "success",
  })
}

func UpdateInfo(c *fiber.Ctx) error {
  var data map[string]string

  if err := c.BodyParser(&data); err != nil {
    return err
  }

  cookie := c.Cookies("jwt")

  id, _ := util.ParseJwt(cookie)

  userId, _ := strconv.Atoi(id)

  user := models.User{
    Id:        uint(userId),
    FirstName: data["first_name"],
    LastName:  data["last_name"],
    Email:     data["email"],
  }

  database.DB.Model(&user).Updates(user)

  return c.JSON(user)
}

func UpdatePassword(c *fiber.Ctx) error {
  var data map[string]string

  if err := c.BodyParser(&data); err != nil {
    return err
  }

  if data["password"] != data["password_confirm"] {
    c.Status(400)
    return c.JSON(fiber.Map{
      "message": "passwords do not match",
    })
  }

  cookie := c.Cookies("jwt")

  id, _ := util.ParseJwt(cookie)

  userId, _ := strconv.Atoi(id)

  user := models.User{
    Id: uint(userId),
  }

  user.SetPassword(data["password"])

  database.DB.Model(&user).Updates(user)

  return c.JSON(user)
}
