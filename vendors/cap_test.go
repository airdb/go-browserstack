package vendors_test

import (
	"testing"

	"github.com/airdb/go-browserstack/vendors"
)

func TestGetAndoirdCapList(t *testing.T) {
	t.Log("start testing...")
	vendors.GetAndoirdCapList()

	t.Log("test finished")
}
