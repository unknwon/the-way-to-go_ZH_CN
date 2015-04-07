# Tideland Common Go Library
#
# Copyright 2009-2011 Frank Mueller / Oldenburg / Germany. 
#
# All rights reserved. Use of this source code is 
# governed by a BSD-style license that can be found
# in the LICENSE file.

include $(GOROOT)/src/Make.inc

TARG=tideland-cgl.googlecode.com/hg
GOFMT=gofmt -w -tabindent -tabwidth=8

GOFILES=\
	cgl.go\
	cglsup.go\
	cglfsm.go\
	cglmon.go\
	cglsmr.go\
	cgltim.go\
	cglsml.go\

include $(GOROOT)/src/Make.pkg

format:
	${GOFMT} ${GOFILES}
	${GOFMT} cgl_test.go

#
# EOF
#
