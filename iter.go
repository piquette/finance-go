package finance

import (
	"context"

	"github.com/piquette/finance-go/form"
)

// ListMeta is the structure that contains the common properties
// of List iterators. The Count property is only populated if the
// total_count include option is passed in (see tests for example).
type ListMeta struct {
	Count uint32
	More  bool
	URL   string
}

// Query is the function used to get a page listing.
type Query func(*form.Values) ([]interface{}, ListMeta, error)

// Iter provides a convenient interface
// for iterating over the elements
// returned from paginated list API calls.
// Successive calls to the Next method
// will step through each item in the list,
// fetching pages of items as needed.
// Iterators are not thread-safe, so they should not be consumed
// across multiple goroutines.
type Iter struct {
	cur    interface{}
	err    error
	meta   ListMeta
	ctx    context.Context
	qs     *form.Values
	query  Query
	values []interface{}
}

// GetIter returns a new Iter for a given query and its options.
func GetIter(ctx *context.Context, qs *form.Values, query Query) *Iter {
	iter := &Iter{}
	iter.query = query
	iter.ctx = *ctx

	q := qs
	if q == nil {
		q = &form.Values{}
	}
	iter.qs = q

	iter.getPage()
	return iter
}

func (it *Iter) getPage() {
	it.values, it.meta, it.err = it.query(it.qs)
}

// Next advances the Iter to the next item in the list,
// which will then be available
// through the Current method.
// It returns false when the iterator stops
// at the end of the list.
func (it *Iter) Next() bool {
	if len(it.values) == 0 && it.meta.More {
		it.getPage()
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

// Err returns the error, if any,
// that caused the Iter to stop.
// It must be inspected
// after Next returns false.
func (it *Iter) Err() error {
	return it.err
}

// Meta returns the list metadata.
func (it *Iter) Meta() *ListMeta {
	return &it.meta
}
