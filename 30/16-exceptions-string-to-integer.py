#!/bin/python

import sys

S = raw_input().strip()

try:
    i = int(S)
    print i
except ValueError:
    print "Bad String"
