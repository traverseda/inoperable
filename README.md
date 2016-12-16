#Goals

Termunlator is an attempt at a rough specification for a frawing protocol
intended to replace ANSI terminal emulator specification.

In "curses" mode, it's intented to be competitive with the "html/css/javascript"
stack, but without the focus on aesthetics or experience. That is, you should be
able to use it to build applications with similar *functionality*, but it's not
intended to be pretty or to be an extention of your marketing efforts.

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
Tthis lets us specify a *very* small minimal set of commands you need to
implement in order to actually render everything cleanly.

If you're more interested in efficiency, you can pull that "polygon()" command
out of the queue before it gets simplified into a series of lines and
triangles. The "letter" tag is for demonstration purposes only, and should
 probably never be used. 


#Stream-mode

Your program needs to output a stream of text, with tags that specify things
like containers and colours.

It's easier to tell then show, so here's an example of what a stream would look
like


```
chunk(#(A single "line" as far as our scrollback is concerned)
  color(primary
    text(hello world)
  )
  color(primary-lc
    text(subheading)
  )
)

```

Instead of using ANSI control codes, we use tags, similar to html.

#Colours

We specify a default "pallette" of colors, similar to the ANSI, but instead of
saying "red" or "blue" we specify when the color should be used, similar to
bootstrap.

Colours include

 * Default-bg (default background, generally a dark colour)
 * default-fg (default foreground, generally a bright colour. Normally used for
   fonts)
 * Primary
 * Success
 * Info
 * Warning
 * Danger

as well as a high-contrast (hc) and low-contrast (lc) modifier for each.

#User-input

Passing user-input to your program.

