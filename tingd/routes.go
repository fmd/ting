package main

func (d* Daemon) Routes() {
    d.Martini.Get("/", func() string {
        return "Hello, world!"
    })
}