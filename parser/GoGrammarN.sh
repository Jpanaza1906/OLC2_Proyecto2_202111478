#!/bin/sh

alias antlr4='java -Xmx500M -cp "./antlr-4.13.0-complete.jar:$CLASSPATH" org.antlr.v4.Tool'

antlr4 -Dlanguage=Go -o ../TswiftGen -package TswiftGen -visitor Tswift_GrammarN.g4