package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"strconv"
)

/*

	下載套件
	啟動本地環境
	設定客戶端
	建立表格
	新增資料
		建立物件struct

	查詢資料
		單一查詢
		範圍查詢
	更新資料
		單一更新
		範圍更新
	刪除資料

	參考資料
	https://hub.docker.com/r/amazon/dynamodb-local/
	https://docs.aws.amazon.com/zh_tw/amazondynamodb/latest/developerguide/DynamoDBLocal.UsageNotes.html
	https://github.com/codacy/aws-dynamodb-local/blob/master/Dockerfile
 	https://gist.github.com/Tamal/02776c3e2db7eec73c001225ff52e827

*/

// list table example
func ListDynamoDB(){

	// 創建客戶端
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.

	// 雲端版本
	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: session.SharedConfigEnable,
	//}))

	// 本地端版本
	creds := credentials.NewStaticCredentials("123", "123", "")
	sess,_ := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:   aws.String("eu-west-1"),
		Endpoint: aws.String("http://localhost:8000")})

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// create the input configuration instance
	input := &dynamodb.ListTablesInput{}

	// Get the list of tables
	result, _ := svc.ListTables(input)

	// print table
	for _, n := range result.TableNames {
		fmt.Println(*n)
	}

}

// 此處寫死，未來須改
func CreateTableInDynamoDB(){

	// 創建客戶端
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.

	// 雲端版本
	//sess := session.Must(session.NewSessionWithOptions(session.Options{
	//	SharedConfigState: session.SharedConfigEnable,
	//}))

	// 本地端版本
	creds := credentials.NewStaticCredentials("123", "123", "")
	sess,_ := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:   aws.String("eu-west-1"),
		Endpoint: aws.String("http://localhost:8000")})

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Year"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Title"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Year"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Title"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.CreateTable(input)
	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Created the table", tableName)
}

// Create struct to hold info about new item
type Movie struct {
	Year   int
	Title  string
	Plot   string
	Rating float64
}

var tableName string = "Movies"

/*
	建立安全Session
	使用Session建立客戶端
	準備Struct
	轉換成DynamoDB需要的MarshalMap格式
	跟Dynamodb申請插入資料請求
	使用客戶端插入資料
*/
func (movie *Movie)Insert(){

	// 本地端版本
	creds := credentials.NewStaticCredentials("123", "123", "")
	sess,_ := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:   aws.String("eu-west-1"),
		Endpoint: aws.String("http://localhost:8000")})

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	keyAttr, _ := dynamodbattribute.MarshalMap(movie)

	// Create item in table Movies
	input := &dynamodb.PutItemInput{
		Item:      keyAttr,
		TableName: aws.String(tableName),
	}

	putItemOutput, _ := svc.PutItem(input)
	fmt.Println(putItemOutput.GoString())
}

/*

	建立安全Session
	使用Session建立客戶端
	輸入Table名
	取出資料
	Unmarshal 成 struct

*/
func (movie *Movie)Read(){

	// 本地端版本
	creds := credentials.NewStaticCredentials("123", "123", "")
	sess,_ := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:   aws.String("eu-west-1"),
		Endpoint: aws.String("http://localhost:8000")})

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	result, _ := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String(strconv.Itoa(movie.Year)),
			},
			"Title": {
				S: aws.String(movie.Title),
			},
		},
	})

	fmt.Println("Year:  ", movie.Year)
	fmt.Println("Title: ", movie.Title)
	fmt.Println("Plot:  ", movie.Plot)
	fmt.Println("Rating:", movie.Rating)

	dynamodbattribute.UnmarshalMap(result.Item, &movie)

}

/*

	建立安全Session
	使用Session建立客戶端
	跟Dynamodb申請更新資料請求
	使用客戶端進行Update

*/
func (movie *Movie)Update(){

	// 本地端版本
	creds := credentials.NewStaticCredentials("123", "123", "")
	sess,_ := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:   aws.String("eu-west-1"),
		Endpoint: aws.String("http://localhost:8000")})

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// Create item in table Movies


	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				N: aws.String(strconv.FormatFloat(movie.Rating,'E',-1,32)),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String(strconv.Itoa(movie.Year)),
			},
			"Title": {
				S: aws.String(movie.Title),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Rating = :r"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}


func (movie *Movie)Delete(){

	// 本地端版本
	creds := credentials.NewStaticCredentials("123", "123", "")
	sess,_ := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:   aws.String("eu-west-1"),
		Endpoint: aws.String("http://localhost:8000")})

	svc := dynamodb.New(sess)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Year": {
				N: aws.String(strconv.Itoa(movie.Year)),
			},
			"Title": {
				S: aws.String(movie.Title),
			},
		},
		TableName: aws.String(tableName),
	}

	_, err := svc.DeleteItem(input)
	if err != nil {
		fmt.Println("Got error calling DeleteItem")
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Deleted '" + movie.Title + "' (" + strconv.Itoa(movie.Year) + ") from table " + tableName)

}

/*
func main() {

	//movie := Movie{
	//	Year:   2015,
	//	Title:  "The Big New Movie",
	//	Plot:   "Nothing happens at all.",
	//	Rating: 0.0,
	//}

	// 瀏覽表格
	// 再使用此Function
	//ListDynamoDB()

	//CreateTableInDynamoDB()

	//ListDynamoDB()


	//movie.Insert()

	//Read
	//movie.Read()

	//Update
	//movie.Rating=4.8
	//movie.Update()
	//
	//movie.Read()
	//fmt.Println(movie)

	//Delete
	//movie.Delete()

}
*/
