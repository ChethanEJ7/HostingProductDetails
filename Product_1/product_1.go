package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/kardianos/osext"
	"tawesoft.co.uk/go/dialog"
)

func main() {
	// Read the product contents of JSON file
	productDetails, err := readProductDetails()
	if err != nil {
		fmt.Println("Error reading product details from JSON file, ", err)
		fmt.Fprintln(os.Stderr, "Oh no, something went wrong: %v\n", err)
		dialog.Alert("Error : Not able to upload the details")
		os.Exit(1)
	} else {
		// validate JSON file contents
		validationErr := validateProductDetails(productDetails)
		if validationErr != nil {
			fmt.Println("Details missing in product-details.json. ", validationErr, "Kindly Update the required information and re-try.")
			fmt.Fprintln(os.Stderr, "Oh no, something went wrong: %v\n", err)
			os.Exit(1)
		}
	}

	// fmt.Print(productDetails)
	setupRestAPICall(productDetails)
	handleRequests()

}
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
func setupRestAPICall(productDetails *ProductDetails) {
	// fmt.Println(productDetails)
	Articles = []Article{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", returnAllArticles)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

//Method to read product details from product-details.json file
func readProductDetails() (*ProductDetails, error) {
	filename, _ := osext.Executable()
	fileDirectory := filepath.Dir(filename)
	jsonFilePath := filepath.Join(fileDirectory, `product-details.json`)
	content, err := ioutil.ReadFile(jsonFilePath)
	if err != nil {
		return nil, err
	}

	productDetails := &ProductDetails{}
	err = json.Unmarshal(content, productDetails)
	if err != nil {
		return nil, err
	}
	return productDetails, nil
}

//Method to validate toolname, artifact Id and entry points provided in tool details
func validateProductDetails(productDetails *ProductDetails) error {

	if (strings.Trim(productDetails.ProductOwner, " ")) == "" {
		return fmt.Errorf("owner is not/incorrectly set in product-details.json. Unable to run the Tool ")
	}
	if (strings.Trim(productDetails.ProductType, " ")) == "" {
		return fmt.Errorf("type is not/incorrectly set in product-details.json. Unable to run the Tool ")
	}
	if (strings.Trim(productDetails.ProductDate, " ")) == "" {
		return fmt.Errorf("date is not/incorrectly set in product-details.json. Unable to run the Tool ")
	}
	if (strings.Trim(productDetails.ProductType, " ")) == "" {
		return fmt.Errorf("type is not/incorrectly set in product-details.json. Unable to run the Tool ")
	}
	// if bool(productDetails.ProductAlbum) {
	// 	return fmt.Errorf("album is not/incorrectly set in product-details.json. Unable to run the Tool ")
	// }
	// if bool(productDetails.ProductVideo) {
	// 	return fmt.Errorf("video is not/incorrectly set in product-details.json. Unable to run the Tool ")
	// }
	// if bool(productDetails.ProductCinematic) {
	// 	return fmt.Errorf("cinematic is not/incorrectly set in product-details.json. Unable to run the Tool ")
	// }
	// if bool(productDetails.ProductDrone) {
	// 	return fmt.Errorf("drone is not/incorrectly set in product-details.json. Unable to run the Tool ")
	// }
	if (strings.Trim(productDetails.ProductEstimation, " ")) == "" {
		return fmt.Errorf("estimation is not/incorrectly set in product-details.json. Unable to run the Tool ")
	}
	if (strings.Trim(productDetails.ProductVenue, " ")) == "" {
		return fmt.Errorf("venue is not/incorrectly set in product-details.json. Unable to run the Tool ")
	}
	return nil
}
