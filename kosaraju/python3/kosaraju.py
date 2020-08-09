def do_kosaraju_scc(edges):
    if edges == None or len(edges) == 0:
        return {}

    g = {}
    for edge in edges:
        from_v = edge[0]
        to_v = edge[1]
        if from_v not in g:
            g[from_v] = []
        if to_v not in g:
            g[to_v] = []
        g[from_v].append(to_v)

    is_vistited1 = {}
    finish_stack = []
    for v in g:
        _dfs1(g, v, is_vistited1, finish_stack)

    rev_g = _rever_g(g)

    is_vistited2 = {}
    componments = {}
    while len(finish_stack) > 0:
        v = finish_stack[-1]
        finish_stack = finish_stack[0:-1]
        componment = []
        _dfs2(rev_g, v, is_vistited2, componment)

        if len(componment) > 0:
            componment.sort()
            componments[componment[0]] = componment
    
    return componments
        




def _dfs1(g, v, is_vistited, finish_stack):
    if v not in is_vistited:
        is_vistited[v] = True
        for to_v in g[v]:
            _dfs1(g, to_v, is_vistited, finish_stack)
        finish_stack.append(v)

def _rever_g(g):
    rev_g = {}
    for from_v in g:
        if from_v not in rev_g:
            rev_g[from_v] = []
        for to_v in g[from_v]:
            if to_v not in rev_g:
                rev_g[to_v] = []
            rev_g[to_v].append(from_v) 
    return rev_g
        
def _dfs2(g, v, is_vistited, componment):
    if v not in is_vistited:
        is_vistited[v] = True
        for to_v in g[v]:
            _dfs2(g, to_v, is_vistited, componment)
        componment.append(v)





