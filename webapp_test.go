package mygowebfw

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRender(t *testing.T) {
	DefComp("comp1", func(a *Assigns) string {
		return `
		<div>
			this is comp1, with assigns: {{.}}
			the number is {{a . "num"}}
		</div>
		`
	})

	DefComp("comp2", func(a *Assigns) string {
		(*a)["a1"] = &Assigns{}
		a1 := (*a)["a1"].(*Assigns)
		if rand.Intn(2) > 0 {
			(*a1)["num"] = "big"
		} else {
			(*a1)["num"] = "small"
		}
		return `
		<div>
			this is comp2
			{{ r "comp1" (a . "a1") }}
			{{ r "comp3" . }}
		</div>
		`
	})

	DefFunc("add", func(a, b int) int {
		return a + b
	})

	DefComp("comp3", func(a *Assigns) string {
		return `
		<p>{{add 4 2}}</p>
		`
	})

	fmt.Println(Render("comp2", nil))
}

func TestRun(t *testing.T) {
	AddPage("/", "comp2", map[string]string{"title": "title"})
	Run(8080)
}
