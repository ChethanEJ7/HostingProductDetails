package main

// ProductDetails - Tool configuration JSON file model structure
type ProductDetails struct {
	ProductOwner            string `json:"owner"`
	ProductType             string `json:"type"`
	ProductDate             string `json:"date"`
	ProductAlbum            bool   `json:"album"`
	ProductVideo            bool   `json:"video"`
	ProductCinematic        bool   `json:"cinematic"`
	ProductDrone            bool   `json:"drone"`
	ProductEstimation       string `json:"estimation"`
	ProductVenue            string `json:"venue"`
	ProductContact          int64  `json:"contact"`
	ProductAlternateContact int64  `json:"alternatecontact"`
}
type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article
