package handlers

import (
	"context"
	"fmt"
	"rogers-software/calipso/models"

	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.ResApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.ResApi
	r.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("method")).(string) {

		}
		//
	case "GET":
		switch ctx.Value(models.Key("method")).(string) {

		}
		//
	case "PUT":
		switch ctx.Value(models.Key("method")).(string) {

		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("method")).(string) {

		}
		//
	}
	r.Message = "Method Invalid"
	return r

}
