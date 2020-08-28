from random import randrange

def do_min_cut(edges):
    if edges == None or len(edges) == 0:
        return 0

    vertex_count = _count_vertex(edges)
    min_cut = len(edges)
    for i in range(vertex_count*vertex_count):
        tmp_min_cut = _do_karger_min_cut(edges)
        if tmp_min_cut < min_cut:
            min_cut = tmp_min_cut

    return min_cut

def _do_karger_min_cut(edges):
    clone_edges = edges.copy()

    max_vertex, vertex_mergeid_map, mergeid_vertex_map = _init_vertex_mergeid_map(edges)

    merge_count = 0
    while merge_count < len(vertex_mergeid_map) - 2:
        pick_idx = randrange(len(clone_edges))
        pick_edge = clone_edges[pick_idx]
        v0 = pick_edge[0]
        m0 = vertex_mergeid_map[v0]
        v1 = pick_edge[1]
        m1 = vertex_mergeid_map[v1]

        if m0 != m1:
            merge_count += 1
            cur_mergeid = merge_count + max_vertex

            for v in mergeid_vertex_map[m0]:
                vertex_mergeid_map[v] = cur_mergeid
            for v in mergeid_vertex_map[m1]:
                vertex_mergeid_map[v] = cur_mergeid

            merge_group0 = mergeid_vertex_map[m0]
            merge_group1 = mergeid_vertex_map[m1]
            merge_group0.extend(merge_group1)

            mergeid_vertex_map[cur_mergeid] = merge_group0
            
            del mergeid_vertex_map[m0]
            del mergeid_vertex_map[m1]

        del clone_edges[pick_idx]

    min_cut = 0
    for edge in clone_edges:
        v0 = edge[0]
        m0 = vertex_mergeid_map[v0]
        v1 = edge[1]
        m1 = vertex_mergeid_map[v1]
        if m0 != m1:
            min_cut+=1

    return min_cut



def _init_vertex_mergeid_map(edges):
    max_vertex = 0
    vertex_mergeid_map = {}
    mergeid_vertex_map = {}

    for edge in edges:
        v0 = edge[0]
        v1 = edge[1]

        vertex_mergeid_map[v0] = v0
        vertex_mergeid_map[v1] = v1

        mergeid_vertex_map[v0] = [v0]
        mergeid_vertex_map[v1] = [v1]

        if v0>max_vertex:
            max_vertex = v0
        if v1>max_vertex:
            max_vertex = v1
    return max_vertex, vertex_mergeid_map, mergeid_vertex_map

def _count_vertex(edges):
    vertex_map = {}
    for edge in edges:
        vertex_map[edge[0]] = True
        vertex_map[edge[1]] = True

    return len(vertex_map)