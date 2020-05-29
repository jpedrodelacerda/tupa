# Tupã #

Tupã is a simple CLI to help you create new projects.

## Introduction ##

Are you tired of creating .gitignore files and LICENSES? If so, here is Tupã.
Tupã makes sure that you don't have to worry about copying that whole text from [gitignore.io](https://gitignore.io/) or setup your project license on Github.

## Instalation ##

You can install `tupa` by running the following:

`$ go get github.com/jpedrodelacerda/tupa`

## Usage ##

To create a project with `tupa`:
`$ tupa init <project name> -i go,emacs,vim,code --author "João de Lacerda" -l mit`

It currently supports the following licenses:
- `apache-2.0`
- `gpl-3.0`
- `lgpl-3.0`
- `mit`
- `unlincense`
- `wtfpl`

Tupã can also check a config file for default parameters:

```yaml
---
author: João de Lacerda
license: mit
```

## Motivation ##

This project was created for educational purposes.
It is simple, but I hope it can be helpful for others.

## License ##

MIT
