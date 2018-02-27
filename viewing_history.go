package main

import (
	"time"
	"github.com/PuerkitoBio/goquery"
	"fmt"
	"io"
	"os"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"github.com/olekukonko/tablewriter"
	"encoding/csv"
)

var timeNowFunc = time.Now
var dateLayout = "2006/01/02"

type ViewingHistory struct {
	Records    []ViewingRecord
	LastUpdate time.Time
}

type ViewingRecord struct {
	Date time.Time
	Title string
}

func (vh *ViewingHistory) LoadFromHTML(html []byte) (error) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))

	if err != nil {
		return fmt.Errorf("failed to parse HTML:%v", err)
	}

	var record ViewingRecord

	doc.Find(viewingHistoryListSelector).Each(func(i int, s *goquery.Selection) {
		date, _ := time.Parse(dateLayout, s.Find(viewingHistoryDateSelector).Text())

		record.Date = date
		record.Title = s.Find(viewingHistoryTitleSelector).Text()

		vh.Records = append(vh.Records, record)
	})

	vh.LastUpdate = timeNowFunc()

	return nil
}

func (vh *ViewingHistory) LoadFromFile(path string) (error) {
	if !vh.ExistData(path) {
		return fmt.Errorf("file does not exists")
	}

	jsonBytes, err := ioutil.ReadFile(path)

	if err != nil {
		return fmt.Errorf("failed to open file:%v", err)
	}

	if err := json.Unmarshal(jsonBytes, vh); err != nil {
		return fmt.Errorf("failed to parse JSON:%v", err)
	}

	return nil
}

func (vh *ViewingHistory) Expire(path string, min int) (bool) {
	if err := vh.LoadFromFile(path); err != nil {
		return true
	}

	now := timeNowFunc()
	last := vh.LastUpdate

	if now.Before(last.Add(time.Duration(min) * time.Minute)) {
		return false
	}

	return true
}

func (vh *ViewingHistory) ExistData(path string) (bool) {
	if _, err := os.Stat(path); err != nil {
		return false
	}

	return true
}

func (vh *ViewingHistory) SaveData(path string) (error) {
	bdata, err := json.Marshal(vh)
	if err != nil {
		return fmt.Errorf("failed to marshal json:%v", err)
	}

	json := []byte(bdata)
	if err := ioutil.WriteFile(path, json, os.ModePerm); err != nil {
		return fmt.Errorf("failed to write file:%v", err)
	}

	return nil
}

func (vh *ViewingHistory) Print(limit int, fromat string, w io.Writer) {
	switch fromat {
	case "csv": vh.CsvPrint(limit, w);
	case "table": vh.TablePrint(limit, w);
	default: vh.SimplePrint(limit, w)
	}
}

func (vh *ViewingHistory) SimplePrint(limit int, w io.Writer) {
	for i, r := range vh.Records {
		if i > limit - 1 {
			break
		}

		fmt.Fprintf(w, "%s\t%s\n", r.Date.Format(dateLayout), r.Title)
	}
}

func (vh *ViewingHistory) CsvPrint(limit int, w io.Writer) {
	writer := csv.NewWriter(w)
	writer.UseCRLF = true
	writer.Write([]string {"view_date", "video_title"})

	for i, r := range vh.Records {
		if i > limit - 1 {
			break
		}

		writer.Write([]string {r.Date.Format(dateLayout), r.Title})
	}

	writer.Flush()
}

func (vh *ViewingHistory) TablePrint(limit int, w io.Writer) {
	writer := tablewriter.NewWriter(w)
	writer.SetHeader([]string{"view_date", "video_title"})

	for i, r := range vh.Records {
		if i > limit - 1 {
			break
		}

		writer.Append([]string {r.Date.Format(dateLayout), r.Title})
	}

	writer.Render()
}
