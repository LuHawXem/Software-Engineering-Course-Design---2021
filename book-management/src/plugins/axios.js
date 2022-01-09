"use strict";

import Vue from 'vue';
import axios from "axios";
import qs from "qs";

// Full config:  https://github.com/axios/axios#request-config
// axios.defaults.baseURL = process.env.baseURL || process.env.apiUrl || '';
// axios.defaults.headers.common['Authorization'] = AUTH_TOKEN;
// axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';

let config = {
  headers: {
    'content-type': 'application/x-www-form-urlencoded'
  }
  // baseURL: process.env.baseURL || process.env.apiUrl || ""
  // timeout: 60 * 1000, // Timeout
  // withCredentials: true, // Check cross-site Access-Control
};

const _axios = axios.create(config);

_axios.interceptors.request.use(
  function(config) {
    // Do something before request is sent
    config.data = qs.stringify(config.data)
    return config;
  },
  function(error) {
    // Do something with request error
    return Promise.reject(error);
  }
);

// Add a response interceptor
_axios.interceptors.response.use(
  function(response) {
    if(response.status >= 200 && response.status < 400) {
      return response.data;
    }
    // Do something with response data
    return response;
  },
  function(error) {
    // Do something with response error
    return Promise.reject(error.response);
  }
);

if(process.env.NODE_ENV=="production") {
  _axios.defaults.baseURL = 'https://bookserver.luhawxem.cn:5000'
}
else {
  _axios.defaults.baseURL = '/api'
}

// eslint-disable-next-line no-unused-vars
Plugin.install = function(Vue, options) {
  Vue.axios = _axios;
  window.axios = _axios;
  Object.defineProperties(Vue.prototype, {
    axios: {
      get() {
        return _axios;
      }
    },
    $axios: {
      get() {
        return _axios;
      }
    },
  });
};

Vue.use(Plugin)

export default Plugin;
