package controllers

import (
	"log"
	"news/database"
	"news/models"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gomodule/redigo/redis"
)

func NewsAdd(c *fiber.Ctx) error {
	// sample json param :
	// {
	// 	"title": "New Title",
	// 	"desc": "Lorem ipsum",
	// 	"status": "draft",
	// 	"tags": "invesment,stock"
	// }

	param := new(models.ParamInputNews)
	if err := c.BodyParser(&param); err != nil {
		return err
	}

	// insert news
	news := new(models.News)
	news.Title = param.Title
	news.Desc = param.Desc
	news.Status = param.Status

	database.DB.Create(&news)

	// insert tags/topics
	tags := strings.Split(param.Tags, ",")
	nt := new(models.NewsTags)
	nt.NewsId = news.Id

	for _, tag := range tags {
		t := new(models.Tags)
		database.DB.Where("name = ?", tag).First(&t)

		nt.TagId = t.Id
		database.DB.Create(&nt)
	}

	return c.SendString("Success insert news")
}

func NewsEdit(c *fiber.Ctx) error {
	n := new(models.News)
	if err := c.BodyParser(&n); err != nil {
		return err
	}

	database.DB.Save(&n)

	return c.SendString("Success modify news")
}

func NewsDelete(c *fiber.Ctx) error {
	// sample json param :
	// {
	// 	"id": 1
	// }

	n := new(models.News)
	if err := c.BodyParser(&n); err != nil {
		return err
	}

	// delete news
	database.DB.Delete(&n)

	// delete news tags
	nt := new(models.NewsTags)
	nt.NewsId = n.Id
	database.DB.Where("news_id = ?", nt.NewsId).Delete(&nt)

	return c.SendString("Success delete news")
}

func NewsList(c *fiber.Ctx) error {
	var err error

	conn, err := redis.Dial("tcp", "0.0.0.0:6379")
	if err != nil {
		log.Panic(err)
	}

	status := c.Params("status")
	topic := c.Params("topic")
	rk := status + topic

	r := []models.ResultSearchNews{}

	cache, err := redis.String(conn.Do("HGET", rk, "value"))

	if cache != "" {
		return c.SendString(cache)
	} else {
		err = database.DB.Table("news").
			Select("news.title, news.desc, news.status, tags.name").
			Joins("join news_tags on news.id = news_tags.news_id").
			Joins("join tags on tags.id = news_tags.tag_id").
			Where("news.status = ? or tags.name = ?", status, topic).
			Order("news.id desc").
			Find(&r).Error

		_, err = conn.Do("HSET", rk, "value", r)
		if err != nil {
			return err
		}
	}

	return c.JSON(r)
}

func TagAdd(c *fiber.Ctx) error {
	t := new(models.Tags)
	if err := c.BodyParser(&t); err != nil {
		return err
	}

	database.DB.Create(&t)

	return c.SendString("Success insert tag")
}

func TagEdit(c *fiber.Ctx) error {
	t := new(models.Tags)
	if err := c.BodyParser(&t); err != nil {
		return err
	}

	database.DB.Save(&t)

	return c.SendString("Success modify tag")
}

func TagDelete(c *fiber.Ctx) error {
	t := new(models.Tags)
	if err := c.BodyParser(&t); err != nil {
		return err
	}

	database.DB.Delete(&t)

	return c.SendString("Success delete tag")
}

func TagList(c *fiber.Ctx) error {
	r := []models.Tags{}
	err := database.DB.Table("tags").Order("name asc").Find(&r).Error
	if err != nil {
		return err
	}

	return c.JSON(r)
}
