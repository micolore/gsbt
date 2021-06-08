package main

import (
	"fmt"
	"github.com/dcu/pdf"
)

func main() {
	content, err := ReadPdf("/Users/kubrick/Documents/org/moppo-kubrick-doc/030-project/lopdeals/斗篷/归档/инструкции.pdf") // Read local pdf file
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	return
}

// ReadPdf  to txt
func ReadPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return "", err
	}
	totalPage := r.NumPage()

	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}

		rows, _ := p.GetTextByRow()
		for _, row := range rows {
			println(">>>> row: ", row.Position)
			for _, word := range row.Content {
				fmt.Print(word.S)
			}
			fmt.Println()
		}
	}
	return "", nil
}
