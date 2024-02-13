package erratum

import (
	"fmt"
)

func Use(opener ResourceOpener, input string) (ret error) {
	resource, err := opener()
	for err != nil && err.Error() == "some error" && err.(TransientError).Err != nil {
		fmt.Println("Resource Unavailable!", err.(TransientError).Err)
		fmt.Println("Retrying!")
		resource, err = opener()
		if err != nil && err.Error() != "some error" {
			return err
		}
	}
	if resource != nil {
		defer resource.Close()
		defer func() {
			//convert to if else
			if err := recover(); err != nil {
				if v, ok := err.(FrobError); ok {
					resource.Defrob(v.DefrobTag)
					ret = err.(FrobError)
				} else {
					ret = err.(error)
				}
			}
		}()
		resource.Frob(input)
	}
	return err
}
