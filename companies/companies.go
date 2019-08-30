package companies

import (
	"fmt"
	"strconv"
	"sync"
)

const errCannotFindCompanyFmt = "cannot find company with ID \"%s\""

// New will return a new instance of companies
func New() *Companies {
	var c Companies
	c.cm = make(map[string]*Company)
	return &c
}

// Companies manages companies
type Companies struct {
	mux sync.RWMutex
	cm  map[string]*Company

	lastIndex int
}

// New will create a new company and return it's ID
func (c *Companies) New(company *Company) (companyID string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	companyID = strconv.Itoa(c.lastIndex)
	c.lastIndex++
	c.cm[companyID] = company
	return
}

// Get will retrieve a company by ID
func (c *Companies) Get(companyID string) (company *Company, err error) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	var ok bool
	if company, ok = c.cm[companyID]; !ok {
		err = fmt.Errorf(errCannotFindCompanyFmt, companyID)
		return
	}

	return
}

// Put will set a company by ID
func (c *Companies) Put(companyID string, company *Company) (err error) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if _, ok := c.cm[companyID]; !ok {
		return fmt.Errorf(errCannotFindCompanyFmt, companyID)
	}

	c.cm[companyID] = company
	return
}

// Delete will remove a company by ID
func (c *Companies) Delete(companyID string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	delete(c.cm, companyID)
}
