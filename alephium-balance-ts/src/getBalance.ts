import fetch from 'node-fetch';

/**
 * Fetches the ALPH balance of a given Alephium address.
 * @param address - The Alephium address.
 * @returns The balance in ALPH.
 */
export async function getBalance(address: string): Promise<number> {
    const response = await fetch(`https://backend.mainnet.alephium.org/addresses/${address}`);
    const data = await response.json();

    if (!response.ok) {
        if (data.error) {
            throw new Error(`API Error: ${data.error}`);
        } else {
            throw new Error(`Error fetching balance: ${response.statusText}`);
        }
    }

    if (!data.balance) {
        throw new Error('Invalid response format');
    }

    const balance = parseFloat(data.balance);
    if (isNaN(balance)) {
        throw new Error('Balance is not a number');
    }

    return balance;
}
