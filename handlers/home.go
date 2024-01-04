package handlers

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Dns struct {
	Url string `json:"url"`
}
type APIResponse struct {
	Status  int        `json:"status"`
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}

func ResolveDNS(c *fiber.Ctx) error {
	var dns Dns
	if err := c.BodyParser(&dns); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}
	if dns.Url == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Url is required",
		})
	}

	url, err := url.Parse(dns.Url)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Error Parsing Url",
		})
	}
	fmt.Println(url.Hostname(), "url")
	// ips, err := net.LookupIP("urenportaal-tst.dakota.nl")
	ips, err := net.LookupIP(url.Hostname())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not get IPs: %v\n", err)
	}

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	fmt.Println(ips, "ips")
	return c.Status(http.StatusCreated).
		JSON(APIResponse{Status: http.StatusCreated, Message: "Successfully resolve dns", Data: &fiber.Map{"data": ips[0]}})
}
