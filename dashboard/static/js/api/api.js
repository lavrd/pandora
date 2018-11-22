const endpoint = "https://127.0.0.1:2004";

class api {
  static MemberCreate(data) {
    return request("POST", `${endpoint}/member/create`, data);
  }

  static MemberFetch(data) {
    return request("POST", `${endpoint}/member/fetch`, data);
  }

  static CertCreate(data) {
    return request("POST", `${endpoint}/cert/issue`, data);
  }

  static CertFetch(data) {
    return request("POST", `${endpoint}/cert/view`, data);
  }

  static CertVerify(data) {
    return request("POST", `${endpoint}/cert/verify`, data);
  }
}
