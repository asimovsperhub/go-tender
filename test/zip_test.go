package test

import (
	"bytes"
	"fmt"
	"github.com/klauspost/compress/zip"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func ZipFiles(filename string, files []string) error {

	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newZipFile.Close()
	//var buff bytes.Buffer
	buff := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buff)
	defer zipWriter.Close()

	//zipWriter := zip.NewWriter(newZipFile)
	//defer zipWriter.Close()

	// Add files to zip
	for _, file := range files {
		if err = AddFileToZip(zipWriter, file); err != nil {
			return err
		}
	}
	ioutil.WriteFile("Hello1.zip", buff.Bytes(), 0777)
	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filename string) error {

	fileToZip, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer fileToZip.Close()

	// Get the file information
	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	// Using FileInfoHeader() above only uses the basename of the file. If we want
	// to preserve the folder structure we can overwrite this with the full path.
	header.Name = filename
	len_name := strings.Split(filename, "/")
	header.Name = len_name[len(len_name)-1]
	// Change to deflate to gain better compression
	// see http://golang.org/pkg/archive/zip/#pkg-constants
	header.Method = zip.Deflate
	log.Println(header.Method, header.Name)
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err
}

func Test_zip(t *testing.T) {
	// List of Files to Zip
	files := []string{"/Users/apple/Desktop/d401921d892841dd8f2e9cb41baca770.pdf"}
	output := "done.zip"
	if err := ZipFiles(output, files); err != nil {
		log.Println(err)
	}
	fmt.Println("Zipped File:", output)
}
func Test_zipData(t *testing.T) {

	// Create a buffer to write our archive to.
	fmt.Println("we are in the zipData function")
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	zipWriter := zip.NewWriter(buf)

	files := []string{"/Users/apple/Desktop/d401921d892841dd8f2e9cb41baca770.pdf", "/Users/apple/Desktop/1.png"}
	for _, file := range files {
		fileToZip, err := os.Open(file)
		defer fileToZip.Close()
		// Get the file information
		info, err := fileToZip.Stat()
		zipFile, err := zipWriter.Create(info.Name())
		if err != nil {
			fmt.Println(err)
		}
		bytes, _ := os.ReadFile(file)
		_, err = zipFile.Write(bytes)
		if err != nil {
			fmt.Println(err)
		}
	}

	// Make sure to check the error on Close.
	err := zipWriter.Close()
	if err != nil {
		fmt.Println(err)
	}

	//write the zipped file to the disk
	ioutil.WriteFile("Hello2.zip", buf.Bytes(), 0777)

}
