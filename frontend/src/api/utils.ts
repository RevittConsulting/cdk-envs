import { AxiosResponse } from 'axios'
import { ApiResponse } from './types'

export async function handleResponse<T>(
  fn: () => Promise<AxiosResponse<T>>,
): Promise<ApiResponse<T>> {

  try {
    const response = await fn()
    const data: T = response.data
    const status: number = response.status
    const errorStatus: boolean = status.toString().substring(0, 1) !== '2'

    return {
      data,
      status,
      error: errorStatus ? data : null,
    }
  } catch (e: any) {
    if (!e.response) {
      console.error(e)
    }

    return {
      data: null,
      status: e.response.status,
      error: e.response.data,
    }
  }
}
