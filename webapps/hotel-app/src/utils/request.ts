import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from "axios";
import toast from "react-hot-toast";

export type ToastType = "error" | "success";

const defaultConfig = {
    timeout: 10000,
    baseURL: import.meta.env.VITE_BASE_API,
};

const defaultToastConfig = {
    error: true,
    success: true,
};

export type DefaultConfig = typeof defaultConfig;
export type ToastConfig = typeof defaultToastConfig;

class Request {
    private instance: AxiosInstance;
    constructor(config: DefaultConfig = defaultConfig) {
        this.instance = axios.create(config);
        this.interceptorsRequest();
        this.interceptorsResponse();
    }

    // 请求拦截
    private interceptorsRequest() {
        this.instance.interceptors.request.use(
            (config) => {
                return config;
            },
            (err) => {
                return Promise.reject(err);
            },
        );
    }

    // 响应拦截
    private interceptorsResponse() {
        this.instance.interceptors.response.use(
            (response) => {
                return response;
            },
            (err) => {
                return Promise.reject(err);
            },
        );
    }

    // Get 请求
    public Get<T>(
        url: string,
        config: AxiosRequestConfig,
        toastConfig: ToastConfig = defaultToastConfig,
    ): Promise<T> {
        // TODO: 处理 catch
        return this.instance
            .get(url, config)
            .then((res) => {
                return this.handleRes(res, toastConfig);
            })
            .catch((error) => {
                this.handleError(error, toastConfig);
            });
    }

    // Post 请求
    public Post<T>(
        url: string,
        config: AxiosRequestConfig,
        toastConfig: ToastConfig = defaultToastConfig,
    ): Promise<T> {
        return this.instance
            .post(url, config.data, config)
            .then((res) => {
                return this.handleRes(res, toastConfig);
            })
            .catch((error) => {
                // axios请求状态非2xx会被catch，还有return Promise.reject(res.data) 。。。
                return this.handleError(error, toastConfig);
            });
    }

    private handleRes(res: AxiosResponse<any, any>, toastConfig: ToastConfig) {
        // 仅有请求状态是2xx并且code=10000才返回正确的数据
        if (res.data?.code === 10000) {
            if (toastConfig.success) {
                this.handleToast(res.data?.msg || "请求成功", "success");
            }
            return res.data;
        }

        if (toastConfig.error) {
            // 处理请求状态码=2xx和业务code!=10000的情况
            let msg = res.data.msg || "请求错误";
            this.handleToast(msg, "error");
        }

        return Promise.reject(res.data);
    }

    private handleToast(msg: string, type: ToastType) {
        toast[type](msg);
    }

    // 请求状态非2xxx、reject状态、其他错误 逻辑处理
    private handleError(error: any, toastConfig: ToastConfig) {
        if (axios.isAxiosError(error)) {
            console.error("Axios error:", error);
            if (toastConfig.error) {
                let msg = error.response?.data?.msg || "请求错误";
                this.handleToast(msg, "error");
            }
        } else {
            // 这里error主要是 Promise.reject(res.data) 返回的数据
            console.error("General error:", error);
        }

        // 必须返回Promise.reject，让调用方也能捕获到错误（是否做额外处理由调用方决定）
        return Promise.reject(error);
    }
}

export const http = new Request();
