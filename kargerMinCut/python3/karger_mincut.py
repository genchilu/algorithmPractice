from random import randrange

def do_min_cut(edges):
    if edges == None or len(edges) == 0:
        return 0

    vertex_count = _count_vertex(edges)
    min_cut = len(edges)
    for i in range(vertex_count*vertex_count):
        cur_min_cut = _karger_min_cut(edges)
        if cur_min_cut < min_cut:
            min_cut = cur_min_cut

    return min_cut


def _count_vertex(edges):
    vertex_set = {}
    for edge in edges:
        if edge[0] not in vertex_set:
            vertex_set[edge[0]] = True
        if edge[1] not in vertex_set:
            vertex_set[edge[1]] = True

    return len(vertex_set)

def _karger_min_cut(edges):
    clone_edges = edges.copy()
    
    [max_vertext, vertex_mergeid_map, mergeid_vertices_map] = init_vertex_mergeid_map(edges)

    merge_count = 0
    while merge_count < len(vertex_mergeid_map) -2:
        picked_idx = randrange(len(clone_edges))
        picked_edge = clone_edges[picked_idx]

        v0 = picked_edge[0]
        mergeid0 = vertex_mergeid_map[v0]

        v1 = picked_edge[1]
        mergeid1 = vertex_mergeid_map[v1]

        if mergeid0 != mergeid1:
            merge_count = merge_count+1
            cur_mergeid = merge_count+max_vertext

            mergeid0_vertices = mergeid_vertices_map[mergeid0]
            mergeid1_vertices = mergeid_vertices_map[mergeid1]

            for v in mergeid0_vertices:
                vertex_mergeid_map[v] = cur_mergeid
            for v in mergeid1_vertices:
                vertex_mergeid_map[v] = cur_mergeid

            mergeid0_vertices.extend(mergeid1_vertices)
            mergeid_vertices_map[cur_mergeid] = mergeid0_vertices

            del mergeid_vertices_map[mergeid0]
            del mergeid_vertices_map[mergeid1]

        del clone_edges[picked_idx]

    min_cut = 0
    for edge in clone_edges:
        mergeid0 = vertex_mergeid_map[edge[0]]
        mergeid1 = vertex_mergeid_map[edge[1]]

        if mergeid0 != mergeid1:
            min_cut = min_cut+1

    return min_cut

def init_vertex_mergeid_map(edges):
    vertex_mergeid_map = {}
    mergeid_vertices_map = {}
    max_vertext = 0

    for edge in edges:
        if edge[0] not in vertex_mergeid_map:
            vertex_mergeid_map[edge[0]] = edge[0]
        if edge[1] not in vertex_mergeid_map:
            vertex_mergeid_map[edge[1]] = edge[1]

        if edge[0] not in mergeid_vertices_map:
            mergeid_vertices_map[edge[0]] = [edge[0]]
        if edge[1] not in mergeid_vertices_map:
            mergeid_vertices_map[edge[1]] = [edge[1]]

        if edge[0] > max_vertext:
            max_vertext = edge[0]
        if edge[1] > max_vertext:
            max_vertext = edge[1]

    return max_vertext, vertex_mergeid_map, mergeid_vertices_map
