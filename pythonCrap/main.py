from settings import *
import fileinput

class inputStream():
    def __init__(self, stream=None):
        self.lineStream = stream
        if not stream:
            self.lineStream=fileinput.input()
        self.runeStream = iter(next(self.lineStream))

    def __iter__(self):
        return self
    def __next__(self):
        return self.next()
    def next(self):
        try:
            return next(self.runeStream)
        except StopIteration:
            self.runeStream = iter(next(self.lineStream))
            return(next(self))

def textTag(runeStream, context):
    counter = 1
    for i in runeStream:
        if i == '(':
            counter = counter+1
        if i == ')':
            counter = counter-1
        if counter == 0:
            break
        print(i, end="")
TAGS['text']=textTag

def line(runeStream, context):
    textTag(runeStream, context)
    print()
TAGS['line']=line

def colorTag(runeStream, context, offset = 0):
    color = []
    oldColor = None
    for i in runeStream:
        if i in (" ",":"):
            break
        color.append(i)
    color = "".join(color)
    if color in COLORS:
        oldColor = context['fg-color']
        context['fg-color'] = color
        print(COLORS[color][offset], end="")
        Router(runeStream, context).run()
        print(COLORS[oldColor][offset], end="")
    else:
        Router(runeStream, context).run()
TAGS['color']=colorTag

def backgroundTag(runeStream, context):
    colorTag(runeStream, context, offset=1)

TAGS['background']=backgroundTag

class Router():
    def __init__(self, runeStream, context):
        self.runeStream = runeStream
        self.context = context
        self.tag = []
    def __iter__(self):
        return self
    def __next__(self):
        return self.next()
    def next(self):
        for i in self.runeStream:
            if i in ("(",):
                break
            self.tag.append(i)
        tagString = "".join(self.tag)
        self.tag = []
        if tagString in TAGS:
           TAGS[tagString](self.runeStream, self.context)
        next(self.runeStream)

    def run(self, root=False):
        counter = 1
        for i in self:
            if i == '(':
                counter = counter+1
            if i == ')':
                counter = counter-1
            if counter == 0 and not root:
                break

if __name__ == "__main__":
    router = Router(inputStream(), CONTEXT).run(root=True)
