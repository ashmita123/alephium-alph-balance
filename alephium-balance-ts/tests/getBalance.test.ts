import { getBalance } from '../src/getBalance';
import fetchMock from 'jest-fetch-mock';

beforeEach(() => {
  fetchMock.resetMocks();
});

describe('getBalance', () => {
  it('should return a number for a valid address', async () => {
    fetchMock.mockResponseOnce(JSON.stringify({ balance: '123.45' }));

    const address = '1H5pLrNnJCPh6snKkp8qodh2Tv3nJ2bYX5zD1AvStZAUG';
    const balance = await getBalance(address);
    expect(balance).toBe(123.45);
  });

  it('should throw an error for an invalid address', async () => {
    fetchMock.mockResponseOnce(JSON.stringify({ error: 'Invalid address' }), { status: 400 });

    const address = 'invalid_address';
    await expect(getBalance(address)).rejects.toThrow('API Error: Invalid address');
  });
});
