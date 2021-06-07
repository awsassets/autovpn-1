package autovpn

import "testing"

var config1 = `Remote Address : 12.13.185.14:80 
Library check >> OpenVPN core 3.git:master linux x86_64 64-bit built on Nov 29 2018 15:36:22
Library check >> Copyright (C) 2012-2017 OpenVPN Inc. All rights reserved.
# ==============================================================================
# Copyright (c) 2016-2020 Technologies 
# Email: contact@ale.com
#
# The MIT License (MIT)
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR # OTHERWISE, ARISING
# FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
# IN THE SOFTWARE.
# ==============================================================================

client
dev tun
proto udp

remote pt.alevpn.com 5060
remote pt.alevpn.com 4569
remote pt.alevpn.com 80
remote pt.alevpn.com 443
remote pt.alevpn.com 1194

remote-random
resolv-retry infinite
nobind
cipher AES-256-CBC
auth SHA512
comp-lzo no
verb 3

setenv CLIENT_CERT 0
tun-mtu 1500
tun-mtu-extra 32
mssfix 1450
persist-key
persist-tun
`
var config2 = `Remote Address : 1.3.85.201:80 
Library check >> OpenVPN core 3.git:master linux x86_64 64-bit built on Nov 19 2019 16:36:22
Library check >> Copyright (C) 2012-2017 OpenVPN Inc. All rights reserved.
# ==============================================================================
# Copyright (c) 2016-2020 Technologies 
# Email: contact@ale.com
#
# The MIT License (MIT)
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR # OTHERWISE, ARISING
# FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
# IN THE SOFTWARE.
# ==============================================================================

client
dev tun
proto udp

remote es.alevpn.com 5060
remote es.alevpn.com 4569
remote es.alevpn.com 80
remote es.alevpn.com 443
remote es.alevpn.com 1194

remote-random
resolv-retry infinite
nobind
cipher AES-256-CBC
auth SHA512
comp-lzo no
verb 3

setenv CLIENT_CERT 0
tun-mtu 1500
tun-mtu-extra 32
mssfix 1450
persist-key
persist-tun
`

func TestGetRemoteAddress(t *testing.T) {
	//test driven tables golang
	var configTest = []struct {
		in  string
		out string
	}{
		{config1, "pt.alevpn.com"},
		{config2, "es.alevpn.com"},
	}
	for _, tt := range configTest {
		t.Run(tt.in, func(t *testing.T) {
			addr, err := getRemoteAddress(tt.in)
			if err != nil {
				t.Error(err)
			}
			if addr != tt.out {
				t.Errorf("got %q, want %q", addr, tt.out)
			}
		})
	}
}
