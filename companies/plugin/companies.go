package main

import (
	"github.com/Hatch1fy/httpserve"
	"github.com/vroomy/common"
	"github.com/vroomy/hello-world/companies"
)

var c *companies.Companies

func init() {
	c = companies.New()
}

// Backend returns the underlying backend to the plugin
func Backend() interface{} {
	return c
}

// OnInit will be called by Vroomie on initialization
func OnInit(p common.Plugins, flags, env map[string]string) (err error) {
	// We currently don't need to initialize anything additional
	return
}

// New will create a new company and return its company ID
func New(ctx *httpserve.Context) httpserve.Response {
	var (
		company companies.Company
		err     error
	)

	if err = ctx.BindJSON(&company); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	companyID := c.New(&company)

	return httpserve.NewTextResponse(200, []byte(companyID))
}

// Get will retrieve a campaign by ID
func Get(ctx *httpserve.Context) httpserve.Response {
	var (
		company *companies.Company
		err     error
	)

	companyID := ctx.Param("companyID")

	if company, err = c.Get(companyID); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewJSONResponse(200, company)
}

// Edit will modify a profile by user ID
func Edit(ctx *httpserve.Context) httpserve.Response {
	var (
		company companies.Company
		err     error
	)

	companyID := ctx.Param("companyID")

	if err = ctx.BindJSON(&company); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	if err = c.Put(companyID, &company); err != nil {
		return httpserve.NewJSONResponse(400, err)
	}

	return httpserve.NewNoContentResponse()
}

// Delete will modify a profile by user ID
func Delete(ctx *httpserve.Context) httpserve.Response {
	companyID := ctx.Param("companyID")
	c.Delete(companyID)
	return httpserve.NewNoContentResponse()
}

// Close will close the plugin
func Close() (err error) {
	// We currently don't offer any persistence, no closing is required
	return
}
