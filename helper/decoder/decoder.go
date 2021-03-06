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

package decoder

import (
	"errors"
	"net/http"
	"strings"

	"github.com/alibaba/ilogtail/helper/decoder/influxdb"
	"github.com/alibaba/ilogtail/helper/decoder/prometheus"
	"github.com/alibaba/ilogtail/helper/decoder/sls"
	"github.com/alibaba/ilogtail/pkg/protocol"
)

// Decoder used to parse buffer to sls logs
type Decoder interface {
	// Decode reader to logs
	Decode(data []byte, req *http.Request) (logs []*protocol.Log, err error)
	ParseRequest(res http.ResponseWriter, req *http.Request, maxBodySize int64) (data []byte, statusCode int, err error)
}

var errDecoderNotFound = errors.New("no such decoder")

// GetDecoder return a new decoder for specific format
func GetDecoder(format string) (Decoder, error) {
	switch strings.TrimSpace(strings.ToLower(format)) {
	case "sls":
		return &sls.Decoder{}, nil
	case "prometheus":
		return &prometheus.Decoder{}, nil
	case "influx", "influxdb":
		return &influxdb.Decoder{}, nil
	}
	return nil, errDecoderNotFound
}
