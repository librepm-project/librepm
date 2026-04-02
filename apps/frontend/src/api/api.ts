import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
import { getToken, removeToken } from '@/lib/cookie';

const apiCall = (): AxiosInstance => {
    const config: AxiosRequestConfig = {
      baseURL: '/api/',
    };

    const accessToken = getToken();
    if (accessToken) {
      config.headers = {
        Authorization: `Bearer ${accessToken}`,
      };
    }
    const instance = axios.create(config);

    instance.interceptors.response.use(
      (response) => response,
      (error) => {
        if (error.response?.status === 401) {
          removeToken();
          window.location.href = '/';
        }
        if (error.response?.status === 403) {
          window.location.href = '/forbidden';
        }
        return Promise.reject(error);
      }
    );

    return instance;
  }

const api = apiCall();

export const get = async <T>(url: string, params?: unknown): Promise<T> => {
    const response = await api.get<T>(url, { params });
    return response.data;
};

export const post = async <T>(url: string, data?: unknown): Promise<T> => {
    const response = await api.post<T>(url, data);
    return response.data;
};

export const put = async <T>(url: string, data?: unknown): Promise<T> => {
    const response = await api.put<T>(url, data);
    return response.data;
};

export const del = async <T>(url: string, data?: unknown): Promise<T> => {
    const response = await api.delete<T>(url, { data });
    return response.data;
};

export default {
    apiCall,
    get,
    post,
    put,
    del,
}
