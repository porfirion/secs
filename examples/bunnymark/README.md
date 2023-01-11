# Bunnymark

It's a copy of [mizu's Bunnymark](https://github.com/sedyh/mizu/tree/master/examples/bunnymark) ported to secs

- Close all other programs for more accurate results.
- Press the left mouse button to add some amount of gophers.
- Adjust the number of gophers that appear at a time with the mouse wheel.
- Increase the number of gophers until the FPS starts dropping below 60 to find out your result.
- To understand that the drop in performance is not a one-off - use the graphs on the left, they show TPS, FPS and the number of objects over a certain time.
- Press the right mouse button to disable batching, this will greatly increase the load, but keep in mind that all measurements were taken without coloring.

![bunnymark-preview](https://user-images.githubusercontent.com/19890545/149235154-52da3044-363e-491a-a25e-80915c5b8df4.gif)

cpu: Intel(R) Core(TM) i7-3770 CPU @ 3.40GHz

Benchmark add component

    Mizu     	  582746	      2128 ns/op	     722 B/op	      12 allocs/op
    secs    	  322435	      3284 ns/op	     845 B/op	       5 allocs/op

Benchmark systems update (N entities)

| N entities |      mizu |                  |    cesc |                  |
|-----------:|----------:|-----------------:|--------:|-----------------:|
|          0 | 4 456 012 |        285 ns/op | 137 958 |      8 717 ns/op |
|        100 |    25 146 |     48 450 ns/op |  29 215 |     41 137 ns/op |
|      1 000 |     2 672 |    473 408 ns/op |   4 498 |    242 906 ns/op |
|     10 000 |       253 |  4 787 131 ns/op |     426 |  2 977 684 ns/op |
|     25 000 |        91 | 12 568 159 ns/op |     121 |  9 791 713 ns/op |
|     50 000 |        46 | 24 936 761 ns/op |      44 | 26 385 829 ns/op |
|    100 000 |        24 | 52 864 111 ns/op |      16 | 71 209 643 ns/op |
