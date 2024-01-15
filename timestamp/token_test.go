// Copyright The Notary Project Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package timestamp

import (
	"fmt"
	"os"
	"testing"

	"github.com/notaryproject/notation-core-go/internal/crypto/oid"
)

func TestParseSignedToken(t *testing.T) {
	timestampToken, err := os.ReadFile("testdata/TimestampTokenWithInvalideContentType.p7s")
	if err != nil {
		t.Fatal("failed to read test timestamp token:", err)
	}
	expectedErrMsg := fmt.Sprintf("unexpected content type: %v", oid.Data)
	_, err = ParseSignedToken(timestampToken)
	if err == nil || err.Error() != expectedErrMsg {
		t.Fatalf("expected error %s, but got %v", expectedErrMsg, err)
	}
}
