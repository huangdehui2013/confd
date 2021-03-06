package dynamodb_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/dynamodb"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDynamoDB_BatchGetItem() {
	svc := dynamodb.New(nil)

	params := &dynamodb.BatchGetItemInput{
		RequestItems: &map[string]*dynamodb.KeysAndAttributes{ // Required
			"Key": &dynamodb.KeysAndAttributes{ // Required
				Keys: []*map[string]*dynamodb.AttributeValue{ // Required
					&map[string]*dynamodb.AttributeValue{ // Required
						"Key": &dynamodb.AttributeValue{ // Required
							B:    []byte("PAYLOAD"),
							BOOL: aws.Boolean(true),
							BS: [][]byte{
								[]byte("PAYLOAD"), // Required
								// More values...
							},
							L: []*dynamodb.AttributeValue{
								&dynamodb.AttributeValue{ // Required
								// Recursive values...
								},
								// More values...
							},
							M: &map[string]*dynamodb.AttributeValue{
								"Key": &dynamodb.AttributeValue{ // Required
								// Recursive values...
								},
								// More values...
							},
							N: aws.String("NumberAttributeValue"),
							NS: []*string{
								aws.String("NumberAttributeValue"), // Required
								// More values...
							},
							NULL: aws.Boolean(true),
							S:    aws.String("StringAttributeValue"),
							SS: []*string{
								aws.String("StringAttributeValue"), // Required
								// More values...
							},
						},
						// More values...
					},
					// More values...
				},
				AttributesToGet: []*string{
					aws.String("AttributeName"), // Required
					// More values...
				},
				ConsistentRead: aws.Boolean(true),
				ExpressionAttributeNames: &map[string]*string{
					"Key": aws.String("AttributeName"), // Required
					// More values...
				},
				ProjectionExpression: aws.String("ProjectionExpression"),
			},
			// More values...
		},
		ReturnConsumedCapacity: aws.String("ReturnConsumedCapacity"),
	}
	resp, err := svc.BatchGetItem(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleDynamoDB_BatchWriteItem() {
	svc := dynamodb.New(nil)

	params := &dynamodb.BatchWriteItemInput{
		RequestItems: &map[string][]*dynamodb.WriteRequest{ // Required
			"Key": []*dynamodb.WriteRequest{ // Required
				&dynamodb.WriteRequest{ // Required
					DeleteRequest: &dynamodb.DeleteRequest{
						Key: &map[string]*dynamodb.AttributeValue{ // Required
							"Key": &dynamodb.AttributeValue{ // Required
								B:    []byte("PAYLOAD"),
								BOOL: aws.Boolean(true),
								BS: [][]byte{
									[]byte("PAYLOAD"), // Required
									// More values...
								},
								L: []*dynamodb.AttributeValue{
									&dynamodb.AttributeValue{ // Required
									// Recursive values...
									},
									// More values...
								},
								M: &map[string]*dynamodb.AttributeValue{
									"Key": &dynamodb.AttributeValue{ // Required
									// Recursive values...
									},
									// More values...
								},
								N: aws.String("NumberAttributeValue"),
								NS: []*string{
									aws.String("NumberAttributeValue"), // Required
									// More values...
								},
								NULL: aws.Boolean(true),
								S:    aws.String("StringAttributeValue"),
								SS: []*string{
									aws.String("StringAttributeValue"), // Required
									// More values...
								},
							},
							// More values...
						},
					},
					PutRequest: &dynamodb.PutRequest{
						Item: &map[string]*dynamodb.AttributeValue{ // Required
							"Key": &dynamodb.AttributeValue{ // Required
								B:    []byte("PAYLOAD"),
								BOOL: aws.Boolean(true),
								BS: [][]byte{
									[]byte("PAYLOAD"), // Required
									// More values...
								},
								L: []*dynamodb.AttributeValue{
									&dynamodb.AttributeValue{ // Required
									// Recursive values...
									},
									// More values...
								},
								M: &map[string]*dynamodb.AttributeValue{
									"Key": &dynamodb.AttributeValue{ // Required
									// Recursive values...
									},
									// More values...
								},
								N: aws.String("NumberAttributeValue"),
								NS: []*string{
									aws.String("NumberAttributeValue"), // Required
									// More values...
								},
								NULL: aws.Boolean(true),
								S:    aws.String("StringAttributeValue"),
								SS: []*string{
									aws.String("StringAttributeValue"), // Required
									// More values...
								},
							},
							// More values...
						},
					},
				},
				// More values...
			},
			// More values...
		},
		ReturnConsumedCapacity:      aws.String("ReturnConsumedCapacity"),
		ReturnItemCollectionMetrics: aws.String("ReturnItemCollectionMetrics"),
	}
	resp, err := svc.BatchWriteItem(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleDynamoDB_CreateTable() {
	svc := dynamodb.New(nil)

	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{ // Required
			&dynamodb.AttributeDefinition{ // Required
				AttributeName: aws.String("KeySchemaAttributeName"), // Required
				AttributeType: aws.String("ScalarAttributeType"),    // Required
			},
			// More values...
		},
		KeySchema: []*dynamodb.KeySchemaElement{ // Required
			&dynamodb.KeySchemaElement{ // Required
				AttributeName: aws.String("KeySchemaAttributeName"), // Required
				KeyType:       aws.String("KeyType"),                // Required
			},
			// More values...
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
			ReadCapacityUnits:  aws.Long(1), // Required
			WriteCapacityUnits: aws.Long(1), // Required
		},
		TableName: aws.String("TableName"), // Required
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			&dynamodb.GlobalSecondaryIndex{ // Required
				IndexName: aws.String("IndexName"), // Required
				KeySchema: []*dynamodb.KeySchemaElement{ // Required
					&dynamodb.KeySchemaElement{ // Required
						AttributeName: aws.String("KeySchemaAttributeName"), // Required
						KeyType:       aws.String("KeyType"),                // Required
					},
					// More values...
				},
				Projection: &dynamodb.Projection{ // Required
					NonKeyAttributes: []*string{
						aws.String("NonKeyAttributeName"), // Required
						// More values...
					},
					ProjectionType: aws.String("ProjectionType"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
					ReadCapacityUnits:  aws.Long(1), // Required
					WriteCapacityUnits: aws.Long(1), // Required
				},
			},
			// More values...
		},
		LocalSecondaryIndexes: []*dynamodb.LocalSecondaryIndex{
			&dynamodb.LocalSecondaryIndex{ // Required
				IndexName: aws.String("IndexName"), // Required
				KeySchema: []*dynamodb.KeySchemaElement{ // Required
					&dynamodb.KeySchemaElement{ // Required
						AttributeName: aws.String("KeySchemaAttributeName"), // Required
						KeyType:       aws.String("KeyType"),                // Required
					},
					// More values...
				},
				Projection: &dynamodb.Projection{ // Required
					NonKeyAttributes: []*string{
						aws.String("NonKeyAttributeName"), // Required
						// More values...
					},
					ProjectionType: aws.String("ProjectionType"),
				},
			},
			// More values...
		},
	}
	resp, err := svc.CreateTable(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleDynamoDB_DeleteItem() {
	svc := dynamodb.New(nil)

	params := &dynamodb.DeleteItemInput{
		Key: &map[string]*dynamodb.AttributeValue{ // Required
			"Key": &dynamodb.AttributeValue{ // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Boolean(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: &map[string]*dynamodb.AttributeValue{
					"Key": &dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Boolean(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		TableName:           aws.String("TableName"), // Required
		ConditionExpression: aws.String("ConditionExpression"),
		ConditionalOperator: aws.String("ConditionalOperator"),
		Expected: &map[string]*dynamodb.ExpectedAttributeValue{
			"Key": &dynamodb.ExpectedAttributeValue{ // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
						B:    []byte("PAYLOAD"),
						BOOL: aws.Boolean(true),
						BS: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						L: []*dynamodb.AttributeValue{
							&dynamodb.AttributeValue{ // Required
							// Recursive values...
							},
							// More values...
						},
						M: &map[string]*dynamodb.AttributeValue{
							"Key": &dynamodb.AttributeValue{ // Required
							// Recursive values...
							},
							// More values...
						},
						N: aws.String("NumberAttributeValue"),
						NS: []*string{
							aws.String("NumberAttributeValue"), // Required
							// More values...
						},
						NULL: aws.Boolean(true),
						S:    aws.String("StringAttributeValue"),
						SS: []*string{
							aws.String("StringAttributeValue"), // Required
							// More values...
						},
					},
					// More values...
				},
				ComparisonOperator: aws.String("ComparisonOperator"),
				Exists:             aws.Boolean(true),
				Value: &dynamodb.AttributeValue{
					B:    []byte("PAYLOAD"),
					BOOL: aws.Boolean(true),
					BS: [][]byte{
						[]byte("PAYLOAD"), // Required
						// More values...
					},
					L: []*dynamodb.AttributeValue{
						&dynamodb.AttributeValue{ // Required
						// Recursive values...
						},
						// More values...
					},
					M: &map[string]*dynamodb.AttributeValue{
						"Key": &dynamodb.AttributeValue{ // Required
						// Recursive values...
						},
						// More values...
					},
					N: aws.String("NumberAttributeValue"),
					NS: []*string{
						aws.String("NumberAttributeValue"), // Required
						// More values...
					},
					NULL: aws.Boolean(true),
					S:    aws.String("StringAttributeValue"),
					SS: []*string{
						aws.String("StringAttributeValue"), // Required
						// More values...
					},
				},
			},
			// More values...
		},
		ExpressionAttributeNames: &map[string]*string{
			"Key": aws.String("AttributeName"), // Required
			// More values...
		},
		ExpressionAttributeValues: &map[string]*dynamodb.AttributeValue{
			"Key": &dynamodb.AttributeValue{ // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Boolean(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: &map[string]*dynamodb.AttributeValue{
					"Key": &dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Boolean(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		ReturnConsumedCapacity:      aws.String("ReturnConsumedCapacity"),
		ReturnItemCollectionMetrics: aws.String("ReturnItemCollectionMetrics"),
		ReturnValues:                aws.String("ReturnValue"),
	}
	resp, err := svc.DeleteItem(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleDynamoDB_DeleteTable() {
	svc := dynamodb.New(nil)

	params := &dynamodb.DeleteTableInput{
		TableName: aws.String("TableName"), // Required
	}
	resp, err := svc.DeleteTable(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleDynamoDB_DescribeTable() {
	svc := dynamodb.New(nil)

	params := &dynamodb.DescribeTableInput{
		TableName: aws.String("TableName"), // Required
	}
	resp, err := svc.DescribeTable(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleDynamoDB_GetItem() {
	svc := dynamodb.New(nil)

	params := &dynamodb.GetItemInput{
		Key: &map[string]*dynamodb.AttributeValue{ // Required
			"Key": &dynamodb.AttributeValue{ // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Boolean(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: &map[string]*dynamodb.AttributeValue{
					"Key": &dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Boolean(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		TableName: aws.String("TableName"), // Required
		AttributesToGet: []*string{
			aws.String("AttributeName"), // Required
			// More values...
		},
		ConsistentRead: aws.Boolean(true),
		ExpressionAttributeNames: &map[string]*string{
			"Key": aws.String("AttributeName"), // Required
			// More values...
		},
		ProjectionExpression:   aws.String("ProjectionExpression"),
		ReturnConsumedCapacity: aws.String("ReturnConsumedCapacity"),
	}
	resp, err := svc.GetItem(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleDynamoDB_ListTables() {
	svc := dynamodb.New(nil)

	params := &dynamodb.ListTablesInput{
		ExclusiveStartTableName: aws.String("TableName"),
		Limit: aws.Long(1),
	}
	resp, err := svc.ListTables(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleDynamoDB_PutItem() {
	svc := dynamodb.New(nil)

	params := &dynamodb.PutItemInput{
		Item: &map[string]*dynamodb.AttributeValue{ // Required
			"Key": &dynamodb.AttributeValue{ // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Boolean(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: &map[string]*dynamodb.AttributeValue{
					"Key": &dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Boolean(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		TableName:           aws.String("TableName"), // Required
		ConditionExpression: aws.String("ConditionExpression"),
		ConditionalOperator: aws.String("ConditionalOperator"),
		Expected: &map[string]*dynamodb.ExpectedAttributeValue{
			"Key": &dynamodb.ExpectedAttributeValue{ // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
						B:    []byte("PAYLOAD"),
						BOOL: aws.Boolean(true),
						BS: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						L: []*dynamodb.AttributeValue{
							&dynamodb.AttributeValue{ // Required
							// Recursive values...
							},
							// More values...
						},
						M: &map[string]*dynamodb.AttributeValue{
							"Key": &dynamodb.AttributeValue{ // Required
							// Recursive values...
							},
							// More values...
						},
						N: aws.String("NumberAttributeValue"),
						NS: []*string{
							aws.String("NumberAttributeValue"), // Required
							// More values...
						},
						NULL: aws.Boolean(true),
						S:    aws.String("StringAttributeValue"),
						SS: []*string{
							aws.String("StringAttributeValue"), // Required
							// More values...
						},
					},
					// More values...
				},
				ComparisonOperator: aws.String("ComparisonOperator"),
				Exists:             aws.Boolean(true),
				Value: &dynamodb.AttributeValue{
					B:    []byte("PAYLOAD"),
					BOOL: aws.Boolean(true),
					BS: [][]byte{
						[]byte("PAYLOAD"), // Required
						// More values...
					},
					L: []*dynamodb.AttributeValue{
						&dynamodb.AttributeValue{ // Required
						// Recursive values...
						},
						// More values...
					},
					M: &map[string]*dynamodb.AttributeValue{
						"Key": &dynamodb.AttributeValue{ // Required
						// Recursive values...
						},
						// More values...
					},
					N: aws.String("NumberAttributeValue"),
					NS: []*string{
						aws.String("NumberAttributeValue"), // Required
						// More values...
					},
					NULL: aws.Boolean(true),
					S:    aws.String("StringAttributeValue"),
					SS: []*string{
						aws.String("StringAttributeValue"), // Required
						// More values...
					},
				},
			},
			// More values...
		},
		ExpressionAttributeNames: &map[string]*string{
			"Key": aws.String("AttributeName"), // Required
			// More values...
		},
		ExpressionAttributeValues: &map[string]*dynamodb.AttributeValue{
			"Key": &dynamodb.AttributeValue{ // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Boolean(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: &map[string]*dynamodb.AttributeValue{
					"Key": &dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Boolean(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		ReturnConsumedCapacity:      aws.String("ReturnConsumedCapacity"),
		ReturnItemCollectionMetrics: aws.String("ReturnItemCollectionMetrics"),
		ReturnValues:                aws.String("ReturnValue"),
	}
	resp, err := svc.PutItem(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleDynamoDB_Query() {
	svc := dynamodb.New(nil)

	params := &dynamodb.QueryInput{
		KeyConditions: &map[string]*dynamodb.Condition{ // Required
			"Key": &dynamodb.Condition{ // Required
				ComparisonOperator: aws.String("ComparisonOperator"), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
						B:    []byte("PAYLOAD"),
						BOOL: aws.Boolean(true),
						BS: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						L: []*dynamodb.AttributeValue{
							&dynamodb.AttributeValue{ // Required
							// Recursive values...
							},
							// More values...
						},
						M: &map[string]*dynamodb.AttributeValue{
							"Key": &dynamodb.AttributeValue{ // Required
							// Recursive values...
							},
							// More values...
						},
						N: aws.String("NumberAttributeValue"),
						NS: []*string{
							aws.String("NumberAttributeValue"), // Required
							// More values...
						},
						NULL: aws.Boolean(true),
						S:    aws.String("StringAttributeValue"),
						SS: []*string{
							aws.String("StringAttributeValue"), // Required
							// More values...
						},
					},
					// More values...
				},
			},
			// More values...
		},
		TableName: aws.String("TableName"), // Required
		AttributesToGet: []*string{
			aws.String("AttributeName"), // Required
			// More values...
		},
		ConditionalOperator: aws.String("ConditionalOperator"),
		ConsistentRead:      aws.Boolean(true),
		ExclusiveStartKey: &map[string]*dynamodb.AttributeValue{
			"Key": &dynamodb.AttributeValue{ // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Boolean(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: &map[string]*dynamodb.AttributeValue{
					"Key": &dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Boolean(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		ExpressionAttributeNames: &map[string]*string{
			"Key": aws.String("AttributeName"), // Required
			// More values...
		},
		ExpressionAttributeValues: &map[string]*dynamodb.AttributeValue{
			"Key": &dynamodb.AttributeValue{ // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Boolean(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: &map[string]*dynamodb.AttributeValue{
					"Key": &dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Boolean(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		FilterExpression:     aws.String("ConditionExpression"),
		IndexName:            aws.String("IndexName"),
		Limit:                aws.Long(1),
		ProjectionExpression: aws.String("ProjectionExpression"),
		QueryFilter: &map[string]*dynamodb.Condition{
			"Key": &dynamodb.Condition{ // Required
				ComparisonOperator: aws.String("ComparisonOperator"), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
						B:    []byte("PAYLOAD"),
						BOOL: aws.Boolean(true),
						BS: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						L: []*dynamodb.AttributeValue{
							&dynamodb.AttributeValue{ // Required
							// Recursive values...
							},
							// More values...
						},
						M: &map[string]*dynamodb.AttributeValue{
							"Key": &dynamodb.AttributeValue{ // Required
							// Recursive values...
							},
							// More values...
						},
						N: aws.String("NumberAttributeValue"),
						NS: []*string{
							aws.String("NumberAttributeValue"), // Required
							// More values...
						},
						NULL: aws.Boolean(true),
						S:    aws.String("StringAttributeValue"),
						SS: []*string{
							aws.String("StringAttributeValue"), // Required
							// More values...
						},
					},
					// More values...
				},
			},
			// More values...
		},
		ReturnConsumedCapacity: aws.String("ReturnConsumedCapacity"),
		ScanIndexForward:       aws.Boolean(true),
		Select:                 aws.String("Select"),
	}
	resp, err := svc.Query(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleDynamoDB_Scan() {
	svc := dynamodb.New(nil)

	params := &dynamodb.ScanInput{
		TableName: aws.String("TableName"), // Required
		AttributesToGet: []*string{
			aws.String("AttributeName"), // Required
			// More values...
		},
		ConditionalOperator: aws.String("ConditionalOperator"),
		ExclusiveStartKey: &map[string]*dynamodb.AttributeValue{
			"Key": &dynamodb.AttributeValue{ // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Boolean(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: &map[string]*dynamodb.AttributeValue{
					"Key": &dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Boolean(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		ExpressionAttributeNames: &map[string]*string{
			"Key": aws.String("AttributeName"), // Required
			// More values...
		},
		ExpressionAttributeValues: &map[string]*dynamodb.AttributeValue{
			"Key": &dynamodb.AttributeValue{ // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Boolean(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: &map[string]*dynamodb.AttributeValue{
					"Key": &dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Boolean(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		FilterExpression:       aws.String("ConditionExpression"),
		IndexName:              aws.String("IndexName"),
		Limit:                  aws.Long(1),
		ProjectionExpression:   aws.String("ProjectionExpression"),
		ReturnConsumedCapacity: aws.String("ReturnConsumedCapacity"),
		ScanFilter: &map[string]*dynamodb.Condition{
			"Key": &dynamodb.Condition{ // Required
				ComparisonOperator: aws.String("ComparisonOperator"), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
						B:    []byte("PAYLOAD"),
						BOOL: aws.Boolean(true),
						BS: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						L: []*dynamodb.AttributeValue{
							&dynamodb.AttributeValue{ // Required
							// Recursive values...
							},
							// More values...
						},
						M: &map[string]*dynamodb.AttributeValue{
							"Key": &dynamodb.AttributeValue{ // Required
							// Recursive values...
							},
							// More values...
						},
						N: aws.String("NumberAttributeValue"),
						NS: []*string{
							aws.String("NumberAttributeValue"), // Required
							// More values...
						},
						NULL: aws.Boolean(true),
						S:    aws.String("StringAttributeValue"),
						SS: []*string{
							aws.String("StringAttributeValue"), // Required
							// More values...
						},
					},
					// More values...
				},
			},
			// More values...
		},
		Segment:       aws.Long(1),
		Select:        aws.String("Select"),
		TotalSegments: aws.Long(1),
	}
	resp, err := svc.Scan(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleDynamoDB_UpdateItem() {
	svc := dynamodb.New(nil)

	params := &dynamodb.UpdateItemInput{
		Key: &map[string]*dynamodb.AttributeValue{ // Required
			"Key": &dynamodb.AttributeValue{ // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Boolean(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: &map[string]*dynamodb.AttributeValue{
					"Key": &dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Boolean(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		TableName: aws.String("TableName"), // Required
		AttributeUpdates: &map[string]*dynamodb.AttributeValueUpdate{
			"Key": &dynamodb.AttributeValueUpdate{ // Required
				Action: aws.String("AttributeAction"),
				Value: &dynamodb.AttributeValue{
					B:    []byte("PAYLOAD"),
					BOOL: aws.Boolean(true),
					BS: [][]byte{
						[]byte("PAYLOAD"), // Required
						// More values...
					},
					L: []*dynamodb.AttributeValue{
						&dynamodb.AttributeValue{ // Required
						// Recursive values...
						},
						// More values...
					},
					M: &map[string]*dynamodb.AttributeValue{
						"Key": &dynamodb.AttributeValue{ // Required
						// Recursive values...
						},
						// More values...
					},
					N: aws.String("NumberAttributeValue"),
					NS: []*string{
						aws.String("NumberAttributeValue"), // Required
						// More values...
					},
					NULL: aws.Boolean(true),
					S:    aws.String("StringAttributeValue"),
					SS: []*string{
						aws.String("StringAttributeValue"), // Required
						// More values...
					},
				},
			},
			// More values...
		},
		ConditionExpression: aws.String("ConditionExpression"),
		ConditionalOperator: aws.String("ConditionalOperator"),
		Expected: &map[string]*dynamodb.ExpectedAttributeValue{
			"Key": &dynamodb.ExpectedAttributeValue{ // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
						B:    []byte("PAYLOAD"),
						BOOL: aws.Boolean(true),
						BS: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						L: []*dynamodb.AttributeValue{
							&dynamodb.AttributeValue{ // Required
							// Recursive values...
							},
							// More values...
						},
						M: &map[string]*dynamodb.AttributeValue{
							"Key": &dynamodb.AttributeValue{ // Required
							// Recursive values...
							},
							// More values...
						},
						N: aws.String("NumberAttributeValue"),
						NS: []*string{
							aws.String("NumberAttributeValue"), // Required
							// More values...
						},
						NULL: aws.Boolean(true),
						S:    aws.String("StringAttributeValue"),
						SS: []*string{
							aws.String("StringAttributeValue"), // Required
							// More values...
						},
					},
					// More values...
				},
				ComparisonOperator: aws.String("ComparisonOperator"),
				Exists:             aws.Boolean(true),
				Value: &dynamodb.AttributeValue{
					B:    []byte("PAYLOAD"),
					BOOL: aws.Boolean(true),
					BS: [][]byte{
						[]byte("PAYLOAD"), // Required
						// More values...
					},
					L: []*dynamodb.AttributeValue{
						&dynamodb.AttributeValue{ // Required
						// Recursive values...
						},
						// More values...
					},
					M: &map[string]*dynamodb.AttributeValue{
						"Key": &dynamodb.AttributeValue{ // Required
						// Recursive values...
						},
						// More values...
					},
					N: aws.String("NumberAttributeValue"),
					NS: []*string{
						aws.String("NumberAttributeValue"), // Required
						// More values...
					},
					NULL: aws.Boolean(true),
					S:    aws.String("StringAttributeValue"),
					SS: []*string{
						aws.String("StringAttributeValue"), // Required
						// More values...
					},
				},
			},
			// More values...
		},
		ExpressionAttributeNames: &map[string]*string{
			"Key": aws.String("AttributeName"), // Required
			// More values...
		},
		ExpressionAttributeValues: &map[string]*dynamodb.AttributeValue{
			"Key": &dynamodb.AttributeValue{ // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Boolean(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					&dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: &map[string]*dynamodb.AttributeValue{
					"Key": &dynamodb.AttributeValue{ // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Boolean(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		ReturnConsumedCapacity:      aws.String("ReturnConsumedCapacity"),
		ReturnItemCollectionMetrics: aws.String("ReturnItemCollectionMetrics"),
		ReturnValues:                aws.String("ReturnValue"),
		UpdateExpression:            aws.String("UpdateExpression"),
	}
	resp, err := svc.UpdateItem(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}

func ExampleDynamoDB_UpdateTable() {
	svc := dynamodb.New(nil)

	params := &dynamodb.UpdateTableInput{
		TableName: aws.String("TableName"), // Required
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			&dynamodb.AttributeDefinition{ // Required
				AttributeName: aws.String("KeySchemaAttributeName"), // Required
				AttributeType: aws.String("ScalarAttributeType"),    // Required
			},
			// More values...
		},
		GlobalSecondaryIndexUpdates: []*dynamodb.GlobalSecondaryIndexUpdate{
			&dynamodb.GlobalSecondaryIndexUpdate{ // Required
				Create: &dynamodb.CreateGlobalSecondaryIndexAction{
					IndexName: aws.String("IndexName"), // Required
					KeySchema: []*dynamodb.KeySchemaElement{ // Required
						&dynamodb.KeySchemaElement{ // Required
							AttributeName: aws.String("KeySchemaAttributeName"), // Required
							KeyType:       aws.String("KeyType"),                // Required
						},
						// More values...
					},
					Projection: &dynamodb.Projection{ // Required
						NonKeyAttributes: []*string{
							aws.String("NonKeyAttributeName"), // Required
							// More values...
						},
						ProjectionType: aws.String("ProjectionType"),
					},
					ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
						ReadCapacityUnits:  aws.Long(1), // Required
						WriteCapacityUnits: aws.Long(1), // Required
					},
				},
				Delete: &dynamodb.DeleteGlobalSecondaryIndexAction{
					IndexName: aws.String("IndexName"), // Required
				},
				Update: &dynamodb.UpdateGlobalSecondaryIndexAction{
					IndexName: aws.String("IndexName"), // Required
					ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
						ReadCapacityUnits:  aws.Long(1), // Required
						WriteCapacityUnits: aws.Long(1), // Required
					},
				},
			},
			// More values...
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Long(1), // Required
			WriteCapacityUnits: aws.Long(1), // Required
		},
	}
	resp, err := svc.UpdateTable(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.StringValue(resp))
}
