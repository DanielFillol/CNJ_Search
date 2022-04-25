package Test

import (
	"github.com/Darklabel91/LegalDoc_Classifier/CSV"
	"github.com/Darklabel91/LegalDoc_Classifier/Classifier"
	"reflect"
	"testing"
)

func TestLegalDocument(t *testing.T) {
	type args struct {
		docName string
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
			got, err := Classifier.LegalDocument(tt.args.docName)
			if (err != nil) != tt.wantErr {
				t.Errorf("LegalDocument() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LegalDocument() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLegalDocumentCSV(t *testing.T) {
	type args struct {
		rawFilePath      string
		separator        rune
		nameResultFolder string
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
			if err := CSV.LegalDocumentCSV(tt.args.rawFilePath, tt.args.separator, tt.args.nameResultFolder); (err != nil) != tt.wantErr {
				t.Errorf("LegalDocumentCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
