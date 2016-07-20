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

##Color

```
red(text(Hello World))
---
SaveColor()
SetColor(#F00)
text(Hello World)
RestoreColor()

```

#Performance

This is my first major project in a statically typed language. It's go largely
because I think I'm less likely to do something *horribly* wrong in golang then
in C.

That being said, it was created with performance in mind. I don't know how
golangs internals work, but it makes heavy use of channels. Hopefully that
minimizes cache mis-predicts, and can create some very fast code.

It doesn't make good use of multithreading, but our routing functions should be
able to handle it at some point in the future.
