package main

import (
	"testing"
	"time"
	"bytes"
	"io/ioutil"
)

const timeLayout = "2006-01-02 15:04:05"
const withRecordJSONPath = "./test/resources/viewing_history/record.json"
const lastUpdateJSONPath = "./test/resources/viewing_history/last_update.json"
const fileDoesNotExistJSONPath = "./test/resources/viewing_history/file_does_not_exist.json"
const dummyHTMLPath = "./test/resources/viewing_history/viewing_history.html"
const expireDurationMinute = 60

// inject time to timeNowFunc
func setNow(t time.Time) {
	timeNowFunc = func() time.Time { return t }
}

func TestViewingHistory_LoadFromHTML_Normal(t *testing.T) {
	vh := &ViewingHistory{}
	b, _ := ioutil.ReadFile(dummyHTMLPath)

	targetTime, _ := time.Parse(timeLayout, "2018-01-01 01:00:01")
	setNow(targetTime)

	err := vh.LoadFromHTML(b)

	if err != nil {
		t.Fatalf("LoadFromHTML fails. Can't parse HTML %v\n", err)
	}
}

func TestViewingHistory_LoadFromFile_Normal(t *testing.T) {
	vh := &ViewingHistory{}
	err := vh.LoadFromFile(withRecordJSONPath)

	if err != nil {
		t.Errorf("LoadFromFile fails. Raise error %v\n", err)
	}
}

func TestViewingHistory_LoadFromFile_RaiseError(t *testing.T) {
	vh := &ViewingHistory{}

	err := vh.LoadFromFile(fileDoesNotExistJSONPath)
	if err == nil {
		t.Errorf("LoadFromFile fails. Not raise error file doesn't exist")
	}
}

func TestViewingHistory_Expire_Expired(t *testing.T) {
	vh := &ViewingHistory{}

	targetTime, _ := time.Parse(timeLayout, "2018-01-01 01:00:01")

	setNow(targetTime)
	// file exists and last update expires
	if !vh.Expire(lastUpdateJSONPath, expireDurationMinute) {
		t.Errorf("Expire fails. Got false\n")
	}

	// file does not exist (this case treats as expire)
	if !vh.Expire(fileDoesNotExistJSONPath, expireDurationMinute) {
		t.Errorf("Expire fails. Got false\n")
	}
}

func TestViewingHistory_Expire_NotExpired(t *testing.T) {
	vh := &ViewingHistory{}

	targetTime, _ := time.Parse(timeLayout, "2018-01-01 00:00:01")

	setNow(targetTime)
	if vh.Expire(lastUpdateJSONPath, expireDurationMinute) {
		t.Errorf("Expire fails. Got true\n")
	}
}

func TestViewingHistory_Exist_DataExists(t *testing.T) {
	vh := &ViewingHistory{}

	if !vh.ExistData(lastUpdateJSONPath) {
		t.Errorf("ExistData fails. Got false\n")
	}
}

func TestViewingHistory_Exist_DataNotExist(t *testing.T) {
	vh := &ViewingHistory{}

	if vh.ExistData(fileDoesNotExistJSONPath) {
		t.Errorf("ExistData fails. Got true\n")
	}
}

func TestViewingHistory_Print_Simple(t *testing.T) {
	targetTime, _ := time.Parse(timeLayout, "2018-01-01 00:00:01")
	setNow(targetTime)

	vh :=&ViewingHistory{
		Records: []ViewingRecord {
			{Date: targetTime, Title: "title1", VideoURL: "url1"},
			{Date: targetTime, Title: "title2", VideoURL: "url2"},
		},
	}

	expected := bytes.NewBufferString("2018/01/01\ttitle1\turl1\n2018/01/01\ttitle2\turl2\n")
	actual := new(bytes.Buffer)

	vh.Print(3, "simple", actual)

	if bytes.Compare(expected.Bytes(), actual.Bytes()) != 0 {
		t.Errorf(
			"Print (format:simple, case:limit > record) fails. expected=%v, actual=%v",
			expected,
			actual)
	}

	actual.Reset()
	vh.Print(2, "simple", actual)

	if bytes.Compare(expected.Bytes(), actual.Bytes()) != 0 {
		t.Errorf(
			"Print (format:simple, case:limit = record) fails. expected=%v, actual=%v",
			expected,
			actual)
	}

	expected = bytes.NewBufferString("2018/01/01\ttitle1\turl1\n")
	actual.Reset()

	vh.Print(1, "simple", actual)

	if bytes.Compare(expected.Bytes(), actual.Bytes()) != 0 {
		t.Errorf(
			"Print (format:simple, case:limit < record) fails. expected=%v, actual=%v",
			expected,
			actual)
	}
}

func TestViewingHistory_Print_CSV(t *testing.T) {
	targetTime, _ := time.Parse(timeLayout, "2018-01-01 00:00:01")
	setNow(targetTime)

	vh :=&ViewingHistory{
		Records: []ViewingRecord {
			{Date: targetTime, Title: "title1", VideoURL: "url1"},
			{Date: targetTime, Title: "title2", VideoURL: "url2"},
		},
	}

	expected := bytes.NewBufferString(
		"view_date,video_title,video_url\r\n2018/01/01,title1,url1\r\n2018/01/01,title2,url2\r\n")

	actual := new(bytes.Buffer)

	vh.Print(3, "csv", actual)

	if bytes.Compare(expected.Bytes(), actual.Bytes()) != 0 {
		t.Errorf(
			"Print (format:csv, case:limit > record) fails. expected=%v, actual=%v",
			expected,
			actual)
	}

	actual.Reset()
	vh.Print(2, "csv", actual)

	if bytes.Compare(expected.Bytes(), actual.Bytes()) != 0 {
		t.Errorf(
			"Print (format:csv, case:limit = record) fails. expected=%v, actual=%v",
			expected,
			actual)
	}

	expected = bytes.NewBufferString("view_date,video_title,video_url\r\n2018/01/01,title1,url1\r\n")
	actual.Reset()

	vh.Print(1, "csv", actual)

	if bytes.Compare(expected.Bytes(), actual.Bytes()) != 0 {
		t.Errorf(
			"Print (format:csv, case:limit < record) fails. expected=%v, actual=%v",
			expected,
			actual)
	}
}
func TestViewingHistory_Print_Table(t *testing.T) {
	targetTime, _ := time.Parse(timeLayout, "2018-01-01 00:00:01")
	setNow(targetTime)

	vh :=&ViewingHistory{
		Records: []ViewingRecord {
			{Date: targetTime, Title: "title1", VideoURL: "url1"},
			{Date: targetTime, Title: "title2", VideoURL: "url2"},
		},
	}

	expected := bytes.NewBufferString(
		"+------------+-------------+-----------+\n" +
		   "| VIEW DATE  | VIDEO TITLE | VIDEO URL |\n" +
		   "+------------+-------------+-----------+\n" +
		   "| 2018/01/01 | title1      | url1      |\n" +
		   "| 2018/01/01 | title2      | url2      |\n" +
		   "+------------+-------------+-----------+\n")

	actual := new(bytes.Buffer)

	vh.Print(3, "table", actual)

	if bytes.Compare(expected.Bytes(), actual.Bytes()) != 0 {
		t.Errorf(
			"Print (format:table, case:limit > record) fails. expected:\n%v\nactual:\n%v",
			expected,
			actual)
	}

	actual.Reset()
	vh.Print(2, "table", actual)

	if bytes.Compare(expected.Bytes(), actual.Bytes()) != 0 {
		t.Errorf(
			"Print (format:table, case:limit = record) fails. expected:\n%v\nactual:\n%v",
			expected,
			actual)
	}

	expected = bytes.NewBufferString(
		"+------------+-------------+-----------+\n" +
			"| VIEW DATE  | VIDEO TITLE | VIDEO URL |\n" +
			"+------------+-------------+-----------+\n" +
			"| 2018/01/01 | title1      | url1      |\n" +
			"+------------+-------------+-----------+\n")
	actual.Reset()

	vh.Print(1, "table", actual)

	if bytes.Compare(expected.Bytes(), actual.Bytes()) != 0 {
		t.Errorf(
			"Print (format:table, case:limit < record) fails. expected:\n%v\nactual:\n%v",
			expected,
			actual)
	}
}