// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DiamondModel undocumented
type DiamondModel int

const (
	// DiamondModelVUnknown undocumented
	DiamondModelVUnknown DiamondModel = 0
	// DiamondModelVAdversary undocumented
	DiamondModelVAdversary DiamondModel = 1
	// DiamondModelVCapability undocumented
	DiamondModelVCapability DiamondModel = 2
	// DiamondModelVInfrastructure undocumented
	DiamondModelVInfrastructure DiamondModel = 3
	// DiamondModelVVictim undocumented
	DiamondModelVVictim DiamondModel = 4
	// DiamondModelVUnknownFutureValue undocumented
	DiamondModelVUnknownFutureValue DiamondModel = 127
)

// DiamondModelPUnknown returns a pointer to DiamondModelVUnknown
func DiamondModelPUnknown() *DiamondModel {
	v := DiamondModelVUnknown
	return &v
}

// DiamondModelPAdversary returns a pointer to DiamondModelVAdversary
func DiamondModelPAdversary() *DiamondModel {
	v := DiamondModelVAdversary
	return &v
}

// DiamondModelPCapability returns a pointer to DiamondModelVCapability
func DiamondModelPCapability() *DiamondModel {
	v := DiamondModelVCapability
	return &v
}

// DiamondModelPInfrastructure returns a pointer to DiamondModelVInfrastructure
func DiamondModelPInfrastructure() *DiamondModel {
	v := DiamondModelVInfrastructure
	return &v
}

// DiamondModelPVictim returns a pointer to DiamondModelVVictim
func DiamondModelPVictim() *DiamondModel {
	v := DiamondModelVVictim
	return &v
}

// DiamondModelPUnknownFutureValue returns a pointer to DiamondModelVUnknownFutureValue
func DiamondModelPUnknownFutureValue() *DiamondModel {
	v := DiamondModelVUnknownFutureValue
	return &v
}
