package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"bufio"
	"image/jpeg"
	"image"
	"image/draw"
	"bytes"
	"encoding/json"
	"strconv"
	"github.com/disintegration/imaging"
	//"./src/imaging"
)

//Parameters:
var PORT = "80"
var URL = "url"
var WIDTH = "width"
var HEIGHT = "height"

func Log(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}

func sizerHandler(res http.ResponseWriter, req *http.Request) {

	log.Printf("Input Req: %s %s\n", req.Host, req.URL.Path)

	err, cause := validateParams(req.URL.Query())
	if err != http.StatusOK {
		wrapJsonError(res, err, cause)
		return
	}
	log.Printf("VALID Req: %s %s   %s\n", req.Host, req.URL.Path, req.URL.Query())

	//Parse param values
	paramValues := req.URL.Query()

	//Get image
	var origImage image.Image
	origImage, err, cause = loadImage(paramValues.Get(URL))
	if err != http.StatusOK {
		wrapJsonError(res, err, cause)
		return
	}
	//Original to file
	saveImageToFile(origImage)

	//Get requested
	userWidth, _ := strconv.Atoi(paramValues.Get(WIDTH))
	userHeight, _ := strconv.Atoi(paramValues.Get(HEIGHT))

	resizedImage := resizeImage(origImage, userWidth, userHeight)

	//Resized to file
	saveImageToFile(resizedImage)

	//Set jpeg into response
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, resizedImage, nil); err != nil {
		wrapJsonError(res, http.StatusInternalServerError, "unable to encode image")
		return
	}

	//Wrap Response  - 200
	if _, err := res.Write(buffer.Bytes()); err != nil {
		wrapJsonError(res, http.StatusInternalServerError, "unable to write image")
		return
	}
	res.Header().Set("Content-Type", "image/jpeg")
	res.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
}

func adjustDimensions(inImage image.Image, userWidth int, userHeight int) (int, int) {

	//Calculate ratio
	ratio := float64(inImage.Bounds().Dx()) / float64(inImage.Bounds().Dy())

	//Adjust best size by ration
	if (int(float64(userHeight)*ratio) > userWidth) { //RONEN???
		return userWidth, int(float64(userWidth) / ratio)
	} else {
		return int(float64(userHeight) * ratio), userHeight //RONEN???
	}
}

func resizeImage(inputImage image.Image, userWidth int, userHeight int) *image.NRGBA {

	if (userWidth >= inputImage.Bounds().Dx() && userHeight >= inputImage.Bounds().Dy()) {
		//pad orig image
		return padImage(inputImage, image.Rect(0, 0, userWidth, userHeight))
	} else {
		//Create inner sized image
		adjWidth, adjHeight := adjustDimensions(inputImage, userWidth, userHeight)
		log.Printf("ADJ   %d  %d  ", adjWidth, adjHeight)
		sizedImage := imaging.Resize(inputImage, adjWidth, adjHeight, imaging.Lanczos)
		return padImage(sizedImage, image.Rect(0, 0, userWidth, userHeight))
	}
}

func padImage(inputImage image.Image, dest image.Rectangle) *image.NRGBA {
	//Pad adjusted to requested
	xDiff := dest.Dx()-inputImage.Bounds().Dx()
	imageWidth := inputImage.Bounds().Dx()
	yDiff := dest.Dy()-inputImage.Bounds().Dy()
	imageHeight := inputImage.Bounds().Dy()

	log.Printf("PADDING   %d   %d", inputImage.Bounds(), dest)
	padImage := image.NewNRGBA(image.Rect(0, 0, dest.Dx(), dest.Dy()))
	draw.Draw(padImage, padImage.Bounds(), image.Black, image.ZP, draw.Src)
	innerRect := image.Rect(xDiff/2, yDiff/2, (xDiff/2) +imageWidth, (yDiff/2)+imageHeight)
	draw.Draw(padImage, innerRect, inputImage, image.Point{0, 0}, draw.Src)

	return padImage
}

func saveImageToFile(myImage image.Image) (error) { //Errors are ignored

	// Prepare parent image where we want to position child image.
	target := image.NewRGBA(myImage.Bounds())
	// Draw white layer.
	draw.Draw(target, target.Bounds(), image.White, image.ZP, draw.Src)
	// Draw child image.
	draw.Draw(target, myImage.Bounds(), myImage, image.Point{0, 0}, draw.Src)

	// Encode to jpeg.
	var imageBuf bytes.Buffer
	err := jpeg.Encode(&imageBuf, target, nil)

	if err != nil {
		return err
	}

	// Write to file.
	//File name
	var fileName bytes.Buffer
	fileName.WriteString("w")
	fileName.WriteString(strconv.Itoa(myImage.Bounds().Max.X))
	fileName.WriteString("_h")
	fileName.WriteString(strconv.Itoa(myImage.Bounds().Max.Y))
	fileName.WriteString("_img.jpeg")

	fo, err := os.Create(fileName.String())
	if err != nil {
		return err
	}
	fw := bufio.NewWriter(fo)

	fw.Write(imageBuf.Bytes())
	return err
}

func loadImage(toTest string) (image.Image, int, string) {
	//Verify URL exists
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return nil, http.StatusBadRequest, "url not found"
	} else {
		resp, err := http.Get(toTest)
		defer resp.Body.Close()
		//Verify URL is an image and decode
		decodedOrigImage, _, err := image.Decode(resp.Body)
		if err != nil {
			return nil, http.StatusBadRequest, "url is not an image"
		}
		return decodedOrigImage, http.StatusOK, ""
	}
}

func validateParams(toTest url.Values) (res int, err string) {
	if toTest.Get(URL) == "" {
		return http.StatusBadRequest, URL + "parameter is required"
	}
	if toTest.Get(WIDTH) == "" {
		return http.StatusBadRequest, WIDTH + " parameter is required"
	}
	if _, err := strconv.Atoi(toTest.Get(WIDTH)); err != nil {
		return http.StatusBadRequest, WIDTH + "  parameter must be a number"
	}
	if toTest.Get(HEIGHT) == "" {
		return http.StatusBadRequest, HEIGHT + " parameter is required"
	}
	if _, err := strconv.Atoi(toTest.Get(HEIGHT)); err != nil {
		return http.StatusBadRequest, HEIGHT + " parameter must be a number"
	}
	return http.StatusOK, ""
}

type handler func(w http.ResponseWriter, r *http.Request)

func verifyGetOnly(h handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h(w, r)
			return
		}
		http.Error(w, "This service serves GET only", http.StatusMethodNotAllowed)
	}
}

func wrapJsonError(res http.ResponseWriter, code int, cause string) {
	log.Println(cause)
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(code)
	_ = json.NewEncoder(res).Encode(cause)
}

func main() {
	http.HandleFunc("/tumbnail", verifyGetOnly(sizerHandler))
	http.ListenAndServe(":"+PORT, Log(http.DefaultServeMux))
}
