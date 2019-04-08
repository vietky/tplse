package marketplace

import (
	"testing"
)

func TestCheckout0(t *testing.T) {
	actual, err := Run("")
	if err != nil {
		t.Error(err)
	}
	if actual != 0 {
		t.Errorf("actual value %v is not equal to 0", actual)
	}
}

func TestCheckout1(t *testing.T) {
	actual, err := Run("001,002,003")
	if err != nil {
		t.Error(err)
	}
	if actual != 66.78 {
		t.Errorf("actual value %v is not equal to 66.78", actual)
	}
}

func TestCheckout2(t *testing.T) {
	actual, err := Run("001,003,001")
	if err != nil {
		t.Error(err)
	}
	if actual != 36.95 {
		t.Errorf("actual value %v is not equal to 36.95", actual)
	}
}
func TestCheckout3(t *testing.T) {
	actual, err := Run("001,002,001,003")
	if err != nil {
		t.Error(err)
	}
	if actual != 73.76 {
		t.Errorf("actual value %v is not equal to 73.76", actual)
	}
}
