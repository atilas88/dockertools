// Simulates the latency of a Go binding and returns a fresh copy of the
// data, so mutations in the UI never touch the mock source.
export function mockResponse<T>(data: T, ms = 250): Promise<T> {
  return new Promise(resolve => setTimeout(() => resolve(structuredClone(data)), ms))
}

export const ago = (seconds: number) => Math.floor(Date.now() / 1000) - seconds
