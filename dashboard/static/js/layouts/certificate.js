class Certificate extends React.Component {
  constructor() {
    super();
    this.state = {
      state: this.STATE.FETCH,
      pending: false,
      error: null,
      success: '',
      data: this.EMPTY_DATA,
      cert: null,
      verifyStatus: this.VERIFY_STATUS.NONE
    };
  }

  VERIFY_STATUS = {
    NONE: 'NONE',
    VERIFIED: 'VERIFIED',
    FAILED: 'FAILED'
  };

  EMPTY_DATA = {
    title: '',
    description: '',
    publicKey: '',
    id: ''
  };

  STATE = {
    FETCH: 'FETCH',
    CREATE: 'CREATE'
  };

  handleState = (e, key) => {
    e.preventDefault();
    this.setState({state: key});
  };

  handleChange = (e) => {
    const target = e.target;
    const name = target.name;
    const value = target.value;
    this.setState({data: {...this.state.data, [name]: value}});
  };

  handleCreate = () => {
    api.CertCreate({
      public_Key: this.state.data.publicKey,
      title: this.state.data.title,
      description: this.state.data.description
    })
      .then(() => this.setState({success: 'Certificate successfully confirmed', pending: false}))
      .catch((error) => this.setState({error: error, pending: false}));
    this.setState({pending: true});
  };

  handleClose = () => {
    this.setState({
      cert: null,
      error: null,
      success: null,
      verifyStatus: this.VERIFY_STATUS.NONE,
      data: this.EMPTY_DATA
    });
  };

  handleVerify = () => {
    api.CertVerify(this.state.cert)
      .then(() => this.setState({verifyStatus: this.VERIFY_STATUS.VERIFIED, pending: false}))
      .catch(() => this.setState({verifyStatus: this.VERIFY_STATUS.FAILED, pending: false}));
    this.setState({pending: true});
  };

  handleFetch = () => {
    api.CertFetch({
      id: this.state.data.id
    })
      .then((cert) => this.setState({cert: cert, pending: false}))
      .catch((error) => this.setState({error: error, pending: false}));
    this.setState({pending: true});
  };

  render() {
    if (this.state.pending) return <Preloader/>;
    if (this.state.error) return <Error error={this.state.error} close={this.handleClose}/>;
    if (this.state.success) return <Alert text={this.state.success} close={this.handleClose}/>;
    if (this.state.cert) return <Cert cert={this.state.cert} verify={this.handleVerify}
                                      close={this.handleClose} verifyStatus={this.state.verifyStatus}/>;

    return (
      <div className="card shadow">
        <div className="card-header">
          <ul className="nav nav-pills nav-fill">
            {
              Object.keys(this.STATE).map((key, index) => (
                <li
                  className="nav-item"
                  key={index}
                >
                  <a
                    className={`nav-link ${this.state.state === key && 'active'}`}
                    onClick={(e) => this.handleState(e, key)}
                    href="#"
                  >
                    {key.capitalize()}
                  </a>
                </li>
              ))
            }
          </ul>
        </div>

        <div className="card-body">
          {
            this.state.state === this.STATE.FETCH ?
              <CertFetch
                submit={this.handleFetch}
                change={this.handleChange}
                id={this.state.data.id}
              /> :
              <CertCreate
                submit={this.handleCreate}
                change={this.handleChange}
                publicKey={this.state.data.publicKey}
                description={this.state.data.description}
                title={this.state.data.title}
              />
          }
        </div>
      </div>
    );
  }
}

const CertMemberCard = ({name, public_key}) => (
  <div className="card">
    <div className="card-body">
      <h5 className="card-title">{name}</h5>
      <small className="text-muted">{public_key}</small>
    </div>
  </div>
);

CertMemberCard.propTypes = {
  name: PropTypes.string.isRequired,
  public_key: PropTypes.string.isRequired
};

const CertCard = ({cert, close, verify, verifyStatus}) => (
  <div className="card shadow text-center">
    <div className="card-header">Certificate</div>

    <div className="card-body">
      <h5 className="card-title">{cert.meta.title}</h5>
      <p className="card-text">{cert.meta.description}</p>
      <p className="card-text">{new Date(cert.meta.timestamp).toString()}</p>

      <div className="row">
        <div className="col-6">
          <CertMemberCard
            name={cert.issuer.name}
            public_key={cert.issuer.public_key.public_key}
          />
        </div>

        <div className="col-6">
          <CertMemberCard
            name={cert.recipient.name}
            public_key={cert.recipient.public_key.public_key}
          />
        </div>
      </div>
    </div>

    <div className="card-footer d-flex align-items-center justify-content-between">
      <button
        onClick={verify}
        className={`btn float-left ${verifyStatus === VERIFY_STATUS.NONE ?
          'btn-primary' : verifyStatus === VERIFY_STATUS.VERIFIED ? 'btn-success' : 'btn-danger' }`}
      >
        <i className="fas fa-check"/>
      </button>

      <button
        className="close"
        onClick={close}
      >
        <span>&times;</span>
      </button>
    </div>
  </div>
);

Cert.propTypes = {
  verifyStatus: PropTypes.string.isRequired,
  verify: PropTypes.func.isRequired,
  cert: PropTypes.object.isRequired,
  close: PropTypes.func.isRequired
};

const CertCreateCard = ({publicKey, title, description, submit, change}) => (
  <section>
    <div className="form-group">
      <label>Public key</label>
      <input
        name="publicKey"
        onChange={change}
        value={publicKey}
        type="text"
        className="form-control"
        placeholder="787c8ef36e46f02a58f014ac7507c27fb29e757d0ca323ffd8d517ec70e3caa9"
      />
    </div>

    <div className="form-group">
      <label>Title</label>
      <input
        name="title"
        onChange={change}
        value={title}
        type="text"
        className="form-control"
        placeholder="Pandora certificate"
      />
    </div>

    <div className="form-group">
      <label>Description</label>
      <textarea
        name="description"
        onChange={change}
        value={description}
        rows={2}
        className="form-control"
        placeholder="For the successful completion of the courses"
      />
    </div>

    <button
      className="btn btn-primary"
      onClick={submit}
    >
      Submit
    </button>
  </section>
);

CertCreateCard.propTypes = {
  title: PropTypes.string.isRequired,
  description: PropTypes.string.isRequired,
  publicKey: PropTypes.string.isRequired,
  change: PropTypes.func.isRequired,
  submit: PropTypes.func.isRequired
};

const CertFetchCard = ({change, id, submit}) => (
  <section>
    <div className="form-group">
      <label>Id</label>
      <input
        name="id"
        onChange={change}
        value={id}
        type="text"
        className="form-control"
        placeholder="068b7dfa-26fa-4716-ba20-cb0943a8486a"
      />
    </div>

    <button
      className="btn btn-primary"
      onClick={submit}
    >
      Submit
    </button>
  </section>
);

CertFetchCard.propTypes = {
  submit: PropTypes.func.isRequired,
  id: PropTypes.string.isRequired,
  change: PropTypes.func.isRequired
};
