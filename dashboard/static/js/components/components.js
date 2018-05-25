const Preloader = () => (
  <div className="lds-ripple">
    <div/>
    <div/>
  </div>
);

const Alert = ({text, close}) => (
  <div className="alert alert-primary alert-dismissible fade show">
    <span>{text}</span>

    <button
      className="close"
      onClick={close}
    >
      <span>&times;</span>
    </button>
  </div>
);

Alert.propTypes = {
  close: PropTypes.func.isRequired,
  text: PropTypes.string.isRequired
};

const Error = ({error, close}) => (
  <div className="alert alert-danger alert-dismissible fade show">
    <h4 className="alert-heading">Oops, some mistake</h4>

    <pre className="text-danger">
{`{
    code: ${error.code}
    status: ${error.status}
    message: ${error.message}
}`}
    </pre>

    <hr/>
    <p className="mb-0">Make sure that you are doing everything right</p>

    <button
      className="close"
      onClick={close}
    >
      <span>&times;</span>
    </button>
  </div>
);

// todo check prop-types everywhere
Error.propTypes = {
  close: PropTypes.func.isRequired,
  error: PropTypes.object.isRequired
};

const Header = () => (
  <header className="row mb-5 d-flex align-items-center">
    <div className="col-md-3">
      <a
        className="display-4 text-dark"
        href="/dashboard"
      >
        Pandora
      </a>
    </div>

    <div className="col-md-9">
      <div className="progress">
        <div className="progress-bar w-100"/>
      </div>
    </div>
  </header>
);

const Footer = () => (
  <footer
    className="d-none d-md-flex pt-5 pb-5 align-items-center justify-content-center fixed-bottom bg-white"
  >
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
    <a
      href="https://github.com/spacelavr/pandora"
      target="_blank"
    >
      <i className="fas fa-heartbeat"/>Pandora
    </a>
  </footer>
);
