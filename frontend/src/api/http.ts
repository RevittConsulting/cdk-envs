import { AxiosRequestConfig } from "axios"
import { ApiResponse } from "./types"
import { handleResponse } from "./utils"
import { axiosInstance } from "./axios-instance"

async function get<T>(
	url: string,
	config?: AxiosRequestConfig
): Promise<ApiResponse<T>> {
	return await handleResponse<T>(async () => await axiosInstance.get(url, config))
}

async function post<T>(
	url: string,
	data?: any,
	config?: AxiosRequestConfig
): Promise<ApiResponse<T>> {
	return await handleResponse<T>(
		async () => await axiosInstance.post(url, data, config)
	)
}

async function patch<T>(
	url: string,
	data?: any,
	config?: AxiosRequestConfig
): Promise<ApiResponse<T>> {
	return await handleResponse<T>(
		async () => await axiosInstance.patch(url, data, config)
	)
}

export const http = {
	get,
	post,
	patch,
}