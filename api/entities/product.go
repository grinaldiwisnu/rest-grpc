package entities

type Product struct {
	tableName struct{} `pg:"products,alias:products"`
	Id        string   `json:"id" pg:"id, pk"`
	Name      string   `json:"name" pg:"name"`
	Price     int64    `json:"price" pg:"price"`
	Quantity  int32    `json:"quantity" pg:"quantity"`
}
