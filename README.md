# gocut

cut cli util written in go

#### single field (default tab delimiter)
```bash
gocut -f 2 sample.tsv

f1
1
6
11
16
21
```
#### single field with custom delimiter
```bash
gocut -f 1 -d , fourchords.csv | head -n 5

Song title
"10000 Reasons (Bless the Lord)"
"20 Good Reasons"
"Adore You"
"Africa"
```
#### list of fields
```bash
gocut -f 1,2 sample.tsv

f0      f1
0       1
5       6
10      11
15      16
20      21
```

```bash
gocut -d , -f "1 2" fourchords.csv | head -n 5

Song title,Artist
"10000 Reasons (Bless the Lord)",Matt Redman and Jonas Myrin
"20 Good Reasons",Thirsty Merc
"Adore You",Harry Styles
"Africa",Toto
```