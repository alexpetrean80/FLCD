package main

import (
	"fmt"
	"github.com/alexpetrean80/FLCD/pif"
	"github.com/alexpetrean80/FLCD/scanner"
	"github.com/alexpetrean80/FLCD/symtable"
)

func runScan(s scanner.Scanner, inFile string) {
	fmt.Println(inFile)
	if errs := s.Scan(inFile); errs != nil {
		for _, err := range errs {
			fmt.Println(err.Error())
		}
	} else {
		fmt.Println("lexically correct")
	}

	s.GetIdentifiers().Output(fmt.Sprintf("%s.d/identifiers", inFile))
	s.GetConstants().Output(fmt.Sprintf("%s.d/constants", inFile))
	s.GetPIF().Output(fmt.Sprintf("./%s.d/pif", inFile))
}

func main() {
	inFile1 := "../test_programs/p1.lang"
	s1 := scanner.New(symtable.New(), symtable.New(), pif.New())
	runScan(s1, inFile1)

	inFile2 := "../test_programs/p2.lang"
	s2 := scanner.New(symtable.New(), symtable.New(), pif.New())
	runScan(s2, inFile2)

	inFile3 := "../test_programs/p3.lang"
	s3 := scanner.New(symtable.New(), symtable.New(), pif.New())
	runScan(s3, inFile3)

	inFileErr := "../test_programs/perror.lang"
	sErr := scanner.New(symtable.New(), symtable.New(), pif.New())
	runScan(sErr, inFileErr)
}
