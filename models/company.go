package models

// Company represents a company document in MongoDB
type Company struct {
	ID      string `json:"id,omitempty" bson:"_id,omitempty"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
	Website string `json:"website" bson:"website"`
}
