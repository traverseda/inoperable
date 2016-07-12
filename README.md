Some thoughts on a terminal with graphics support, and some attempts to get
familiar with golang.

It follows unix philosophy. The basic idea is that there's a stream of text,
and we pipe it through multiple functions, "simplyfying" it.

```
text(Hello world!)
polygon([0,0],[0,1],[1,1],[1,0])#a square

```

Could be transformed into something like

```
letter(H)letter(e)letter(l)letter(l)letter(o)letter( )
letter(w)letter(o)letter(l)letter(d)letter(!)
line(0,0,0,1)
line(0,1,1,1)
line(1,1,1,0)
line(1,0,0,0)
```

Which would then get passed to a very simple opengl renderer.

Of course there's no reason not to pass polygons directly to the renderer, but
this lets us specify a *very* small minimal set of commands you need to
implement in order to actually render everything cleanly.

If you're more interested in efficiency, you can pull that "polygon()" command
out of the queue before it gets simplified into a series of lines and
triangles.

---

The internal flow is hopefully equally simple.

A stream of commands gets sent to a "router" function. It checks the name
("name(blehblehbleh)") against a list of filters.

When it matches, the function transorms the input and appends it to the
begining or end of the buffer.

```
red(text(Hello World))

```

to

```

saveColor()
color(#FF0000)
text(Hello World)
restoreColor()

```

Some other stuff about a filter that stores function calls.

How do we make history/scrollback work?
