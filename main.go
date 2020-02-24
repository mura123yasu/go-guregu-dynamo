package main

import (
	"fmt"
	"os"

	"github.com/guregu/dynamo"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// Item is struct for DynamoDB
type Item struct {
	MyHashKey  string
	MyRangeKey int
	Text       string
}

func main() {
	//////////////////////
	// クライアントの設定
	dynamoDbRegion := os.Getenv("AWS_REGION")
	disableSsl := false

	// DynamoDB Localを利用する場合はEndpointのURLを設定する
	dynamoDbEndpoint := os.Getenv("DYNAMO_ENDPOINT")
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	// デフォルトでは東京リージョンを指定
	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	db := dynamo.New(session.New(), &aws.Config{
		Region:     aws.String(dynamoDbRegion),
		Endpoint:   aws.String(dynamoDbEndpoint),
		DisableSSL: aws.Bool(disableSsl),
	})

	table := db.Table("MyFirstTable")

	//////////////////////
	// 単純なCRUD - Create
	item := Item{
		MyHashKey:  "MyHash",
		MyRangeKey: 1,
		Text:       "My First Text",
	}

	err := table.Put(item).Run()
	if err != nil {
		fmt.Printf("Failed to put item[%v]\n", err)
	}

	//////////////////////
	// 単純なCRUD - Read
	var readResult Item
	err = table.Get("MyHashKey", item.MyHashKey).Range("MyRangeKey", dynamo.Equal, item.MyRangeKey).One(&readResult)
	if err != nil {
		fmt.Printf("Failed to get item[%v]\n", err)
	}

	//////////////////////
	// 単純なCRUD - Update
	var updateResult Item
	text := "My Second Text"
	err = table.Update("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).Set("Text", text).Value(&updateResult)
	if err != nil {
		fmt.Printf("Failed to update item[%v]\n", err)
	}

	//////////////////////
	// 単純なCRUD - Delete
	err = table.Delete("MyHashKey", item.MyHashKey).Range("MyRangeKey", item.MyRangeKey).Run()
	if err != nil {
		fmt.Printf("Failed to delete item[%v]\n", err)
	}

	//////////////////////
	// Conditional Update
}
