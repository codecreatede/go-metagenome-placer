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

	type alignmentID struct {
		id string
	}

	type alignmentSeq struct {
		seq string
	}

	type magStruct struct {
		start int
		end   int
		id    string
	}

	fOpen, err := os.Open(*fastafile)
	if err != nil {
		log.Fatal(err)
	}
	fRead := bufio.NewScanner(fOpen)
	alignmentIDStore := []alignmentID{}
	alignmentSeqStore := []alignmentSeq{}
	for fRead.Scan() {
		line := fRead.Text()
		if strings.HasPrefix(string(line), ">") {
			alignmentIDStore = append(alignmentIDStore, alignmentID{
				id: strings.Replace(string(line), ">", "", -1),
			})
		}
		if !strings.HasPrefix(string(line), ">") {
			alignmentSeqStore = append(alignmentSeqStore, alignmentSeq{
				seq: string(line),
			})
		}
	}
	csvOpen, err := os.Open(*takecsv)
	if err != nil {
		log.Fatal(err)
	}
	csvRead := bufio.NewScanner(csvOpen)
	magStore := []magStruct{}
	for csvRead.Scan() {
		line := csvRead.Text()
		firstline, _ := strconv.Atoi(strings.Split(string(line), ",")[0])
		secondline, _ := strconv.Atoi(strings.Split(string(line), ",")[1])
		magStore = append(magStore, magStruct{
			start: firstline,
			end:   secondline,
			id:    strings.Split(string(line), ",")[2],
		})
	}

	magExtractID := []string{}
	magExtractSeq := []string{}

	for i := range alignmentIDStore {
		for j := range magStore {
			if alignmentIDStore[i].id == magStore[j].id {
				magExtractID = append(magExtractID, alignmentIDStore[i].id)
				magExtractSeq = append(
					magExtractSeq,
					alignmentSeqStore[i].seq[magStore[j].start:magStore[j].end],
				)
			}
		}
	}

	for i := range magExtractID {
		fmt.Println(">", magExtractID[i], "\n", magExtractSeq[i])
	}
}
