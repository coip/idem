# idem

Hi there!

This is a quick n dirty means to provide a lib in support of a story similar to:

### As a Database Administrator, I would like to know specific details regarding applications connecting to and running queries against my Database:
 - the binary name
 - the binarys latest commit SHA @ build/compilation time
 - the host executing said binary


-----

`pkg idem` is meant to allow for ~DX akin to `var dsn = "ServerName= ...; Username= ...; etc ...; AppName=" + idem.GetCurrentMainGopher() ...`

see [v0.0.2](https://github.com/coip/idem/releases/tag/v0.0.2)
