package ultimateguitar

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
)

// API constants
const ugAPIEndpoint = "https://api.ultimate-guitar.com/api/v1"
const ugUserAgent = "UGT_ANDROID/4.11.1 (Pixel; 8.1.0)"

// Default headers
var ugHeaders = map[string]string{
	"Accept-Charset": "utf-8",
	"Accept":         "application/json",
	"User-Agent":     ugUserAgent,
	"Connection":     "close",
}

// Scraper struct
type Scraper struct {
	Client   *http.Client
	DeviceID string
	APIKey   string
}

// Generates a new device id for the scraper instances. This value is used in the request headers and to generate X-UG-API-KEY.
func (s *Scraper) generateDeviceID() {
	raw := make([]byte, 16)
	_, err := rand.Read(raw)
	if err != nil {
		log.Fatal(err)
	}
	s.DeviceID = fmt.Sprintf("%x", raw)[:16]
}

// Generate the X-UG-API-KEY for this request
func (s *Scraper) generateAPIKey() string {

	t := time.Now().UTC()

	// Because Go doesn't like to let us properly format dates...
	var yearVal string
	var monthVal string
	var dayVal string

	if t.Year() < 10 {
		yearVal = fmt.Sprintf("0%d", t.Year())
	} else {
		yearVal = fmt.Sprintf("%d", t.Year())
	}
	if t.Month() < 10 {
		monthVal = fmt.Sprintf("0%d", t.Month())
	} else {
		monthVal = fmt.Sprintf("%d", t.Month())
	}
	if t.Day() < 10 {
		dayVal = fmt.Sprintf("0%d", t.Day())
	} else {
		dayVal = fmt.Sprintf("%d", t.Day())
	}

	formattedDate := fmt.Sprintf("%s-%s-%s:%d", yearVal, monthVal, dayVal, t.Hour())
	payload := fmt.Sprintf("%s%s%s", s.DeviceID, formattedDate, "createLog()")
	rawAPI := []byte(payload)
	hashed := md5.Sum(rawAPI)
	return fmt.Sprintf("%x", hashed)
}

// SetProxy - Set a proxy for this scraper instance. Call again with SetProxy("") to remove.
func (s *Scraper) SetProxy(proxy string) {
	if len(proxy) > 1 {
		proxyStr := proxy
		proxyURL, _ := url.Parse(proxyStr)
		transport := &http.Transport{
			Proxy:           http.ProxyURL(proxyURL),
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		s.Client.Transport = transport
	} else {
		s.Client.Transport = &http.Transport{}
	}
}

// New Scraper instance
func New() Scraper {
	s := Scraper{
		Client: &http.Client{},
	}
	s.generateDeviceID()
	s.generateAPIKey()
	return s
}
