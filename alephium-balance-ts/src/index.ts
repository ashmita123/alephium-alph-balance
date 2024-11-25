import { getBalance } from './getBalance';

const address = process.argv[2];

if (!address) {
    console.error('Please provide an Alephium address.');
    process.exit(1);
}

getBalance(address)
    .then(balance => {
        console.log(`Balance of ${address}: ${balance} ALPH`);
    })
    .catch(error => {
        console.error(`Error: ${error.message}`);
        process.exit(1);
    });
