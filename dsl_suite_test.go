// Copyright 2015 Peter Goetz
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

package pegomock_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/petergtz/pegomock"

	"testing"
)

func TestPegomock(t *testing.T) {
	RegisterFailHandler(Fail)
	RegisterMockFailHandler(func(message string, callerSkip ...int) { panic(message) })
	// RegisterMockFailHandler(Fail)
	RunSpecs(t, "Pegomock DSL Suite")
}
