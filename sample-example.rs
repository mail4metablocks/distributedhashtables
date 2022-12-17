use libp2p::kad::record::{self, Key};
use libp2p::kad::{store::MemoryStore, Kademlia};
use std::env;
use std::time::Duration;

#[tokio::main]
async fn main() {
    // Parse command line arguments
    let key = env::args().nth(1).unwrap();
    let value = env::args().nth(2).unwrap();

    // Create a Kademlia DHT with a memory store
    let mut dht = Kademlia::new(env::args().nth(3).unwrap().parse().unwrap(), MemoryStore::new());

    // Insert the (key, value) pair into the DHT
    dht.store(Key::new(key.as_bytes()), value.into(), Duration::from_secs(3600)).await;

    // Look up the value for the given key
    let record = dht.get(Key::new(key.as_bytes())).await.unwrap();
    let value = record::value(&record).unwrap();
    println!("Retrieved value: {}", value);
}
