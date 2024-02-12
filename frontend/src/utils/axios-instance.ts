import axios from 'axios';

const API_URL = "http://localhost:80/api/v1";

export const axiosInstance = axios.create({
  baseURL: API_URL,
});