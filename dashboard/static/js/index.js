const request = (method, url, body) => {
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

  return fetch(url, opts)
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
  return this.charAt(0).toUpperCase() + this.slice(1);
};

const Preloader = () => (
  <div className="lds-ripple">
    <div/>
    <div/>
  </div>
);

const Error = ({error}) => (
  <div className="alert alert-danger">
    <h4 className="alert-heading">Error</h4>

    <code>
      {`
        {
          "code": ${error.code},\n
          "status": ${error.status},\n
          "message": ${error.message}
        }
      `}
    </code>

    <hr/>
    <p className="mb-0">Check the server or input data and try again</p>
  </div>
);

Error.propTypes = {
  error: PropTypes.object.isRequired
};

class Certificate extends React.Component {
  constructor() {
    super();
    this.state = {};
  }

  render() {
    return (
      <Preloader/>
    );
  }
}

class AccountCreate extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      email: '',
      name: ''
    };
  }

  handleChange = (e) => {
    const target = e.target;
    const name = target.name;
    const value = target.value;

    this.setState({[name]: value, error: null});
  };

  render() {
    return (
      <section>
        <div className="form-group">
          <label>Email</label>
          <input
            name="email"
            onChange={this.handleChange}
            value={this.state.email}
            type="email"
            className="form-control"
            placeholder="Enter email"
          />
        </div>

        <div className="form-group">
          <label>Name</label>
          <input
            name="name"
            onChange={this.handleChange}
            value={this.state.name}
            type="text"
            className="form-control"
            placeholder="Sergey Lavrentev"
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

AccountCreate.propTypes = {
  submit: PropTypes.func.isRequired
};

class AccountFetch extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      publicKey: ''
    };
  }

  handleChange = (e) => {
    const target = e.target;
    const name = target.name;
    const value = target.value;

    this.setState({[name]: value, error: null});
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
            placeholder="Public key"
          />
        </div>

        <button
          className="btn btn-primary"
          onClick={this.handleSubmit}
        >
          Submit
        </button>
      </section>);
  }
}

AccountFetch.propTypes = {
  submit: PropTypes.func.isRequired
};

class Account extends React.Component {
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

  handleState = () => {
    const state = this.state.state === this.STATE.FETCH ? this.STATE.CREATE : this.STATE.FETCH;
    this.setState({state: state});
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
      <section>
        <button
          onClick={this.handleState}
          className="btn btn-primary mb-5"
        >
          {
            this.STATE[this.state.state].toLowerCase().capitalize()
          }
        </button>

        {
          this.state.state === this.STATE.FETCH ?
            <AccountFetch submit={this.handleFetch}/> : <AccountCreate submit={this.handleCreate}/>
        }
      </section>
    );
  }
}

class App extends React.Component {
  constructor() {
    super();
    this.state = {};
  }

  render() {
    return (
      <section className="container-fluid">
        <section className="container mt-5">
          <div className="row">
            <div className="col-8">
              <Certificate/>
            </div>

            <div className="col-4">
              <Account/>
            </div>
          </div>
        </section>
      </section>
    );
  }
}

ReactDOM.render(
  <App/>,
  document.getElementById('root')
);
