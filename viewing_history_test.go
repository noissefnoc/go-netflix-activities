package main

import (
	"testing"
	"time"
)

const timeLayout = "2006-01-02 15:04:05"
const withRecordJSONPath = "./test/resources/viewing_history/record.json"
const lastUpdateJSONPath = "./test/resources/viewing_history/last_update.json"
const fileDoesNotExistJSONPath = "./test/resources/viewing_history/file_does_not_exist.json"
const expireDurationMinute = 60

// inject time to timeNowFunc
func setNow(t time.Time) {
	timeNowFunc = func() time.Time { return t }
}

func TestLoadFromFileNormal(t *testing.T) {
	vh := &ViewingHistory{}
	err := vh.LoadFromFile(withRecordJSONPath)

	if err != nil {
		t.Errorf("LoadFromFile fails. Raise error %v\n", err)
	}
}

func TestLoadFromFileRaiseError(t *testing.T) {
	vh := &ViewingHistory{}

	err := vh.LoadFromFile(fileDoesNotExistJSONPath)
	if err == nil {
		t.Errorf("LoadFromFile fails. Not raise error file doesn't exist")
	}
}

func TestExpireExpired(t *testing.T) {
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

func TestExpireNotExpired(t *testing.T) {
	vh := &ViewingHistory{}

	targetTime, _ := time.Parse(timeLayout, "2018-01-01 00:00:01")

	setNow(targetTime)
	if vh.Expire(lastUpdateJSONPath, expireDurationMinute) {
		t.Errorf("Expire fails. Got true\n")
	}
}

func TestExistDataExists(t *testing.T) {
	vh := &ViewingHistory{}

	if !vh.ExistData(lastUpdateJSONPath) {
		t.Errorf("ExistData fails. Got false\n")
	}
}

func TestExistDataNotExist(t *testing.T) {
	vh := &ViewingHistory{}

	if vh.ExistData(fileDoesNotExistJSONPath) {
		t.Errorf("ExistData fails. Got true\n")
	}
}
