package handlers

import "github.com/CloudyKit/jet"

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)
