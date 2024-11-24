package products

import (
	"context"
	"log"
	"testing"

	"github.com/ardwiinoo/online-shop/external/database"
	"github.com/ardwiinoo/online-shop/infra/response"
	"github.com/ardwiinoo/online-shop/internal/config"
	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)

	if err != nil {
		panic(err)
	}

	db, err := database.ConnectPostgres(config.Cfg.DB)

	if err != nil {
		panic(err)
	}

	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateProduct_Success(t *testing.T) {
	req := CreateProductRequestPayload{
		Name: "Baju baru",
		Stock: 10,
		Price: 10_000,
	}

	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}

func TestCreateProduct_Fail(t *testing.T) {
	t.Run("Failed Create Product, name is required", func(t *testing.T) {
		req := CreateProductRequestPayload{
			Name: "",
			Stock: 10,
			Price: 10_000,
		}

		err := svc.CreateProduct(context.Background(), req)
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductRequired, err)
	})
}

func TestListProduct_Success(t *testing.T) {
	pagination := ListProductRequestPayload{
		Cursor: 0,
		Size: 10,
	}

	products, err := svc.ListProducts(context.Background(), pagination)
	require.Nil(t, err)
	require.NotNil(t, products)
	log.Printf("%+v", products)
}