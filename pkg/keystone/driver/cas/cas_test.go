// Copyright 2019 Yunion
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

package cas

import (
	"encoding/xml"
	"testing"
)

func TestXmlUnmarshal(t *testing.T) {
	xmlstr := `<cas:serviceResponse xmlns:cas='http://www.yale.edu/tp/cas'>
    <cas:authenticationSuccess>
        <cas:user>casuser</cas:user>
    </cas:authenticationSuccess>
</cas:serviceResponse>
<cas:serviceResponse xmlns:cas='http://www.yale.edu/tp/cas'>
    <cas:authenticationSuccess>
        <cas:user>casuser</cas:user>
        <cas:attributes>
            <cas:credentialType>UsernamePasswordCredential</cas:credentialType>
            <cas:isFromNewLogin>false</cas:isFromNewLogin>
            <cas:authenticationDate>2019-09-05T12:40:08.014Z[UTC]</cas:authenticationDate>
            <cas:authenticationMethod>AcceptUsersAuthenticationHandler</cas:authenticationMethod>
            <cas:successfulAuthenticationHandlers>AcceptUsersAuthenticationHandler</cas:successfulAuthenticationHandlers>
            <cas:longTermAuthenticationRequestTokenUsed>false</cas:longTermAuthenticationRequestTokenUsed>
            </cas:attributes>
    </cas:authenticationSuccess>
</cas:serviceResponse>`
	casresp := SCASServiceResponse{}
	err := xml.Unmarshal([]byte(xmlstr), &casresp)
	if err != nil {
		t.Errorf("fail to unmarshal %s", err)
	} else {
		t.Logf("%#v", casresp)
	}
}

func TestFetchAttribute(t *testing.T) {
	cases := []struct {
		Xml  string
		Key  string
		Want string
	}{
		{
			Xml: `<cas:serviceResponse xmlns:cas='http://www.yale.edu/tp/cas'>
    <cas:authenticationSuccess>
        <cas:user>casuser</cas:user>
        <cas:proj>casproj</cas:proj>
    </cas:authenticationSuccess>
</cas:serviceResponse>`,
			Key:  "cas:proj",
			Want: "casproj",
		},
		{
			Xml: `<?xml version="1.0" encoding="UTF-8"?>
			<cas:serviceResponse xmlns:cas="http://www.yale.edu/tp/cas"><cas:authenticationSuccess><cas:user>lcftest0416</cas:user><cas:proj>周凌测试无线公司1112342</cas:proj></cas:authenticationSuccess></cas:serviceResponse>`,
			Key:  "cas:proj",
			Want: "周凌测试无线公司1112342",
		},
	}
	for _, c := range cases {
		got := fetchAttribute([]byte(c.Xml), c.Key)
		if got != c.Want {
			t.Errorf("want %s got %s", c.Want, got)
		}
	}
}
