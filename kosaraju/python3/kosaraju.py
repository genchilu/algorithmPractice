def do_kosaraju_scc(edges):

    g = {}
    for edge in edges:
        if edge[0] not in g:
            g[edge[0]] = []
        if edge[1] not in g:
            g[edge[1]] = []
        g[edge[0]].append(edge[1])

    is_vistited = {}
    finish_stack = []
    
    for vertex in g:
        _dfs1(g, vertex, is_vistited, finish_stack)

    gt = _reverse_graph(g)

    is_vistited2 = {}
    componments = {}
    while len(finish_stack) > 0:
        componment = []
        _dfs2(gt, finish_stack.pop(), is_vistited2, componment)
        
        if(len(componment) > 0):
            componment.sort()
            componments[componment[0]] = componment


    return componments


def _reverse_graph(g):
    gt = {}
    for from_vertex in g:
        if from_vertex not in gt:
            gt[from_vertex] = []
        for to_vertex in g[from_vertex]:
            if to_vertex not in gt:
                gt[to_vertex] = []
            gt[to_vertex].append(from_vertex)
    return gt


def _dfs1(g, cur_vertex, is_vistited, finish_stack):
    if cur_vertex not in is_vistited or not is_vistited[cur_vertex]:
        is_vistited[cur_vertex] = True

        for neighbor in g[cur_vertex]:
            _dfs1(g, neighbor, is_vistited, finish_stack)
        
        finish_stack.append(cur_vertex)

def _dfs2(gt, cur_vertex, is_vistited, componment):
    if cur_vertex not in is_vistited or not is_vistited[cur_vertex]:
        is_vistited[cur_vertex] = True
        componment.append(cur_vertex)
        for neighbor in gt[cur_vertex]:
            _dfs2(gt, neighbor, is_vistited, componment)


