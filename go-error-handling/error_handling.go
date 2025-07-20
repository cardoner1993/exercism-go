package erratum

import (
	"fmt"
)

func Use(opener ResourceOpener, input string) (finalErr error) {
	// Step 1: Keep trying to open the resource until success or non-transient error
	var resource Resource
	var err error

	for {
		resource, err = opener()
		if err == nil {
			break // Successfully opened
		}

		// Check if it's a TransientError
		if _, isTransient := err.(TransientError); isTransient {
			continue // Keep trying
		}

		// Non-transient error, return it
		return err
	}

	// Step 2: Handle panics and ensure proper cleanup order
	defer func() {
		if r := recover(); r != nil {
			// Check if the panic is a FrobError
			if frobErr, ok := r.(FrobError); ok {
				// Call Defrob FIRST, then Close
				resource.Defrob(frobErr.defrobTag)
				resource.Close()
				finalErr = frobErr // Return the original FrobError
			} else {
				// Other type of panic, convert to error
				resource.Close()
				// If the panic value is already an error, use it directly
				if panicErr, ok := r.(error); ok {
					finalErr = panicErr
				} else {
					finalErr = fmt.Errorf("panic occurred: %v", r)
				}
			}
		} else {
			// No panic, just close normally
			resource.Close()
		}
	}()

	// Step 3: Call Frob on the resource
	resource.Frob(input) // Frob doesn't return an error, but may panic

	return nil // If we get here, everything succeeded
}
