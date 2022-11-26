package internal

import (
	"reflect"
	"testing"
)

func Test_addDefaultLabel(t *testing.T) {
	t.Run("should add default label when not any default label in map", func(t *testing.T) {
		actualLabels := make(map[string]string)

		expectedAuthor := "author for test"
		expectedEmail := "author@test-not-real.com"

		addDefaultLabel(actualLabels, expectedAuthor, expectedEmail)

		expectedLabels := map[string]string{
			"commitAuthor": expectedAuthor,
			"commitEmail":  expectedEmail,
			"reviewAuthor": expectedAuthor,
			"reviewEmail":  expectedEmail,
			"status":       OPEN.string(),
		}

		if !reflect.DeepEqual(expectedLabels, actualLabels) {
			t.Errorf("can't add default labels into labels")
		}
	})

	t.Run("should not overwrite value of label given collection that was existed default label", func(t *testing.T) {
		type defaultLabel struct {
			labelName  string
			labelValue string
			want       *map[string]string
		}

		oneAuthor := "author for test"
		oneEmail := "author@test-not-real.com"

		anotherAuthor := "another author for test"
		anotherEmail := "another-author@test-not-real.com"

		tests := []defaultLabel{
			{
				labelName:  "commitAuthor",
				labelValue: anotherAuthor,
				want: &map[string]string{
					"commitAuthor": oneAuthor,
					"commitEmail":  oneEmail,
					"reviewAuthor": oneAuthor,
					"reviewEmail":  oneEmail,
					"status":       OPEN.string(),
				},
			},
			{
				labelName:  "commitEmail",
				labelValue: anotherEmail,
				want: &map[string]string{
					"commitAuthor": oneAuthor,
					"commitEmail":  oneEmail,
					"reviewAuthor": oneAuthor,
					"reviewEmail":  oneEmail,
					"status":       OPEN.string(),
				},
			},
			{
				labelName:  "reviewAuthor",
				labelValue: anotherAuthor,
				want: &map[string]string{
					"commitAuthor": oneAuthor,
					"commitEmail":  oneEmail,
					"reviewAuthor": anotherAuthor,
					"reviewEmail":  oneEmail,
					"status":       OPEN.string(),
				},
			},
			{
				labelName:  "reviewEmail",
				labelValue: anotherEmail,
				want: &map[string]string{
					"commitAuthor": oneAuthor,
					"commitEmail":  oneEmail,
					"reviewAuthor": oneAuthor,
					"reviewEmail":  anotherEmail,
					"status":       OPEN.string(),
				},
			},
			{
				labelName:  "status",
				labelValue: "CLOSE",
				want: &map[string]string{
					"commitAuthor": oneAuthor,
					"commitEmail":  oneEmail,
					"reviewAuthor": oneAuthor,
					"reviewEmail":  oneEmail,
					"status":       "CLOSE",
				},
			},
		}

		for _, tt := range tests {
			actualLabels := map[string]string{
				tt.labelName: tt.labelValue,
			}

			addDefaultLabel(actualLabels, oneAuthor, oneEmail)

			if !reflect.DeepEqual(*tt.want, actualLabels) {
				t.Errorf(
					"default label labelValue is wrong: %s - %s - %s",
					tt.labelName, (*tt.want)[tt.labelName], actualLabels[tt.labelName],
				)
			}
		}
	})
}
