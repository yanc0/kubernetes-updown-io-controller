package updown

import (
	"github.com/antoineaugusti/updown"
	"github.com/yanc0/kubernetes-updown-io-controller/api/types/v1alpha1"
	"log"
	"time"
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
			c := updown.CheckItem{
				URL: check.Spec.URL,
			}
			_, _, err := client.Check.Add(c)
			if err == nil {
				log.Printf("[INFO] Check %s successfully added", check.Spec.URL)
				lastUpdate = time.Time{} // reset cache
			}
		}
		found = false
	}
}
