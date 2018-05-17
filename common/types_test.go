package common

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestStringToActivityID(t *testing.T) {
	input := "1512189195897392628|1872552297_OMGETH"
	expectedOutput := ActivityID{1512189195897392628, "1872552297_OMGETH"}
	activityID, err := StringToActivityID(input)
	if err != nil {
		t.Fatalf("Expected convert successfully but got error: %v", err)
	} else {
		if activityID != expectedOutput {
			t.Fatalf("Expected %v, got %v", expectedOutput, activityID)
		}
	}
}

func TestStringToActivityIDError(t *testing.T) {
	input := "1512189195897392628_1872552297_OMGETH"
	_, err := StringToActivityID(input)
	if err == nil {
		t.Fatalf("Expected to return error")
	}
}

func TestActivityIDStringable(t *testing.T) {
	id := ActivityID{1512189195897392628, "1872552297_OMGETH"}
	expectedOutput := "1512189195897392628|1872552297_OMGETH"
	output := fmt.Sprintf("%s", id)
	if output != expectedOutput {
		t.Fatalf("Expected %s, got %s", expectedOutput, output)
	}
}

func TestActivityIDToJSON(t *testing.T) {
	id := ActivityID{1512189195897392628, "1872552297_OMGETH"}
	expectedOutput := `"1512189195897392628|1872552297_OMGETH"`
	b, err := json.Marshal(id)
	output := string(b)
	if err != nil {
		t.Fatalf("Expected convert successfully but got error: %v", err)
	} else {
		if output != expectedOutput {
			t.Fatalf("Expected %v, got %v", expectedOutput, output)
		}
	}
}

func TestJSONToActivityID(t *testing.T) {
	expectedOutput := ActivityID{1512189195897392628, "1872552297_OMGETH"}
	input := `"1512189195897392628|1872552297_OMGETH"`
	output := ActivityID{}
	err := json.Unmarshal([]byte(input), &output)
	if err != nil {
		t.Fatalf("Expected convert successfully but got error: %v", err)
	} else {
		if output != expectedOutput {
			t.Fatalf("Expected %v, got %v", expectedOutput, output)
		}
	}
}

type testNewExchangeInfo struct {
	exchange *ExchangeInfo
}

func (t *testNewExchangeInfo) GetInfo() (ExchangeInfo, error) {
	return *t.exchange, nil
}

func TestNewExchangeInfo(t *testing.T) {

	tn := testNewExchangeInfo{exchange: NewExchangeInfo()}
	exchangeInfo1, _ := tn.GetInfo()
	t.Logf("exchange info 1: %p", &exchangeInfo1.mu)
	exchangeInfo2, _ := tn.GetInfo()
	t.Logf("exchange info 2: %p", &exchangeInfo2.mu)
	exchangeInfo1.mu.Lock()
	exchangeInfo2.mu.Lock()
	exchangeInfo1.mu.Unlock()
	exchangeInfo2.mu.Unlock()
}
