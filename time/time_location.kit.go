package timeutil

import "time"

// Location wraps time.Location.
type Location time.Location

// ToLocation *time.Location to *Location
func ToLocation(location *time.Location) *Location {
	return (*Location)(location)
}

// TimeLocation time.Location
func (l *Location) TimeLocation() *time.Location {
	return (*time.Location)(l)
}

// Set implements pflag/flag.Value
func (l *Location) Set(s string) error {
	location, err := ParseLocation(s)
	if err != nil {
		return err
	}

	*l = *location

	return nil
}

// Type implements pflag.Value
func (l *Location) Type() string {
	return "location"
}

// String string
func (l Location) String() string {
	location := (time.Location)(l)
	return location.String()
}

// MarshalYAML implements the yaml.Marshaler interface.
func (l Location) MarshalYAML() (interface{}, error) {
	return l.String(), nil
}

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (l *Location) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	location, err := ParseLocation(s)
	if err != nil {
		return err
	}

	*l = *location

	return nil
}

// ParseLocation parses a string into a time.Location
func ParseLocation(locationStr string) (*Location, error) {
	location, err := time.LoadLocation(locationStr)
	if err != nil {
		return nil, err
	}
	return (*Location)(location), nil
}
