package algo2

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestDomainCount(t *testing.T) {

	cpdomains := []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}

	actualstrs := subdomainVisits(cpdomains)
	sort.Strings(actualstrs)
	expectedstrs := []string{"1 intel.mail.com", "5 org", "5 wiki.org", "50 yahoo.com", "900 google.mail.com", "901 mail.com", "951 com"}

	for _, elem := range actualstrs {
		fmt.Println(elem)
	}

	if !reflect.DeepEqual(actualstrs, expectedstrs) {
		t.Errorf("Expected strings are not same as actual strings")
	}
}
