
import axios, { AxiosInstance, AxiosRequestConfig } from "axios"

const defaultConfig = {
    timeout: 10000,
    baseUrl: ""
}

export type DefaultConfig = typeof defaultConfig

class Request {
    private instance: AxiosInstance
    constructor(config: DefaultConfig = defaultConfig) {
        this.instance = axios.create(config)
        this.interceptorsRequest()
        this.interceptorsResponse()
    }

    // 请求拦截
    private interceptorsRequest() {
        this.instance.interceptors.request.use(
            config => {
                return config
            },
            err => {
                return Promise.reject(err)
            }
        )
    }

    // 响应拦截
    private interceptorsResponse() {
        this.instance.interceptors.response.use(
            response => {
                return response
            },
            err => {
                return Promise.reject(err)
            }
        )
    }

    // Get 请求
    public Get<T>(url: string, params: AxiosRequestConfig):Promise<T> {
        // TODO: 处理 catch
        return this.instance.get(url, params).then(res=>res.data).catch()
    }

    // Post 请求
    public Post<T>(url: string, params: AxiosRequestConfig):Promise<T> {
        // TODO: 处理 catch
        return this.instance.get(url, params).then(res=>res.data).catch()
    }
}

export const http = new Request()
