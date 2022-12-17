import Libp2p from 'libp2p';
import Kademlia from 'libp2p-kad-dht';
import { Key } from 'libp2p-kad-dht/src/record';

async function main() {
  // Create a libp2p node with the Kademlia DHT module
  const node = await Libp2p.create({
    modules: {
      transport: [],
      streamMuxer: [],
      connEncryption: [],
      dht: Kademlia,
    },
  });

  // Start the node
  await node.start();

  // Insert a (key, value) pair into the DHT
  const key = new Key('my-key');
  const value = Buffer.from('my-value');
  await node.dht.put(key, value);

  // Look up the value for the given key
  const record = await node.dht.get(key);
  const retrievedValue = record.value;
  console.log(`Retrieved value: ${retrievedValue}`);

  // Stop the node
  await node.stop();
}

main();
