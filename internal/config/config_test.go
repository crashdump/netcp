package config

import (
	"os"
	"os/user"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
	cfg *Config
}

var cfgDefaults = map[string]interface{}{
	"is_default":          "thing",
	"sub.is_default":      "foo",
	"nested.sub.sub.test": "bar",
}

type testCfgHarness struct {
	key   string
	value interface{}
}

var testCfgString1 = testCfgHarness{"string.one", "a"}
var testCfgString2 = testCfgHarness{"string.two.three", "b"}
var testCfgInt = testCfgHarness{"int.one", 1}
var testCfgBool = testCfgHarness{"true", true}
var testCfgFileYaml = `---
string:
  one: a
  two:
    three: b
int:
  one: 1
true: true
`

//var testCfgNewBool = testCfgHarness{"this_is_new", true}

func writeMockConfigFile(t *testing.T) {
	user, err := user.Current()
	if err != nil {
		t.Fatal(err)
	}

	path := user.HomeDir + "/.netcp"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0700)
	}
	f, err := os.Create(path + "/omnitestapp.testenv.yaml")
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.Write([]byte(testCfgFileYaml))
	if err != nil {
		t.Fatal(err)
	}
	f.Close()
}

func TestConfigTestSuite(t *testing.T) {
	writeMockConfigFile(t)

	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) SetupTest() {
	var err error
	ts.cfg, err = New("testapp", "testenv", cfgDefaults)
	ts.NoError(err)

	err = ts.cfg.Load()
	ts.NoError(err)
}

func (ts *TestSuite) TestConfig_GetStringFromFile() {
	tests := []struct {
		name    string
		fields  string
		want    interface{}
		wantErr bool
	}{
		{
			name:    "valid-testCfgString1",
			fields:  testCfgString1.key,
			want:    testCfgString1.value,
			wantErr: false,
		},
		{
			name:    "valid-testCfgString2",
			fields:  testCfgString2.key,
			want:    testCfgString2.value,
			wantErr: false,
		},
		{
			name:    "valid-testCfgBool",
			fields:  testCfgBool.key,
			want:    testCfgBool.value,
			wantErr: false,
		},
		{
			name:    "valid-testCfgInt",
			fields:  testCfgInt.key,
			want:    testCfgInt.value,
			wantErr: false,
		},
		{
			name:    "valid-testCfgInt",
			fields:  testCfgInt.key,
			want:    testCfgInt.value,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			var val interface{}
			ts.Run(tt.name, func() {
				switch tt.want.(type) {
				case string:
					val = ts.cfg.GetString(tt.fields)
				case int:
					val = ts.cfg.GetInt(tt.fields)
				case bool:
					val = ts.cfg.GetBool(tt.fields)
				}
			})
			if tt.wantErr {
				ts.Empty(val)
			} else {
				ts.Equal(tt.want, val)
			}
		})
	}
}

func (ts *TestSuite) TestConfig_GetStringFromDefaults() {
	tests := []struct {
		name    string
		fields  string
		want    interface{}
		wantErr bool
	}{
		{
			name:   "has-default1",
			fields: "is_default",
			want:   "thing",
		},
		{
			name:   "has-default2",
			fields: "sub.is_default",
			want:   "foo",
		},
		{
			name:   "has-default3",
			fields: "nested.sub.sub.test",
			want:   "bar",
		},
		{
			name:   "does-not-have-default",
			fields: "xyz123",
			want:   "",
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			var val interface{}
			ts.Run(tt.name, func() {
				val = ts.cfg.GetString(tt.fields)
			})
			ts.Equal(tt.want, val)
		})
	}
}

func (ts *TestSuite) TestConfig_GetString() {
	tests := []struct {
		name    string
		fields  string
		want    interface{}
		wantErr bool
	}{
		{
			name:    "valid-testCfgString1",
			fields:  testCfgString1.key,
			want:    testCfgString1.value,
			wantErr: false,
		},
		{
			name:    "valid-testCfgString2",
			fields:  testCfgString2.key,
			want:    testCfgString2.value,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			var val string
			ts.Run(tt.name, func() {
				val = ts.cfg.GetString(tt.fields)
			})
			if tt.wantErr {
				ts.Empty(val)
			} else {
				ts.Equal(tt.want, val)
			}
		})
	}
}

func (ts *TestSuite) TestConfig_GetInt() {
	tests := []struct {
		name    string
		fields  string
		want    interface{}
		wantErr bool
	}{
		{
			name:    "valid-testCfgInt",
			fields:  testCfgInt.key,
			want:    testCfgInt.value,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			var val int
			ts.Run(tt.name, func() {
				val = ts.cfg.GetInt(tt.fields)
			})
			if tt.wantErr {
				ts.Empty(val)
			} else {
				ts.Equal(tt.want, val)
			}
		})
	}
}

func (ts *TestSuite) TestConfig_GetBool() {
	tests := []struct {
		name    string
		fields  string
		want    interface{}
		wantErr bool
	}{
		{
			name:    "valid-testCfgBool",
			fields:  testCfgBool.key,
			want:    testCfgBool.value,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		ts.Run(tt.name, func() {
			var val bool
			ts.Run(tt.name, func() {
				val = ts.cfg.GetBool(tt.fields)
			})
			if tt.wantErr {
				ts.Empty(val)
			} else {
				ts.Equal(tt.want, val)
			}
		})
	}
}

//func (ts *TestSuite) TestConfig_Save() {
//	type fields struct {
//		viper viper.Viper
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		wantErr bool
//	}{
//		{
//			name: "",
//			fields: fields{
//				viper: viper.Viper{},
//			},
//			wantErr: false,
//		},
//	}
//	for _, tt := range tests {
//		ts.Run(tt.name, func() {
//			var err error
//			ts.Run(tt.name, func() {
//				err = ts.cfg.Save()
//			})
//			if tt.wantErr {
//				ts.Error(err)
//			} else {
//				ts.NoError(err)
//			}
//		})
//	}
//}