# chronicling-america-ui

User interface based on the chronicling america api.

> About the Site and API « Chronicling America « Library of Congress  
> https://chroniclingamerica.loc.gov/about/api/#search

![Imgur](https://i.imgur.com/KLVMw8L.png)

# How to use

```
# Run
$ ./buildtool -m run
[ INFO ] >>> update assets.go ...
[ INFO ] >>> start appliaction ...
INFO[2020-03-20T07:44:16+09:00]/Users/user/Develop/go/chronicling-america-ui/main.go:59 main.main() ====== Start Application ======
INFO[2020-03-20T07:44:19+09:00]/Users/user/Develop/go/chronicling-america-ui/main.go:33 main.main.func1() ====== Exit Application ======


# Build application
$ ./buildtool -m build -o /tmp -t mac
[ INFO ] Build start for MacOS
[ INFO ] >>> make output directory ...
[ INFO ] >>> update assets.go ...
[ INFO ] >>> build application ...
[ INFO ] >>> success : /tmp/ChroniclingAmericaUI.app
```
