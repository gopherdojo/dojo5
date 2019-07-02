## 課題3-2: godl - Simple Downloader written in Go

#### Usage

```
$ make build
$ ./godl -url url [-r set num of goroutine] [-o output filename]
```

##### Option

| flag | Usage |
|:----|:---- 
| `-url url` | Set URL for download |
| `-r 2` | Set number of goroutine [default: 2] |
| `-o filename` | Set output filename [default: output] |

### Test

```
$ make test
```

