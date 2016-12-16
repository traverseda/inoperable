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
    'xReal':None,
    'yReal':None,
}

def object():
    pass

class WatcherDict(dict):

    def __init__(self, *args, **kwargs):
        super(WatcherDict, self).__setitem__(*arg,**kwargs)
        self.onUpdate = None

    def __setitem__(self, item, value):
        if self.onUpdate:
            self.onUpdate(self, item, value)
        super(WatcherDict, self).__setitem__(item, value)

class flexBox():
    def __init__(self, parent, *arg, **kwargs):
        self.context=WatcherDict
        self.context.update(flexBoxContext)
        self.context.update(kwargs)
        self.context.onUpdate = self.onUpdate
        self.parent=parent

    def onUpdate(self, d, item, value):
        pass

    def draw(self):
        pass
    def layout(self):
        pass
