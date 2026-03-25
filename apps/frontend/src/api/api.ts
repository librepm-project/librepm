import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';

const apiCall = (): AxiosInstance => {
    const config: AxiosRequestConfig = {
      baseURL: '/api/',
    };

    const accessToken = window.localStorage.getItem("accessToken");
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
          window.localStorage.removeItem("accessToken");
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