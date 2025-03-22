# Motivation

I was a little worry that I used to have the same password and its so easy to know what it was, so I make this little program to generate a more secure one.

## Dependencies

You must have a go version > 1.23.0 for some hash algorithms. I dont add it like a dependency because I think its a little invasive if you work with Golang.

## Use

Its easy to use, you have to run `make all` and then run `.\myExecutable algorithm yourPassword`.

You can see the available algorithms with `.\myExecutable -l` and if you forget this file and the above line, you can run `./myExecutable -help`.

By default the executable name is `createYourPassword`

PS: for the HMAC_SHA256 algorithm, you can pass a key, I left one random but you can change it if you want :D
