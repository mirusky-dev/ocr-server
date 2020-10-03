package handlers

import (
	"bytes"
	"io"

	"github.com/gofiber/fiber/v2"
	"github.com/otiai10/gosseract/v2"
)

// OCR make a OCR from multipart-form file
func OCR(c *fiber.Ctx) error {
	f, err := c.FormFile("image")
	if err != nil {
		return c.SendString(err.Error())
	}
	ff, err := f.Open()
	if err != nil {
		return c.SendString(err.Error())
	}
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, ff); err != nil {
		return c.SendString(err.Error())
	}
	client := c.Context().UserValue("tess-client").(*gosseract.Client)
	client.SetImageFromBytes(buf.Bytes())
	bb, err := client.GetBoundingBoxes(gosseract.RIL_TEXTLINE)
	if err != nil {
		return c.SendString(err.Error())
	}
	// TODO: Make a binding for Microsoft like OCR result
	// o, err := client.HOCRText()
	// if err != nil {
	// 	return c.SendString(err.Error())
	// }
	return c.JSON(bb)
}
