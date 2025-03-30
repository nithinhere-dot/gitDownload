package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func main() {
	//https://github.com/nithinhere-dot/E-Cleanse.git
	repo := flag.String("Repo", "", "owner/repository")

	dir, err := os.Getwd() //for the cuurent directory
	check(err)

	dest := flag.String("Dest", dir, "Destiantion")

	flag.Parse() //for reading

	UrlBuilt(*repo, *dest) //the url is built

}

// check will accpet an error and if it is not nil it will panic
func check(err error) {
	if err != nil {
		panic(err)
	}

}

// UrlBuilt accpets a string which will have the owner and repo name
// in this function we will built an url to download a zip file
func UrlBuilt(repo, dest string) {

	if repo == "" { //checking if string is empty
		fmt.Println("Sorry bud Give a repository and owner name")
	}

	u := url.URL{ //used for to construct the url from the string
		Scheme: "https",
		Host:   "github.com",
		Path:   fmt.Sprintf("%s/archive/refs/heads/main.zip", repo),
	}
	err := downloadFile(u.String(), dest)
	check(err)
	fmt.Println("Url Is:", u.String())

	err = os.Remove("Downloaded.zip")
	check(err)
}

// downloadFile will accept a url and an string
// create folder if not exist other wise the downloaded file
// with the name of repo will be stored in the folder
func downloadFile(urlstr, dest string) error {
	resp, err := http.Get(urlstr)
	check(err)
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file:%s", resp.Status)
	}

	err = os.MkdirAll(dest, os.ModeAppend) //will create the directory if not exist
	if err != nil {
		return fmt.Errorf("failed to create diectory:%v", err)
	}

	outFile, err := os.Create("Downloaded.zip") //create a file named Downloaded.zip
	if err != nil {
		return fmt.Errorf("error in dowloading file:%v", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body) //copy the contents from website into downloaded file
	if err != nil {
		return fmt.Errorf("error in zip file:%v", err)
	}

	zipReader, err := zip.OpenReader("Downloaded.zip") //opening the zip file
	if err != nil {
		return fmt.Errorf("error opening ZIP file:%v", err)
	}
	defer zipReader.Close()

	for _, file := range zipReader.File { //iterates through extracted file
		extractedFilePath := filepath.Join(dest, file.Name)

		if file.FileInfo().IsDir() { //creates if directories needed needed
			os.MkdirAll(extractedFilePath, os.ModePerm)
			continue
		}
		//open file inside ZIP
		srcFile, _ := file.Open()
		defer srcFile.Close()

		//create file inside diectory
		dstFile, _ := os.Create(extractedFilePath)
		defer dstFile.Close()

		//copy all contents from zip to destiantion
		io.Copy(dstFile, srcFile)
	}
	return nil
}
