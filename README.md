# metagenome-placer

- given a metagenomics MAGs file, which should be the CSV file showing the start and the end position of the starting of the MAGs from the assembled genome.
- given a fasta file of the assembled metagenome. 
- it will extract all the assembled MAGs from the metagenomes. 

```
awk '/^>/ {printf("\n%s\n",$0);next; } { printf("%s",$0);}  \
                         END {printf("\n");}' fastafile > faster.fasta

[gauravsablok@fedora]~/Desktop/codecreatede/go-metagenome-placer% \
go run main.go -csvfile ./sample-files/mags.csv \ 
                -fastafile ./sample-files/multi.fasta
> chr11:66478458
 CCCAGTACAT
> chr11:66478458
 TGTCTAGCCTGGACTGCCGT
> chr11:66478458
 CGCCTCTATGCCTACCACCTGTCCCGTGCCGCCTGGTACG
> chr11:66478458
 CAGACCTCCCCTGAGGCCCCCTACATCTAT

```

Gaurav Sablok
