'''
Some tests on a layout-creator
'''

flexBoxContext = {
    'direction':'row',
    'wrap':'scroll',
    'justify':'start',
    'align':'center',
    'xMax':None,
    'yMax':None,
    'xMin':None,
    'yMin':None,
    'xReal':0,
    'yReal':0,
    'children':[],
    'parent':None,
    'depth':0
}

depth=0

def object():
    pass

class WatcherDict(dict):

    def __init__(self, *args, **kwargs):
        super(WatcherDict, self).__init__(*args,**kwargs)
        self.onUpdate = None

    def __setitem__(self, item, value):
        if self.onUpdate:
            self.onUpdate(self, item, value)
        super(WatcherDict, self).__setitem__(item, value)

class FlexBox():
    def __init__(self, *args, **kwargs):
        self.context=WatcherDict()
        self.context.onUpdate = self.onUpdate
        self.context.update(flexBoxContext)
        self.context.update(kwargs)
        for child in self.context['children']:
            child.context['parent'] = self

    def onUpdate(self, d, item, value):
        pass

    def draw(self):
        print("{depth}: {xReal}".format(**self.context))
        for child in self.context['children']:
            child.draw()

    def layout(self):
        global depth
        depth=depth+1
        for child in self.context['children']:
            child.layout()
            child.context['depth']=depth
            self.context['xReal']= self.context['xReal']+child.context['xReal']
            self.context['yReal']= self.context['xReal']+child.context['yReal']
         
        return self

r = FlexBox(children=[FlexBox(xReal=20),FlexBox(xReal=20)])

r.layout().draw()

