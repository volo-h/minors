package controllers

import (
  "github.com/volo-h/minors/database"
  "github.com/volo-h/minors/models"
  // "github.com/volo-h/minors/middlewares"

  "github.com/gofiber/fiber/v2"
  "strconv"
)

func AllUsers(c *fiber.Ctx) error {
  // var users []models.User
  // // database.DB.Find(&users)
  // database.DB.Preload("Role").Find(&users)
  // return c.JSON(users)

  // if err := middlewares.IsAuthorized(c, "users"); err != nil {
  //   return err
  // }

  page, _ := strconv.Atoi(c.Query("page", "1"))

  return c.JSON(models.Paginate(database.DB, &models.User{}, page))
}

func CreateUser(c *fiber.Ctx) error {
  // if err := middlewares.IsAuthorized(c, "users"); err != nil {
  //   return err
  // }

  var user models.User

  // TODO: rewrite without using temp: CustomUser, 
  // because receive error: "json: cannot unmarshal string into Go struct field User.role_id of type uint64"

  type CustomUser struct {
    Id        uint  
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
    Email     string `json:"email"`
    Password  []byte `json:"password"`
    RoleId    string `json:"role_id"`
  }

  custom_user := new(CustomUser)
  c.BodyParser(&custom_user)
  user.FirstName = custom_user.FirstName
  user.LastName = custom_user.LastName
  user.Email = custom_user.Email
  roleid, _ := strconv.ParseUint(custom_user.RoleId, 10, 32)
  user.RoleId = roleid

  // user := new([]models.User)
  // user = &custom_user

  // if err := c.BodyParser(&user); err != nil {
  //   return err
  // }

  user.SetPassword("1234")

  database.DB.Create(&user)

  return c.JSON(user)
}

func GetUser(c *fiber.Ctx) error {
  // if err := middlewares.IsAuthorized(c, "users"); err != nil {
  //   return err
  // }

  id, _ := strconv.Atoi(c.Params("id"))

  user := models.User{
    Id: uint(id),
  }

  database.DB.Preload("Role").Find(&user)
  // database.DB.Find(&user)

  return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
  // if err := middlewares.IsAuthorized(c, "users"); err != nil {
  //   return err
  // }

  id, _ := strconv.Atoi(c.Params("id"))

  user := models.User{
    Id: uint(id),
  }

  if err := c.BodyParser(&user); err != nil {
    return err
  }

  database.DB.Model(&user).Updates(user)

  return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
  // if err := middlewares.IsAuthorized(c, "users"); err != nil {
  //   return err
  // }

  id, _ := strconv.Atoi(c.Params("id"))

  user := models.User{
    Id: uint(id),
  }

  database.DB.Delete(&user)

  return nil
}
