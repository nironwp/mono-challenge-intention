package dto

type IntentProduct struct {
	Title       string  `bson:"title" valid:"required" json:"title"`
	Price       float64 `bson:"price" valid:"required,float" json:"price"`
	Category    string  `bson:"category" valid:"required" json:"category"`
	Description string  `bson:"description" valid:"required" json:"description"`
	Image       string  `bson:"image" valid:"required" json:"image"`
}
