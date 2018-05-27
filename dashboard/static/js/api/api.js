const endpoint = 'https://127.0.0.1:2004';
const route_cert = 'cert';
const route_member = 'member';

class api {
  static MemberCreate(data) {
    return request('POST', `${endpoint}/${route_member}/create`, data);
  }

  static MemberFetch(data) {
    return request('POST', `${endpoint}/${route_member}/fetch`, data);
  }

  static CertCreate(data) {
    return request('POST', `${endpoint}/${route_cert}/issue`, data);
  }

  static CertFetch(data) {
    return request('POST', `${endpoint}/${route_cert}/view`, data);
  }

  static CertVerify(data) {
    return request('POST', `${endpoint}/${route_cert}/verify`, data);
  }
}
