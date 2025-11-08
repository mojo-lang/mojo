// Copyright 2021 Mojo-lang.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logs

type Config struct {
	Level        string                 `json:"level" default:"debug"`
	Encode       string                 `json:"encode" default:"console"`
	LevelPort    int                    `json:"levelPort" default:"0"`
	LevelPattern string                 `json:"levelPattern" default:""`
	Output       string                 `json:"output" default:"console"`
	InitFields   map[string]interface{} `json:"initFields"`
	File         FileSinkConfig         `json:"file"`
}

type FileSinkConfig struct {
	Path       string `json:"path" default:"./logs/app.log"`
	MaxSize    int    `json:"maxSize" default:"100"` // megabytes
	MaxBackups int    `json:"maxBackups" default:"10"`
	MaxAge     int    `json:"maxAge" default:"30"` // days
	Encode     string `json:"encode"  default:"json"`
	Compress   bool   `json:"compress" default:"false"`
}

func NewConfig() *Config {
	return &Config{
		Level:  "debug",
		Encode: "console",
		Output: "console",
	}
}
