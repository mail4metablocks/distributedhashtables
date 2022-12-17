# Distributed HashTables

Distributed Hash Tables (DHTs) are a type of distributed system that provides a lookup service similar to a hash table: (key, value) pairs are stored in the DHT, and any participating node can efficiently retrieve the value associated with a given key. DHTs are a distributed data structure that is designed to scale to a large number of nodes and to handle frequent node arrivals, departures, and failures.

In a DHT, each node is responsible for storing and forwarding (key, value) pairs within a certain range of keys. The mapping of keys to nodes is called the key-to-node mapping or the routing table. The key-to-node mapping is defined by a hash function, which maps keys to nodes in a way that is (ideally) evenly distributed across the nodes in the DHT.

One of the main benefits of DHTs is that they allow for efficient and decentralized data storage and retrieval. Because each node is responsible for storing and forwarding (key, value) pairs within a certain range of keys, there is no need for a central server or index to handle lookups. This makes DHTs highly scalable and resilient to failures, as the loss of any single node will not disrupt the system as a whole.

There are several different types of DHTs, including Chord, Kademlia, and Pastry. Each type of DHT has its own specific characteristics and trade-offs, but they all share the basic principles of distributed data storage and retrieval described above                         

## Implementation
The examples creates a Kademlia DHT with a memory store, inserts a (key, value) pair into the DHT, and then retrieves the value for the given key. The key and value are passed as command line arguments, and the local node's address is also passed as a command line argument.

The client initiates a lookup operation by sending a lookup(key) request to the DHT. The DHT nodes then use a recursive find_node(key) operation to search for the node responsible for the given key. When the responsible node is found, the client sends a get_value(key) request to retrieve the value for the given key. The responsible node responds with the value, which is then returned to the client.

### Sequence diagram

![image](https://user-images.githubusercontent.com/117555665/208231144-cd74a039-1f44-443d-9419-b0ecdbfb2938.png)
