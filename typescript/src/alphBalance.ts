import { NodeProvider } from '@alephium/web3';

/**
 * This function will get the ALPH balance for Alephium address.
 * @param address - Alephium address that we want balance for.
 * @param nodeUrl - This is the URL of Alephium node to connect. Default is the local node.
 * @returns The ALPH balance as a number.
 * @throws Error if the balance cannot be fetched.
 */
export async function getAlphBalance(address: string, nodeUrl: string = 'http://localhost:22973'): Promise<number> {
    const nodeProvider = new NodeProvider(nodeUrl);
    try {
        const balance = await nodeProvider.address.getBalance({ address });
        return balance.attoAlph / 1e18; // Change attoAlph to ALPH unit
    } catch (error) {
        throw new Error(`Failed to get ALPH balance: ${error}`);
    }
}
