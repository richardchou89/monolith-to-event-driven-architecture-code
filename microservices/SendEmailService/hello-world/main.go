package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	_ "github.com/lib/pq"
)

func handler(ctx context.Context, event events.SQSEvent) error {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbSSLMode := os.Getenv("DB_SSL_MODE")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", dbHost, dbPort, dbUser, dbPass, dbName, dbSSLMode)

	sql.Open("postgres", connStr)

	//send email
	for _, record := range event.Records {
		if *record.MessageAttributes["eventType"].StringValue == "send_email" {
			//send email
		} else if *record.MessageAttributes["eventType"].StringValue == "archive_email" {
			//archive email
		}
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
