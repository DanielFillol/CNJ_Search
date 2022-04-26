package Test

import (
	"github.com/Darklabel91/LegalDoc_Classifier/CSV"
	"github.com/Darklabel91/LegalDoc_Classifier/Classifier"
	"reflect"
	"testing"
)

func TestSearchCNJ(t *testing.T) {
	type args struct {
		searchString string
		searchType   int
	}
	tests := []struct {
		name    string
		args    args
		want    Classifier.FinalData
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Classifier.SearchCNJ(tt.args.searchString, tt.args.searchType)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchCNJ() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchCNJ() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLegalDocumentCSV(t *testing.T) {
	type args struct {
		rawFilePath      string
		separator        rune
		nameResultFolder string
		searchType       int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CSV.LegalDocumentCSV(tt.args.rawFilePath, tt.args.separator, tt.args.nameResultFolder, tt.args.searchType); (err != nil) != tt.wantErr {
				t.Errorf("LegalDocumentCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
