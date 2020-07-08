def do_kosaraju_scc(edges):
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
        _dfs_round1(g, v, is_vistited1, finish_stack)

    rev_g = _reverse_g(g)


    componments = {}
    is_vistited2 = {}
    while(len(finish_stack) > 0):
        v = finish_stack.pop()
        componment = []
        _dfs_round2(rev_g, v, is_vistited2, componment)
        if len(componment) > 0:
            componment.sort()
            componments[componment[0]] = componment

    return componments

def _dfs_round1(g, cur_v, is_vistited, finish_stack):
    if cur_v not in is_vistited:
        is_vistited[cur_v] = True
        for v in g[cur_v]:
            _dfs_round1(g, v, is_vistited, finish_stack)
        finish_stack.append(cur_v)

def _reverse_g(g):
    rev_g = {}
    for from_v in g:
        if from_v not in rev_g:
            rev_g[from_v] = []
            
        for to_v in g[from_v]:
            if to_v not in rev_g:
                rev_g[to_v] = []
            rev_g[to_v].append(from_v)
    return rev_g

def _dfs_round2(g, cur_v, is_vistited, componment):
    if cur_v not in is_vistited:
        is_vistited[cur_v] = True
        for v in g[cur_v]:
            _dfs_round2(g, v, is_vistited, componment)
        componment.append(cur_v)






