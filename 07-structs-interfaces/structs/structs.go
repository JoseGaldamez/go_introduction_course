package structs

type Product struct {
	ID    uint     `json:"id"`
	Name  string   `json:"name"`
	Type  Type     `json:"type"`
	Count uint     `json:"count"`
	Price float64  `json:"price"`
	Tags  []string `json:"tags"`
}

type Type struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}

func (product Product) TotalPrice() float64 {
	return float64(product.Count) * product.Price
}

func (product *Product) SetName(name string) {
	product.Name = name
}

func (product *Product) SetTags(tags ...string) {
	product.Tags = append(product.Tags, tags...)
}

type Car struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

func (car *Car) AddProducts(products ...Product) {
	car.Products = append(car.Products, products...)
}

func (car Car) Total() float64 {
	var total float64

	for _, product := range car.Products {
		total += product.TotalPrice()
	}

	return total
}

func NewCar(userID uint) Car {
	return Car{UserID: userID}
}
