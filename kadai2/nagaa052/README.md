gocon
=== 
image convert command by golang

### Usage

```
Usage:
  gocon [option] <directory path>
Options:
  -d string
    	Convert output directory. (default "out")
  -f string
    	Convert image format. The input format is [In]:[Out]. image is jpeg|png. (default "jpeg:png")
```

#### Example
```
$ gocon -f jpeg:png -d output image/jpegs
```
