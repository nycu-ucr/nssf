/*
 * NSSF Testing Utility
 */

package test

import (
	"reflect"

	. "github.com/nycu-ucr/openapi/models"
)

func CheckAuthorizedNetworkSliceInfo(target AuthorizedNetworkSliceInfo, expectList []AuthorizedNetworkSliceInfo) bool {
	for _, expectElement := range expectList {
		if reflect.DeepEqual(target, expectElement) {
			return true
		}
	}
	return false
}
