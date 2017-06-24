all: target.go

target.go:
	../../macro_preprocessor.py -m macros.go -i source.go -o target.go

clean:
	$(RM) target.go
