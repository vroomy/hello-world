package companies

import (
	"github.com/vroomy/common"
)

var (
	p Plugin
)

type Plugin struct {
	c *Companies
}

// Methods to match plugins.Plugin interface below

// Backend returns the underlying backend to the plugin
func (p *Plugin) Backend() interface{} {
	return p.c
}

// Load will be called by Vroomie on load
func (p *Plugin) Init(env map[string]string) (err error) {
	// We currently don't need to initialize anything additional
	return
}

// Init will be called by Vroomie on initialization
func (p *Plugin) Load() (err error) {
	// We currently don't need to initialize anything additional
	return
}

// Close will close the plugin
func (p *Plugin) Close() (err error) {
	// We currently don't offer any persistence, no closing is required
	return
}

// Handlers below

// New will create a new company and return its company ID
func (p *Plugin) New(ctx common.Context) {
	var (
		company Company
		err     error
	)

	if err = ctx.Bind(&company); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	companyID := p.c.New(&company)
	ctx.WriteJSON(400, companyID)
}

// Get will retrieve a campaign by ID
func (p *Plugin) Get(ctx common.Context) {
	var (
		company *Company
		err     error
	)

	companyID := ctx.Param("companyID")

	if company, err = p.c.Get(companyID); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteJSON(200, company)
}

// Edit will modify a profile by user ID
func (p *Plugin) Edit(ctx common.Context) {
	var (
		company Company
		err     error
	)

	companyID := ctx.Param("companyID")

	if err = ctx.Bind(&company); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	if err = p.c.Put(companyID, &company); err != nil {
		ctx.WriteJSON(400, err)
		return
	}

	ctx.WriteNoContent()
}

// Delete will modify a profile by user ID
func (p *Plugin) Delete(ctx common.Context) {
	companyID := ctx.Param("companyID")
	p.c.Delete(companyID)
	ctx.WriteNoContent()
}
