package bankserv

import (
	"github.com/dottics/dutil"
	"github.com/google/uuid"
	"net/url"
)

// CreateItem creates a new Item for a transaction based on the item data passed
// to the function.
func (s *Service) CreateItem(i Item) (Item, dutil.Error) {
	// set path
	s.serv.URL.Path = "/item"
	// marshal payload
	p, e := dutil.MarshalReader(i)
	if e != nil {
		return Item{}, e
	}
	// do request
	r, e := s.serv.NewRequest("POST", s.serv.URL.String(), nil, p)
	if e != nil {
		return Item{}, e
	}

	type Data struct {
		Item `json:"item"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode the response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return Item{}, e
	}

	if r.StatusCode != 201 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Item{}, e
	}
	return res.Data.Item, nil
}

// UpdateItem updates an Item based on the item data passed to the function.
func (s *Service) UpdateItem(i Item) (Item, dutil.Error) {
	// set path
	s.serv.URL.Path = "/item/-"
	// marshal payload
	p, e := dutil.MarshalReader(i)
	if e != nil {
		return Item{}, e
	}

	// do request
	r, e := s.serv.NewRequest("PUT", s.serv.URL.String(), nil, p)
	if e != nil {
		return Item{}, e
	}

	type Data struct {
		Item `json:"item"`
	}
	res := struct {
		Data   `json:"data"`
		Errors dutil.Errors `json:"errors"`
	}{}
	// decode response
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return Item{}, e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return Item{}, e
	}

	return res.Data.Item, nil
}

// DeleteItem deletes a specific item based on the UUID of the item that is
// passed. It returns nil if the item delete was successful and return an error
// if any error has occurred.
func (s *Service) DeleteItem(UUID uuid.UUID) dutil.Error {
	// set path
	s.serv.URL.Path = "/item/-"
	// set query string
	qs := url.Values{"uuid": {UUID.String()}}
	s.serv.URL.RawQuery = qs.Encode()
	// do request
	r, e := s.serv.NewRequest("DELETE", s.serv.URL.String(), nil, nil)
	if e != nil {
		return e
	}

	res := struct {
		Errors dutil.Errors `json:"errors"`
	}{}
	_, e = s.serv.Decode(r, &res)
	if e != nil {
		return e
	}

	if r.StatusCode != 200 {
		e := &dutil.Err{
			Status: r.StatusCode,
			Errors: res.Errors,
		}
		return e
	}
	return nil
}
