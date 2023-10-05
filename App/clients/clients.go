package clients

import (
	"database/sql"
	"fmt"
	"restArchitecture/mikail/App/models"
)

type ProductClient struct {
	db *sql.DB
}

type ProductClientRepo interface {
	InsertProductList(request *models.ProductList) error
	SortDateProductList(order string) ([]*models.SortProductResponse, error)
	SortPriceProductList(order string) ([]*models.SortProductResponse, error)
	SortNameProductList(order string) ([]*models.SortProductResponse, error)
}

func NewProductClient(db *sql.DB) *ProductClient {
	return &ProductClient{db: db}
}

func (m *ProductClient) InsertProductList(request *models.ProductList) error {
	_, err := m.db.Exec("INSERT INTO product_services.product_list "+
		"(product_name, product_desc, product_price, product_quantity, date_created) VALUES (?, ?, ?, ?, ?)",
		request.ProductName, request.ProductDesc, request.ProductPrice, request.ProductQuantity, request.CreatedDate)
	if err != nil {
		return err
	}
	return nil
}

func (m *ProductClient) SortDateProductList(order string) ([]*models.SortProductResponse, error) {
	query := fmt.Sprintf(`SELECT id, product_name, product_desc,product_price, product_quantity from product_services.product_list order by date_created %s`, order)
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []*models.SortProductResponse

	for rows.Next() {
		var product models.SortProductResponse
		err := rows.Scan(&product.Id, &product.ProductName, &product.ProductDesc, &product.ProductPrice, &product.ProductQuantity)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (m *ProductClient) SortPriceProductList(order string) ([]*models.SortProductResponse, error) {
	query := fmt.Sprintf(`SELECT id, product_name, product_desc,product_price, product_quantity from product_services.product_list order by product_price %s`, order)
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []*models.SortProductResponse

	for rows.Next() {
		var product models.SortProductResponse
		err := rows.Scan(&product.Id, &product.ProductName, &product.ProductDesc, &product.ProductPrice, &product.ProductQuantity)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func (m *ProductClient) SortNameProductList(order string) ([]*models.SortProductResponse, error) {
	query := fmt.Sprintf(`SELECT id, product_name, product_desc,product_price, product_quantity from product_services.product_list order by product_name %s`, order)
	rows, err := m.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []*models.SortProductResponse

	for rows.Next() {
		var product models.SortProductResponse
		err := rows.Scan(&product.Id, &product.ProductName, &product.ProductDesc, &product.ProductPrice, &product.ProductQuantity)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}
