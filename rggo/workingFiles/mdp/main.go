package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!DOCTYPE html>
<html>
  <head>
    <meta http-equiv="content-type" content="text/html"; charset=utf-8">
    <title>Markdown Preview Tool</title>
  </head>
  <body>
  `

	footer = `
  </body>
<html>
`
)
