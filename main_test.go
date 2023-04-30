package main

import "testing"

func TestToFixedUrl(t *testing.T) {
	fixedUrl := toFixedUrl("/aboutus.html", "http://Deepti.com/")
	if fixedUrl != "http://Deepti.com/aboutus.html" {
		t.Error("toFixedUrl did not get expected href")
	}

	mailToUrl := toFixedUrl("mailto:deeptimayeemaharana06@gmail.com", "http://Deepti.com/")
	if mailToUrl != "http://Deepti.com/" {
		t.Error("expected baseUrl instead of mailto link")
	}

	telephoneUrl := toFixedUrl("tel://1234567890", "http://Deepti.com/")
	if telephoneUrl != "http://Deepti.com/" {
		t.Error("expected baseUrl instead of telephone link")
	}
}
