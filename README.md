# graph-auto
A quick hack to test graph automorphism using ASP/SAT

A model in ASP that would be equivalent to a SAT model, so clasp should be similar to a normal SAT solver. 



```
./graph-iso-asp.sh <graph1> <graph2> <timelimit in seconds>

```

Prerequisites: 
--

* Gringo(version 3.) and Clasp  from http://potassco.sourceforge.net/ 
* zsh, gcc, golang compiler (https://golang.org/dl/). 
* compile both convert.go and graph.c/h 

Example
---

Two graphs that are isomorphic: 

```
./graph-iso-asp.sh CHH/B/CHH_cc\(1-1\)_22.B33  CHH/B/CHH_cc\(1-1\)_22.B46 100  
clasp version 3.1.4
Reading from stdin
Solving...
Answer: 1
map(a21,b10) map(a20,b19) map(a19,b9) map(a18,b20) map(a17,b7) map(a16,b1) map(a15,b21) map(a14,b8) map(a13,b18) map(a12,b2) map(a11,b4) map(a10,b0) map(a9,b12) map(a8,b11) map(a7,b13) map(a6,b3) map(a5,b5) map(a4,b16) map(a3,b15) map(a2,b6) map(a1,b14) map(a0,b17)
SATISFIABLE

Models       : 1+    
Calls        : 1
Time         : 0.131s (Solving: 0.01s 1st Model: 0.01s Unsat: 0.00s)
CPU Time     : 0.100s
Threads      : 10       (Winner: 3)

Choices      : 586     
Conflicts    : 176      (Analyzed: 176)
<etc>
```

```
./graph-iso-asp.sh CHH/B/CHH_cc\(1-1\)_22.B33  CHH/C/CHH_cc\(1-1\)_22.C78  100
clasp version 3.1.4
Reading from stdin
Solving...
UNSATISFIABLE

Models       : 0     
Calls        : 1
Time         : 0.144s (Solving: 0.01s 1st Model: 0.00s Unsat: 0.01s)
CPU Time     : 0.180s
Threads      : 10       (Winner: 8)

Choices      : 1894    
Conflicts    : 1046     (Analyzed: 1041)
Restarts     : 5        (Average: 208.20 Last: 180)
Problems     : 8        (Average Length: 0.00 Splits: 0)
<etc>

```


Tests on CHH instances
--------

Taken from https://sites.google.com/site/giconauto/home/benchmarks

Results with option basic invariant (degree of nodes): -di

size |   pairing  |  result        |   total    |    search  |  clauses   |    conflicts 
-----|------------|----------------|------------|------------|------------|-------------
22   |   B25-B49  |  SATISFIABLE   |   0.119s   |    0.00s   |  62788     |    73 
22   |   B25-B50  |  SATISFIABLE   |   0.130s   |    0.01s   |  62788     |    96 
22   |   B49-B50  |  SATISFIABLE   |   0.109s   |    0.00s   |  62788     |    45 
22   |   B25-C01  |  UNSATISFIABLE |   0.091s   |    0.01s   |  62788     |    659 
22   |   B25-C14  |  UNSATISFIABLE |   0.100s   |    0.01s   |  62788     |    909 
22   |   B25-C55  |  UNSATISFIABLE |   0.077s   |    0.01s   |  62788     |    632 
22   |   B49-C01  |  UNSATISFIABLE |   0.111s   |    0.01s   |  62788     |    665 
22   |   B49-C14  |  UNSATISFIABLE |   0.095s   |    0.01s   |  62788     |    677 
22   |   B49-C55  |  UNSATISFIABLE |   0.084s   |    0.01s   |  62788     |    633 
22   |   B50-C01  |  UNSATISFIABLE |   0.126s   |    0.01s   |  62788     |    976 
22   |   B50-C14  |  UNSATISFIABLE |   0.115s   |    0.01s   |  62788     |    998 
22   |   B50-C55  |  UNSATISFIABLE |   0.087s   |    0.01s   |  62788     |    671 
22   |   C01-C14  |  SATISFIABLE   |   0.124s   |    0.01s   |  62788     |    150 
22   |   C01-C55  |  SATISFIABLE   |   0.106s   |    0.00s   |  62788     |    43 
22   |   C14-C55  |  SATISFIABLE   |   0.139s   |    0.00s   |  62788     |    209 
44   |   B14-B67  |  SATISFIABLE   |   0.405s   |    0.01s   |  321420    |    394 
44   |   B14-B72  |  SATISFIABLE   |   0.444s   |    0.01s   |  321420    |    314 
44   |   B67-B72  |  SATISFIABLE   |   0.455s   |    0.02s   |  321420    |    605 
44   |   B14-C39  |  UNSATISFIABLE |   0.457s   |    0.02s   |  321420    |    535 
44   |   B14-C52  |  UNSATISFIABLE |   0.432s   |    0.02s   |  321420    |    438 
44   |   B14-C83  |  UNSATISFIABLE |   0.406s   |    0.02s   |  321420    |    409 
44   |   B67-C39  |  UNSATISFIABLE |   0.392s   |    0.02s   |  321420    |    511 
44   |   B67-C52  |  UNSATISFIABLE |   0.465s   |    0.03s   |  321420    |    492 
44   |   B67-C83  |  UNSATISFIABLE |   0.398s   |    0.03s   |  321420    |    781 
44   |   B72-C39  |  UNSATISFIABLE |   0.466s   |    0.02s   |  321420    |    405 
44   |   B72-C52  |  UNSATISFIABLE |   0.428s   |    0.02s   |  321420    |    337 
44   |   B72-C83  |  UNSATISFIABLE |   0.455s   |    0.01s   |  321420    |    423 
44   |   C39-C52  |  SATISFIABLE   |   0.397s   |    0.02s   |  321420    |    338 
44   |   C39-C83  |  SATISFIABLE   |   0.449s   |    0.02s   |  321420    |    60 
44   |   C52-C83  |  SATISFIABLE   |   0.406s   |    0.01s   |  321420    |    227 
88   |   B16-B44  |  SATISFIABLE   |   3.469s   |    0.08s   |  2349816   |    1170 
88   |   B16-B75  |  SATISFIABLE   |   3.520s   |    0.10s   |  2349816   |    1452 
88   |   B44-B75  |  SATISFIABLE   |   3.534s   |    0.12s   |  2349816   |    1312 
88   |   B16-C32  |  UNSATISFIABLE |   3.423s   |    0.12s   |  2349816   |    3037 
88   |   B16-C47  |  UNSATISFIABLE |   3.471s   |    0.13s   |  2349816   |    1745 
88   |   B16-C86  |  UNSATISFIABLE |   3.478s   |    0.11s   |  2349816   |    2616 
88   |   B44-C32  |  UNSATISFIABLE |   3.414s   |    0.12s   |  2349816   |    2938 
88   |   B44-C47  |  UNSATISFIABLE |   3.598s   |    0.12s   |  2349816   |    2707 
88   |   B44-C86  |  UNSATISFIABLE |   3.728s   |    0.13s   |  2349816   |    3513 
88   |   B75-C32  |  UNSATISFIABLE |   3.569s   |    0.10s   |  2349816   |    1523 
88   |   B75-C47  |  UNSATISFIABLE |   3.449s   |    0.09s   |  2349816   |    1786 
88   |   B75-C86  |  UNSATISFIABLE |   3.523s   |    0.11s   |  2349816   |    1844 
88   |   C32-C47  |  SATISFIABLE   |   3.441s   |    0.08s   |  2349816   |    761 
88   |   C32-C86  |  SATISFIABLE   |   3.441s   |    0.07s   |  2349816   |    899 
88   |   C47-C86  |  SATISFIABLE   |   3.432s   |    0.10s   |  2349816   |    1629 
132  |   B55-B78  |  SATISFIABLE   |   13.777s  |    0.20s   |  8070900   |    2515 
132  |   B55-B97  |  SATISFIABLE   |   13.603s  |    0.29s   |  8070900   |    4142 
132  |   B78-B97  |  SATISFIABLE   |   13.891s  |    0.33s   |  8070900   |    3893 
132  |   B55-C39  |  UNSATISFIABLE |   14.017s  |    0.37s   |  8070900   |    7612 
132  |   B55-C68  |  UNSATISFIABLE |   13.743s  |    0.36s   |  8070900   |    6762 
132  |   B55-C69  |  UNSATISFIABLE |   13.657s  |    0.28s   |  8070900   |    5622 
132  |   B78-C39  |  UNSATISFIABLE |   13.602s  |    0.32s   |  8070900   |    6251 
132  |   B78-C68  |  UNSATISFIABLE |   13.755s  |    0.34s   |  8070900   |    7881 
132  |   B78-C69  |  UNSATISFIABLE |   13.921s  |    0.39s   |  8070900   |    8064 
132  |   B97-C39  |  UNSATISFIABLE |   13.675s  |    0.39s   |  8070900   |    6903 
132  |   B97-C68  |  UNSATISFIABLE |   13.574s  |    0.37s   |  8070900   |    8921 
132  |   B97-C69  |  UNSATISFIABLE |   14.022s  |    0.39s   |  8070900   |    6860 
132  |   C39-C68  |  SATISFIABLE   |   13.681s  |    0.34s   |  8070900   |    4962 
132  |   C39-C69  |  SATISFIABLE   |   13.751s  |    0.29s   |  8070900   |    3388 
132  |   C68-C69  |  SATISFIABLE   |   13.421s  |    0.26s   |  8070900   |    3812 
198  |   B04-B17  |  SATISFIABLE   |   55.982s  |    0.74s   |  28218492  |    8129 
198  |   B04-B23  |  SATISFIABLE   |   56.462s  |    0.73s   |  28218492  |    9133 
198  |   B17-B23  |  SATISFIABLE   |   55.846s  |    0.71s   |  28218492  |    6952 
198  |   B04-C56  |  UNSATISFIABLE |   59.341s  |    1.75s   |  28218492  |    20758 
198  |   B04-C65  |  UNSATISFIABLE |   56.824s  |    1.39s   |  28218492  |    20915 
198  |   B04-C91  |  UNSATISFIABLE |   57.690s  |    1.85s   |  28218492  |    31561 
198  |   B17-C56  |  UNSATISFIABLE |   59.129s  |    1.88s   |  28218492  |    32136 
198  |   B17-C65  |  UNSATISFIABLE |   57.470s  |    1.98s   |  28218492  |    27111 
198  |   B17-C91  |  UNSATISFIABLE |   57.570s  |    1.76s   |  28218492  |    26085 
198  |   B23-C56  |  UNSATISFIABLE |   57.742s  |    1.83s   |  28218492  |    26578 
198  |   B23-C65  |  UNSATISFIABLE |   57.137s  |    1.60s   |  28218492  |    26569 
198  |   B23-C91  |  UNSATISFIABLE |   57.585s  |    1.78s   |  28218492  |    25471 
198  |   C56-C65  |  SATISFIABLE   |   56.593s  |    0.80s   |  28218492  |    6974 
198  |   C56-C91  |  SATISFIABLE   |   57.739s  |    1.35s   |  28218492  |    14336 
198  |   C65-C91  |  SATISFIABLE   |   57.673s  |    0.79s   |  28218492  |    8145 
264  |   B26-B51  |  SATISFIABLE   |   152.608s |    1.98s   |  65807592  |    11928 
264  |   B26-B81  |  SATISFIABLE   |   145.708s |    1.50s   |  65807592  |    7610 
264  |   B51-B81  |  SATISFIABLE   |   141.619s |    1.42s   |  65807592  |    8245 
264  |   B26-C47  |  UNSATISFIABLE |   141.340s |    3.61s   |  65807592  |    35273 
264  |   B26-C80  |  UNSATISFIABLE |   140.027s |    3.29s   |  65807592  |    30496 
264  |   B26-C86  |  UNSATISFIABLE |   143.356s |    4.18s   |  65807592  |    39207 
264  |   B51-C47  |  UNSATISFIABLE |   145.615s |    4.01s   |  65807592  |    33849 
264  |   B51-C80  |  UNSATISFIABLE |   142.769s |    4.25s   |  65807592  |    37485 
264  |   B51-C86  |  UNSATISFIABLE |   141.805s |    3.80s   |  65807592  |    32986 
264  |   B81-C47  |  UNSATISFIABLE |   144.315s |    3.53s   |  65807592  |    28316 
264  |   B81-C80  |  UNSATISFIABLE |   141.080s |    3.30s   |  65807592  |    30864 
264  |   B81-C86  |  UNSATISFIABLE |   139.179s |    3.38s   |  65807592  |    32983 
264  |   C47-C80  |  SATISFIABLE   |   139.732s |    2.48s   |  65807592  |    18889 
264  |   C47-C86  |  SATISFIABLE   |   138.521s |    2.91s   |  65807592  |    22521 
264  |   C80-C86  |  SATISFIABLE   |   138.158s |    1.88s   |  65807592  |    15135 
352  |   B49-B54  |  SATISFIABLE   |   374.714s |    6.72s   |  156708064 |    34223 
352  |   B49-B98  |  SATISFIABLE   |   380.360s |    7.68s   |  156708064 |    38267 
352  |   B54-B98  |  SATISFIABLE   |   387.728s |    6.21s   |  156708064 |    30496 
352  |   B49-C00  |  UNSATISFIABLE |   379.292s |    11.83s  |  156708064 |    66753 
352  |   B49-C32  |  UNSATISFIABLE |   383.929s |    14.42s  |  156708064 |    85964 
352  |   B49-C98  |  UNSATISFIABLE |   384.351s |    14.46s  |  156708064 |    90083 
352  |   B54-C00  |  UNSATISFIABLE |   391.097s |    10.58s  |  156708064 |    69682 
352  |   B54-C32  |  UNSATISFIABLE |   380.315s |    11.18s  |  156708064 |    69916 
352  |   B54-C98  |  UNSATISFIABLE |   391.107s |    10.58s  |  156708064 |    65030 
352  |   B98-C00  |  UNSATISFIABLE |   381.900s |    10.55s  |  156708064 |    68459 
352  |   B98-C32  |  UNSATISFIABLE |   384.794s |    11.85s  |  156708064 |    82158 
352  |   B98-C98  |  UNSATISFIABLE |   385.514s |    11.13s  |  156708064 |    76809 
352  |   C00-C32  |  SATISFIABLE   |   384.620s |    8.44s   |  156708064 |    45919 
352  |   C00-C98  |  SATISFIABLE   |   381.059s |    9.50s   |  156708064 |    51873 
352  |   C32-C98  |  SATISFIABLE   |   399.930s |    18.19s  |  156708064 |    112825 
440  |   B32-B41  |  SATISFIABLE   |   842.242s |    16.72s  |  307745880 |    56092 
440  |   B32-B85  |  SATISFIABLE   |   846.223s |    24.03s  |  307745880 |    97558 
440  |   B41-B85  |  SATISFIABLE   |   822.549s |    15.90s  |  307745880 |    52127 
440  |   B32-C17  |  UNSATISFIABLE |   828.400s |    31.15s  |  307745880 |    132292 
440  |   B32-C76  |  UNSATISFIABLE |   824.018s |    28.63s  |  307745880 |    124904 
440  |   B32-C79  |  UNSATISFIABLE |   836.383s |    30.48s  |  307745880 |    130574 
440  |   B41-C17  |  UNSATISFIABLE |   838.830s |    36.21s  |  307745880 |    159638 
440  |   B41-C76  |  UNSATISFIABLE |   834.742s |    31.69s  |  307745880 |    136376 
440  |   B41-C79  |  UNSATISFIABLE |   837.989s |    37.68s  |  307745880 |    172121 
440  |   B85-C17  |  UNSATISFIABLE |   838.358s |    26.75s  |  307745880 |    121884 
440  |   B85-C76  |  UNSATISFIABLE |   837.481s |    31.01s  |  307745880 |    135286 
440  |   B85-C79  |  UNSATISFIABLE |   824.427s |    28.26s  |  307745880 |    129540 
440  |   C17-C76  |  SATISFIABLE   |   806.934s |    14.26s  |  307745880 |    50806 
440  |   C17-C79  |  SATISFIABLE   |   832.818s |    18.96s  |  307745880 |    74127 
440  |   C76-C79  |  SATISFIABLE   |   862.868s |    40.67s  |  307745880 |    195965 
550  |   B00-B04  |  UNSATISFIABLE |   1467.159s|    0.00s   |  602575350 |    1 
550  |   B00-B80  |  UNSATISFIABLE |   1424.958s|    0.00s   |  602575350 |    1 
550  |   B04-B80  |  UNSATISFIABLE |   1425.298s|    0.00s   |  602575350 |    1 
550  |   B00-C27  |  UNSATISFIABLE |   1445.255s|    0.00s   |  602575350 |    1 
550  |   B00-C43  |  UNSATISFIABLE |   1410.052s|    0.00s   |  602575350 |    1 
550  |   B00-C70  |  UNSATISFIABLE |   1456.411s|    0.00s   |  602575350 |    1 
550  |   B04-C27  |  UNSATISFIABLE |   1434.254s|    0.00s   |  602575350 |    1 
550  |   B04-C43  |  UNSATISFIABLE |   1431.797s|    0.00s   |  602575350 |    1 
550  |   B04-C70  |  UNSATISFIABLE |   1420.585s|    0.00s   |  602575350 |    1 
550  |   B80-C27  |  UNSATISFIABLE |   1437.137s|    0.00s   |  602575350 |    1 
550  |   B80-C43  |  UNSATISFIABLE |   1427.937s|    0.00s   |  602575350 |    1 
550  |   B80-C70  |  UNSATISFIABLE |   1412.361s|    0.00s   |  602575350 |    1 
550  |   C27-C43  |  UNSATISFIABLE |   1592.838s|    0.00s   |  602575350 |    1 
550  |   C27-C70  |  UNSATISFIABLE |   1464.332s|    0.00s   |  602575350 |    1 
550  |   C43-C70  |  UNSATISFIABLE |   1430.969s|    0.00s   |  602575350 |    1 
660  |   B06-B07  |  UNSATISFIABLE |   2768.681s|    0.00s   |  1046039700|    1 
660  |   B06-B54  |  UNSATISFIABLE |   2710.003s|    0.00s   |  1046039700|    1 
660  |   B07-B54  |  UNSATISFIABLE |   2700.501s|    0.00s   |  1046039700|    1 


Results with option advanced depth invariant: -dd2

size |   pairing  |  result        |   total    |    search  |  clauses   |    conflicts 
-----|------------|----------------|------------|------------|------------|-------------
22 | B27-B44 | SATISFIABLE | 0.000s | 0.00s | 5061 | 2 |
22 | B27-C52 | UNSATISFIABLE | 0.000s | 0.00s | 203 | 1 |
22 | B27-C61 | UNSATISFIABLE | 0.000s | 0.00s | 203 | 1 |
22 | B44-C52 | UNSATISFIABLE | 0.000s | 0.00s | 203 | 1 |
22 | B44-C61 | UNSATISFIABLE | 0.000s | 0.00s | 203 | 1 |
22 | C52-C61 | SATISFIABLE | 0.000s | 0.00s | 5061 | 2 |
44 | B52-B80 | SATISFIABLE | 0.010s | 0.00s | 11074 | 20 |
44 | B52-C03 | UNSATISFIABLE | 0.000s | 0.00s | 439 | 1 |
44 | B52-C86 | UNSATISFIABLE | 0.000s | 0.00s | 439 | 1 |
44 | B80-C03 | UNSATISFIABLE | 0.000s | 0.00s | 439 | 1 |
44 | B80-C86 | UNSATISFIABLE | 0.000s | 0.00s | 439 | 1 |
44 | C03-C86 | SATISFIABLE | 0.070s | 0.00s | 40978 | 6 |
88 | B01-B20 | SATISFIABLE | 0.150s | 0.00s | 89066 | 24 |
88 | B01-C24 | UNSATISFIABLE | 0.000s | 0.00s | 887 | 1 |
88 | B01-C25 | UNSATISFIABLE | 0.000s | 0.00s | 887 | 1 |
88 | B20-C24 | UNSATISFIABLE | 0.000s | 0.00s | 887 | 1 |
88 | B20-C25 | UNSATISFIABLE | 0.000s | 0.00s | 887 | 1 |
88 | C24-C25 | SATISFIABLE | 0.270s | 0.00s | 151442 | 41 |
132 | B54-B99 | SATISFIABLE | 0.880s | 0.00s | 409850 | 27 |
132 | B54-C42 | UNSATISFIABLE | 0.000s | 0.00s | 1447 | 1 |
132 | B54-C68 | UNSATISFIABLE | 0.000s | 0.00s | 1447 | 1 |
132 | B99-C42 | UNSATISFIABLE | 0.000s | 0.00s | 1447 | 1 |
132 | B99-C68 | UNSATISFIABLE | 0.000s | 0.00s | 1447 | 1 |
132 | C42-C68 | SATISFIABLE | 0.640s | 0.00s | 315674 | 38 |
198 | B04-B30 | SATISFIABLE | 3.440s | 0.01s | 1413452 | 54 |
198 | B04-C07 | UNSATISFIABLE | 0.000s | 0.00s | 2221 | 1 |
198 | B04-C55 | UNSATISFIABLE | 0.000s | 0.00s | 2221 | 1 |
198 | B30-C07 | UNSATISFIABLE | 0.000s | 0.00s | 2221 | 1 |
198 | B30-C55 | UNSATISFIABLE | 0.000s | 0.00s | 2221 | 1 |
198 | C07-C55 | SATISFIABLE | 2.870s | 0.00s | 1129196 | 31 |
264 | B55-B70 | SATISFIABLE | 7.760s | 0.01s | 2706818 | 93 |
264 | B55-C19 | UNSATISFIABLE | 0.000s | 0.00s | 3223 | 1 |
264 | B55-C51 | UNSATISFIABLE | 0.000s | 0.00s | 3223 | 1 |
264 | B70-C19 | UNSATISFIABLE | 0.000s | 0.00s | 3223 | 1 |
264 | B70-C51 | UNSATISFIABLE | 0.000s | 0.00s | 3223 | 1 |
264 | C19-C51 | SATISFIABLE | 8.590s | 0.01s | 2896898 | 67 |
352 | B19-B71 | SATISFIABLE | 23.860s | 0.04s | 6518146 | 169 |
352 | B19-C36 | UNSATISFIABLE | 0.000s | 0.00s | 4423 | 1 |
352 | B19-C84 | UNSATISFIABLE | 0.000s | 0.00s | 4423 | 1 |
352 | B71-C36 | UNSATISFIABLE | 0.000s | 0.00s | 4423 | 1 |
352 | B71-C84 | UNSATISFIABLE | 0.000s | 0.00s | 4423 | 1 |
352 | C36-C84 | SATISFIABLE | 22.170s | 0.03s | 6772162 | 61 |
440 | B01-B04 | SATISFIABLE | 49.650s | 0.03s | 14538850 | 68 |
440 | B01-C25 | UNSATISFIABLE | 0.000s | 0.00s | 6007 | 1 |
440 | B01-C90 | UNSATISFIABLE | 0.000s | 0.00s | 6007 | 1 |
440 | B04-C25 | UNSATISFIABLE | 0.000s | 0.00s | 6007 | 1 |
440 | B04-C90 | UNSATISFIABLE | 0.000s | 0.00s | 6007 | 1 |
440 | C25-C90 | SATISFIABLE | 45.080s | 0.04s | 13584994 | 131 |
550 | B06-B60 | SATISFIABLE | 99.300s | 0.07s | 28682752 | 79 |
550 | B06-C31 | UNSATISFIABLE | 0.000s | 0.00s | 7757 | 1 |
550 | B06-C86 | UNSATISFIABLE | 0.000s | 0.00s | 7757 | 1 |
550 | B60-C31 | UNSATISFIABLE | 0.000s | 0.00s | 7757 | 1 |
550 | B60-C86 | UNSATISFIABLE | 0.000s | 0.00s | 7757 | 1 |
550 | C31-C86 | SATISFIABLE | 97.360s | 0.07s | 27091264 | 157 |
660 | B52-B64 | SATISFIABLE | 169.870s | 0.11s | 47120042 | 151 |
660 | B52-C33 | UNSATISFIABLE | 0.000s | 0.00s | 10087 | 1 |
660 | B52-C94 | UNSATISFIABLE | 0.000s | 0.00s | 10087 | 1 |
660 | B64-C33 | UNSATISFIABLE | 0.000s | 0.00s | 10087 | 1 |
660 | B64-C94 | UNSATISFIABLE | 0.000s | 0.00s | 10087 | 1 |
660 | C33-C94 | SATISFIABLE | 165.490s | 0.11s | 47597834 | 160 |
792 | B29-B45 | SATISFIABLE | 327.200s | 0.26s | 82095554 | 289 |
792 | B29-C00 | UNSATISFIABLE | 0.010s | 0.00s | 12535 | 1 |
792 | B29-C48 | UNSATISFIABLE | 0.010s | 0.00s | 12535 | 1 |
792 | B45-C00 | UNSATISFIABLE | 0.010s | 0.00s | 12535 | 1 |
792 | B45-C48 | UNSATISFIABLE | 0.010s | 0.00s | 12535 | 1 |
792 | C00-C48 | SATISFIABLE | 313.920s | 0.21s | 82669250 | 577 |
924 | B38-B64 | SATISFIABLE | 554.990s | 0.49s | 141589970 | 226 |
924 | B38-C33 | UNSATISFIABLE | 0.010s | 0.00s | 15799 | 1 |
924 | B38-C98 | UNSATISFIABLE | 0.010s | 0.00s | 15799 | 1 |
924 | B64-C33 | UNSATISFIABLE | 0.010s | 0.00s | 15799 | 1 |
924 | B64-C98 | UNSATISFIABLE | 0.010s | 0.00s | 15799 | 1 |
924 | C33-C98 | SATISFIABLE | 618.250s | 0.33s | 138241970 | 196 |
1078 | B44-B69 | SATISFIABLE | 964.390s | 0.42s | 226255444 | 141 |
1078 | B44-C26 | UNSATISFIABLE | 0.010s | 0.00s | 19117 | 1 |
1078 | B44-C98 | UNSATISFIABLE | 0.010s | 0.00s | 19117 | 1 |
1078 | B69-C26 | UNSATISFIABLE | 0.010s | 0.00s | 19117 | 1 |
1078 | B69-C98 | UNSATISFIABLE | 0.020s | 0.00s | 19117 | 1 |
1078 | C26-C98 | SATISFIABLE | 994.120s | 1.83s | 221566516 | 5959 |

Tests on Funkybee Instances. 
--------

Tested on instances from http://funkybee.narod.ru/graphs.htm

```
t01 SATISFIABLE CPU Time : 0.010s
t02 SATISFIABLE CPU Time : 0.020s
t03 SATISFIABLE CPU Time : 0.090s
t04 SATISFIABLE CPU Time : 0.120s
t05 SATISFIABLE CPU Time : 0.340s
t06 UNSATISFIABLE CPU Time : 0.190s
t07 UNSATISFIABLE CPU Time : 0.070s
t08 UNSATISFIABLE CPU Time : 0.790s
t09 UNSATISFIABLE CPU Time : 9.960s
t10 UNSATISFIABLE CPU Time : 0.010s
t11 UNSATISFIABLE CPU Time : 4.100s
t12 UNSATISFIABLE CPU Time : 0.830s
t13 UNSATISFIABLE CPU Time : 125.310s
t14 UNSATISFIABLE CPU Time : 33.790s
t15 UNSATISFIABLE CPU Time : 4.230s
t16 UNSATISFIABLE CPU Time : 9.780s
t17 UNSATISFIABLE CPU Time : 9.450s
t18 UNSATISFIABLE CPU Time : 9.180s
t19 UNSATISFIABLE CPU Time : 9.090s
t20 UNSATISFIABLE CPU Time : 0.040s
t21 UNSATISFIABLE CPU Time : 2.270s
t22 UNSATISFIABLE CPU Time : 1.780s
t23 UNSATISFIABLE CPU Time : 1.870s
t24 UNSATISFIABLE CPU Time : 0.420s
t25 UNSATISFIABLE CPU Time : 0.040s
t26 UNSATISFIABLE CPU Time : 5.340s
t27 UNSATISFIABLE CPU Time : 5.250s
t28 UNSATISFIABLE CPU Time : 5.220s
t29 UNSATISFIABLE CPU Time : 0.040s
t30 UNSATISFIABLE CPU Time : 0.030s
t31 UNSATISFIABLE CPU Time : 0.000s
t32 UNSATISFIABLE CPU Time : 2.770s
t33 SATISFIABLE CPU Time : 0.020s
t34 UNSATISFIABLE CPU Time : 1.170s
t35 SATISFIABLE CPU Time : 1.100s
t36 SATISFIABLE CPU Time : 1.020s
t37 UNSATISFIABLE CPU Time : 0.190s
t38 UNSATISFIABLE CPU Time : 0.000s
t39 UNSATISFIABLE CPU Time : 0.000s
```
size | pairing | result | total | search | clauses | conflicts |
