module github.com/EvanJGunn/YAMSG/cli

go 1.21.3

replace github.com/EvanJGunn/YAMSG/lib => ../lib

require (
    github.com/EvanJGunn/YAMSG/lib/emath v1.0
    github.com/EvanJGunn/YAMSG/lib/minecraft v1.0
)