// For full tutorial go to this url: "https://medium.com/@IndianGuru/using-google-geocoding-and-street-view-image-apis-with-go-b67bb4841ff0"
package geowebgap

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"appengine"
	"appengine/urlfetch"
)
// This structype is for the Response(or Resullt). Note that there is no struct type for the Request
type Response struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64
				Lng float64
			}
		}
	}
}

func init() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/showimage", showimage)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, rootForm)
}

const rootForm = `
  <!DOCTYPE html>
    <html>
      <head>
        <meta charset="utf-8">
        <title>Accept Address</title>
      </head>
      <body>
        <h1>Accept Address</h1>
        <p>Please enter your address:</p>
        <form action="/showimage" method="post" accept-charset="utf-8">
	  <input type="text" name="str" value="Type address..." id="str">
	  <input type="submit" value=".. and see the image!">
        </form>
      </body>
    </html>
`

func check(e error, str string) {
	if e != nil {
		log.Fatal(str, " ", e)
		return
	}
}

var upperTemplate = template.Must(template.New("showimage").Parse(upperTemplateHTML))

func showimage(w http.ResponseWriter, r *http.Request) {
	// Sample address "1600 Amphitheatre Parkway, Mountain View, CA"
	addr := r.FormValue("str")

	// QueryEscape escapes the addr string so
	// it can be safely placed inside a URL query
	safeAddr := url.QueryEscape(addr)

	// Geocoding API
	fullUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/geocode/json?address=%s", safeAddr)

	// One needs an appengine.Context to interact with external
	// services, including email and urlfetch.
	c := appengine.NewContext(r)

	// Function Client returns an *http.Client
	client := urlfetch.Client(c)

	resp, err := client.Get(fullUrl)
	check(err, "Get:")

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	var res Response

	// We generate the latitude and longitude using "The Google Geocoding API".
	// Geocoding is the process of converting an address (like
	// "1600 Amphitheatre Parkway, Mountain View, CA") into its geographic
	// coordinates (like latitude 37.423021 and longitude -122.083739).
	// Use json.Decode or json.Encode for reading or writing streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Println(err)
	}

	// lat, lng as float64
	lat := res.Results[0].Geometry.Location.Lat
	lng := res.Results[0].Geometry.Location.Lng

	//safeLatLng := url.QueryEscape(fmt.Sprintf("%.13f,%.13f", lat, lng))

	// %.13f is used to convert float64 to a string
	queryUrl := fmt.Sprintf("http://maps.googleapis.com/maps/api/streetview?size=600x300&location=%.13f,%.13f", lat, lng)

	tempErr := upperTemplate.Execute(w, queryUrl)
	check(tempErr, "Execute:")
}

const upperTemplateHTML = ` 
<!DOCTYPE html>
  <html>
    <head>
      <meta charset="utf-8">
      <title>Display Image</title>
    </head>
    <body>
      <h3>Image at your Address</h3>
      <img src="{{html .}}" alt="Image" />
    </body>
  </html>
`
