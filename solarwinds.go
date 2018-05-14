package solarwinds

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"strconv"

	"golang.org/x/net/html/charset"
)

// Solarwinds is a client for the data extraction API of Solarwinds RMM.
type Solarwinds struct {
	Host string
	Key  string
}

func (api Solarwinds) get(values url.Values, dst interface{}) error {
	values.Set("apikey", api.Key)
	apiurl := (&url.URL{
		Scheme:   "https",
		Host:     api.Host,
		Path:     "/api/",
		RawQuery: values.Encode(),
	}).String()

	resp, err := http.Get(apiurl)
	if err != nil {
		return err
	}

	dec := xml.NewDecoder(resp.Body)
	dec.CharsetReader = charset.NewReaderLabel
	err = dec.Decode(&dst)
	return err
}

// ListClients returns all clients.
func (api Solarwinds) ListClients() ([]Client, error) {
	body := struct {
		Items []Client `xml:"items>client"`
	}{}

	err := api.get(url.Values{
		"service": []string{"list_clients"},
	}, &body)
	if err != nil {
		return nil, err
	}

	return body.Items, nil
}

// ListSites retuns all sites for a client.
func (api Solarwinds) ListSites(clientid int) ([]Site, error) {
	body := struct {
		Items []Site `xml:"items>site"`
	}{}

	err := api.get(url.Values{
		"service":  []string{"list_sites"},
		"clientid": []string{strconv.Itoa(clientid)},
	}, &body)
	if err != nil {
		return nil, err
	}

	return body.Items, nil
}

// ListServers returns all server monitoring devices for a site (including top level asset information if available).
func (api Solarwinds) ListServers(siteid int) ([]Server, error) {
	body := struct {
		Items []Server `xml:"items>server"`
	}{}

	err := api.get(url.Values{
		"service": []string{"list_servers"},
		"siteid":  []string{strconv.Itoa(siteid)},
	}, &body)
	if err != nil {
		return nil, err
	}

	return body.Items, nil
}

// ListWorkstations returns all workstation monitoring devices for a site (including top level asset information if available).
func (api Solarwinds) ListWorkstations(siteid int) ([]Workstation, error) {
	body := struct {
		Items []Workstation `xml:"items>workstation"`
	}{}

	err := api.get(url.Values{
		"service": []string{"list_workstations"},
		"siteid":  []string{strconv.Itoa(siteid)},
	}, &body)
	if err != nil {
		return nil, err
	}

	return body.Items, nil
}

// ListDeviceSoftware returns a list of installed software for a device.
func (api Solarwinds) ListDeviceSoftware(deviceid int) ([]Software, error) {
	body := struct {
		Items []Software `xml:"software>item"`
	}{}

	err := api.get(url.Values{
		"service":  []string{"list_device_asset_details"},
		"deviceid": []string{strconv.Itoa(deviceid)},
	}, &body)
	if err != nil {
		return nil, err
	}

	return body.Items, nil
}

// ListChecks returns a list of all checks for a device.
func (api Solarwinds) ListChecks(deviceid int) ([]Check, error) {
	body := struct {
		Items []Check `xml:"items>check"`
	}{}

	err := api.get(url.Values{
		"service":  []string{"list_checks"},
		"deviceid": []string{strconv.Itoa(deviceid)},
	}, &body)
	if err != nil {
		return nil, err
	}

	return body.Items, nil
}

// GetFormattedCheckOutput returns formatted Dashboard More Information firstline result of check (error or otherwise).
func (api Solarwinds) GetFormattedCheckOutput(checkid int) (string, error) {
	body := struct {
		Output string `xml:"formatted_output"`
	}{}

	err := api.get(url.Values{
		"service": []string{"get_formatted_check_output"},
		"checkid": []string{strconv.Itoa(checkid)},
	}, &body)
	if err != nil {
		return "", err
	}

	return body.Output, nil
}
