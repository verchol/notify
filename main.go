package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig"
)

//import hook manifest .
//add hook phase
//add image , command and arguments
/*
   "helm.sh/hook": post-upgrade
   "helm.sh/hook-delete-policy": hook-failed
   "helm.sh/hook-delete-policy": hook-succeeded
*/
const tmpl = `

{{- define "helm.hook" -}}
"helm.sh/hook": {{ .Values.CI.Hooks.Events| quote  }}
"helm.sh/hook-delete-policy": {{.Values.CI.Hooks.Policy| quote }} 
{{- end -}}

apiVersion: v1
kind: Pod
metadata:
  name: "t1-{{ uuidv4 }}-test-connection"
  annotations:
    "helm.sh/hook": {{ .Values.CI.Hooks.Events| quote  }}
    "helm.sh/hook-delete-policy": {{.Values.CI.Hooks.Policy| quote }} 
 
spec:
 containers:
    - name: cf
      env:
        - name : CF_API_KEY
          value : {{.Values.CI.Hooks.Token}}
      image: {{printf "%v" .Values.CI.Hooks.Image  }}
      {{- if .Values.CI.Hooks.Cmd}} 
          {{-  .Values.CI.Hooks.Cmd  |  putCommand  |nindent 6 }}
      {{- end}}
      args: {{.Values.CI.Hooks.Args}}
    `

func tpl(t string, vals interface{}, out io.Writer) error {
	tt, err := template.New("_").Funcs(sprig.TxtFuncMap()).Funcs(template.FuncMap{
		"putCommand": func(cmd []string) string {
			//    fmt.Printf("\n[CMD]- %s\n", cmd)
			if len(cmd) == 0 {
				return ""
			}
			return fmt.Sprintf("command: %v", cmd)
		},
	}).Parse(t)
	if err != nil {
		return err
	}
	return tt.Execute(out, vals)
}

func main() {

	image := flag.String("image", "codefresh/cli", "image where hook will be executed")
	cmd := flag.String("cmd", "", "hook command")
	args := flag.String("args", "", "hook args")
	pipeline := flag.String("pip", "", "pipeline ")
	apitoken := flag.String("apitoken", "", "context for count")
	pipArgs := flag.String("pipArgs", "", "context for count")
	hooks := flag.String("hooks", "", "list of events for count")

	output := flag.String("output", "", "where to located hook file")
	flag.Parse()

	if *hooks == "" {

		//*args = "--cfconfig /cf/.cfconfig  run 'test/helm3"
		*hooks = fmt.Sprintf("%s, %s", "post-install", "post-upgrade")

		//os.Exit(1)
	}

	var writer io.Writer
	if *output == "" {
		writer = os.Stdout
	} else {
		f, err := os.Create(*output)
		if err != nil {
			panic(err)
		}
		writer = f
	}
	var argsList []string
	var cmdList []string
	if *args == "" {

		//*args = "--cfconfig /cf/.cfconfig  run 'test/helm3"
		args := fmt.Sprintf("run %s %s", *pipeline, *pipArgs)
		argsList = strings.Split(args, " ")
		//os.Exit(1)
	} else {
		argsList = strings.Split(*args, " ")
	}

	if *cmd != "" {
		cmdList = strings.Split(*cmd, " ")
	} else {
		//cmdList = []string{"a", "b", "c", "d"}
	}
	type Hooks struct {
		Image  string
		Cmd    []string
		Args   []string
		Token  string
		Events string
		Policy string
	}
	type CI struct {
		Hooks
	}
	type Values struct {
		CI
	}
	type Input struct {
		Values
	}
	v := Input{
		Values{CI{Hooks{
			Image:  *image,
			Cmd:    cmdList,
			Args:   argsList,
			Token:  *apitoken,
			Events: *hooks,
			Policy: "hook-failed, hook-succeded",
		}}}}

	err := tpl(tmpl, v, writer)
	if err != nil {
		panic(err)
	}
}
