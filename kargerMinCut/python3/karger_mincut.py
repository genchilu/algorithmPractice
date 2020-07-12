from random import randrange

def do_min_cut(edges):
    if edges == None or len(edges) == 0:
        return 0

    vertex_count = _count_vertex(edges)
    
    min_cut = len(edges)
    for i in range(vertex_count*vertex_count):
        tmp_min_cut = _karger_min_cut(edges)
        if tmp_min_cut < min_cut:
            min_cut = tmp_min_cut

    return min_cut

def _karger_min_cut(edges):
    clone_edges = edges.copy()

    vertex_mergeidx_map, mergeidx_vertex_map, max_vertex = _init_vertex_mergeidx_map(clone_edges)
    merge_count = 0
    while merge_count < (len(vertex_mergeidx_map) -2):
        pick_idx = randrange(len(clone_edges))
        pick_edge = clone_edges[pick_idx]
        v0 = pick_edge[0]
        v1 = pick_edge[1]
        mergeidx0 = vertex_mergeidx_map[v0]
        mergeidx1 = vertex_mergeidx_map[v1]

        if mergeidx0 != mergeidx1 :
            merge_count=merge_count+1
            cur_mergeidx_idx = merge_count + max_vertex

            for v in mergeidx_vertex_map[mergeidx0]:
                vertex_mergeidx_map[v] = cur_mergeidx_idx
            for v in mergeidx_vertex_map[mergeidx1]:
                vertex_mergeidx_map[v] = cur_mergeidx_idx

            mergeidx_group0 = mergeidx_vertex_map[mergeidx0]
            mergeidx_group1 = mergeidx_vertex_map[mergeidx1]

            mergeidx_group0.extend(mergeidx_group1)
            mergeidx_vertex_map[cur_mergeidx_idx] = mergeidx_group0

            del mergeidx_vertex_map[mergeidx0]
            del mergeidx_vertex_map[mergeidx1]

        del clone_edges[pick_idx]

    min_cut = 0
    for edge in clone_edges:
        v0 = edge[0]
        v1 = edge[1]
        if vertex_mergeidx_map[v0] != vertex_mergeidx_map[v1]:
            min_cut = min_cut+1

    return min_cut



def _init_vertex_mergeidx_map(edges):
    vertex_mergeidx_map = {}
    mergeidx_vertex_map ={}
    max_vertex = 0

    for edge in edges:
        v0 = edge[0]
        v1 = edge[1]

        if v0 not in vertex_mergeidx_map:
            vertex_mergeidx_map[v0] = v0
        if v1 not in vertex_mergeidx_map:
            vertex_mergeidx_map[v1] = v1
        
        if v0 not in mergeidx_vertex_map:
            mergeidx_vertex_map[v0] = [v0]
        if v1 not in mergeidx_vertex_map:
            mergeidx_vertex_map[v1] = [v1]

        if v0 > max_vertex:
            max_vertex = v0
        if v1 > max_vertex:
            max_vertex = v1

    return vertex_mergeidx_map, mergeidx_vertex_map, max_vertex

def _count_vertex(edges):
    vertex_set = set()
    for edge in edges:
        vertex_set.add(edge[0])
        vertex_set.add(edge[1])
    return len(vertex_set)