# Motivation

I was a little worry that I used to have the same password and its so easy to know what it was, so I make this little program to generate a more secure one.

## Dependencies

You must have a go version > 1.23.0 for some hash algorithms. I dont add it like a dependency because I think its a little invasive if you work with Golang.

## Use

Its easy to use, you have to run `make all` and then run 

`./myExecutable algorithm yourPassword maxLength optional[numberOfRotations]`.

Explanation of program's args:

`algorithm` is one of the algorithms available in this program, you can see they running 

`./myExecutable -l`

`yourPassword` its ... the password you want to hash and making it more secure!

`maxLength` is the length of your password, I add it because usually the social medias and stuffs like that, add a limit to the length of your password, so you can set it directly.

`numberOfRotations` is an optional arg if you want to rotate your hash password so its a way to make a new combination for your password.

If you forget this file and the above lines, you can run `./myExecutable -h`.

By default the executable name is `createYourPassword`

PS: for the HMAC_SHA256 algorithm, you can pass a key, I left one random but you can change it if you want :D
