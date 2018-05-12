class api {
  static MemberCreate(data) {
    return request('POST', 'http://localhost:2004/account/create', data);
  }

  static MemberFetch(data) {
    return request('POST', 'http://localhost:2004/account/fetch', data);
  }
}
