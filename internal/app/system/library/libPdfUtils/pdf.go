package libPdfUtils

import (
	"bytes"
	"fmt"
	"github.com/gen2brain/go-fitz"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/ledongthuc/pdf"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func ParsePdf(srcFilePath string) (string, string, error) {
	doc, err := fitz.New(srcFilePath)
	if err != nil {
		return "", "", err
	}

	defer doc.Close()

	tmpDir, err := ioutil.TempDir("/tmp", "fitz")
	if err != nil {
		return "", "", err
	}

	//// Extract pages as images
	//for n := 0; n < doc.NumPage(); n++ {
	//	img, err := doc.Image(n)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.jpg", n)))
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	err = jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	f.Close()
	//}

	//Extract pages as text
	txtContent := ""
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			panic(err)
		}

		//f, err := os.Create(filepath.Join(tmpDir, fmt.Sprintf("test%03d.txt", n)))
		//if err != nil {
		//	panic(err)
		//}
		//
		//_, err = f.WriteString(text)
		//if err != nil {
		//	panic(err)
		//}
		//
		//f.Close()
		txtContent += text
	}

	// Extract pages as html
	fileName := gfile.Name(srcFilePath)
	htmlPath := filepath.Join(tmpDir, fmt.Sprintf("%s.html", fileName))
	f, err := os.Create(htmlPath)
	if err != nil {
		return "", "", err

	}

	//f.WriteString("<!DOCTYPE html>\n<html>\n<head>\n<style>\nbody{background-color:slategray}\ndiv{position:relative;background-color:white;margin:1em auto;box-shadow:1px 1px 8px -2px black}\np{position:absolute;white-space:pre;margin:0}\n</style>\n</head>\n<body>")
	for n := 0; n < 1; n++ {
		html, err := doc.HTML(n, false)
		if err != nil {
			return "", "", err
		}

		_, err = f.WriteString(html)
		if err != nil {
			return "", "", err
		}

	}
	//f.WriteString("</body>\n</html>")
	defer f.Close()

	return htmlPath, txtContent, nil
}

func ReadPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	// remember close file
	defer f.Close()
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

func ConvertWordToPdf(serverFile, workDir string) (string, error) {
	cmd := exec.Command("soffice", "--headless", "--invisible", "--convert-to", "pdf", serverFile, "--outdir", workDir)
	data, err := cmd.Output()
	if err != nil {
		fmt.Println("convert failed: ", err)
		return "", err
	}

	return string(data), nil
}

func ReadContent(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ConvertPdfToWord(serverFile, workDir string) (string, error) {
	//libreoffice soffice
	cmd := exec.Command("soffice", "--headless", "--invisible", "--infilter=writer_pdf_import", "--convert-to", "docx", serverFile, "--outdir", workDir)
	data, err := cmd.Output()
	if err != nil {
		fmt.Println("convert failed: ", err)
		return "", err
	}

	return string(data), nil
}
