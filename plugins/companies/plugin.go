package plugin

import (
	"log"

	"github.com/vroomy/common"
	"github.com/vroomy/vroomy"

	"github.com/vroomy/hello-world/libs/companies"
)

var (
	p Plugin
)

func init() {
	if err := vroomy.Register("companies", &p); err != nil {
		log.Fatal(err)
	}
}

type Plugin struct {
	vroomy.BasePlugin

	c *companies.Companies
}

// Methods to match plugins.Plugin interface below

// Load ensures Teams Database is built and open for access
func (p *Plugin) Load(env vroomy.Environment) (err error) {
	p.c = companies.New()
	return
}

// Backend exposes this plugin's data layer to other plugins
func (p *Plugin) Backend() interface{} {
	return p.c
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
		company companies.Company
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
		company *companies.Company
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
		company companies.Company
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
