const Preloader = () => (
  <div className="lds-ripple">
    <div/>
    <div/>
  </div>
);

const Alert = ({text}) => (
  <div className="alert alert-primary" role="alert">{text}</div>
);

Alert.propTypes = {
  text: PropTypes.string.isRequired
};

const Error = ({error}) => (
  <div className="alert alert-danger">
    <h4 className="alert-heading">Error</h4>

    <code className="text-danger">
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

const Header = () => (
  <header className="row mb-5 d-flex align-items-center">
    <div className="col-md-3">
      <h1 className="display-4">Pandora</h1>
    </div>

    <div className="col-md-9">
      <div className="progress">
        <div className="progress-bar w-100"/>
      </div>
    </div>
  </header>
);

const Footer = () => (
  <footer className="mt-5 text-center">
    <i className="fas fa-university"/>
    <i className="fas fa-plus"/>
    <i className="fas fa-user-graduate"/>
    <i className="fas fa-caret-right"/>
    <i className="fas fa-money-check"/>
    <i className="fas fa-caret-right"/>
    <i className="fas fa-box"/>
    <i className="fas fa-link"/>
    <i className="fas fa-box"/>
    <i className="fas fa-caret-right"/>
    <i className="fas fa-handshake"/>
    <i className="fas fa-caret-right"/>
    <i className="fas fa-heartbeat"/>Pandora
  </footer>
);
