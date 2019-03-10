#!/usr/bin/env python3
# https://www.hackerearth.com/problem/algorithm/chandu-and-his-gurudevs-1/description/
from itertools import repeat

class Vertex:
    def __init__(self, node):
        self.id = node
        self.adjacent = {}

    def __str__(self):
        return str(self.id) + ' adjacent: ' + str([x.id for x in self.adjacent])

    def add_neighbor(self, neighbor, weight=0):
        self.adjacent[neighbor] = weight

    def get_connections(self):
        return self.adjacent.keys()

    def get_id(self):
        return self.id

    def get_weight(self, neighbor):
        return self.adjacent[neighbor]

class Graph:
    def __init__(self):
        self.vert_dict = {}
        self.num_vertices = 0

    def __iter__(self):
        return iter(self.vert_dict.values())

    def add_vertex(self, node):
        self.num_vertices = self.num_vertices + 1
        new_vertex = Vertex(node)
        self.vert_dict[node] = new_vertex
        return new_vertex

    def get_vertex(self, n):
        if n in self.vert_dict:
            return self.vert_dict[n]
        else:
            return None

    def add_edge(self, frm, to, cost = 0):
        if frm not in self.vert_dict:
            self.add_vertex(frm)
        if to not in self.vert_dict:
            self.add_vertex(to)

        self.vert_dict[frm].add_neighbor(self.vert_dict[to], cost)
        self.vert_dict[to].add_neighbor(self.vert_dict[frm], cost)

    def get_vertices(self):
        return self.vert_dict.keys()

def max_path_weight(node1, node2, previous_node = None):
    '''
        In a given graph g, find maximum edge weight in the path from a to b
    '''
    max_weight = 0
    if [previous_node] == node1.get_connections():
        return -1
    if node2 in node1.get_connections():
        return node1.get_weight(node2)
    for node in node1.get_connections():
        if node == previous_node:
            continue
        weight = max_path_weight(node, node2, node1)
        if -1 == weight:
            continue
        max_weight = node1.get_weight(node)
        if max_weight < weight:
            max_weight = weight
    return max_weight


if __name__ == '__main__':

    # number of testcases
    t = int(input())
    for i in repeat(None, t):
        g = Graph()

        # number of nodes
        N = int(input())
        # for i in range(1,N+1):
        #    g.add_vertex(i)

        # add edges
        for i in repeat(None, N-1):
            edge = list(map(int, input().split()))
            g.add_edge(edge[0],edge[1], edge[2])

        sum = 0
        for i in range(1, N+1):
            for j in range(i+1, N+1):
                weight = max_path_weight(g.vert_dict[i], g.vert_dict[j])
                # print("({},{}) = {}".format(i, j, weight))
                sum += weight
        print(sum)

'''
    for v in g:
        for w in v.get_connections():
            vid = v.get_id()
            wid = w.get_id()
            print( '( %s , %s, %3d)'  % ( vid, wid, v.get_weight(w)))

    for v in g:
        print( 'g.vert_dict[%s]=%s' %(v.get_id(), g.vert_dict[v.get_id()]))
'''
