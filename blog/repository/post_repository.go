package repository

import (
	"blog/model"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type PostRepository struct {
	post  *model.Post
	posts *[]model.Post
}

var (
	tableName = os.Getenv("DYNAMODB_TABLE_NAME")
)

func NewPostRepository() *PostRepository {
	pr := new(PostRepository)
	pr.post = model.NewPost()

	return pr
}

func (pr *PostRepository) GetOne(id string) (model.Post, error) {
	result, err := DB.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberN{Value: id},
		},
	})

	if err != nil {
		return *pr.post, err
	}

	if err := attributevalue.Unmarshal(&types.AttributeValueMemberM{Value: result.Item}, pr.post); err != nil {
		return *pr.post, err
	}

	return *pr.post, nil
}

func (pr *PostRepository) GetAll() ([]model.Post, error) {
	pr.posts = &[]model.Post{}

	result, err := DB.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		return nil, err
	}

	if err := attributevalue.UnmarshalListOfMaps(result.Items, pr.posts); err != nil {
		return nil, err
	}

	return *pr.posts, nil
}

func (pr *PostRepository) Post(post *model.Post) (bool, error) {
	_, err := DB.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"Id":       &types.AttributeValueMemberN{Value: fmt.Sprint(post.ID)},
			"Title":    &types.AttributeValueMemberS{Value: post.Title},
			"Desc":     &types.AttributeValueMemberS{Value: post.Desc},
			"Content":  &types.AttributeValueMemberS{Value: post.Content},
			"Tag":      &types.AttributeValueMemberS{Value: post.Tag},
			"Locale":   &types.AttributeValueMemberS{Value: post.Locale},
			"PostedAt": &types.AttributeValueMemberS{Value: post.PostedAt},
		},
	})

	if err != nil {
		return false, err
	}

	return true, nil
}

func (pr *PostRepository) Edit(post *model.Post) (bool, error) {
	update := expression.UpdateBuilder{}.Set(expression.Name("Title"), expression.Value(post.Title))
	update = update.Set(expression.Name("Desc"), expression.Value(post.Desc))
	update = update.Set(expression.Name("Content"), expression.Value(post.Content))
	update = update.Set(expression.Name("Tag"), expression.Value(post.Tag))
	update = update.Set(expression.Name("Locale"), expression.Value(post.Locale))
	update = update.Set(expression.Name("PostedAt"), expression.Value(post.PostedAt))

	expr, err := expression.NewBuilder().WithUpdate(update).Build()

	if err != nil {
		return false, err
	}
	updateInput := &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberN{Value: fmt.Sprint(post.ID)},
		},
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		UpdateExpression:          expr.Update(),
	}

	_, err = DB.UpdateItem(context.Background(), updateInput)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (pr *PostRepository) Delete(id string) (bool, error) {
	_, err := DB.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]types.AttributeValue{
			"Id": &types.AttributeValueMemberN{Value: id},
		},
	})

	if err != nil {
		return false, err
	}

	return true, nil
}
