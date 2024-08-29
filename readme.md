# How to benchmark in Go

Write an implementation of `wc` and then benchmark that to see if i can make it go faster

## Running

```sh
go test -bench=. -benchmem ./uniq
```

## Output Overview

```sh
➜  go-benchmark git:(main) ✗ go test -bench=. -benchmem ./uniq

goos: darwin
goarch: arm64
pkg: github.com/griggsca91/gobenchmarkexample/uniq
BenchmarkUniq-8            36658             33032 ns/op           42118 B/op         27 allocs/op
PASS
ok      github.com/griggsca91/gobenchmarkexample/uniq   1.728s
```

1st Column is the benchmark test that was ran
2nd Column is the number of times that the benchmark loop ran
3rd Column is the amount of time each loop took to run
4th Column is the amount of bytes per operation
5th column is the amount of allocations per operation

### Thoughts

Can we reduce it?

## Get CPU Profile

```sh
go test -bench=. -benchmem ./uniq -cpuprofile=cpuprofile.out -memprofile=memprofile.out
go tool pprof cpuprofile.out
```

use `top` to show the top functions
use `list <function>` to get a more detailed version

```sh
➜  go-benchmark git:(main) ✗ go tool pprof cpuprofile.out
File: uniq.test
Type: cpu
Time: Aug 29, 2024 at 2:32pm (EDT)
Duration: 1.31s, Total samples = 1.17s (89.26%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 980ms, 83.76% of 1170ms total
Showing top 10 nodes out of 88
      flat  flat%   sum%        cum   cum%
     260ms 22.22% 22.22%      430ms 36.75%  runtime.mapaccess2_faststr
     160ms 13.68% 35.90%      160ms 13.68%  indexbytebody
     140ms 11.97% 47.86%      140ms 11.97%  aeshashbody
     110ms  9.40% 57.26%      300ms 25.64%  strings.genSplit
      90ms  7.69% 64.96%      890ms 76.07%  github.com/griggsca91/gobenchmarkexample/uniq.Uniq
      60ms  5.13% 70.09%       60ms  5.13%  runtime.usleep
      50ms  4.27% 74.36%       50ms  4.27%  runtime.pthread_cond_wait
      40ms  3.42% 77.78%       40ms  3.42%  runtime.pthread_cond_signal
      40ms  3.42% 81.20%       40ms  3.42%  runtime.pthread_kill
      30ms  2.56% 83.76%       60ms  5.13%  runtime.scanobject
(pprof)
```

```sh
(pprof) list Uniq
Total: 1.17s
ROUTINE ======================== github.com/griggsca91/gobenchmarkexample/uniq.BenchmarkUniq in /Users/chris/Documents/go-benchmark/uniq/uniq_test.go
         0      890ms (flat, cum) 76.07% of Total
         .          .     35:func BenchmarkUniq(t *testing.B) {
         .          .     36:   var result string
         .          .     37:   for i := 0; i < t.N; i++ {
         .      890ms     38:           result = Uniq(bigContent)
         .          .     39:   }
         .          .     40:   sink = result
         .          .     41:}
ROUTINE ======================== github.com/griggsca91/gobenchmarkexample/uniq.Uniq in /Users/chris/Documents/go-benchmark/uniq/uniq.go
      90ms      890ms (flat, cum) 76.07% of Total
         .          .      5:func Uniq(content []byte) string {
         .          .      6:   cache := make(map[string]bool)
         .          .      7:
         .          .      8:   var output []string
         .          .      9:
      50ms      360ms     10:   for _, l := range strings.Split(string(content), "\n") {
         .      430ms     11:           if _, ok := cache[l]; !ok {
      40ms       50ms     12:                   output = append(output, l)
         .       40ms     13:                   cache[l] = true
         .          .     14:           }
         .          .     15:   }
         .          .     16:
         .       10ms     17:   return strings.Join(output, "\n")
         .          .     18:}
(pprof)
```

Can do the same thing for memory

```sh
➜  go-benchmark git:(main) ✗ go tool pprof memprofile.out
File: uniq.test
Type: alloc_space
Time: Aug 29, 2024 at 2:35pm (EDT)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top
Showing nodes accounting for 788.54MB, 99.51% of 792.44MB total
Dropped 19 nodes (cum <= 3.96MB)
Showing top 10 nodes out of 16
      flat  flat%   sum%        cum   cum%
  395.58MB 49.92% 49.92%   782.35MB 98.73%  github.com/griggsca91/gobenchmarkexample/uniq.Uniq
  373.41MB 47.12% 97.04%   373.41MB 47.12%  strings.genSplit
   13.36MB  1.69% 98.73%    13.36MB  1.69%  strings.(*Builder).grow
    6.20MB  0.78% 99.51%     6.20MB  0.78%  os.ReadFile
         0     0% 99.51%   782.35MB 98.73%  github.com/griggsca91/gobenchmarkexample/uniq.BenchmarkUniq
         0     0% 99.51%     6.20MB  0.78%  github.com/griggsca91/gobenchmarkexample/uniq.getTestFileContent (inline)
         0     0% 99.51%     6.20MB  0.78%  github.com/griggsca91/gobenchmarkexample/uniq.init
         0     0% 99.51%     6.70MB  0.84%  runtime.doInit (inline)
         0     0% 99.51%     6.70MB  0.84%  runtime.doInit1
         0     0% 99.51%     7.85MB  0.99%  runtime.main
(pprof) list Uniq
Total: 792.44MB
ROUTINE ======================== github.com/griggsca91/gobenchmarkexample/uniq.BenchmarkUniq in /Users/chris/Documents/go-benchmark/uniq/uniq_test.go
         0   782.35MB (flat, cum) 98.73% of Total
         .          .     35:func BenchmarkUniq(t *testing.B) {
         .          .     36:   var result string
         .          .     37:   for i := 0; i < t.N; i++ {
         .   782.35MB     38:           result = Uniq(bigContent)
         .          .     39:   }
         .          .     40:   sink = result
         .          .     41:}
ROUTINE ======================== github.com/griggsca91/gobenchmarkexample/uniq.Uniq in /Users/chris/Documents/go-benchmark/uniq/uniq.go
  395.58MB   782.35MB (flat, cum) 98.73% of Total
         .          .      5:func Uniq(content []byte) string {
         .          .      6:   cache := make(map[string]bool)
         .          .      7:
         .          .      8:   var output []string
         .          .      9:
  130.10MB   503.51MB     10:   for _, l := range strings.Split(string(content), "\n") {
         .          .     11:           if _, ok := cache[l]; !ok {
  149.10MB   149.10MB     12:                   output = append(output, l)
  116.38MB   116.38MB     13:                   cache[l] = true
         .          .     14:           }
         .          .     15:   }
         .          .     16:
         .    13.36MB     17:   return strings.Join(output, "\n")
         .          .     18:}
(pprof)
```

```

## Running a specific benchmark

## Comparing Benchmarks

## Odd findings

### Sinks

## References

<https://codingchallenges.fyi/challenges/challenge-wc>
<https://gobyexample.com/testing-and-benchmarking>
<https://dave.cheney.net/2013/06/30/how-to-write-benchmarks-in-go>
<https://pkg.go.dev/testing#hdr-Benchmarks>
