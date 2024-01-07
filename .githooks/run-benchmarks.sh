bazel run //src/board:board_test -- --test.bench=. 2>&1 | grep Benchmark
bazel run //src/chess:chess_test -- --test.bench=. 2>&1 | grep Benchmark
bazel run //src/peace:peace_test -- --test.bench=. 2>&1 | grep Benchmark