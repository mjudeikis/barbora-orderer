package brb

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/mjudeikis/barbora-orderer/pkg/hargo"
)

func usage() {
	fmt.Fprint(flag.CommandLine.Output(), "usage: \n")
	fmt.Fprintf(flag.CommandLine.Output(), "       %s {har_file_name} \n", os.Args[0])
	flag.PrintDefaults()
}

type orderer struct {
	har *hargo.Har
}

func Run(path string) error {
	har, err := parseHarFile(path)
	if err != nil {
		return err
	}
	o := orderer{
		har: har,
	}

	// get cart
	client := &http.Client{
		Timeout: time.Duration(time.Second * 5),
	}
	req, err := http.NewRequest("GET", "https://www.barbora.lt/api/eshop/v1/cart/deliveries", nil)
	if err != nil {
		return err
	}

	o.addHeader(req)
	o.addCookies(req)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	deliveriesList := deliveries{}
	err = json.Unmarshal(data, &deliveriesList)
	if err != nil {
		return err
	}

	for _, delivery := range deliveriesList.Deliveries[0].Params.Matrix {
		log.Printf("checking %s (%s)", delivery.Day, delivery.ID)
		for _, windows := range delivery.Hours {
			//log.Printf("hour %v (%s) is %v", windows.Hour, windows.ID, windows.Available)
			// we want booze
			if windows.Available && windows.Hour != "08 - 09" && windows.Hour != "09 - 10" {
				log.Printf("found - trying %s - %s", windows.Hour, delivery.Day)
				err := o.reserve(delivery.ID, windows.ID)
				if err != nil {
					return err
				}
			}
		}
	}
	//err = o.reserve("2020-03-24", "f600b1ab-7f03-45f3-8413-e542d3eda37b")
	//if err != nil {
	//	log.Fatal(err)
	//}
	return nil
}

func parseHarFile(path string) (*hargo.Har, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	har := &hargo.Har{}
	err = json.Unmarshal(data, har)
	if err != nil {
		return nil, err
	}

	return har, err
}

func (o *orderer) reserve(dayID, hourID string) error {
	q := url.Values{}
	q.Add("dayID", dayID)
	q.Add("hourID", hourID)
	q.Add("isExpressDeliveryTimeslot", "false")

	client := &http.Client{
		Timeout: time.Duration(time.Second * 5),
	}
	req, err := http.NewRequest("POST", "https://www.barbora.lt/api/eshop/v1/cart/ReserveDeliveryTimeSlot", strings.NewReader(q.Encode()))
	if err != nil {
		return err
	}
	spew.Dump(req.Body)

	o.addCookies(req)
	o.addHeader(req)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		log.Println("reserved")

		os.Exit(0)
	}

	return nil
}

func (o *orderer) addHeader(req *http.Request) {
	req.Header.Add("Authorization", "Basic YXBpa2V5OlNlY3JldEtleQ==")
	for _, h := range o.har.Log.Entries[0].Request.Headers {
		req.Header.Add(h.Name, h.Value)
	}
}

func (o *orderer) addCookies(req *http.Request) {
	for _, h := range o.har.Log.Entries[0].Request.Cookies {
		c := &http.Cookie{
			Name:  h.Name,
			Value: h.Value,
		}
		req.AddCookie(c)
	}
}
