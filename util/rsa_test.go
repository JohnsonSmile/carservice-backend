package util_test

import (
	"carservice/infra/config"
	"carservice/infra/logger"
	"carservice/util"
	"encoding/base64"
	"testing"
)

func TestMain(m *testing.M) {
	logger.Initialize(true)
	config.Initialize(true)
	m.Run()
}

// 太长了，这里就不使用rsa了

func TestRsa(t *testing.T) {
	// car:highway:(base64(rsa(highway:cityid:regionid:placeid)))
	// car:charge:(base64(rsa(charge:cityid:regionid:placeid)))
	// car:park:(base64(rsa(park:cityid:regionid:placeid)))
	text := "highway:start:1:1:1001"
	bcryptText, err := util.RsaEncrypt([]byte(text))
	if err != nil {
		t.Errorf("rsa encrypt failed: %s\n", err)
		return
	}
	base64EncodedText := base64.RawStdEncoding.EncodeToString(bcryptText)
	t.Logf("base64 encoded text: %s\n", base64EncodedText)

	cypherText, err := base64.RawStdEncoding.DecodeString(base64EncodedText)
	if err != nil {
		t.Errorf("base64 decode failed: %s\n", err)
		return
	}

	originalText, err := util.RsaDecrypt(cypherText)
	if err != nil {
		t.Errorf("rsa decrypt failed: %s\n", err)
		return
	}
	t.Logf("original text: %s\n", originalText)

}

func TestRsaBase64(t *testing.T) {
	text := "highway:start:1:1:1001"
	cypherText, err := util.RsaEncryptBase64(text)
	if err != nil {
		t.Errorf("rsa encrypt failed: %s\n", err)
		return
	}
	t.Logf("cypher text: %s\n", cypherText)
	originalText, err := util.RsaDecryptBase64(cypherText)
	if err != nil {
		t.Errorf("rsa decrypt failed: %s\n", err)
		return
	}
	t.Logf("original text: %s\n", originalText)
}

/*
=== RUN   TestGenerateQRCodeData
=== RUN   TestGenerateQRCodeData/car:highway:start
    rsa_test.go:118: targetText: car:highway:SWFRgbG8cSmB0hRkuzCoX5axVZh2nwzJfQVpcIH7vEY3caCQJdhWtzB5yh92yihZKPbkHp7tdcVPkN1u+VRoVX+LzS9ew4su10hm678w0iIGs13YgCGBym3dkrkx4Mv374ut5DnWojYEn/qelAxcobywS9C/gkMm2VaSdT0MK9ztB76EC1jI6QxsKQuCninSK6aEsjqScUu/dAaCF8c78w8+a9SeeghZQkZGDOAASTT+ZCy+TpbBT3tkft8+fG5qiQ2L/un1bLOeec+Ys/TKbTZSzGJFvxsHv7xyBiWNC5P0rNTt1bAIx3Px3OtCVdVAcyslNu3gnR0IWzWgT2jNgA
=== RUN   TestGenerateQRCodeData/car:highway:end
    rsa_test.go:118: targetText: car:highway:hVRuFb4Cin0+X9eauNzcb7s9Sq1hUdl2Mt/Mhz+vXJGIXzOdwBLPGiSA5+ZqKIJ6EeSsj4bFzdtAmumKu6nQyRpheysHWRCUbPbM3UKhx+OIFWVo5PGkE4rmov0vl6kwthM/IR/zH1LNjPi+mhqsk9T1GW/vqcSLtuYyQ+B3g10NhycpePYkdPS9wakRGsPEqHkJ+SWYTK2ndtppj7BgFrOfuDuvicGPd/5UIZly0hI0WGk/XUkLl2c9CADSzhT4AVR4dFZYB7+Q14nXNte7TOFUi5gnGj1TUtnoaXK9ZzX6Su68BC/TocMnTyATKXPqmVT5A9aMbWfDl5xJDmNAUQ
=== RUN   TestGenerateQRCodeData/car:charge:start
    rsa_test.go:118: targetText: car:charge:yM8B7JeXpgRGsF01vPsBIapAyXrRh4SrSRMNtmBy9R+0C3xC+3OmuvomtIFICTJBJm9H/VWVy8gO7UZztGvwPv/QBHtIYR44OpFulBlhdW4BQZY7QRVX/7239rx6D8NCHl2N5beA81FEf8geKNmVYdXan0LlKZGEDUJXiV+3WikcK3u0Ni41oeKn+aIED4+zAIJO9/ipBKyRX53k+gpnRaDmUKJdBczNnnHGtQWWu5zx5OJ4QC3esRPLVXCrrdPpJ0K9YDVrsETrhVKk+fUQQjpI5ZguLAbmgZDWLUYR+6ftxv8EmEqFAf/tOYmGFLJ4KD87Zp0//VTFEK8pxvNJhg
=== RUN   TestGenerateQRCodeData/car:charge:end
    rsa_test.go:118: targetText: car:charge:bOH5911MAEsOTW/sKfeivSg37WEwVKX3nalzhIDQDHE2YpN7rvMZbkofiJU7eQ/Hl3lvpbM5edzmg4Wqq8abNBoeTtYE1EOSLxMz4+FHQLdjJ8Pfb26S/ezMtvOFj7GEJ/qIZarASDqqCVM80LYhTWcRMd2UfPdy//TTmQw8pt0hR/K279ubG54qkTpFuTa/SHMRrtn82lNPlpLiy55O/o2RkYZoJUsHEmKjZ/lygXjWNw0Xs4UEgT9qCNht8bjbdYrO0ltzucbd8Wpanq857mbkmt2d1gVojiC+f9ev/H59oUEqWfyfbw9zV5Xx2pVNFyo7Rn2fL9Hy4pJc9RZkaw
=== RUN   TestGenerateQRCodeData/car:park:start
    rsa_test.go:118: targetText: car:park:LNw14PW3N1T5sadu42aDw9E51yvtMNOF8dFCkGikRBorAtU/ne9h/91Fca83MVwzWpgiGkcBoTFRxj7iLOVjZa12QmnRht54UCBDUVWuy/uChCldOTnY7oRmtDIh7BeOy0WZYY4Y4Fti2RQuI0k9HyfkgQ6XBuggwRCgTtIYGm5Swygs3kej2hNrTKVH8TvK4H6yrxiCpJd115TnOMD3jsSXtQC1O4Nzbfmh6V/49hgMTw1v6ClAf8G3Au7osBPSyUQaOdSYVoyAKwYAUUU/SDrS/Q8FOQZjfVv9VVkkO9Z1KYxWAX9F1Bz5HOyqYWv2AKiWwGH/wkkiHAhKpU3Kbg
=== RUN   TestGenerateQRCodeData/car:park:end
    rsa_test.go:118: targetText: car:park:XHURYx4MEZa+STXLJJAMaHld/9aEIuH/X0qof3wUTLBJEAVC6xrCD/YRV7H4xwB0jVCqiMC3DAhAjjfTAFnXI9Y0yGQ0tC7mx/sCBz6xC3g/8Nv8SI/AR8udFH81gRboIWOy7WWm/+7ON3FlVms4M/kSIVOL1eDD8BNQwd3skv/HXthEZs2VhjHtzUwMZI6uUgo1HmS9vHVICJtuJM77PsjFUXoH9r16xN4a+TMDWofpOBl1hCidEkPTiIULV1OE1WPWOs17gOmFFvAoyySMmDaeu0cABKW31xvnK/jVfnCzw5iTiC9SJvd2spOrkYdhUdE6pvqAKCquH4+tJLaJbQ

*/

func TestGenerateQRCodeData(t *testing.T) {
	// car:highway:(base64(rsa(highway:start:cityid:regionid:placeid)))
	// car:charge:(base64(rsa(charge:start:cityid:regionid:placeid)))
	// car:park:(base64(rsa(park:start:cityid:regionid:placeid)))

	tests := []struct {
		name     string
		text     string
		codeType string
	}{
		{
			name:     "car:highway:start",
			text:     "car:highway:start:1:1:1001",
			codeType: "car:highway",
		},
		{
			name:     "car:highway:end",
			text:     "car:highway:end:1:1:1011",
			codeType: "car:highway",
		},
		{
			name:     "car:charge:start",
			text:     "car:charge:start:1:1:11",
			codeType: "car:charge",
		},
		{
			name:     "car:charge:end",
			text:     "car:charge:end:1:1:11",
			codeType: "car:charge",
		},
		{
			name:     "car:park:start",
			text:     "car:park:start:1:1:21",
			codeType: "car:park",
		},
		{
			name:     "car:park:end",
			text:     "car:park:end:1:1:21",
			codeType: "car:park",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bcryptText, err := util.RsaEncryptBase64(tt.text)
			if err != nil {
				t.Errorf("rsa encrypt failed: %s\n", err)
				return
			}
			targetText := tt.codeType + ":" + bcryptText
			t.Logf("targetText: %s\n", targetText)
		})
	}

}

/*
=== RUN   TestGenerateQRCodeText
=== RUN   TestGenerateQRCodeText/car:highway:start
    rsa_test.go:194: targetText: car:highway:start:Y2FyOmhpZ2h3YXk6c3RhcnQ6MToxOjEwMDE
=== RUN   TestGenerateQRCodeText/car:highway:end
    rsa_test.go:194: targetText: car:highway:end:Y2FyOmhpZ2h3YXk6ZW5kOjI6MjoxMDEx
=== RUN   TestGenerateQRCodeText/car:charge:start
    rsa_test.go:194: targetText: car:charge:start:Y2FyOmNoYXJnZTpzdGFydDoxOjE6MTE
=== RUN   TestGenerateQRCodeText/car:charge:end
    rsa_test.go:194: targetText: car:charge:end:Y2FyOmNoYXJnZTplbmQ6MToxOjEx
=== RUN   TestGenerateQRCodeText/car:park:start
    rsa_test.go:194: targetText: car:park:start:Y2FyOnBhcms6c3RhcnQ6MToxOjIx
=== RUN   TestGenerateQRCodeText/car:park:end
    rsa_test.go:194: targetText: car:park:end:Y2FyOnBhcms6ZW5kOjE6MToyMQ
*/

func TestGenerateQRCodeText(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		codeType string
	}{
		{
			name:     "car:highway:start",
			text:     "car:highway:start:1:1:1001",
			codeType: "car:highway",
		},
		{
			name:     "car:highway:end",
			text:     "car:highway:end:2:2:1011",
			codeType: "car:highway",
		},
		{
			name:     "car:charge:start",
			text:     "car:charge:start:1:1:11",
			codeType: "car:charge",
		},
		{
			name:     "car:charge:end",
			text:     "car:charge:end:1:1:11",
			codeType: "car:charge",
		},
		{
			name:     "car:park:start",
			text:     "car:park:start:1:1:21",
			codeType: "car:park",
		},
		{
			name:     "car:park:end",
			text:     "car:park:end:1:1:21",
			codeType: "car:park",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bcryptText := base64.RawStdEncoding.EncodeToString([]byte(tt.text))
			targetText := tt.name + ":" + bcryptText
			t.Logf("targetText: %s\n", targetText)
		})
	}
}

func TestDemo(t *testing.T) {
	s, err := util.RsaDecryptBase64("Y2FyOmhpZ2h3YXk6c3RhcnQ6MToxOjEwMDE")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(s)
}
