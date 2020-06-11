from random import randrange

def do_min_cut(edges):
    if edges == None or len(edges) == 0:
        return 0

    vertex_count = _caculate_vertex_count(edges)

    min_cut = len(edges)
    for i in range(vertex_count):
        for j in range(vertex_count):
            tmp_min_cut = _karger_min_cut(edges, vertex_count)
            if tmp_min_cut < min_cut:
                min_cut = tmp_min_cut

    return min_cut

def _caculate_vertex_count(edges):
    vertex_set = set()
    for edge in edges:
        if edge[0] not in vertex_set:
            vertex_set.add(edge[0])
        
        if edge[1] not in vertex_set:
            vertex_set.add(edge[1])
    
    return len(vertex_set)

def _karger_min_cut(edges, vertex_count):
    tmp_edges = edges.copy()
    [vertex_merged_idx_map, rev_vertex_merged_idx_map] = _init_vertex_merged_idx_map(edges)

    merge_count = 0
    while merge_count < (vertex_count-2):
        picked_idx = randrange(len(tmp_edges))
        picked_edge = tmp_edges[picked_idx]

        v1 = tmp_edges[picked_idx][0]
        cur_v1_idx = vertex_merged_idx_map[v1]

        v2 = tmp_edges[picked_idx][1]
        cur_v2_idx = vertex_merged_idx_map[v2]

        if cur_v1_idx != cur_v2_idx:
            cur_idx = vertex_count+merge_count

            for v in rev_vertex_merged_idx_map[cur_v1_idx]:
                vertex_merged_idx_map[v] = cur_idx

            for v in rev_vertex_merged_idx_map[cur_v2_idx]:
                vertex_merged_idx_map[v] = cur_idx

            idx_vertex_group = []
            idx_vertex_group.extend(rev_vertex_merged_idx_map[cur_v1_idx])
            idx_vertex_group.extend(rev_vertex_merged_idx_map[cur_v2_idx])     

            rev_vertex_merged_idx_map[cur_idx] = idx_vertex_group           
            
            del rev_vertex_merged_idx_map[cur_v1_idx]
            del rev_vertex_merged_idx_map[cur_v2_idx]

            merge_count = merge_count+1

        del tmp_edges[picked_idx]

    min_cut = 0
    for edge in tmp_edges:
        v1 = edge[0]
        v2 = edge[1]
        if vertex_merged_idx_map[v1] != vertex_merged_idx_map[v2]:
            min_cut = min_cut+1

    return min_cut

def _init_vertex_merged_idx_map(edges):
    vertex_merged_idx_map = {}
    rev_vertex_merged_idx_map = {}

    for edge in edges:
        v1 = edge[0]
        v2 = edge[1]
        
        if v1 not in vertex_merged_idx_map:
            vertex_merged_idx_map[v1] = v1
        if v1 not in rev_vertex_merged_idx_map:
            rev_vertex_merged_idx_map[v1] = [v1]

        if v2 not in vertex_merged_idx_map:
            vertex_merged_idx_map[v2] = v2
        if v2 not in rev_vertex_merged_idx_map:
            rev_vertex_merged_idx_map[v2] = [v2]
    
    return vertex_merged_idx_map, rev_vertex_merged_idx_map
