import axios, { Method } from 'axios';

// 接口返回数据格式
export interface HTTPResponse<T> {
  code: number;
  data: T;
  status: boolean;
  message: string;
}

export interface Ping {
  Msg: string;
  Ip: string;
  RequestId: string;
}

// 接口基础地址
export const ApiBaseURL = '/api';
export const accessTokenKey = 'ACCESS_TOKEN';

export class ApiRequest {
  static async request<T>(
    url: string,
    method: Method,
    param: Record<string, unknown>,
    body: Record<string, unknown>
  ): Promise<HTTPResponse<T>> {
    try {
      // 获取access_token
      const accessToken = localStorage.getItem(accessTokenKey);
      const { data } = await axios.request({
        baseURL: ApiBaseURL,
        url: url,
        method: method,
        headers: {
          'Content-Type': 'application/json',
        },
        params: {
          access_token: accessToken,
          ...param,
        },
        data: {
          ...body,
        },
      });
      const responseData = data as HTTPResponse<T>;
      if (responseData.code < 200 || responseData.code > 300) {
        // http code 错误
        throw new Error(`response code ${responseData.code} error`);
      }
      if (!responseData.status) {
        // 接口操作失败
        throw new Error(
          `response status is false, message: ${responseData.message}`
        );
      }
      return responseData;
    } catch (error) {
      console.error(`http request fail: ${error}`);
      throw error;
    }
  }
}
