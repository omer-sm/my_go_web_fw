# concept
user defines components which are things that receive *assigns and return html.
components are defined with webapp.DefComp (webapp is a singleton)
DefComp receives a name and a render func
render func: receives *assigns and returns a (assigns, html template (string))

components can do {{"component" | r}} to render another component in them
or {{"component" | r a b c etc}} to give the component assigns
assigns can be accessed like {{a . "assign name"}}

server start receives webapp and port

maybe: POST routes, live updates