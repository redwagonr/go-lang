package convh

import "fmt"

//Celsius ...
type Celsius float64

//Fahrenheit ...
type Fahrenheit float64

func (c Celsius) String() string    { return fmt.Sprintf("%g C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g F", f) }
