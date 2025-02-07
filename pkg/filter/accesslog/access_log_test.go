/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package accesslog

import (
	"os"
	"testing"
	"time"
)

import (
	"github.com/stretchr/testify/assert"
)

import (
	"github.com/apache/dubbo-go-pixiu/pkg/common/constant"
	"github.com/apache/dubbo-go-pixiu/pkg/logger"
)

func TestAccessLog_Write_to_file(t *testing.T) {
	msg := "this is test msg"

	filePath := "dubbo-go-pixiu/logs/dubbo-go-access"
	accessLogWriter := &AccessLogWriter{AccessLogDataChan: make(chan AccessLogData, constant.LogDataBuffer)}
	accessLogWriter.Write()
	accessLogWriter.Writer(AccessLogData{AccessLogMsg: msg, AccessLogConfig: AccessLogConfig{OutPutPath: filePath}})
	time.Sleep(3e9)
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			logger.Warnf("can not create dir: %s, %v", filePath, err)
		}
	}
	assert.FileExists(t, filePath, nil)
}
