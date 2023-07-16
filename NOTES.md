# Perf

Start

```
Benchmark_GenMovesStart-16                        	  125871	      9042 ns/op
Benchmark_GenMovesStart_D4-16                     	  106768	     11348 ns/op
Benchmark_GenMovesStart_D4_E5-16                  	   99714	     12061 ns/op
Benchmark_GenMovesStart_D4_E5_DXE5-16             	   71361	     15536 ns/op
Benchmark_GenMovesStart_D4_E5_DXE5_D6-16          	   95257	     13324 ns/op
Benchmark_GenMovesStart_D4_E5_DXE5_D6_E4-16       	   65515	     17574 ns/op
Benchmark_GenMovesStart_D4_E5_DXE5_D6_E4_D5-16    	   65025	     15564 ns/op
```

After

```
Benchmark_GenMovesStart-16                        	  200866	      5861 ns/op
Benchmark_GenMovesStart_D4-16                     	  197935	      5881 ns/op
Benchmark_GenMovesStart_D4_E5-16                  	  143457	      8759 ns/op
Benchmark_GenMovesStart_D4_E5_DXE5-16             	  144585	      8446 ns/op
Benchmark_GenMovesStart_D4_E5_DXE5_D6-16          	  132219	      9226 ns/op
Benchmark_GenMovesStart_D4_E5_DXE5_D6_E4-16       	  125929	      9818 ns/op
Benchmark_GenMovesStart_D4_E5_DXE5_D6_E4_D5-16    	  110126	     10976 ns/op
Benchmark_GenMove_Interesting-16                  	  141762	      8727 ns/op


Benchmark_GenMovesStart-16                        	  185793	      6573 ns/op
Benchmark_GenMovesStart_D4-16                     	  182054	      6283 ns/op
Benchmark_GenMovesStart_D4_E5-16                  	  132356	      9005 ns/op
Benchmark_GenMovesStart_D4_E5_DXE5-16             	  128660	      8907 ns/op
Benchmark_GenMovesStart_D4_E5_DXE5_D6-16          	  118315	      9790 ns/op
Benchmark_GenMovesStart_D4_E5_DXE5_D6_E4-16       	  114098	     10352 ns/op
Benchmark_GenMovesStart_D4_E5_DXE5_D6_E4_D5-16    	   96637	     12435 ns/op
Benchmark_GenMove_Interesting-16                  	  128460	      9240 ns/op
Benchmark_GenMove_FromStart_Depth4-16             	      13	  90422801 ns/op	   2685839 nodes	    206603 nodes/op	    1552 B/op	       7 allocs/op

```

