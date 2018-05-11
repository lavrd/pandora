const App = () => (
  <section className="container-fluid">
    <section className="container mt-5">
      <Header/>

      <div className="row">
        <div className="col-md-7 mb-5">
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
