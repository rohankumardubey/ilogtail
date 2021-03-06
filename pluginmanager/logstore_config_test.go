// Copyright 2021 iLogtail Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pluginmanager

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/alibaba/ilogtail/pkg/logger"
	"github.com/alibaba/ilogtail/plugins/input"
	"github.com/alibaba/ilogtail/plugins/processor/regex"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestLogstroreConfig(t *testing.T) {
	suite.Run(t, new(logstoreConfigTestSuite))
}

type logstoreConfigTestSuite struct {
	suite.Suite
}

func (s *logstoreConfigTestSuite) BeforeTest(suiteName, testName string) {
	logger.Infof(context.Background(), "========== %s %s test start ========================", suiteName, testName)
	LogtailConfig = make(map[string]*LogstoreConfig)
}

func (s *logstoreConfigTestSuite) AfterTest(suiteName, testName string) {
	logger.Infof(context.Background(), "========== %s %s test end =======================", suiteName, testName)

}

func (s *logstoreConfigTestSuite) TestPluginGlobalConfig() {
	str := `{
		"global": {
			"InputIntervalMs": 9999,
			"AggregatIntervalMs": 369,
			"FlushIntervalMs": 323,
			"DefaultLogQueueSize": 21,
			"DefaultLogGroupQueueSize": 31
		},
		"inputs" : [
			{
				"type" : "metric_http",
				"detail" : {
					"Addresses" : [
						"http://config.sls.aliyun.com"
					],
					"IncludeBody" : true
				}
			},
			{
				"type" : "service_canal",
				"detail" : {
					"Host" : "localhost",
					"Password" : "111",
					"User" : "111",
					"DataBases" : ["zc"],
					"IncludeTables" : [".*\\..*"],
					"ServerID" : 12312,
					"TextToString" : true
				}
			}
		],
		"flushers" : [
			{
				"type" : "flusher_stdout",
				"detail" : {

				}
			}
		]
	}`
	s.NoError(LoadMockConfig("project", "logstore", "1", str), "load config fail")
	s.Equal(len(LogtailConfig), 1)
	s.Equal(LogtailConfig["1"].ConfigName, "1")
	config := LogtailConfig["1"]
	s.Equal(config.ProjectName, "project")
	s.Equal(config.LogstoreName, "logstore")
	s.Equal(config.LogstoreKey, int64(666))
	// InputIntervalMs          int
	// AggregatIntervalMs       int
	// FlushIntervalMs          int
	// DefaultLogQueueSize      int
	// DefaultLogGroupQueueSize int
	s.Equal(config.GlobalConfig.AggregatIntervalMs, 369)
	s.Equal(config.GlobalConfig.DefaultLogGroupQueueSize, 31)
	s.Equal(config.GlobalConfig.DefaultLogQueueSize, 21)
	s.Equal(config.GlobalConfig.InputIntervalMs, 9999)
	s.Equal(config.GlobalConfig.FlushIntervalMs, 323)
	s.Equal(config.MetricPlugins[0].Interval, time.Duration(9999)*time.Millisecond)
	s.Equal(config.AggregatorPlugins[0].Interval, time.Duration(369)*time.Millisecond)
	s.Equal(config.FlusherPlugins[0].Interval, time.Duration(323)*time.Millisecond)
}

func (s *logstoreConfigTestSuite) TestLoadConfig() {
	s.NoError(LoadMockConfig("project", "logstore", "1"))
	s.NoError(LoadMockConfig("project", "logstore", "3"))
	s.NoError(LoadMockConfig("project", "logstore", "2"))
	s.Equal(len(LogtailConfig), 3)
	s.Equal(LogtailConfig["1"].ConfigName, "1")
	s.Equal(LogtailConfig["2"].ConfigName, "2")
	s.Equal(LogtailConfig["3"].ConfigName, "3")
	for i := 0; i < 3; i++ {
		config := LogtailConfig[strconv.Itoa(i+1)]
		s.Equal(config.ProjectName, "project")
		s.Equal(config.LogstoreName, "logstore")
		s.Equal(config.LogstoreKey, int64(666))
		s.Equal(len(config.MetricPlugins), 0)
		s.Equal(len(config.ServicePlugins), 1)
		s.Equal(len(config.ProcessorPlugins), 1)
		s.Equal(len(config.AggregatorPlugins), 1)
		s.Equal(len(config.FlusherPlugins), 2)
		// global config
		s.Equal(config.GlobalConfig, &LogtailGlobalConfig)

		// check plugin inner info
		reg, ok := config.ProcessorPlugins[0].Processor.(*regex.ProcessorRegex)
		s.True(ok)
		// "SourceKey": "content",
		// "Regex": "Active connections: (\\d+)\\s+server accepts handled requests\\s+(\\d+)\\s+(\\d+)\\s+(\\d+)\\s+Reading: (\\d+) Writing: (\\d+) Waiting: (\\d+).*",
		// "Keys": [
		// 	"connection",
		// 	"accepts",
		// 	"handled",
		// 	"requests",
		// 	"reading",
		// 	"writing",
		// 	"waiting"
		// ],
		// "FullMatch": true,
		// "NoKeyError": true,
		// "NoMatchError": true,
		// "KeepSource": true
		s.True(reg.KeepSource)
		s.True(reg.NoKeyError)
		s.True(reg.NoMatchError)
		s.True(reg.KeepSource)
		s.Equal(reg.SourceKey, "content")
		s.Equal(reg.Regex, "Active connections: (\\d+)\\s+server accepts handled requests\\s+(\\d+)\\s+(\\d+)\\s+(\\d+)\\s+Reading: (\\d+) Writing: (\\d+) Waiting: (\\d+).*")
		s.Equal(reg.Keys, []string{
			"connection",
			"accepts",
			"handled",
			"requests",
			"reading",
			"writing",
			"waiting",
		})
	}
}

func Test_hasDockerStdoutInput(t *testing.T) {
	{
		plugins := map[string]interface{}{
			"inputs": []interface{}{
				map[string]interface{}{
					"type": input.ServiceDockerStdoutPluginName,
				},
			},
		}
		require.True(t, hasDockerStdoutInput(plugins))
	}

	{
		// No field inputs.
		plugins := map[string]interface{}{
			"processors": []interface{}{},
		}
		require.False(t, hasDockerStdoutInput(plugins))
	}

	{
		// Empty inputs.
		plugins := map[string]interface{}{
			"inputs": []interface{}{},
		}
		require.False(t, hasDockerStdoutInput(plugins))
	}

	{
		// Invalid inputs.
		plugins := map[string]interface{}{
			"inputs": "xxxx",
		}
		require.False(t, hasDockerStdoutInput(plugins))
	}
}
