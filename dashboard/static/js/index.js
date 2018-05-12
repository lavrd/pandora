const App = () => (
  <section className="container-fluid">
    <section className="container mt-5">
      <Header/>

      <div className="row mb-5">
        <div className="col-md-7 mb-3">
          <Certificate/>
        </div>

        <div className="col-md-5">
          <Account/>
        </div>
      </div>

      <Footer/>
    </section>
  </section>
);

ReactDOM.render(
  App(),
  document.getElementById('root')
);
