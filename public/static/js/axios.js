var Axios = (function () {
  var service = axios.create({
    headers: {},
    withCredentials: true,
    timeout: 30000
  });

  // 请求拦截器
  service.interceptors.request.use(
    (config) => {
      const token = localStorage.getItem("__token__");

      if (/^\/api\/login/.test(config.url)) {
        // 登录
      }

      if (token && config["headers"]) {
        config["headers"]["Authorization"] = `Basic ${token}`;
      }
      return config;
    },
    (error) => {
      // error.message = "服务器异常，请联系管理员！";
      // 错误抛到业务代码
      return Promise.reject(error);
    }
  );

  // 响应拦截器
  service.interceptors.response.use(
    (response) => {
      const status = response.status;
      const config = response.config;

      if (
        /^\/api\/login/.test(config.url) &&
        status === 200 &&
        response.data.code === 0
      ) {
        localStorage.setItem("__token__", response.data.data);
      }

      if (status < 200 || (status >= 300 && status != 401 && status != 500)) {
        // 处理http错误，抛到业务代码
        return response.data;
      } else if (status == 200) {
        return response.data;
      } else if (status == 500) {
        return response.data;
      }
    },
    (error) => {
      if (error.response.status === 401) {
        location.href = "/login";
      }

      // err.message = "请求超时或服务器异常，请检查网络或联系管理员！";
      return Promise.reject(error);
    }
  );

  window.Axios = service;
  return service;
})(window);
