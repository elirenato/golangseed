package filters

import (
	"github.com/revel/revel"
)

// HeaderFilter adds common security headers
// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Set("Access-Control-Allow-Origin", "*")
	c.Response.Out.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE");
	c.Response.Out.Header().Set("Access-Control-Allow-Headers", "Content-type, Authorization, Access-Control-Allow-Origin, Access-Control-Allow-Methods, Accept-Language");
	// c.Response.Out.Header().Add("Access-Control-Allow-Credentials", "true");
	// Stop here if its Preflighted OPTIONS request
	if c.Request.Method == "OPTIONS" {
		return
	}
	fc[0](c, fc[1:]) // Execute the next filter stage.
}
