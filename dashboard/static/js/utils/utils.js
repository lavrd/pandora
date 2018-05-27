const request = (method, endpoint, body) => {
  let headers = {};
  if (!!body) {
    headers['Content-Type'] = 'application/json';
  }

  let opts = {};
  opts.method = method;
  opts.headers = headers;
  if (!!body) {
    opts.body = JSON.stringify(body);
  }

  return fetch(endpoint, opts)
    .then(response => {
      if (response.status >= 200 && response.status < 300) {
        return response.json().then((res) => {
          return res;
        }).catch(() => {
          return response;
        });
      }
      return response.json().then((e) => {
        throw e;
      });
    });
};

String.prototype.capitalize = function () {
  return this.charAt(0).toUpperCase() + this.slice(1).toLowerCase();
};

const VERIFY_STATUS = {
  NONE: 'NONE',
  VERIFIED: 'VERIFIED',
  FAILED: 'FAILED'
};

const STATE = {
  FETCH: 'FETCH',
  CREATE: 'CREATE'
};
