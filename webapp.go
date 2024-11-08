package mygowebfw

import (
	"strings"
	"text/template"
)

type Assigns map[string]any

type RenderFunc func(*Assigns) string

type WebApp struct {
	components map[string]RenderFunc
	funcs      template.FuncMap
}

var webapp *WebApp
var w = App()

// returns / makes the webapp singleton
func App() *WebApp {
	if webapp == nil {
		webapp = &WebApp{
			components: make(map[string]RenderFunc),
			funcs: template.FuncMap{
				"a": func(m map[string]any, key string) any {
					return m[key]
				},
			},
		}
	}
	return webapp
}

// adds a component to the webapp
func DefComp(name string, renderFunc RenderFunc) {
	if w.components[name] != nil {
		panic("multiple component with name '" + name + "' were attempted to be defined in the webapp.")
	}
	w.components[name] = renderFunc
}

// adds a function that can be used in templates to the webapp
func DefFunc(name string, f any) {
	if w.funcs[name] != nil {
		panic("multiple functions with name '" + name + "' were attempted to be defined in the webapp.")
	}
	w.funcs[name] = f
}

// renders the component and returns its html
func Render(componentName string, assigns *Assigns) string {
	if w.funcs["r"] == nil {
		w.funcs["r"] = Render
	}
	if assigns == nil {
		assigns = &Assigns{}
	}
	renderFunc := w.components[componentName]
	if renderFunc == nil {
		panic("component with name " + componentName + " not found")
	}
	return renderComponent(assigns, renderFunc)
}

// defines a component of the html of the component inside the root page layout with the specified title and description
func DefPage(pageName string, opts map[string]string) {
	if opts["title"] == "" {
		opts["title"] = "Web App"
	}
	DefComp("_page_"+pageName, func(a *Assigns) string {
		return `
		<!DOCTYPE html>
		<html lang="en">

		<head>
		<meta name="description" content="` + opts["description"] + `" />
		<meta charset="utf-8">
		<title>` + opts["title"] + `</title>
		<meta name="viewport" content="width=device-width, initial-scale=1">
		</head>

		<body>
			{{r "` + pageName + `" .}}
		</body>
		</html>
		`
	})
}

// returns the parsed html of the func
func renderComponent(assigns *Assigns, renderFunc RenderFunc) string {
	// execute renderfunc to get the template
	raw := renderFunc(assigns)
	// parse template with funcmap that has r
	tmpl, err := template.New("comp").Funcs(w.funcs).Parse(raw)
	if err != nil {
		panic(err)
	}
	// return parsed template
	var sb strings.Builder
	err = tmpl.Execute(&sb, assigns)
	if err != nil {
		panic(err)
	}
	return sb.String()
}
