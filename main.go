package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Author Gaurav Sablok
Universitat Potsdam
Date 2024-9-15

A program for extracting the MAGs from the assembled metagenomes with the prediction of the start and the end of the MAGs.
It takes the fasta file of the assembled metagenome given that you have a linerarized fasta and then it takes the csv file
indicating the start and the end of the MAGs and then extracts those MAGs from the specified fasta and gives you a fasta fastafile
for the MAGs.


*/

func main() {
	takecsv := flag.String("csvfile", "path to the csv file", "file")
	fastafile := flag.String("fastafile", "path to the fasta file", "file")
	flag.Parse()
	if len(*takecsv) == 0 || len(*fastafile) == 0 {
		fmt.Println("path to the csv and the fasta file needs to be there and defined")
	}

	type csvStruct struct {
		start int
		end   int
		id    string
	}

	type fastaID struct {
		id string
	}

	type fastaSeq struct {
		seq string
	}

	csvOpen, err := os.Open(*takecsv)
	if err != nil {
		log.Fatal(err)
	}
	csvRead := bufio.NewScanner(csvOpen)
	csvLoad := []csvStruct{}
	for csvRead.Scan() {
		line := csvRead.Text()
		csvLoad = append(csvLoad, csvStruct{
			start: strconv.Atoi(strings.Split(string(line), ",")[0]),
			end:   strconv.Atoi(strings.Split(string(line), ",")[1]),
			id:    strings.Split(string(string(line)), ",")[2],
		})
	}

	openFasta, err := os.Open(*fastafile)
	if err != nil {
		log.Fatal(err)
	}
	takeFasta := bufio.NewScanner(openFasta)
	fasID := []fastaID{}
	fasSeq := []fastaSeq{}
	extractMAGSeq := []string{}
	extractMAGID := []string{}
	for takeFasta.Scan() {
		line := takeFasta.Text()
		if strings.HasPrefix(string(line), ">") {
			fasID = append(fasID, fastaID{
				id: string(line),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			fasSeq = append(fasSeq, fastaSeq{
				seq: string(line),
			})
		}

		for i := range fasID {
			for j := range csvLoad {
				if fasID[i].id == csvLoad[j].id {
					extractMAGSeq = append(
						extractMAGSeq,
						fasSeq[i].seq[csvLoad[j].start:csvLoad[j].end],
					)
					extractMAGID = append(extractMAGID, fasID[i].id)
				}
			}
		}
	}

	for i := range extractMAGSeq {
		fmt.Println(">", extractMAGID[i], "\n", extractMAGSeq[i])
	}
}
