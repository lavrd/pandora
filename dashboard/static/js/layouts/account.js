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
              <AccountFetch submit={this.handleFetch}/> :
              <AccountCreate submit={this.handleCreate}/>
          }
        </div>
      </div>
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
    this.setState({[name]: value});
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
            placeholder="thomas@mail.system"
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
            placeholder="Thomas Hobbs"
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

AccountCreate.proppb = {
  submit: Proppb.func.isRequired
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

        <button
          className="btn btn-primary"
          onClick={this.handleSubmit}
        >
          Submit
        </button>
      </section>
    );
  }
}

AccountFetch.proppb = {
  submit: Proppb.func.isRequired
};
