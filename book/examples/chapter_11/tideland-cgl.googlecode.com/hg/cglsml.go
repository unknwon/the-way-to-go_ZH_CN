/*
	Tideland Common Go Library - Simple Markup Language

	Copyright (C) 2011 Frank Mueller / Oldenburg / Germany

	Redistribution and use in source and binary forms, with or
	modification, are permitted provided that the following conditions are
	met:

	Redistributions of source code must retain the above copyright notice, this
	list of conditions and the following disclaimer.

	Redistributions in binary form must reproduce the above copyright notice,
	this list of conditions and the following disclaimer in the documentation
	and/or other materials provided with the distribution.

	Neither the name of Tideland nor the names of its contributors may be
	used to endorse or promote products derived from this software without
	specific prior written permission.

	THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
	AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
	IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
	ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
	LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
	CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
	SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
	INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
	CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
	ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF
	THE POSSIBILITY OF SUCH DAMAGE.
*/

package cgl

//--------------------
// IMPORTS
//--------------------

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

//--------------------
// PROCESSOR
//--------------------

// Processor interface.
type Processor interface {
	OpenTag(tag []string)
	CloseTag(tag []string)
	Text(text string)
}

//--------------------
// NODE
//--------------------

// The node.
type Node interface {
	Len() int
	ProcessWith(p Processor)
}

//--------------------
// TAG NODE
//--------------------

// The tag node.
type TagNode struct {
	tag      []string
	children []Node
}

// Create a new tag node.
func NewTagNode(tag string) *TagNode {
	tmp := strings.ToLower(tag)

	if !validIdentifier(tmp) {
		return nil
	}

	tn := &TagNode{
		tag:      strings.Split(tmp, ":"),
		children: make([]Node, 0),
	}

	return tn
}

// Append a new tag.
func (tn *TagNode) AppendTag(tag string) *TagNode {
	n := NewTagNode(tag)

	if n != nil {
		tn.children = append(tn.children, n)
	}

	return n
}

// Append a tag node.
func (tn *TagNode) AppendTagNode(n *TagNode) *TagNode {
	tn.children = append(tn.children, n)

	return n
}

// Append a text node.
func (tn *TagNode) AppendText(text string) *TextNode {
	n := NewTextNode(text)

	tn.children = append(tn.children, n)

	return n
}

// Append a tagged text node.
func (tn *TagNode) AppendTaggedText(tag, text string) *TagNode {
	n := NewTagNode(tag)

	if n != nil {
		n.AppendText(text)

		tn.children = append(tn.children, n)
	}

	return n
}

// Append a text node.
func (tn *TagNode) AppendTextNode(n *TextNode) *TextNode {
	tn.children = append(tn.children, n)

	return n
}

// Return the len of the tag node (aka number of children).
func (tn *TagNode) Len() int {
	return len(tn.children)
}

// Process the node.
func (tn *TagNode) ProcessWith(p Processor) {
	p.OpenTag(tn.tag)

	for _, child := range tn.children {
		child.ProcessWith(p)
	}

	p.CloseTag(tn.tag)
}

// Return the node as a string.
func (tn *TagNode) String() string {
	buf := bytes.NewBufferString("")
	spp := NewSmlWriterProcessor(buf, true)

	tn.ProcessWith(spp)

	return buf.String()
}

//--------------------
// TEXT NODE
//--------------------

// The text node.
type TextNode struct {
	text string
}

// Create a new text node.
func NewTextNode(text string) *TextNode {
	return &TextNode{text}
}

// Return the len of the text node.
func (tn *TextNode) Len() int {
	return len(tn.text)
}

// Process the node.
func (tn *TextNode) ProcessWith(p Processor) {
	p.Text(tn.text)
}

// Return the node as a string.
func (tn *TextNode) String() string {
	return tn.text
}

//--------------------
// PRIVATE FUNCTIONS
//--------------------

// Check an identifier (tag or id).
func validIdentifier(id string) bool {
	for _, c := range id {
		if c < 'a' || c > 'z' {
			if c < '0' || c > '9' {
				if c != '-' && c != ':' {
					return false
				}
			}
		}
	}

	return true
}

//--------------------
// SML READER
//--------------------

// Control values.
const (
	ctrlText = iota
	ctrlSpace
	ctrlOpen
	ctrlClose
	ctrlEscape
	ctrlTag
	ctrlEOF
	ctrlInvalid
)

// Node read modes.
const (
	modeInit = iota
	modeTag
	modeText
)

// Reader for SML.
type SmlReader struct {
	reader *bufio.Reader
	index  int
	root   *TagNode
	error  os.Error
}

// Create the reader.
func NewSmlReader(reader io.Reader) *SmlReader {
	// Init the reader.

	sr := &SmlReader{
		reader: bufio.NewReader(reader),
		index:  -1,
	}

	node, ctrl := sr.readNode()

	switch ctrl {
	case ctrlClose:
		sr.root = node
		sr.error = nil
	case ctrlEOF:
		msg := fmt.Sprintf("eof too early at index %v", sr.index)

		sr.error = os.NewError(msg)
	case ctrlInvalid:
		msg := fmt.Sprintf("invalid rune at index %v", sr.index)

		sr.error = os.NewError(msg)
	}

	return sr
}

// Return the root tag node.
func (sr *SmlReader) RootTagNode() (*TagNode, os.Error) {
	return sr.root, sr.error
}

// Read a node.
func (sr *SmlReader) readNode() (*TagNode, int) {
	var node *TagNode
	var buffer *bytes.Buffer

	mode := modeInit

	for {
		rune, ctrl := sr.readRune()

		sr.index++

		switch mode {
		case modeInit:
			// Before the first opening bracket.
			switch ctrl {
			case ctrlEOF:
				return nil, ctrlEOF
			case ctrlOpen:
				mode = modeTag
				buffer = bytes.NewBufferString("")
			}
		case modeTag:
			// Reading a tag.
			switch ctrl {
			case ctrlEOF:
				return nil, ctrlEOF
			case ctrlTag:
				buffer.WriteRune(rune)
			case ctrlSpace:
				if buffer.Len() == 0 {
					return nil, ctrlInvalid
				}

				node = NewTagNode(buffer.String())
				buffer = bytes.NewBufferString("")
				mode = modeText
			case ctrlClose:
				if buffer.Len() == 0 {
					return nil, ctrlInvalid
				}

				node = NewTagNode(buffer.String())

				return node, ctrlClose
			default:
				return nil, ctrlInvalid
			}
		case modeText:
			// Reading the text including the subnodes following
			// the space after the tag or id.
			switch ctrl {
			case ctrlEOF:
				return nil, ctrlEOF
			case ctrlOpen:
				text := strings.TrimSpace(buffer.String())

				if len(text) > 0 {
					node.AppendText(text)
				}

				buffer = bytes.NewBufferString("")

				sr.reader.UnreadRune()

				subnode, subctrl := sr.readNode()

				if subctrl == ctrlClose {
					// Correct closed subnode.

					node.AppendTagNode(subnode)
				} else {
					// Error while reading the subnode.

					return nil, subctrl
				}
			case ctrlClose:
				text := strings.TrimSpace(buffer.String())

				if len(text) > 0 {
					node.AppendText(text)
				}

				return node, ctrlClose
			case ctrlEscape:
				rune, ctrl = sr.readRune()

				if ctrl == ctrlOpen || ctrl == ctrlClose || ctrl == ctrlEscape {
					buffer.WriteRune(rune)

					sr.index++
				} else {
					return nil, ctrlInvalid
				}
			default:
				buffer.WriteRune(rune)
			}
		}
	}

	return nil, ctrlEOF
}

// Read a rune.
func (sr *SmlReader) readRune() (rune, control int) {
	var size int

	rune, size, sr.error = sr.reader.ReadRune()

	switch {
	case size == 0:
		return rune, ctrlEOF
	case rune == '{':
		return rune, ctrlOpen
	case rune == '}':
		return rune, ctrlClose
	case rune == '^':
		return rune, ctrlEscape
	case rune >= 'a' && rune <= 'z':
		return rune, ctrlTag
	case rune >= 'A' && rune <= 'Z':
		return rune, ctrlTag
	case rune >= '0' && rune <= '9':
		return rune, ctrlTag
	case rune == '-':
		return rune, ctrlTag
	case rune == ':':
		return rune, ctrlTag
	case unicode.IsSpace(rune):
		return rune, ctrlSpace
	}

	return rune, ctrlText
}

//--------------------
// SML WRITER PROCESSOR
//--------------------

// Processor for writing SML.
type SmlWriterProcessor struct {
	writer      *bufio.Writer
	prettyPrint bool
	indentLevel int
}

// Create a new SML writer processor.
func NewSmlWriterProcessor(writer io.Writer, prettyPrint bool) *SmlWriterProcessor {
	swp := &SmlWriterProcessor{
		writer:      bufio.NewWriter(writer),
		prettyPrint: prettyPrint,
		indentLevel: 0,
	}

	return swp
}

// Open a tag.
func (swp *SmlWriterProcessor) OpenTag(tag []string) {
	swp.writeIndent(true)

	swp.writer.WriteString("{")
	swp.writer.WriteString(strings.Join(tag, ":"))
}

// Close a tag.
func (swp *SmlWriterProcessor) CloseTag(tag []string) {
	swp.writer.WriteString("}")

	if swp.prettyPrint {
		swp.indentLevel--
	}

	swp.writer.Flush()
}

// Write a text.
func (swp *SmlWriterProcessor) Text(text string) {
	ta := strings.Replace(text, "^", "^^", -1)
	tb := strings.Replace(ta, "{", "^{", -1)
	tc := strings.Replace(tb, "}", "^}", -1)

	swp.writeIndent(false)

	swp.writer.WriteString(tc)
}

// Write an indent in case of pretty print.
func (swp *SmlWriterProcessor) writeIndent(increase bool) {
	if swp.prettyPrint {
		if swp.indentLevel > 0 {
			swp.writer.WriteString("\n")
		}

		for i := 0; i < swp.indentLevel; i++ {
			swp.writer.WriteString("\t")
		}

		if increase {
			swp.indentLevel++
		}
	} else {
		swp.writer.WriteString(" ")
	}
}

/*
	EOF
*/
