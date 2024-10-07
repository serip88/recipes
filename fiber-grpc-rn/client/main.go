// ⚡️ Fiber is an Express inspired web framework written in Go with ☕️
// 🤖 Github Repository: https://github.com/gofiber/fiber
// 📌 API Documentation: https://docs.gofiber.io

package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	proto "github.com/serip88/recipes/protogen/service/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := proto.NewCommonServiceClient(conn)

	// g := gin.Default()
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/add/:a/:b", func(c *fiber.Ctx) error {
		a, err := strconv.ParseUint(c.Params("a"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid argument A",
			})
		}
		b, err := strconv.ParseUint(c.Params("b"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid argument B",
			})
		}
		req := &proto.Request{A: int64(a), B: int64(b)}
		//B test
		req.Id = "1234566"
		res := &proto.Response{
			Result: 0,
			User:   &proto.User{},
		}
		if res1, err := client.GetUser(context.Background(), req); err == nil {
			res.User = res1.User
		}
		//E test
		if res2, err := client.Add(context.Background(), req); err == nil {
			res.Result = res2.Result
			return c.JSON(fiber.Map{"status": fiber.StatusOK, "message": "data found", "data": res})
			// return c.Status(fiber.StatusOK).JSON(fiber.Map{
			// 	"result": fmt.Sprint(res.Result),
			// })
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})

	})

	app.Get("/mult/:a/:b", func(c *fiber.Ctx) error {
		a, err := strconv.ParseUint(c.Params("a"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid argument A",
			})
		}
		b, err := strconv.ParseUint(c.Params("b"), 10, 64)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid argument B",
			})
		}
		req := &proto.Request{A: int64(a), B: int64(b)}
		if res, err := client.Multiply(context.Background(), req); err == nil {
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"result": fmt.Sprint(res.Result),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})

	})
	app.Get("/user/:a/:b", func(c *fiber.Ctx) error {
		// a := c.Params("a")
		// b := c.Params("b")

		req := &proto.Request{A: int64(0), B: int64(0)}
		req.Id = "1234566"
		if res, err := client.GetUser(context.Background(), req); err == nil {
			return c.JSON(fiber.Map{"status": fiber.StatusOK, "message": "data found", "data": res})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})

	})
	log.Fatal(app.Listen(":3001"))
}
