//go:build e2e
// +build e2e

package grpc

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"

	productService "github.com/mehdihadeli/go-ecommerce-microservices/internal/services/catalogwriteservice/internal/shared/grpc/genproto"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/catalogwriteservice/internal/shared/test_fixtures/integration"
)

type productGrpcServiceE2eTests struct {
	*integration.IntegrationTestSharedFixture
}

func TestProductGrpcServiceEndToEnd(t *testing.T) {
	suite.Run(
		t,
		&productGrpcServiceE2eTests{
			IntegrationTestSharedFixture: integration.NewIntegrationTestSharedFixture(t),
		},
	)
}

func (c *productGrpcServiceE2eTests) Test_Should_Create_Product_With_Valid_Data_In_DB() {
	ctx := context.Background()

	request := &productService.CreateProductReq{
		Price:       gofakeit.Price(100, 1000),
		Name:        gofakeit.Name(),
		Description: gofakeit.AdjectiveDescriptive(),
	}

	res, err := c.ProductServiceClient.CreateProduct(ctx, request)
	c.Require().NoError(err)
	c.NotEmpty(res.ProductId)
}

func (c *productGrpcServiceE2eTests) Test_Should_Return_Data_With_Valid_Id() {
	ctx := context.Background()
	id := c.Items[0].ProductId.String()

	res, err := c.ProductServiceClient.GetProductById(
		ctx,
		&productService.GetProductByIdReq{ProductId: id},
	)

	fmt.Println(err)
	fmt.Println(res)
	c.Require().NoError(err)
	c.NotNil(res.Product)
	c.Equal(res.Product.ProductId, id)
}
