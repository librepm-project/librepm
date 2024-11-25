import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';

const apiCall = (): AxiosInstance => {
    const config: AxiosRequestConfig = {
      baseURL: '/api/',
    };

    const accessToken = window.localStorage.getItem("accessToken");
    if (accessToken) {
      config.headers = {
        Authorization: `${accessToken}`,
      };
    }
    const instance = axios.create(config);

    instance.interceptors.response.use(
      (response) => response,
      (response) => {
        if (response.code === "ERR_BAD_REQUEST") {
          window.localStorage.removeItem("accessToken");
          window.location.href = "/";
        }
        return response;
      }
    );

    return instance;
  }


  export default {
    apiCall
  }