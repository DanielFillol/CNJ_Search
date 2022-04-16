package LegalDoc_Classifier

import (
	"github.com/Darklabel91/LegalDoc_Classifier"
	"github.com/Darklabel91/LegalDoc_Classifier/Structs"
	"reflect"
	"testing"
)

func TestDocClassifier(t *testing.T) {
	type args struct {
		identifier   string
		searchString string
		searchType   int
	}
	tests := []struct {
		name  string
		args  args
		want  Structs.FinalData
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := LegalDoc_Classifier.DocClassifier(tt.args.identifier, tt.args.searchString, tt.args.searchType)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DocClassifier() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("DocClassifier() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestDocClassifierCSV(t *testing.T) {
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
			if err := LegalDoc_Classifier.DocClassifierCSV(tt.args.rawFilePath, tt.args.separator, tt.args.nameResultFolder, tt.args.searchType); (err != nil) != tt.wantErr {
				t.Errorf("DocClassifierCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
