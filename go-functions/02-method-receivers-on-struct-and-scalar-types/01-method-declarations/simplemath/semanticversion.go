package simplemath

import "fmt"

// SemanticVersion : defines a type struct
type SemanticVersion struct {
	major, minor, patch int
}

// NewSemanticVersion : this function creates a receiver object which defined methods are executed on.
func NewSemanticVersion(major, minor, patch int) SemanticVersion {
	return SemanticVersion{
		major: major,
		minor: minor,
		patch: patch,
	}
}

// String : this method prints the version
func (sv SemanticVersion) String() string {
	return fmt.Sprintf("%d.%d.%d", sv.major, sv.minor, sv.patch)
}

// IncrementMajor : this method increments and prints the major version
func (sv *SemanticVersion) IncrementMajor() {
	sv.major++
}

// IncrementMinor : this method increments and prints the minor version
func (sv *SemanticVersion) IncrementMinor() {
	sv.minor++
}

// IncrementPatch : this method increments and prints the patch version
func (sv *SemanticVersion) IncrementPatch() {
	sv.patch++
}
