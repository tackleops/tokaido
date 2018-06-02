package unisontmpl

var PRFTemplateStr = `# Generated by TOKAIDO
root = {{.ProjectPath}}
root = socket://localhost:{{.UnisonPort}}

# Paths to synchronize
path = .

# Some regexps specifying names and paths to ignore
ignore = Name node_modules
ignore = Name .git

# Log actions to the terminal
log = true

# Other options
prefer = newer
repeat = watch
auto = true`
