package iter

import (
	"github.com/piquette/finance-go/form"
)

// Query is the function used to get a response listing.
type Query = func(*form.Values) (interface{}, []interface{}, error)

// Iter provides a convenient interface
// for iterating over the elements
// returned from paginated list API calls.
// Successive calls to the Next method
// will step through each item in the list.
// Iterators are not thread-safe, so they should not be consumed
// across multiple goroutines.
type Iter struct {
	meta   interface{}
	cur    interface{}
	err    error
	values []interface{}
}

// NewE returns a iter wrapping an error.
func NewE(e error) *Iter {
	iter := &Iter{}
	iter.err = e
	return iter
}

// New returns a new instance of Iter for a given query and its options.
func New(qs *form.Values, query Query) *Iter {
	iter := &Iter{}

	q := qs
	if q == nil {
		q = &form.Values{}
	}

	iter.meta, iter.values, iter.err = query(q)
	return iter
}

// Next advances the Iter to the next item in the list,
// which will then be available
// through the Current method.
// It returns false when the iterator stops
// at the end of the list.
func (it *Iter) Next() bool {

	if it.values == nil {
		return false
	}

	if len(it.values) == 0 {
		return false
	}

	it.cur = it.values[0]
	it.values = it.values[1:]
	return true
}

// Current returns the most recent item
// visited by a call to Next.
func (it *Iter) Current() interface{} {
	return it.cur
}

// Meta returns the the meta data
// associated with the query.
func (it *Iter) Meta() interface{} {
	return it.meta
}

// Err returns the error, if any,
// that caused the Iter to stop.
// It must be inspected
// after Next returns false.
func (it *Iter) Err() error {
	return it.err
}

// Count returns the list count.
func (it *Iter) Count() int {
	return len(it.values)
}
