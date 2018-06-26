package updown

import (
	"log"
	"time"

	"github.com/antoineaugusti/updown"
	"github.com/yanc0/kubernetes-updown-io-controller/api/types/v1alpha1"
)

var lastUpdate time.Time
var preHeat uint
var updownChecks []updown.Check
var err error

func Sync(checks []*v1alpha1.Check, apiKey string) {
	client := updown.NewClient(apiKey, nil)
	if time.Since(lastUpdate) > 10*time.Second {
		updownChecks, _, err = client.Check.List()
		if err != nil {
			log.Println(err)
			return
		}
		lastUpdate = time.Now()
	}

	// No checks can mean that the informer is not ready to
	// distribute checks, wait for 6 secondes
	if len(checks) == 0 && preHeat < 3 {
		time.Sleep(2 * time.Second)
		preHeat = preHeat + 1
		return
	}
	for _, uCheck := range updownChecks {
		found := false
		for _, check := range checks {
			if !found && check.Spec.URL == uCheck.URL {
				found = true
				if updateNeeded(uCheck, *check) {
					_, _, err = client.Check.Update(uCheck.Token, toCheckItem(check))
					if err == nil {
						log.Printf("[INFO] Check %s (%s) successfully updated", uCheck.URL, uCheck.Token)
						lastUpdate = time.Time{} // reset cache
					}
				}
			}
		}
		if !found {
			deleted, _, _ := client.Check.Remove(uCheck.Token)
			if deleted {
				log.Printf("[INFO] Check %s (%s) successfully deleted", uCheck.URL, uCheck.Token)
				lastUpdate = time.Time{} // reset cache
			}
		}
		found = false
	}

	for _, check := range checks {
		found := false
		for _, uCheck := range updownChecks {
			if !found && check.Spec.URL == uCheck.URL {
				found = true
			}
		}
		if !found {
			c := toCheckItem(check)
			_, _, err := client.Check.Add(c)
			if err == nil {
				log.Printf("[INFO] Check %s successfully added", check.Spec.URL)
				lastUpdate = time.Time{} // reset cache
			}
		}
		found = false
	}
}

func toCheckItem(c *v1alpha1.Check) updown.CheckItem {
	c.LoadDefaults()

	customHeaders := make(map[string]string)
	for _, ch := range c.Spec.CustomHeaders {
		customHeaders[ch.Key] = ch.Value
	}

	return updown.CheckItem{
		URL:               c.Spec.URL,
		Alias:             c.Spec.Alias,
		Apdex:             c.Spec.ApdexT,
		Enabled:           c.Spec.Enabled,
		Period:            c.Spec.Period,
		Published:         c.Spec.Published,
		StringMatch:       c.Spec.StringMatch,
		DisabledLocations: c.Spec.DisabledLocations,
		CustomHeaders:     customHeaders,
	}
}

func updateNeeded(uCheck updown.Check, check v1alpha1.Check) bool {
	if uCheck.Alias != check.Spec.Alias {
		return true
	}
	if uCheck.Apdex != check.Spec.ApdexT {
		return true
	}
	if uCheck.Period != check.Spec.Period {
		return true
	}
	if uCheck.Published != check.Spec.Published {
		return true
	}
	if uCheck.StringMatch != check.Spec.StringMatch {
		return true
	}

	if len(uCheck.DisabledLocations) != len(check.Spec.DisabledLocations) {
		return true
	}

	for _, ch := range check.Spec.CustomHeaders {
		if _, ok := uCheck.CustomHeaders[ch.Key]; !ok {
			return true
		}
		if uCheck.CustomHeaders[ch.Key] != ch.Value {
			return true
		}
	}

	return false
}
