import { axiosInstance } from '@/utils/axios-instance'
import { Count, KeyValuePairString } from "@/types/buckets";

export const getBuckets = async (): Promise<string[]> => {
  try {
    const response = await axiosInstance.get<string[]>("/buckets");
    if (response.status === 200) {
      return response.data;
    }
  } catch (error) {
    console.error(error);
    return [];
  }
  return [];
};

export const getCount = async (bucket: string): Promise<Count> => {
  try {
    const response = await axiosInstance.get<Count>(`/buckets/${bucket}/count`);
    if (response.status === 200) {
      return response.data;
    }
  } catch (error) {
    console.error(error);
    return { count: 0 };
  }
  return { count: 0 };
}

export const getData = async (): Promise<string[]> => {
  try {
    const response = await axiosInstance.get<string[]>("/buckets/data");
    if (response.status === 200) {
      return response.data;
    }
  } catch (error) {
    console.error(error);
    return [];
  }
  return [];
}

export const changeDbSource = async (file: string): Promise<void> => {
  try {
    const response = await axiosInstance.post("/buckets", { path: file });
    if (response.status === 200) {
      return;
    }
  } catch (error) {
    console.error(error);
  }
}

export const loadPages = async (bucket: string, pageNumber: number, resultsNumber: number): Promise<KeyValuePairString[]> => {
  try {
    const response = await axiosInstance.get<KeyValuePairString[]>(`/buckets/${bucket}/pages/${pageNumber}/${resultsNumber}`);
    if (response.status === 200) {
      return response.data;
    }
  } catch (error) {
    console.error(error);
    return [];
  }
  return [];
}

export const loadKeys = async (bucket: string, key: string): Promise<KeyValuePairString[]> => {
  try {
    const response = await axiosInstance.get<KeyValuePairString[]>(`/buckets/${bucket}/keys/${key}`);
    console.log(response);
    if (response.status === 200) {
      return response.data;
    }
  } catch (error) {
    console.error(error);
    return [];
  }
  return [];
}