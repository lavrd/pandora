class api {
  static MemberCreate(data) {
    return request('POST', 'http://localhost:2004/account/create', data);
  }

  static MemberFetch(data) {
    return request('POST', 'http://localhost:2004/account/fetch', data);
  }

  static CertCreate(data) {
    return request('POST', 'http://localhost:2004/cert/issue', data);
  }

  static CertFetch(data) {
    return request('POST', 'http://localhost:2004/cert/view', data);
  }

  static CertVerify(data) {
    return request('POST', 'http://localhost:2004/cert/verify', data);
  }
}
