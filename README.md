# CF Prompt plugin
CF Prompt plugin formats a string containing placeholders, so you can show it in your prompt.

## Usage

```
$ cf prompt "You're targetting org %o and space %s"
You're targetting org FOO and space BAR
```

To use this in your shell prompt, export the PS1 variable in your ~/.bashrc (or equivalent):

```
export PS1="\$(cf prompt '%o/%s (%a)') $ "
```

## Installation
```
go build -o cf-prompt
cf install-plugin -f ./cf-prompt
```
