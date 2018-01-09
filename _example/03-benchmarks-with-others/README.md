# Benchmarks
Benchmarks with others

## Installation
```
$ go get github.com/apcera/termtables
$ go get github.com/gosuri/uitable
```

## Show
```
$ cd $GOPATH/src/github.com/alexeyco/simpletable/_example/03-benchmarks-with-others
$ go run main.go
```
Result:
```
github.com/alexeyco/simpletable:
+----+------------------+--------------+-----------------------------+------+
| #  |       NAME       |    PHONE     |            EMAIL            | QTTY |
+----+------------------+--------------+-----------------------------+------+
|  1 | Newton G. Goetz  | 252-585-5166 | NewtonGGoetz@dayrep.com     |   10 |
|  2 | Rebecca R. Edney | 865-475-4171 | RebeccaREdney@armyspy.com   |   12 |
|  3 | John R. Jackson  | 810-325-1417 | JohnRJackson@armyspy.com    |   15 |
|  4 | Ron J. Gomes     | 217-450-8568 | RonJGomes@rhyta.com         |   25 |
|  5 | Penny R. Lewis   | 870-794-1666 | PennyRLewis@rhyta.com       |    5 |
|  6 | Sofia J. Smith   | 770-333-7379 | SofiaJSmith@armyspy.com     |    3 |
|  7 | Karlene D. Owen  | 231-242-4157 | KarleneDOwen@jourrapide.com |   12 |
|  8 | Daniel L. Love   | 978-210-4178 | DanielLLove@rhyta.com       |   44 |
|  9 | Julie T. Dial    | 719-966-5354 | JulieTDial@jourrapide.com   |    8 |
| 10 | Juan J. Kennedy  | 908-910-8893 | JuanJKennedy@dayrep.com     |   16 |
+----+------------------+--------------+-----------------------------+------+
|                                                           Subtotal |  150 |
+----+------------------+--------------+-----------------------------+------+
github.com/apcera/termtables:
+----+------------------+--------------+-----------------------------+------+
| #  | NAME             | PHONE        | EMAIL                       | QTTY |
+----+------------------+--------------+-----------------------------+------+
| 1  | Newton G. Goetz  | 252-585-5166 | NewtonGGoetz@dayrep.com     | 10   |
| 2  | Rebecca R. Edney | 865-475-4171 | RebeccaREdney@armyspy.com   | 12   |
| 3  | John R. Jackson  | 810-325-1417 | JohnRJackson@armyspy.com    | 15   |
| 4  | Ron J. Gomes     | 217-450-8568 | RonJGomes@rhyta.com         | 25   |
| 5  | Penny R. Lewis   | 870-794-1666 | PennyRLewis@rhyta.com       | 5    |
| 6  | Sofia J. Smith   | 770-333-7379 | SofiaJSmith@armyspy.com     | 3    |
| 7  | Karlene D. Owen  | 231-242-4157 | KarleneDOwen@jourrapide.com | 12   |
| 8  | Daniel L. Love   | 978-210-4178 | DanielLLove@rhyta.com       | 44   |
| 9  | Julie T. Dial    | 719-966-5354 | JulieTDial@jourrapide.com   | 8    |
| 10 | Juan J. Kennedy  | 908-910-8893 | JuanJKennedy@dayrep.com     | 16   |
|    |                  |              | Subtotal                    | 150  |
+----+------------------+--------------+-----------------------------+------+

github.com/gosuri/uitable:
#       NAME                    PHONE           EMAIL                           QTTY
1       Newton G. Goetz         252-585-5166    NewtonGGoetz@dayrep.com         10
2       Rebecca R. Edney        865-475-4171    RebeccaREdney@armyspy.com       12
3       John R. Jackson         810-325-1417    JohnRJackson@armyspy.com        15
4       Ron J. Gomes            217-450-8568    RonJGomes@rhyta.com             25
5       Penny R. Lewis          870-794-1666    PennyRLewis@rhyta.com           5
6       Sofia J. Smith          770-333-7379    SofiaJSmith@armyspy.com         3
7       Karlene D. Owen         231-242-4157    KarleneDOwen@jourrapide.com     12
8       Daniel L. Love          978-210-4178    DanielLLove@rhyta.com           44
9       Julie T. Dial           719-966-5354    JulieTDial@jourrapide.com       8
10      Juan J. Kennedy         908-910-8893    JuanJKennedy@dayrep.com         16
                                                Subtotal                        150
```

## Run benchmarks
```
$ cd $GOPATH/src/github.com/alexeyco/simpletable/_example/03-benchmarks-with-others/bench
$ go test -bench=. -benchmem
```
## Results
Intel Core i5 (3470), 8Gb DDR3 memory, Windows 10 results:
```
BenchmarkSimpletable-4             10000            186476 ns/op           45651 B/op       1330 allocs/op
BenchmarkTermtables-4              10000            130461 ns/op           27632 B/op       1164 allocs/op
BenchmarkUITable-4                 10000            229150 ns/op           39425 B/op       1073 allocs/op
PASS
ok      github.com/alexeyco/simpletable/_example/03-benchmarks-with-others/bench        5.559s
```
