package endpoints

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"

	customEcho "github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/http/custom_echo"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/logger"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/messaging/bus"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/catalogs/read_service/internal/products/delivery"
	getProductByIdV1 "github.com/mehdihadeli/go-ecommerce-microservices/internal/services/catalogs/read_service/internal/products/features/get_product_by_id/v1/endpoints"
	getProductsV1 "github.com/mehdihadeli/go-ecommerce-microservices/internal/services/catalogs/read_service/internal/products/features/getting_products/v1/endpoints"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/catalogs/read_service/internal/products/features/searching_products/v1/endpoints"
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/services/catalogs/read_service/internal/shared/contracts"
)

func ConfigProductsEndpoints(
	routeBuilder *customEcho.RouteBuilder,
	bus bus.Bus,
	metrics *contracts.CatalogsMetrics,
	validator *validator.Validate,
	logger logger.Logger,
) {
	configV1Endpoints(routeBuilder, bus, metrics, validator, logger)
}

func configV1Endpoints(
	routeBuilder *customEcho.RouteBuilder,
	bus bus.Bus,
	metrics *contracts.CatalogsMetrics,
	validator *validator.Validate,
	logger logger.Logger,
) {
	routeBuilder.RegisterGroupFunc("/api/v1", func(v1 *echo.Group) {
		group := v1.Group("/products")
		productEndpointBase := delivery.NewProductEndpointBase(
			logger,
			validator,
			group,
			bus,
			metrics,
		)

		// GetProducts
		getProductsEndpoint := getProductsV1.NewGetProductsEndpoint(productEndpointBase)
		getProductsEndpoint.MapRoute()

		// SearchProducts
		searchProductsEndpoint := endpoints.NewSearchProductsEndpoint(productEndpointBase)
		searchProductsEndpoint.MapRoute()

		// GetProductById
		getProductByIdEndpoint := getProductByIdV1.NewGetProductByIdEndpoint(productEndpointBase)
		getProductByIdEndpoint.MapRoute()
	})
}
