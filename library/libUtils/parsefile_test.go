package libUtils

import (
	"archive/zip"
	"bytes"
	"code.sajari.com/docconv"
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"log"
	"testing"
)

/*
	git地址：https://github.com/sajari/docconv#dependencies
	依赖项
	go get code.sajari.com/docconv/docd@latest
	ocr:
	git地址：https://github.com/otiai10/gosseract/tree/v2.2.4
	go get -t github.com/otiai10/gosseract
	依赖项：需要安装tesseract


	go get -tags ocr code.sajari.com/docconv/...



	读pdf问题：exec: "pdftotext": executable file not found in $PATH
	服务端依赖：poppler-utils wv unrtf tidy


*/

//
func Test_P(t *testing.T) {
	res, err := docconv.ConvertPath("/Users/apple/Desktop/d401921d892841dd8f2e9cb41baca770.pdf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

//func Test_Ocr(t *testing.T) {
//	client := gosseract.NewClient()
//	defer client.Close()
//	client.SetImage("/Users/apple/Desktop/1.png")
//	text, _ := client.Text()
//	fmt.Println(text)
//}

// 写pdf
func Test_ToPdf(t *testing.T) {
	pdf := gofpdf.New("P", "mm", "A4", "/Users/apple/work/go/src/project/tenderBack/resource/ttf/")
	titleStr := "标题"
	// pdf.SetFont("dejavu", "", 14)
	pdf.AddUTF8Font("PingFang", "", "PingFang Heavy.ttf")
	pdf.SetTitle(titleStr, false)
	//pdf.SetAuthor("Jules Verne", false)
	pdf.SetHeaderFunc(func() {
		// Arial bold 15
		pdf.SetFont("PingFang", "", 15)
		// Calculate width of title and position
		wd := pdf.GetStringWidth(titleStr) + 6
		pdf.SetX((210 - wd) / 2)
		// Colors of frame, background and text
		pdf.SetDrawColor(0, 80, 180)
		pdf.SetFillColor(230, 230, 0)
		pdf.SetTextColor(220, 50, 50)
		// Thickness of frame (1 mm)
		pdf.SetLineWidth(1)
		// Title
		pdf.CellFormat(wd, 9, titleStr, "1", 1, "C", true, 0, "")
		// Line break
		pdf.Ln(10)
	})
	pdf.SetFooterFunc(func() {
		// Position at 1.5 cm from bottom
		pdf.SetY(-15)
		// Arial italic 8
		pdf.SetFont("PingFang", "", 8)
		// Text color in gray
		pdf.SetTextColor(128, 128, 128)
		// Page number
		pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()),
			"", 0, "C", false, 0, "")
	})
	chapterTitle := func(chapNum int, titleStr string) {
		// 	// Arial 12
		pdf.SetFont("PingFang", "", 12)
		// Background color
		pdf.SetFillColor(200, 220, 255)
		// Title
		pdf.CellFormat(0, 6, fmt.Sprintf("发布时间 : %s", titleStr),
			"", 1, "L", true, 0, "")
		// Line break
		pdf.Ln(4)
	}
	chapterBody := func(txtStr string) {
		// Read text file
		//txtStr, err := ioutil.ReadFile(fileStr)
		//if err != nil {
		//	pdf.SetError(err)
		//}
		// Times 12
		pdf.SetFont("PingFang", "", 12)
		// Output justified text
		pdf.MultiCell(0, 5, string(txtStr), "", "", false)
		// Line break
		pdf.Ln(-1)
		// Mention in italics
		pdf.SetFont("PingFang", "", 0)
		pdf.Cell(0, 5, "(end of excerpt)")
	}
	printChapter := func(chapNum int, titleStr, fileStr string) {
		pdf.AddPage()
		chapterTitle(chapNum, titleStr)
		chapterBody(fileStr)
	}
	printChapter(1, "发布时间", "内容")
	printChapter(2, "发布时间", "内容")
	// err := pdf.OutputFileAndClose("./test.pdf")
	var buf bytes.Buffer
	err := pdf.Output(&buf)
	//log.Println(buf.Bytes())
	if err != nil {
		log.Println("Error generating PDF: ", err)
	}
}

func Test_word(t *testing.T) {
	// var zipFile *os.File
	// defer zipFile.Close()
	var buff bytes.Buffer
	zipWriter := zip.NewWriter(&buff)
	defer zipWriter.Close()
	partFile, _ := zipWriter.Create("word/document.xml")
	partFile.Write([]byte(makeDocumentXMLFile("测试代码")))

}
func makeDocumentXMLFile(txt string) string {
	return `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<w:document xmlns:wpc="http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas" xmlns:cx="http://schemas.microsoft.com/office/drawing/2014/chartex" xmlns:cx1="http://schemas.microsoft.com/office/drawing/2015/9/8/chartex" xmlns:cx2="http://schemas.microsoft.com/office/drawing/2015/10/21/chartex" xmlns:cx3="http://schemas.microsoft.com/office/drawing/2016/5/9/chartex" xmlns:cx4="http://schemas.microsoft.com/office/drawing/2016/5/10/chartex" xmlns:cx5="http://schemas.microsoft.com/office/drawing/2016/5/11/chartex" xmlns:cx6="http://schemas.microsoft.com/office/drawing/2016/5/12/chartex" xmlns:cx7="http://schemas.microsoft.com/office/drawing/2016/5/13/chartex" xmlns:cx8="http://schemas.microsoft.com/office/drawing/2016/5/14/chartex" xmlns:mc="http://schemas.openxmlformats.org/markup-compatibility/2006" xmlns:aink="http://schemas.microsoft.com/office/drawing/2016/ink" xmlns:am3d="http://schemas.microsoft.com/office/drawing/2017/model3d" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships" xmlns:m="http://schemas.openxmlformats.org/officeDocument/2006/math" xmlns:v="urn:schemas-microsoft-com:vml" xmlns:wp14="http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing" xmlns:wp="http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing" xmlns:w10="urn:schemas-microsoft-com:office:word" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main" xmlns:w14="http://schemas.microsoft.com/office/word/2010/wordml" xmlns:w15="http://schemas.microsoft.com/office/word/2012/wordml" xmlns:w16cex="http://schemas.microsoft.com/office/word/2018/wordml/cex" xmlns:w16cid="http://schemas.microsoft.com/office/word/2016/wordml/cid" xmlns:w16="http://schemas.microsoft.com/office/word/2018/wordml" xmlns:w16se="http://schemas.microsoft.com/office/word/2015/wordml/symex" xmlns:wpg="http://schemas.microsoft.com/office/word/2010/wordprocessingGroup" xmlns:wpi="http://schemas.microsoft.com/office/word/2010/wordprocessingInk" xmlns:wne="http://schemas.microsoft.com/office/word/2006/wordml" xmlns:wps="http://schemas.microsoft.com/office/word/2010/wordprocessingShape" mc:Ignorable="w14 w15 w16se w16cid w16 w16cex wp14">
    <w:body>
        <w:p w14:paraId="396385E8" w14:textId="6BF766E2" w:rsidR="00947F74" w:rsidRDefault="006640DB">
            <w:r>
                <w:rPr>
                    <w:rFonts w:hint="eastAsia" />
                </w:rPr>
                <w:t>` + txt + `</w:t>
            </w:r>
            <w:r>
                <w:t xml:space="preserve"></w:t>
            </w:r>
        </w:p>
        <w:sectPr w:rsidR="00947F74">
            <w:pgSz w:w="11906" w:h="16838" />
            <w:pgMar w:top="1440" w:right="1800" w:bottom="1440" w:left="1800" w:header="851" w:footer="992" w:gutter="0" />
            <w:cols w:space="425" />
            <w:docGrid w:type="lines" w:linePitch="312" />
        </w:sectPr>
    </w:body>
</w:document>
	`
}
