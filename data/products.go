package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

// Product defines the structure for an API product
//Since encoding/json package can only marshal/unmarshal exported fields, we need to make sure that the fields in our struct are exported. This means that the first letter of the field name must be capitalized.
//To get the JSON tags to work, we need to make sure that the field names in our struct match the JSON field names exactly.
//We can use the omitempty option to tell the JSON package to omit the field from the encoded JSON if the field has an empty value.

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products is a defining slice of Product
type Products []*Product

// ToJSON serializes the contents of the collection to JSON
func (p *Products) ToJSON(w io.Writer) error {
	// NewEncoder requires an io.Writer interface to write the JSON to (in this case, the http.ResponseWriter)
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// FromJSON is a Method on type Product (slice of products)
func (p *Product) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(p)
}

//GetProducts - Returns the product list
func GetProducts() Products {
	return productList
}

//AddProduct - Adds a product to  our struct Product
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

//UpdateProduct - Updates a product in our struct Product
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	productList[pos] = p
	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, fmt.Errorf("Product with ID '%v' not found", id)
}

// ErrProductNotFound is the Standard error message for a product not found
var ErrProductNotFound = fmt.Errorf("Product not found")

// Increments the Product ID by one
func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		SKU:         "abc323",
		Description: "Frothy milky coffee",
		Price:       2.45,
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		SKU:         "fjd34",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		CreatedOn:   time.Now().UTC().String(),

		UpdatedOn: time.Now().UTC().String(),
	},
}
