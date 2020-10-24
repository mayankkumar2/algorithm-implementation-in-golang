// Dijkstras algorithm - C++ implementation

// TO BE PORTED TO GO

#include <iostream>
#include <vector>

class node {
public:
    int distance = INT16_MAX;
    int id;
    bool is_visited_flag;
    explicit node(int _id) :  id(_id) {}

    bool is_visited(){
        return is_visited_flag;
    }

    void set_visited(bool value) {
        is_visited_flag = value;
    }

    bool updateDistance(int d) {
//        std::cout << id << " " << distance<< " "<< d << "|\n";
        if (distance > d) {
            distance = d;
            return true;
        } else {
            return false;
        }
    }

    struct node_pointer_cmp {
        bool operator()(node *p, node *p2) {
            return p->distance > p2->distance;
        }
    };

};


using namespace std;
int main() {

    vector<node> q;
    vector<node*> pq;
    const int NODE_COUNT = 7;
    vector<vector<int>> adj_list {
            {1,2},
            {0,2,4},
            {0,1,3},
            {2},
            {1,5,6},
            {4,6},
            {4,5},
    };

    vector<vector<int>> dist_list {
            {10,80},
            {10,6,70},
            {80, 6, 70},
            {70},
            {20, 10, 50},
            {50,5},
            {10, 5}

    };

    for (int i = 0; i < NODE_COUNT; ++i ) q.push_back(node(i));

    for (auto& e: q) pq.push_back(&e);

    q[0].updateDistance(0);
    make_heap(pq.begin(), pq.end(), node::node_pointer_cmp());


    for (int i = 0 ; i < NODE_COUNT; ++i) {
        auto top = pq.front();
//        std::cout << "top" << " " << top->distance << " " << top->id << endl;
        if (top->is_visited()) {
            pop_heap(pq.begin(), pq.end(), node::node_pointer_cmp());
            pq.pop_back();
            continue;
        }
        top->set_visited(true);
        auto& list_order = adj_list[top->id];
        auto& dist_order = dist_list[top->id];

        for (int j = 0 ; j < list_order.size(); ++j) {

//            cout << "dist: " << top->distance << " = " << dist_order[j]<< " = " << list_order[j] << " = " <<  endl;
            q[list_order[j]].updateDistance(
                    dist_order[j] + top->distance);
        }
        pop_heap(pq.begin(), pq.end(), node::node_pointer_cmp());
        pq.pop_back();

    }

    for (auto& e: q) std::cout << e.id << " " << e.distance << endl;




    return 0;
}
