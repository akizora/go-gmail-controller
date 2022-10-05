package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
)

type SheetClient struct {
	srv           *sheets.Service
	spreadsheetID string
}

func NewSheetClient(ctx context.Context, spreadsheetID string) (*SheetClient, error) {
	fmt.Println("NewSheetClient!!!")
	b, err := ioutil.ReadFile("secret.json")
	if err != nil {
		return nil, err
	}
    // read & write permission
	jwt, err := google.JWTConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return nil, err
	}
	srv, err := sheets.New(jwt.Client(ctx))
	if err != nil {
		return nil, err
	}
	return &SheetClient{
		srv:           srv,
		spreadsheetID: spreadsheetID,
	}, nil
}

func (s *SheetClient) Get(range_ string) ([][]interface{}, error) {
	resp, err := s.srv.Spreadsheets.Values.Get(s.spreadsheetID, range_).Do()
	if err != nil {
		return nil, err
	}
	return resp.Values, nil
}

func main() {
	fmt.Println("main!!!")
	ctx := context.Background()
	client, err := NewSheetClient(ctx, os.Getenv("SPREAD_SHEET_ID"))
	if err != nil {
		panic(err)
	}
	fmt.Println(os.Getenv("SPREAD_SHEET_ID"))
	fmt.Println("データ取得")
	values, err := client.Get("'応募者リスト'!A1:C4")
	if err != nil {
		panic(err)
	}
	for _, row := range values {
		fmt.Println(row)
	}
}

func (s *SheetClient) Update(range_ string, values [][]interface{}) error {
	_, err := s.srv.Spreadsheets.Values.Update(s.spreadsheetID, range_, &sheets.ValueRange{
		Values: values,
	}).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		return err
	}
	return nil
}

func (s *SheetClient) Append(values [][]interface{}) error {
	_, err := s.srv.Spreadsheets.Values.Append(s.spreadsheetID, "シート1", &sheets.ValueRange{
		Values: values,
	}).ValueInputOption("USER_ENTERED").InsertDataOption("INSERT_ROWS").Do()
	if err != nil {
		return err
	}
	return nil
}
