package main

func getProducts() ([]Product, error) {
	products := []Product{}
	rows, err := db.Query("SELECT id, name, price from products")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID, &p.Name, &p.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func getProduct(id int) (Product, error) {
	var p Product
	row := db.QueryRow("SELECT id, name, price FROM products WHERE id = $1;", id)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func createProduct(product *Product) error {
	_, err := db.Exec(
		"INSERT INTO products(name, price) VALUES ($1, $2);",
		product.Name,
		product.Price,
	)
	return err
}

func updateProduct(id int, product *Product) (Product, error) {
	var p Product
	row := db.QueryRow(
		"UPDATE products SET name=$1, price=$2 WHERE id=$3 RETURNING id, name, price;",
		product.Name,
		product.Price,
		id,
	)
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	if err != nil {
		return Product{}, err
	}
	return p, err
}

func deleteProduct(id int) error {
	_, err := db.Exec(
		"DELETE FROM products WHERE id=$1;",
		id,
	)
	return err
}
