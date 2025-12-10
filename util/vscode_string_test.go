package util_test

import (
	"taoey/obsidian-vscode-snippet/util"
	"testing"
)

var NoNeedConvertSpecialCharForTest = []string{
	"TM_SELECTED_TEXT",
	"TM_CURRENT_LINE",
	"TM_CURRENT_WORD",
	"TM_LINE_INDEX",
	"TM_LINE_NUMBER",
	"TM_FILENAME",
	"TM_FILENAME_BASE",
	"TM_DIRECTORY",
	"TM_FILEPATH",
	"RELATIVE_FILEPATH",
	"CLIPBOARD",
	"WORKSPACE_NAME",
	"WORKSPACE_FOLDER",
	"CURSOR_INDEX",
	"CURSOR_NUMBER",
	"CURRENT_YEAR",
	"CURRENT_YEAR_SHORT",
	"CURRENT_MONTH",
	"CURRENT_MONTH_NAME",
	"CURRENT_MONTH_NAME_SHORT",
	"CURRENT_DATE",
	"CURRENT_DAY_NAME",
	"CURRENT_DAY_NAME_SHORT",
	"CURRENT_HOUR",
	"CURRENT_MINUTE",
	"CURRENT_SECOND",
	"CURRENT_SECONDS_UNIX",
	"CURRENT_TIMEZONE_OFFSET",
}

// go test -v -timeout 30s -run ^TestEscapeSpecialChars$ taoey/obsidian-vscode-snippet/util
func TestEscapeSpecialChars(t *testing.T) {
	type args struct {
		code                     string
		noNeedConvertSpecialChar []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				code:                     "echo $HOME ${HOME1}",
				noNeedConvertSpecialChar: NoNeedConvertSpecialCharForTest,
			},
			want: "echo \\$HOME \\${HOME1}",
		},
		{
			name: "test time-1",
			args: args{
				code:                     "@create_time: ${CURRENT_YEAR}-${CURRENT_MONTH}-${CURRENT_DATE} ${CURRENT_HOUR}:${CURRENT_MINUTE}",
				noNeedConvertSpecialChar: NoNeedConvertSpecialCharForTest,
			},
			want: "@create_time: ${CURRENT_YEAR}-${CURRENT_MONTH}-${CURRENT_DATE} ${CURRENT_HOUR}:${CURRENT_MINUTE}",
		},
		{
			name: "test time-2",
			args: args{
				code:                     "@create_time: $CURRENT_YEAR-$CURRENT_MONTH-$CURRENT_DATE $CURRENT_HOUR:$CURRENT_MINUTE",
				noNeedConvertSpecialChar: NoNeedConvertSpecialCharForTest,
			},
			want: "@create_time: $CURRENT_YEAR-$CURRENT_MONTH-$CURRENT_DATE $CURRENT_HOUR:$CURRENT_MINUTE",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.EscapeSpecialChars(tt.args.code, tt.args.noNeedConvertSpecialChar); got != tt.want {
				t.Errorf("EscapeSpecialChars() = %v, want %v", got, tt.want)
			} else {
				t.Logf("test_case=[%v], pass", tt.name)
			}
		})
	}
}
