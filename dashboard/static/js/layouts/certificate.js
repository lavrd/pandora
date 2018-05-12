class Certificate extends React.Component {
  constructor() {
    super();
    this.state = {
      state: this.STATE.FETCH,
      pending: false,
      error: null
    };
  }

  STATE = {
    FETCH: 'FETCH',
    CREATE: 'CREATE'
  };

  handleState = (e, key) => {
    e.preventDefault();
    this.setState({state: key});
  };

  handleCreate = () => {
    this.setState({pending: true});
  };

  handleFetch = () => {
    this.setState({pending: true});
  };

  render() {
    if (this.state.pending) return <Preloader/>;
    if (this.state.error) return <Error error={this.state.error}/>;

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
              <CertificateFetch submit={this.handleFetch}/> :
              <CertificateCreate submit={this.handleCreate}/>
          }
        </div>
      </div>
    );
  }
}

class CertificateCreate extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      publicKey: '',
      title: '',
      description: ''
    };
  }

  handleChange = (e) => {
    const target = e.target;
    const name = target.name;
    const value = target.value;
    this.setState({[name]: value});
  };

  render() {
    return (
      <section>
        <div className="form-group">
          <label>Public key</label>
          <input
            name="publicKey"
            onChange={this.handleChange}
            value={this.state.publicKey}
            type="text"
            className="form-control"
            placeholder="787c8ef36e46f02a58f014ac7507c27fb29e757d0ca323ffd8d517ec70e3caa9"
          />
        </div>

        <div className="form-group">
          <label>Title</label>
          <input
            name="title"
            onChange={this.handleChange}
            value={this.state.title}
            type="text"
            className="form-control"
            placeholder="Pandora certificate"
          />
        </div>

        <div className="form-group">
          <label>Description</label>
          <textarea
            name="description"
            onChange={this.handleChange}
            value={this.state.description}
            rows={2}
            className="form-control"
            placeholder="For the successful completion of the courses"
          />
        </div>

        <button
          className="btn btn-primary"
          onClick={this.props.submit}
        >
          Submit
        </button>
      </section>
    );
  }
}

CertificateCreate.proppb = {
  submit: Proppb.func.isRequired
};

class CertificateFetch extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: ''
    };
  }

  handleChange = (e) => {
    const target = e.target;
    const name = target.name;
    const value = target.value;
    this.setState({[name]: value});
  };

  render() {
    return (
      <section>
        <div className="form-group">
          <label>Id</label>
          <input
            name="id"
            onChange={this.handleChange}
            value={this.state.id}
            type="text"
            className="form-control"
            placeholder="068b7dfa-26fa-4716-ba20-cb0943a8486a"
          />
        </div>

        <button
          className="btn btn-primary"
          onClick={this.props.submit}
        >
          Submit
        </button>
      </section>
    );
  }
}

CertificateFetch.proppb = {
  submit: Proppb.func.isRequired
};
