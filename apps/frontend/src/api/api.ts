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
          window.location.href = "/login";
        }
        return Promise.reject(error);
      }
    );

    return instance;
  }


  export default {
    apiCall
  }
