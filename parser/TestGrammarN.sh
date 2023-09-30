#!/bin/bash
# antlr -o TestRig Tswift_GrammarN.g4
java -jar ./antlr-4.13.0-complete.jar -o TestRig Tswift_GrammarN.g4
javac -cp ./antlr-4.13.0-complete.jar -g -Xlint TestRig/Tswift_GrammarN*.java
java -cp ./antlr-4.13.0-complete.jar:./TestRig org.antlr.v4.gui.TestRig Tswift_GrammarN s -gui ./../entrada.txt