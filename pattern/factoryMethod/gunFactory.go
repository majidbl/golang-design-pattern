package factoryMethod

import "fmt"
import "design_pattern/domain/factoryMethod"

func GetGun(gunType string) (factoryMethod.Gun, error) {
	if gunType == "ak47" {
		return NewAk47(), nil
	}
	if gunType == "musket" {
		return NewMusket(), nil
	}
	return nil, fmt.Errorf("wrong gun type passed")
}
